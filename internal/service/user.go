package service

import (
	abi "binanceexchange_user/abi"
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

// PullUserDeposit 模式1 用户充值总入口, 单一协程处理，遍历所有用户
func (b *BinanceUserService) PullUserDeposit(ctx context.Context, req *v1.PullUserDepositRequest) (*v1.PullUserDepositReply, error) {
	var (
		users []string
		err   error
	)

	users, err = pullUsersByDeposit()
	for _, v := range users {
		var (
			balance string
			cost    uint64
			open    bool
		)
		balance, cost, open, err = pullUserDepositForInfo(v)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = b.buc.SetUserBalanceAndUser(ctx, v, balance, cost, open, 1)
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

		contractAddress := common.HexToAddress("0x57f76DFe8f197b5bAF205170e9b8260Ab4e3Cbc2")
		instance, err := abi.NewDeposit(contractAddress, client)
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

func pullUserDepositForInfo(address string) (string, uint64, bool, error) {
	var (
		balance string
		cost    uint64
		open    bool
	)

	url1 := "https://bsc-dataseed4.binance.org/"

	for i := 0; i < 5; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			url1 = "https://bsc-dataseed1.bnbchain.org"
			continue
			//return usdtAmount, err
		}

		contractAddress := common.HexToAddress("0x57f76DFe8f197b5bAF205170e9b8260Ab4e3Cbc2")
		instance, err := abi.NewDeposit(contractAddress, client)
		if err != nil {
			fmt.Println(err)
			return balance, cost, open, err
		}

		var (
			tmp  *big.Int
			tmp2 *big.Int
			tmp3 bool
		)
		tmp, err = instance.UserdepositForUsdtAmount(&bind.CallOpts{}, common.HexToAddress(address))
		if err != nil {
			return balance, cost, open, err
		}

		tmp2, err = instance.UserCost(&bind.CallOpts{}, common.HexToAddress(address))
		if err != nil {
			return balance, cost, open, err
		}

		contractAddress2 := common.HexToAddress("0xCD329BEb4574Ddf10411F71383FA0d313A65A737")
		instance2, err := abi.NewDepositOpen(contractAddress2, client)
		if err != nil {
			fmt.Println(err)
			return balance, cost, open, err
		}

		tmp3, err = instance2.UserOpen(&bind.CallOpts{}, common.HexToAddress(address))
		if err != nil {
			return balance, cost, open, err
		}

		if "0" != tmp.String() {
			balance = tmp.String()
		}

		if 0 < tmp2.Uint64() {
			cost = tmp2.Uint64()
		}

		open = tmp3

		break
	}

	return balance, cost, open, nil
}

// PullUserDeposit2 .
func (b *BinanceUserService) PullUserDeposit2(ctx context.Context, req *v1.PullUserDepositRequest) (*v1.PullUserDepositReply, error) {
	var (
		users []string
		err   error
	)

	users, err = pullUsersByStakeTfi()
	for _, v := range users {
		var (
			balance string
			cost    uint64
			open    bool
		)
		balance, cost, open, err = pullUserStakeTfiForInfo(v)
		if err != nil {
			fmt.Println(err)
			continue
		}

		err = b.buc.SetUserBalanceAndUser(ctx, v, balance, cost, open, 2)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}

	return nil, nil
}

func pullUsersByStakeTfi() ([]string, error) {
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

		contractAddress := common.HexToAddress("0xEB1D3e6995E6B65618790F53ad8aC482b8Ba283d")
		instance, err := abi.NewStakeTfi(contractAddress, client)
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

func pullUserStakeTfiForInfo(address string) (string, uint64, bool, error) {
	var (
		balance string
		cost    uint64
		open    bool
		err     error
	)

	url1 := "https://bsc-dataseed4.binance.org/"

	for i := 0; i < 5; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			url1 = "https://bsc-dataseed1.bnbchain.org"
			continue
			//return usdtAmount, err
		}

		contractAddress := common.HexToAddress("0xEB1D3e6995E6B65618790F53ad8aC482b8Ba283d")
		instance, err := abi.NewStakeTfi(contractAddress, client)
		if err != nil {
			fmt.Println(err)
			return balance, cost, open, err
		}

		var (
			tmp  *big.Int
			tmp2 *big.Int
			tmp3 bool
		)
		tmp3, err = instance.UserOpen(&bind.CallOpts{}, common.HexToAddress(address))
		if err != nil {
			return balance, cost, open, err
		}

		tmp, err = instance.UserTfiAmount(&bind.CallOpts{}, common.HexToAddress(address))
		if err != nil {
			return balance, cost, open, err
		}

		tmp2, err = instance.UserCost(&bind.CallOpts{}, common.HexToAddress(address))
		if err != nil {
			return balance, cost, open, err
		}

		if "0" != tmp.String() {
			balance = tmp.String()
		}

		if 0 < tmp2.Uint64() {
			cost = tmp2.Uint64()
		}

		open = tmp3
		break
	}

	return balance, cost, open, err
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
		instance, err := abi.NewUserCredentialsBsc(tokenAddress, client)
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

func (b *BinanceUserService) BindTrader(ctx context.Context, req *v1.BindTraderRequest) (*v1.BindTraderReply, error) {
	return b.buc.BindTrader(ctx)
}

func (b *BinanceUserService) ListenTraderAndUserOrder(ctx context.Context, req *v1.ListenTraderAndUserOrderRequest) (*v1.ListenTraderAndUserOrderReply, error) {
	return b.buc.ListenTraders(ctx, req)
}

func (b *BinanceUserService) OrderHandle(ctx context.Context, req *v1.OrderHandleRequest) (*v1.OrderHandleReply, error) {
	return b.buc.OrderHandle(ctx, req)
}

func (b *BinanceUserService) Analyze(ctx context.Context, req *v1.AnalyzeRequest) (*v1.AnalyzeReply, error) {
	return b.buc.Analyze(ctx, req)
}
