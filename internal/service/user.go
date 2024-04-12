package service

import (
	abi "binanceexchange_user/abi"
	v1 "binanceexchange_user/api/binanceexchange_user/v1"
	"binanceexchange_user/internal/biz"
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strconv"
	"strings"
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

		//contractAddress2 := common.HexToAddress("0xCD329BEb4574Ddf10411F71383FA0d313A65A737")
		//instance2, err := abi.NewDepositOpen(contractAddress2, client)
		//if err != nil {
		//	fmt.Println(err)
		//	return balance, cost, open, err
		//}

		//tmp3, err = instance2.UserOpen(&bind.CallOpts{}, common.HexToAddress(address))
		//if err != nil {
		//	return balance, cost, open, err
		//}

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
		//tmp3, err = instance.UserOpen(&bind.CallOpts{}, common.HexToAddress(address))
		//if err != nil {
		//	return balance, cost, open, err
		//}

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
			continue
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

func (b *BinanceUserService) OrderHandleTwo(ctx context.Context, req *v1.OrderHandleRequest) (*v1.OrderHandleReply, error) {
	return b.buc.OrderHandleTwo(ctx, req)
}

func (b *BinanceUserService) Analyze(ctx context.Context, req *v1.AnalyzeRequest) (*v1.AnalyzeReply, error) {
	return b.buc.Analyze(ctx, req)
}

func (b *BinanceUserService) CloseOrderAfterBind(ctx context.Context, req *v1.CloseOrderAfterBindRequest) (*v1.CloseOrderAfterBindReply, error) {
	return b.buc.CloseOrderAfterBind(ctx, req)
}

func (b *BinanceUserService) CloseOrderAfterBindTwo(ctx context.Context, req *v1.CloseOrderAfterBindRequest) (*v1.CloseOrderAfterBindReply, error) {
	return b.buc.CloseOrderAfterBindTwo(ctx, req)
}

func (b *BinanceUserService) InitOrderAfterBind(ctx context.Context, req *v1.InitOrderAfterBindRequest) (*v1.InitOrderAfterBindReply, error) {
	return b.buc.InitOrderAfterBind(ctx, req)
}

func (b *BinanceUserService) InitOrderAfterBindTwo(ctx context.Context, req *v1.InitOrderAfterBindRequest) (*v1.InitOrderAfterBindReply, error) {
	return b.buc.InitOrderAfterBindTwo(ctx, req)
}

func (b *BinanceUserService) OverOrderAfterBind(ctx context.Context, req *v1.OverOrderAfterBindRequest) (*v1.OverOrderAfterBindReply, error) {
	return b.buc.OverOrder(ctx, req)
}

func (b *BinanceUserService) OverOrderAfterBindTwo(ctx context.Context, req *v1.OverOrderAfterBindRequest) (*v1.OverOrderAfterBindReply, error) {
	return b.buc.OverOrderTwo(ctx, req)
}

func (b *BinanceUserService) AdminOverOrderAfterBind(ctx context.Context, req *v1.OverOrderAfterBindRequest) (*v1.OverOrderAfterBindReply, error) {
	return b.buc.AdminOverOrder(ctx, req)
}

func (b *BinanceUserService) AdminOverOrderAfterBindTwo(ctx context.Context, req *v1.OverOrderAfterBindRequest) (*v1.OverOrderAfterBindReply, error) {
	return b.buc.AdminOverOrderTwo(ctx, req)
}

// PullTradingBoxOpen .
func (b *BinanceUserService) PullTradingBoxOpen(ctx context.Context, req *v1.PullTradingBoxOpenRequest) (*v1.PullTradingBoxOpenReply, error) {
	var (
		terms          []uint64
		boxTotalAmount map[uint64]uint64
		balance        map[uint64]float64
		tokenIds       []uint64
		boxAmount      map[uint64]uint64
		tradingBoxOpen map[uint64]*biz.TradingBoxOpen
		incomeTokenIds []uint64
		incomeAmounts  []string
		err            error
	)

	terms, boxTotalAmount, err = pullBoxTerm()
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	if 0 >= len(terms) {
		return nil, nil
	}

	// todo 每期一个账户
	balance = make(map[uint64]float64, 0)
	for k, vTerms := range terms {
		if 0 == k {
			var tmpBalance float64
			tmpBalance, err = b.buc.GetTermBinanceCurrentBalance(ctx, 17)
			if err != nil {
				fmt.Println(err)
				return nil, nil
			}

			if 0 < tmpBalance {
				balance[vTerms] = tmpBalance
			}
		}
	}

	tokenIds, boxAmount, err = pullTradingBoxOpen()
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	tradingBoxOpen, err = b.buc.GetTradingBoxOpenMap()
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	incomeTokenIds = make([]uint64, 0)
	incomeAmounts = make([]string, 0)
	for _, vTokenIds := range tokenIds {
		if _, ok := boxAmount[vTokenIds]; !ok {
			continue
		}

		if _, ok := tradingBoxOpen[vTokenIds]; ok { // 存在数据
			continue
		}

		tmpLastTerm := uint64(0)
		for _, vTerms := range terms {
			if _, ok := boxTotalAmount[vTerms]; !ok {
				continue
			}

			if _, ok := balance[vTerms]; !ok {
				continue
			}

			if tmpLastTerm < vTokenIds && vTokenIds <= vTerms { // 本期区间
				// 写入数据
				tmpRate := float64(boxAmount[vTokenIds]) / float64(boxTotalAmount[vTerms])
				tmpIncome := balance[vTerms] * tmpRate
				err = b.buc.InsertTradingBoxOpen(ctx, vTokenIds, boxAmount[vTokenIds], boxTotalAmount[vTerms], tmpRate, balance[vTerms], tmpIncome)
				if nil != err {
					fmt.Println(err)
				}

				incomeTokenIds = append(incomeTokenIds, vTokenIds)
				incomeAmounts = append(incomeAmounts, float64ToString(tmpIncome))
				break
			}

			tmpLastTerm = vTerms
		}

	}

	if 0 < len(incomeAmounts) {
		err = setIncomeTradingBox(incomeTokenIds, incomeAmounts)
		if nil != err {
			fmt.Println("记录开盒数据请求区块链失败", err, incomeTokenIds[0], incomeTokenIds[len(incomeTokenIds)-1])
			return nil, err
		}
	}

	return nil, nil
}

func float64ToString(input float64) string {
	// 将浮点数转换为字符串
	str := strconv.FormatFloat(input, 'f', -1, 64)
	// 找到小数点的位置
	dotIndex := strings.Index(str, ".")
	if dotIndex == -1 {
		// 如果没有小数点，则在字符串末尾添加 18 个零
		return str + strings.Repeat("0", 18)
	}
	// 如果有小数点，则截取小数点前面的部分和后面的部分
	intPart := str[:dotIndex]
	decimalPart := str[dotIndex+1:]
	// 如果小数部分的长度不足 18 位，则在末尾补齐零
	if len(decimalPart) < 18 {
		decimalPart += strings.Repeat("0", 18-len(decimalPart))
	}
	// 将整数部分和小数部分拼接起来并返回
	return intPart + decimalPart
}

func pullTradingBoxOpen() ([]uint64, map[uint64]uint64, error) {
	var (
		openBox   []uint64
		boxAmount map[uint64]uint64
		err       error
	)

	openBox = make([]uint64, 0)
	boxAmount = make(map[uint64]uint64, 0)

	url1 := "https://bsc-dataseed4.binance.org/"

	for i := 0; i < 5; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			url1 = "https://bsc-dataseed1.bnbchain.org"
			continue
			//return usdtAmount, err
		}

		contractAddress := common.HexToAddress("0x7a0E6153d0FEB55B83c977e064A64C4b4A8a76b0")
		instance, err := abi.NewTradingBox(contractAddress, client)
		if err != nil {
			fmt.Println(err)
			return openBox, boxAmount, err
		}

		var (
			tmp *big.Int
		)

		tmp, err = instance.TotalSupply(&bind.CallOpts{})
		if err != nil {
			return openBox, boxAmount, err
		}

		if 0 >= tmp.Uint64() {
			return openBox, boxAmount, err
		}

		for tokenId := int64(1); tokenId <= tmp.Int64(); tokenId++ {
			var (
				tmp2 *big.Int
				tmp3 *big.Int
			)

			tmp3, err = instance.BoxTokenAmount(&bind.CallOpts{}, big.NewInt(tokenId))
			if err != nil {
				continue
			}

			lengthToKeep := len(tmp3.String()) - 15 // todo 每个盒子业务默认100u

			if lengthToKeep > 0 {
				amountTmpStr := tmp3.String()[:lengthToKeep]
				var (
					tmp4 int64
				)
				tmp4, err = strconv.ParseInt(amountTmpStr, 10, 64)
				if nil != err || 0 >= tmp4 {
					continue
				}

				boxAmount[uint64(tokenId)] = uint64(tmp4)
			} else {
				continue
			}

			tmp2, err = instance.OpenTimeStamp(&bind.CallOpts{}, big.NewInt(tokenId))
			if err != nil {
				continue
			}

			if 0 < tmp2.Uint64() {
				openBox = append(openBox, uint64(tokenId))
			}
		}

		break
	}

	return openBox, boxAmount, err
}

func pullBoxTerm() ([]uint64, map[uint64]uint64, error) {
	var (
		terms      []uint64
		termAmount map[uint64]uint64
		err        error
	)

	url1 := "https://bsc-dataseed4.binance.org/"

	for i := 0; i < 5; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			url1 = "https://bsc-dataseed1.bnbchain.org"
			continue
			//return usdtAmount, err
		}

		contractAddress := common.HexToAddress("0xDbcbFc8c250Fad5CAe8b6f08d5273AA366f218c0")
		instance, err := abi.NewTradingAdmin(contractAddress, client)
		if err != nil {
			fmt.Println(err)
			return terms, termAmount, err
		}

		var (
			tmp []*big.Int
		)

		tmp, err = instance.GetTradingTerm(&bind.CallOpts{})
		if err != nil {
			return terms, termAmount, err
		}

		if 0 >= len(tmp) {
			return terms, termAmount, err
		}

		terms = make([]uint64, 0)
		termAmount = make(map[uint64]uint64, 0)
		for _, vTmp := range tmp {
			var (
				tmp3 *big.Int
			)

			tmp3, err = instance.GetTermAmount(&bind.CallOpts{}, vTmp)
			if err != nil {
				continue
			}

			lengthToKeep := len(tmp3.String()) - 15 // todo 每个盒子业务默认100u

			if lengthToKeep > 0 {
				amountTmpStr := tmp3.String()[:lengthToKeep]
				var (
					tmp4 int64
				)
				tmp4, err = strconv.ParseInt(amountTmpStr, 10, 64)
				if nil != err || 0 >= tmp4 {
					continue
				}

				termAmount[vTmp.Uint64()] = uint64(tmp4)
			} else {
				continue
			}

			terms = append(terms, vTmp.Uint64())
		}

		break
	}

	return terms, termAmount, err
}

// SettleTradingBoxOpen .
func (b *BinanceUserService) SettleTradingBoxOpen(ctx context.Context, req *v1.SettleTradingBoxOpenRequest) (*v1.SettleTradingBoxOpenReply, error) {
	//var (
	//	terms          []uint64
	//	tradingBoxOpen map[uint64]*biz.TradingBoxOpen
	//	err            error
	//)
	//
	//terms, _, err = pullBoxTerm()
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, nil
	//}
	//
	//tradingBoxOpen, err = b.buc.GetTradingBoxOpenMapByStatus(1)
	//if err != nil {
	//	fmt.Println(err)
	//	return nil, nil
	//}
	//
	//for tokenId, vTradingBoxOpen := range tradingBoxOpen {
	//	// tokenId从1开始，和 totalSupply是相等的，因为不能燃烧，燃烧需要update_role权限也不会放开
	//	tmpLastTerm := uint64(0)
	//	for _, vTerms := range terms {
	//		if tmpLastTerm >= tokenId || tokenId > vTerms { // 非本期区间
	//			continue
	//		}
	//
	//		tmpLastTerm = vTerms
	//	}
	//}

	return nil, nil
}

func setIncomeTradingBox(tokenIds []uint64, amounts []string) error {
	url1 := "https://bsc-dataseed4.binance.org/"

	for i := 0; i < 5; i++ {
		client, err := ethclient.Dial(url1)
		if err != nil {
			url1 = "https://bsc-dataseed1.bnbchain.org"
			continue
			//return usdtAmount, err
		}

		contractAddress := common.HexToAddress("0x7a0E6153d0FEB55B83c977e064A64C4b4A8a76b0")
		instance, err := abi.NewTradingBox(contractAddress, client)
		if err != nil {
			fmt.Println(err)
			return err
		}

		var authUser *bind.TransactOpts

		var privateKey *ecdsa.PrivateKey
		privateKey, err = crypto.HexToECDSA("")
		if err != nil {
			fmt.Println(err)
			return err
		}

		authUser, err = bind.NewKeyedTransactorWithChainID(privateKey, new(big.Int).SetInt64(56))
		if err != nil {
			fmt.Println(err)
			return err
		}

		var (
			tmpTokenIds       []*big.Int
			tmpWithdrawAmount []*big.Int
		)

		for _, vTmpTokenIds := range tokenIds {
			tmpTokenIds = append(tmpTokenIds, new(big.Int).SetUint64(vTmpTokenIds))
		}

		for _, vAmounts := range amounts {
			tmp1, _ := new(big.Int).SetString(vAmounts, 10)
			tmpWithdrawAmount = append(tmpWithdrawAmount, tmp1)
		}

		_, err = instance.SetIncomeAny(&bind.TransactOpts{
			From:     authUser.From,
			Signer:   authUser.Signer,
			GasLimit: 0,
		}, tmpTokenIds, tmpWithdrawAmount)
		if err != nil {
			return err
		}

		break
	}

	return nil
}
