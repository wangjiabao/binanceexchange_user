package service

import (
	v1 "binanceexchange_user/api/binanceexchange_user/v1"
	"binanceexchange_user/internal/biz"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"time"
)

// BinanceUserService is a BinanceData service .
type BinanceUserService struct {
	v1.UnimplementedBinanceUserServer

	buc *biz.BinanceUserUsecase
}

// NewBinanceDataService new a BinanceData service.
func NewBinanceDataService(buc *biz.BinanceUserUsecase) *BinanceUserService {
	return &BinanceUserService{buc: buc}
}

func (b *BinanceUserService) SetUser(ctx context.Context, req *v1.SetUserRequest) (*v1.SetUserReply, error) {
	return nil, nil
}

func (b *BinanceUserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	return b.buc.GetUser(ctx, req)
}

func (b *BinanceUserService) PullUserStatus(ctx context.Context, req *v1.PullUserStatusRequest) (*v1.PullUserStatusReply, error) {
	var (
		err   error
		users []*biz.LhBinanceUser
	)
	users, err = b.buc.GetUsers()
	if nil != err {
		fmt.Println(err)
		return &v1.PullUserStatusReply{}, nil
	}

	for _, user := range users {
		var (
			res        float64
			userStatus *biz.LhBinanceUserStatus
		)

		userStatus, err = b.buc.GetUserStatus(user.ID)
		if nil != err {
			fmt.Println(err)
			continue
		}

		// 代币1
		res, err = pullStakeUserInfo(user.Address, "0xf6E73f9dF438Bf59D647812DD9506678Ccd07236") // tdc
		if nil != err {                                                                          // 本次查询异常
			fmt.Println(err)
			continue
		}

		// 代币1未开启 检查代币2
		if res < 0 {
			res, err = pullStakeUserInfo(user.Address, "0x0CA25ef27823356B314fBc57a32181f2A6a285e8") // ttdc
			if nil != err {                                                                          // 本次查询异常
				fmt.Println(err)
				continue
			}

			// 代币2也未开启
			if res < 0 {
				if nil != userStatus && "open" == userStatus.Status {
					fmt.Println("all close", user)
					_, err = b.buc.UpdateUserStatusClose(ctx, user.ID)
				}
			}
		}

		if 0 < res {
			// open
			if nil == userStatus {
				_, err = b.buc.InsertUserStatus(ctx, user.ID, res)
			} else if "close" == userStatus.Status || res > userStatus.BaseMoney || res < userStatus.BaseMoney {
				_, err = b.buc.UpdateUserStatusOpen(ctx, user.ID, res)
			}
		} else if 0 == res {
			// close
			if nil == userStatus {
				continue
			}

			if "close" == userStatus.Status {
				continue
			}

			fmt.Println("close")
			_, err = b.buc.UpdateUserStatusClose(ctx, user.ID)
		}

		if nil != err {
			fmt.Println(err)
			continue
		}
	}

	return b.buc.PullUserStatus(ctx, req)
}

func pullStakeUserInfo(address string, addressToken string) (float64, error) {
	var (
		usdtAmount       float64 = -1
		err              error
		client           *ethclient.Client
		instance         *Stake
		currentOpenToken common.Address
	)

	url1 := "https://bsc-dataseed4.binance.org/"

	for i := 0; i < 5; i++ {
		client, err = ethclient.Dial(url1)
		if err != nil {
			url1 = "https://bsc-dataseed1.bnbchain.org"
			continue
			//return usdtAmount, err
		}

		tokenAddress := common.HexToAddress("0x02226c139F83425CE0ac9EC1611Bf1728B99D4cF")
		instance, err = NewStake(tokenAddress, client)
		if err != nil {
			continue
			//return usdtAmount, err
		}

		// 检测是否开启当前开启的
		currentOpenToken, err = instance.UserCurrentOpen(
			&bind.CallOpts{},
			common.HexToAddress(address),
		)
		if err != nil {
			fmt.Println(err)
			continue
		}

		// 非认可代币
		if "0xf6E73f9dF438Bf59D647812DD9506678Ccd07236" != currentOpenToken.String() && "0x0CA25ef27823356B314fBc57a32181f2A6a285e8" != currentOpenToken.String() {
			return 0, nil
		}

		// 当前代币非本次查询的
		if addressToken != currentOpenToken.String() {
			break
		}

		var bal *big.Int
		bal, err = instance.UserMaxTime(
			&bind.CallOpts{},
			common.HexToAddress(address),
			common.HexToAddress(addressToken),
		)

		if err != nil {
			fmt.Println(err)
			continue
		}

		tmpNow := time.Now().UTC().Add(8 * time.Hour)
		if 0 >= bal.Int64() || tmpNow.Unix() >= bal.Int64() {
			// close
			return 0, nil
		}

		var bal2 *big.Int
		bal2, err = instance.UserUsdtAmount(
			&bind.CallOpts{},
			common.HexToAddress(address),
			common.HexToAddress(addressToken),
		)

		if err != nil {
			fmt.Println(err)
			continue
		}

		usdtAmount, _ = bal2.Float64()
		break
	}

	return usdtAmount, err
}

func (b *BinanceUserService) PullUserCredentialsBsc(ctx context.Context, req *v1.PullUserCredentialsBscRequest) (*v1.PullUserCredentialsBscReply, error) {
	var (
		users []string
		err   error
	)

	users, err = pullUsersBySystemRole()
	if nil != err {
		fmt.Println(err)
		return nil, err
	}

	for _, v := range users {
		var (
			apiKey    string
			apiSecret string
		)

		apiKey, apiSecret, err = pullUserCredentialsBscBySystemRole(v)
		if nil != err {
			fmt.Println(err)
			return nil, err
		}

		err = b.buc.SetUser(ctx, v, apiKey, apiSecret)
		if nil != err {
			fmt.Println(err)
			return nil, err
		}
	}

	return nil, nil
}

func pullUsersBySystemRole() ([]string, error) {
	var (
		users []string
	)

	url1 := "https://bsc-dataseed4.binance.org/"

	for i := 0; i < 5; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			url1 = "https://bsc-dataseed1.bnbchain.org"
			continue
			//return usdtAmount, err
		}

		tokenAddress := common.HexToAddress("0x4526Dbc1a86624f9A9bf99726F946278BCFb2e2B")
		instance, err := NewUserCredentialsBsc(tokenAddress, client)
		if err != nil {
			fmt.Println(err)
			return users, err
		}
		var (
			addresses []common.Address
		)

		addresses, err = instance.GetUsersBySystemRole(&bind.CallOpts{
			From: common.HexToAddress(""),
		})
		if err != nil {
			return users, err
		}

		for _, v := range addresses {
			users = append(users, v.String())
		}

		break
	}

	return users, nil
}

func pullUserCredentialsBscBySystemRole(address string) (string, string, error) {
	var (
		apiKey    string
		apiSecret string
	)

	url1 := "https://bsc-dataseed4.binance.org/"

	for i := 0; i < 5; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			url1 = "https://bsc-dataseed1.bnbchain.org"
			continue
			//return usdtAmount, err
		}

		tokenAddress := common.HexToAddress("0x4526Dbc1a86624f9A9bf99726F946278BCFb2e2B")
		instance, err := NewUserCredentialsBsc(tokenAddress, client)
		if err != nil {
			fmt.Println(err)
			return apiKey, apiSecret, err
		}

		apiKey, apiSecret, err = instance.GetUserCredentialsBscBySystemRole(&bind.CallOpts{
			From: common.HexToAddress(""),
		}, common.HexToAddress(address))
		if err != nil {
			return apiKey, apiSecret, err
		}

		break
	}

	return apiKey, apiSecret, nil
}
