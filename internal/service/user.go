package service

import (
	v1 "binanceexchange_user/api/binanceexchange_user/v1"
	"binanceexchange_user/internal/biz"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
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
	return b.buc.SetUser(ctx, req)
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
		res, err = pullStakeUserInfo(user.Address)
		if nil != err {
			fmt.Println(err)
			continue
		}

		if res < 0 {
			continue
		}

		userStatus, err = b.buc.GetUserStatus(user.ID)
		if nil != err {
			fmt.Println(err)
			continue
		}

		if 0 < res {
			// open
			if nil == userStatus {
				_, err = b.buc.InsertUserStatus(ctx, user.ID, res)
			} else if res > userStatus.BaseMoney || res < userStatus.BaseMoney {
				_, err = b.buc.UpdateUserStatusOpen(ctx, user.ID, res)
			}
		} else {
			// close
			if nil == userStatus {
				continue
			}

			_, err = b.buc.UpdateUserStatusClose(ctx, user.ID)
		}

		if nil != err {
			fmt.Println(err)
			continue
		}
	}

	return b.buc.PullUserStatus(ctx, req)
}

func pullStakeUserInfo(address string) (float64, error) {
	var (
		usdtAmount float64 = -1
	)

	url1 := "https://bsc-dataseed4.binance.org/"

	for i := 0; i < 5; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			url1 = "https://bsc-dataseed1.bnbchain.org"
			continue
			//return usdtAmount, err
		}

		tokenAddress := common.HexToAddress("0x4E29c650a0c793A8e39B8E1234D49acDf06e4292")
		instance, err := NewStake(tokenAddress, client)
		if err != nil {
			continue
			//return usdtAmount, err
		}

		bal, err := instance.UserMaxTime(
			&bind.CallOpts{},
			common.HexToAddress(address),
			common.HexToAddress("0xf1a03B357849Cf0Fec27f8D9731a48aC0205A63D"),
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

		bal2, err := instance.UserUsdtAmount(
			&bind.CallOpts{},
			common.HexToAddress(address),
			common.HexToAddress("0xf1a03B357849Cf0Fec27f8D9731a48aC0205A63D"),
		)

		if err != nil {
			fmt.Println(err)
			continue
		}

		usdtAmount, _ = bal2.Float64()
		break
	}

	return usdtAmount, nil
}
