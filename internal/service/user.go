package service

import (
	deposit "binanceexchange_user/abi"
	v1 "binanceexchange_user/api/binanceexchange_user/v1"
	"binanceexchange_user/internal/biz"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
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

// GetUser 客户端查看用户信息接口，主要是api能否下单使用的反馈
func (b *BinanceUserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	return b.buc.GetUser(ctx, req)
}

// PullUserDeposit 用户充值总入口, 单一协程处理，遍历所有用户
func (b *BinanceUserService) PullUserDeposit(ctx context.Context, req *v1.PullUserDepositRequest) (*v1.PullUserDepositReply, error) {
	var (
		users []string
		err   error
	)

	users, err = pullUsersByDeposit()
	for _, v := range users {
		var balance string
		balance, err = pullUserDepositForUsdtAmount(v)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = b.buc.SetUserBalanceAndUser(ctx, v, balance)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	return nil, nil
}

func pullUsersByDeposit() ([]string, error) {
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

		contractAddress := common.HexToAddress("0x8fBdCc0cDD99CE35c1935D8a69286e32b0FA6e2b")
		instance, err := deposit.NewDeposit(contractAddress, client)
		if err != nil {
			fmt.Println(err)
			return users, err
		}
		var (
			addresses []common.Address
		)

		addresses, err = instance.GetUsers(&bind.CallOpts{})
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

func pullUserDepositForUsdtAmount(address string) (string, error) {
	var (
		balance string
	)

	url1 := "https://bsc-dataseed4.binance.org/"

	for i := 0; i < 5; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			url1 = "https://bsc-dataseed1.bnbchain.org"
			continue
			//return usdtAmount, err
		}

		contractAddress := common.HexToAddress("0x8fBdCc0cDD99CE35c1935D8a69286e32b0FA6e2b")
		instance, err := deposit.NewDeposit(contractAddress, client)
		if err != nil {
			fmt.Println(err)
			return balance, err
		}

		var (
			tmp *big.Int
		)
		tmp, err = instance.UserdepositForUsdtAmount(&bind.CallOpts{
			From: common.HexToAddress(""),
		}, common.HexToAddress(address))
		if err != nil {
			return balance, err
		}

		if "0" != tmp.String() {
			balance = tmp.String()
		}

		break
	}

	return balance, nil
}

// PullUserCredentialsBsc 拉取用户api_key api_secret
func (b *BinanceUserService) PullUserCredentialsBsc(ctx context.Context, req *v1.PullUserCredentialsBscRequest) (*v1.PullUserCredentialsBscReply, error) {
	var (
		users []*biz.User
		err   error
	)

	users, err = b.buc.GetUsers()
	if nil != err {
		fmt.Println(err)
		return nil, err
	}

	for _, v := range users {
		var (
			apiKey    string
			apiSecret string
		)

		apiKey, apiSecret, err = pullUserCredentialsBscBySystemRole(v.Address)
		if nil != err {
			fmt.Println(err)
			return nil, err
		}

		err = b.buc.UpdateUser(ctx, v, apiKey, apiSecret)
		if nil != err {
			fmt.Println(err)
			return nil, err
		}
	}

	return nil, nil
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
