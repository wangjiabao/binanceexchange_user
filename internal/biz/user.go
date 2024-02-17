package biz

import (
	v1 "binanceexchange_user/api/binanceexchange_user/v1"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

type User struct {
	ID               uint64
	Address          string
	PlayType         uint64
	ApiStatus        uint64
	ApiKey           string
	ApiSecret        string
	BindTraderStatus uint64
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type UserBalance struct {
	ID        uint64
	UserId    uint64
	Balance   string
	Cost      uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserBalanceRecord struct {
	ID        uint64
	UserId    uint64
	Amount    string
	Balance   string
	Tx        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserAmount struct {
	ID        uint64
	UserId    uint64
	Amount    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserAmountRecord struct {
	ID        uint64
	UserId    uint64
	OrderId   uint64
	Amount    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Trader struct {
	ID        uint64
	Status    uint64
	Amount    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserBindTrader struct {
	ID        uint64
	UserId    uint64
	TraderId  uint64
	Amount    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserOrder struct {
	ID            uint64
	UserId        uint64
	TraderId      uint64
	ClientOrderId string
	OrderId       string
	Symbol        string
	Side          string
	PositionSide  string
	Quantity      float64
	Price         float64
	TraderQty     float64
	OrderType     string
	ClosePosition string
	CumQuote      float64
	ExecutedQty   float64
	AvgPrice      float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Symbol struct {
	ID                uint64
	Symbol            string
	QuantityPrecision int64
}

type BinanceUserRepo interface {
	InsertUser(ctx context.Context, User *User) (*User, error)
	UpdateUser(ctx context.Context, userId uint64, apiKey string, apiSecret string) (bool, error)
	UpdateUserBindTraderStatus(ctx context.Context, userId uint64) (bool, error)
	InsertUserBalance(ctx context.Context, userBalance *UserBalance) (bool, error)
	UpdatesUserBalance(ctx context.Context, userId uint64, balance string, cost uint64) (bool, error)
	InsertUserBalanceRecord(ctx context.Context, userBalance *UserBalanceRecord) (bool, error)
	InsertUserAmount(ctx context.Context, userAmount *UserAmount) (bool, error)
	UpdatesUserAmount(ctx context.Context, userId uint64, amount int64) (bool, error)
	InsertUserAmountRecord(ctx context.Context, userAmount *UserAmountRecord) (bool, error)
	InsertUserBindTrader(ctx context.Context, userId uint64, traderId uint64, amount uint64) (*UserBindTrader, error)
	InsertUserOrder(ctx context.Context, order *UserOrder) (*UserOrder, error)
	GetUsers() ([]*User, error)
	GetUsersByUserIds(userIds []uint64) (map[uint64]*User, error)
	GetUsersByBindUserStatus() ([]*User, error)
	GetUserByAddress(ctx context.Context, address string) (*User, error)
	GetUserByApiKeyAndApiSecret(key string, secret string) (*User, error)
	GetUserBalance(ctx context.Context, userId uint64) (*UserBalance, error)
	GetUserAmount(ctx context.Context, userId uint64) (*UserAmount, error)
	GetUserBalanceByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*UserBalance, error)
	GetUserAmountByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*UserAmount, error)
	GetTradersOrderByAmountDesc() ([]*Trader, error)
	GetTraders() ([]*Trader, error)
	GetUserBindTraderByUserId(userId uint64) ([]*UserBindTrader, error)
	GetUserBindTraderByTraderIds(traderIds []uint64) (map[uint64][]*UserBindTrader, error)
	GetSymbol() (map[string]*Symbol, error)
	GetUserOrderByUserTraderIdAndSymbol(userId uint64, traderId uint64, symbol string) ([]*UserOrder, error)
	GetUserOrderByUserIdAndSymbolAndPositionSide(userId uint64, symbol string, positionSide string) ([]*UserOrder, error)
}

// BinanceUserUsecase is a BinanceData usecase.
type BinanceUserUsecase struct {
	binanceUserRepo BinanceUserRepo
	tx              Transaction
	log             *log.Helper
}

// NewBinanceDataUsecase new a BinanceData usecase.
func NewBinanceDataUsecase(binanceUserRepo BinanceUserRepo, tx Transaction, logger log.Logger) *BinanceUserUsecase {
	return &BinanceUserUsecase{binanceUserRepo: binanceUserRepo, tx: tx, log: log.NewHelper(logger)}
}

func (b *BinanceUserUsecase) SetUserBalanceAndUser(ctx context.Context, address string, balance string, cost uint64, playType uint64) error {
	var (
		user            *User
		userBalance     *UserBalance
		err             error
		lastUserBalance string // 上次用户余额
	)

	user, err = b.binanceUserRepo.GetUserByAddress(ctx, address)
	if nil != err {
		return err
	}

	// 初始化 用户和余额
	if nil == user {
		if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
			user, err = b.binanceUserRepo.InsertUser(ctx, &User{
				Address:  address,
				PlayType: playType, // 初始化玩法，模式
			})
			if nil != err {
				return err
			}

			if nil == user {
				return errors.New(500, "CREATE_USER_ERROR", "未发现创建的用户")
			}

			_, err = b.binanceUserRepo.InsertUserBalance(ctx, &UserBalance{
				UserId:  user.ID,
				Balance: balance,
				Cost:    cost,
			})
			if nil != err {
				return err
			}

			_, err = b.binanceUserRepo.InsertUserAmount(ctx, &UserAmount{
				UserId: user.ID,
			})
			if nil != err {
				return err
			}

			_, err = b.binanceUserRepo.InsertUserBalanceRecord(ctx, &UserBalanceRecord{
				UserId:  user.ID,
				Amount:  balance,
				Balance: lastUserBalance,
			})
			if nil != err {
				return err
			}

			return nil
		}); err != nil {
			b.log.Error(err)
			return err
		}

	} else {
		// 模式对应才给充值，初始化用户时已经决定了，条件很关键
		if playType != user.PlayType {
			return nil
		}

		userBalance, err = b.binanceUserRepo.GetUserBalance(ctx, user.ID)
		if nil != err {
			return err
		}

		// 无变化
		if userBalance.Balance == balance && userBalance.Cost == cost {
			return nil
		}

		// 有变化
		tmpAmount := new(big.Int)
		if userBalance.Balance != balance {
			// 上次余额
			tmpLastUserBalance := new(big.Int)
			tmpLastUserBalance.SetString(userBalance.Balance, 10)
			// 本次余额
			tmpCurrentUserBalance := new(big.Int)
			tmpCurrentUserBalance.SetString(balance, 10)
			// 增长
			tmpAmount.Sub(tmpCurrentUserBalance, tmpLastUserBalance)
		}

		tmpCost := userBalance.Cost
		if tmpCost != cost {
			tmpCost = cost
		}

		// 更新余额
		if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
			_, err = b.binanceUserRepo.UpdatesUserBalance(ctx, user.ID, balance, tmpCost)
			if nil != err {
				return err
			}

			if userBalance.Balance != balance {
				_, err = b.binanceUserRepo.InsertUserBalanceRecord(ctx, &UserBalanceRecord{
					UserId:  user.ID,
					Amount:  tmpAmount.String(),
					Balance: userBalance.Balance,
				})
				if nil != err {
					return err
				}
			}

			return nil
		}); err != nil {
			b.log.Error(err)
			return err
		}
	}

	return nil
}

func (b *BinanceUserUsecase) UpdateUser(ctx context.Context, user *User, apiKey string, apiSecret string) error {
	var (
		user2  *User
		symbol map[string]*Symbol
		err    error
	)

	user2, err = b.binanceUserRepo.GetUserByApiKeyAndApiSecret(apiKey, apiSecret)
	if nil != err {
		return err
	}

	if nil == user2 || (user2.ID == user.ID) { // api不存在 或 地址和api指向相同ID(同一条记录)
		if apiKey != user.ApiKey || apiSecret != user.ApiSecret { // api_key或api_secret发生了变化
			if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
				_, err = b.binanceUserRepo.UpdateUser(ctx, user.ID, apiKey, apiSecret)
				if nil != err {
					return err
				}

				return nil
			}); err != nil {
				b.log.Error(err)
				return err
			}

			symbol, err = b.binanceUserRepo.GetSymbol()
			if nil != err {
				return err
			}

			for k, _ := range symbol {
				_, err = requestBinanceLeverAge(k, int64(20), apiKey, apiSecret)
				if nil != err {
					continue
				}
			}
		}
	}

	return nil
}

func (b *BinanceUserUsecase) GetUsers() ([]*User, error) {
	return b.binanceUserRepo.GetUsers()
}

func (b *BinanceUserUsecase) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	var (
		user        *User
		userBalance *UserBalance
		userAmount  *UserAmount
		err         error
	)

	if 0 >= len(req.Address) || 300 < len(req.Address) {
		return &v1.GetUserReply{}, nil
	}

	user, err = b.binanceUserRepo.GetUserByAddress(ctx, req.Address)
	if nil != err {
		return &v1.GetUserReply{}, nil
	}

	if nil == user {
		return &v1.GetUserReply{}, nil
	}

	userBalance, err = b.binanceUserRepo.GetUserBalance(ctx, user.ID)
	if nil != err {
		return &v1.GetUserReply{}, nil
	}

	userAmount, err = b.binanceUserRepo.GetUserAmount(ctx, user.ID)
	if nil != err {
		return &v1.GetUserReply{}, nil
	}
	return &v1.GetUserReply{Status: int64(user.ApiStatus), Play: int64(user.PlayType), Balance: userBalance.Balance, Amount: userAmount.Amount}, err
}

func (b *BinanceUserUsecase) BindTrader(ctx context.Context) (*v1.BindTraderReply, error) {
	var (
		users       []*User
		userBalance map[uint64]*UserBalance
		traders     []*Trader
		err         error
	)

	// 获取未绑定交易员的用户
	users, err = b.binanceUserRepo.GetUsersByBindUserStatus()
	if nil != err {
		return nil, err
	}
	if 0 == len(users) {
		return nil, nil
	}

	// 获取用户cost
	userIds := make([]uint64, 0)
	for _, vUsers := range users {
		userIds = append(userIds, vUsers.ID)
	}
	userBalance, err = b.binanceUserRepo.GetUserBalanceByUserIds(ctx, userIds)
	if nil != err {
		return nil, err
	}

	// 按amount排序好的交易员
	traders, err = b.binanceUserRepo.GetTradersOrderByAmountDesc()
	if nil != err {
		return nil, err
	}
	if 0 == len(traders) {
		return nil, nil
	}

	// 1. 算法遍历，轮询用户
	// 2. 轮询交易员（查询时已经按amount desc排序）
	// 3. 用户cost的百分之30>=当前交易员的amount字段，建立绑定关系。
	// 4. 轮询结束时充足者不予理会。
	for _, vUsers := range users {
		if _, ok := userBalance[vUsers.ID]; !ok {
			continue
		}

		// 初始化
		tmpCost := userBalance[vUsers.ID].Cost
		var limitAmount uint64
		bindTrader := make(map[uint64]*Trader, 0)

		// 第一轮
		for _, vTraders := range traders {
			if 0 >= tmpCost {
				break
			}

			if 0 >= vTraders.Amount {
				continue
			}

			if tmpCost*30/100 >= vTraders.Amount {
				// 绑定
				if _, ok := bindTrader[vTraders.ID]; ok {
					continue
				}

				bindTrader[vTraders.ID] = vTraders
				tmpCost -= vTraders.Amount

				limitAmount = vTraders.Amount // 最后的限制金额
			}
		}

		// 第二轮，跳过分配限制的额度，剩下的按顺序分配
		for _, vTraders := range traders {
			if 0 >= tmpCost {
				break
			}

			if 0 < limitAmount && vTraders.Amount > limitAmount {
				continue
			}

			if 0 >= vTraders.Amount {
				continue
			}

			// 绑定
			if _, ok := bindTrader[vTraders.ID]; ok {
				continue
			}

			bindTrader[vTraders.ID] = vTraders
			tmpCost -= vTraders.Amount
		}

		if 0 >= len(bindTrader) {
			continue
		}
		// 写入
		if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
			for k, v := range bindTrader {
				_, err = b.binanceUserRepo.InsertUserBindTrader(ctx, vUsers.ID, k, v.Amount)
				if nil != err {
					return err
				}
			}

			_, err = b.binanceUserRepo.UpdateUserBindTraderStatus(ctx, vUsers.ID)
			if nil != err {
				return err
			}

			return nil
		}); err != nil {
			b.log.Error(err)
		}
	}

	return nil, nil
}

func (b *BinanceUserUsecase) TestLeverAge(ctx context.Context, req *v1.TestLeverAgeRequest) (*v1.TestLeverAgeReply, error) {
	var (
		users  map[uint64]*User
		symbol map[string]*Symbol
		err    error
	)
	userIds := make([]uint64, 0)
	userIds = append(userIds, 3, 4)
	users, err = b.binanceUserRepo.GetUsersByUserIds(userIds)

	for _, v := range users {
		symbol, err = b.binanceUserRepo.GetSymbol()
		if nil != err {
			continue
		}

		for k, _ := range symbol {
			_, err = requestBinanceLeverAge(k, int64(20), v.ApiKey, v.ApiSecret)
			if nil != err {
				continue
			}
		}
	}

	return nil, nil
}

func (b *BinanceUserUsecase) TestOrder(ctx context.Context, req *v1.TestOrderRequest) (*v1.TestOrderReply, error) {
	var (
		users map[uint64]*User
		order *BinanceOrder
		err   error
	)
	userIds := make([]uint64, 0)
	userIds = append(userIds, 3, 4)
	users, err = b.binanceUserRepo.GetUsersByUserIds(userIds)

	for _, v := range users {
		if v.ID == 4 {
			order, err = requestBinanceOrderInfo("LINKUSDT", 30667218504, v.ApiKey, v.ApiSecret)
			if nil != err {
				fmt.Println(err, v)
				continue
			}
			fmt.Println(order)
			order, err = requestBinanceOrderInfo("LINKUSDT", 30668742220, v.ApiKey, v.ApiSecret)
			if nil != err {
				fmt.Println(err, v)
				continue
			}
			fmt.Println(order)
		}

		if v.ID == 3 {
			order, err = requestBinanceOrderInfo("LINKUSDT", 30667218505, v.ApiKey, v.ApiSecret)
			if nil != err {
				fmt.Println(err, v)
				continue
			}
			fmt.Println(order)
			order, err = requestBinanceOrderInfo("LINKUSDT", 30668742221, v.ApiKey, v.ApiSecret)
			if nil != err {
				fmt.Println(err, v)
				continue
			}
			fmt.Println(order)
		}
	}

	return nil, nil
}

type OrderData struct {
	Coin  string
	Type  string
	Price string
	Side  string
	Qty   string
}

func (b *BinanceUserUsecase) ListenTraders(ctx context.Context, req *v1.ListenTraderAndUserOrderRequest) (*v1.ListenTraderAndUserOrderReply, error) {
	var (
		wg             sync.WaitGroup
		traderIds      []uint64
		userBindTrader map[uint64][]*UserBindTrader
		userIds        []uint64
		users          map[uint64]*User
		userBalance    map[uint64]*UserBalance
		userAmount     map[uint64]*UserAmount
		symbol         map[string]*Symbol
		err            error
	)

	traderIds = make([]uint64, 0)
	for _, vOrders := range req.SendBody.Orders {
		traderIds = append(traderIds, vOrders.Uid)
	}

	userBindTrader, err = b.binanceUserRepo.GetUserBindTraderByTraderIds(traderIds)
	if nil != err {
		return &v1.ListenTraderAndUserOrderReply{
			Status: "err",
		}, err
	}

	userIds = make([]uint64, 0)
	for _, vUserBindTrader := range userBindTrader {
		for _, vVUserBindTrader := range vUserBindTrader {
			userIds = append(userIds, vVUserBindTrader.UserId)
		}
	}

	if 0 >= len(userIds) {
		return &v1.ListenTraderAndUserOrderReply{
			Status: "ok",
		}, nil
	}

	// 获取用户信息，余额信息，收益信息
	users, err = b.binanceUserRepo.GetUsersByUserIds(userIds)
	if nil != err {
		return &v1.ListenTraderAndUserOrderReply{
			Status: "err",
		}, err
	}

	userBalance, err = b.binanceUserRepo.GetUserBalanceByUserIds(ctx, userIds)
	if nil != err {
		return &v1.ListenTraderAndUserOrderReply{
			Status: "err",
		}, err
	}

	userAmount, err = b.binanceUserRepo.GetUserAmountByUserIds(ctx, userIds)
	if nil != err {
		return &v1.ListenTraderAndUserOrderReply{
			Status: "err",
		}, err
	}

	symbol, err = b.binanceUserRepo.GetSymbol()
	if nil != err {
		return &v1.ListenTraderAndUserOrderReply{
			Status: "err",
		}, err
	}

	for _, vOrders := range req.SendBody.Orders {
		if _, ok := userBindTrader[vOrders.Uid]; !ok {
			continue
		}

		for _, vOrdersData := range vOrders.Data {
			for _, vUserBindTrader := range userBindTrader[vOrders.Uid] {
				if _, ok := users[vUserBindTrader.UserId]; !ok {
					continue
				}

				if _, ok := userBalance[vUserBindTrader.UserId]; !ok {
					continue
				}

				if _, ok := userAmount[vUserBindTrader.UserId]; !ok {
					continue
				}

				// 判断是开单还是关单，sell long 关多 buy short 关空
				if ("SELL" == vOrdersData.Side && "LONG" == vOrdersData.Type) || ("BUY" == vOrdersData.Side && "SHORT" == vOrdersData.Type) {

				} else if ("SELL" == vOrdersData.Side && "SHORT" == vOrdersData.Type) || ("BUY" == vOrdersData.Side && "LONG" == vOrdersData.Type) {
					// 模式，开单余额判断
					if 1 == users[vUserBindTrader.UserId].PlayType {
						// 精度按代币18位，截取小数点后到5位计算
						var balanceTmp int64
						lengthToKeep := len(userBalance[vUserBindTrader.UserId].Balance) - 13

						if lengthToKeep > 0 {
							balanceTmpStr := userBalance[vUserBindTrader.UserId].Balance[:lengthToKeep]
							balanceTmp, err = strconv.ParseInt(balanceTmpStr, 10, 64)
							if nil != err || 0 >= balanceTmp {
								continue
							}
						} else {
							continue
						}

						// 余额不足，1u的收益，要10000u的余额
						if userAmount[vUserBindTrader.UserId].Amount > balanceTmp*10000 {
							continue
						}
					} else if 2 == users[vUserBindTrader.UserId].PlayType {
						// 余额不足，10u的收益，要1u的余额
						// 精度按代币18位，截取小数点后到5位计算
						var balanceTmp int64
						lengthToKeep := len(userBalance[vUserBindTrader.UserId].Balance) - 13

						if lengthToKeep > 0 {
							balanceTmpStr := userBalance[vUserBindTrader.UserId].Balance[:lengthToKeep]
							balanceTmp, err = strconv.ParseInt(balanceTmpStr, 10, 64)
							if nil != err || 0 >= balanceTmp {
								continue
							}
						} else {
							continue
						}

						// 余额不足，10u的收益，要1u的余额 todo 1tfi按1000算
						if userAmount[vUserBindTrader.UserId].Amount > balanceTmp*1000 {
							continue
						}

					} else {
						continue
					}
				} else {
					continue
				}

				// 精度
				if _, ok := symbol[vOrdersData.Symbol]; !ok {
					continue
				}

				// 发送订单
				wg.Add(1) // 启动一个goroutine就登记+1
				go b.userOrderGoroutine(ctx, &wg, &OrderData{
					Coin:  vOrdersData.Symbol,
					Type:  vOrdersData.Type,
					Price: vOrdersData.Price,
					Side:  vOrdersData.Side,
					Qty:   vOrdersData.Qty,
				}, vOrders.BaseMoney, users[vUserBindTrader.UserId], vUserBindTrader, symbol[vOrdersData.Symbol].QuantityPrecision)
			}

		}

	}

	wg.Wait() // 等待所有登记的goroutine都结束
	return &v1.ListenTraderAndUserOrderReply{
		Status: "ok",
	}, nil
}

// 用户下单
func (b *BinanceUserUsecase) userOrderGoroutine(ctx context.Context, wg *sync.WaitGroup, order *OrderData, amount string, user *User, userBindTrader *UserBindTrader, quantityPrecision int64) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println(order, user, quantityPrecision, userBindTrader, amount)
	var (
		binanceOrder  *BinanceOrder
		side          string
		orderType     = "MARKET"
		positionSide  string
		quantity      string
		qty           float64
		price         float64
		traderAmount  float64
		currentOrders []*UserOrder
		historyOrders []*UserOrder
		currentOrder  *UserOrder
		insertOrder   *UserOrder
		income        int64
		err           error
	)

	// 新订单数据
	currentOrder = &UserOrder{
		UserId:        userBindTrader.UserId,
		TraderId:      userBindTrader.TraderId,
		Symbol:        order.Coin,
		Side:          "",
		PositionSide:  "",
		Quantity:      0,
		Price:         0,
		TraderQty:     0,
		OrderType:     orderType,
		ClosePosition: "",
		CumQuote:      0,
		ExecutedQty:   0,
		AvgPrice:      0,
	}

	qty, err = strconv.ParseFloat(order.Qty, 64)
	if nil != err {
		fmt.Println(err)
		return
	}
	currentOrder.TraderQty = qty

	price, err = strconv.ParseFloat(order.Price, 64)
	if nil != err {
		fmt.Println(err)
		return
	}
	currentOrder.Price = price

	traderAmount, err = strconv.ParseFloat(amount, 64)
	if nil != err {
		fmt.Println(err)
		return
	}

	if "LONG" == order.Type {
		positionSide = "LONG"
	} else if "SHORT" == order.Type {
		positionSide = "SHORT" // 空
	} else {
		fmt.Println("err position side")
		return
	}
	currentOrder.PositionSide = positionSide

	// 订单统计
	currentOrders, err = b.binanceUserRepo.GetUserOrderByUserTraderIdAndSymbol(userBindTrader.UserId, userBindTrader.TraderId, order.Coin)
	if nil != err {
		fmt.Println(err)
		return
	}

	quantityFloat := float64(userBindTrader.Amount) * qty / traderAmount
	var historyQuantityFloat float64
	if ("SELL" == order.Side && "LONG" == order.Type) || ("BUY" == order.Side && "SHORT" == order.Type) {
		side = order.Side
		// 查出用户的BUY单的币的数量，在对应的trader下，超过了BUY不能SELL todo 使用数据库量太大以后
		if 0 >= len(currentOrders) {
			return
		}

		// 多的部分不管，按剩余的数量关 todo 少的部分另一个程序解决
		for _, vCurrentOrders := range currentOrders {
			if ("SELL" == vCurrentOrders.Side && "SHORT" == vCurrentOrders.PositionSide) || ("BUY" == vCurrentOrders.Side && "LONG" == vCurrentOrders.PositionSide) {
				historyQuantityFloat += vCurrentOrders.ExecutedQty
			} else if ("SELL" == vCurrentOrders.Side && "LONG" == vCurrentOrders.PositionSide) || ("BUY" == vCurrentOrders.Side && "SHORT" == vCurrentOrders.PositionSide) {
				historyQuantityFloat -= vCurrentOrders.ExecutedQty
			}
		}

		// 开单历史数量不足了
		if 0 >= historyQuantityFloat {
			fmt.Println("trader的开单数量小于等于关单数量了", userBindTrader.UserId, userBindTrader.TraderId, historyQuantityFloat)
			return
		}

		// 超过了净开单数量
		if quantityFloat > historyQuantityFloat {
			quantityFloat = historyQuantityFloat
		}

	} else if ("SELL" == order.Side && "SHORT" == order.Type) || ("BUY" == order.Side && "LONG" == order.Type) {
		side = order.Side // 买
	} else {
		fmt.Println("err order side")
		return
	}
	currentOrder.Side = side

	if 0 >= quantityPrecision {
		quantity = fmt.Sprintf("%d", int64(quantityFloat))
	} else {
		quantity = strconv.FormatFloat(quantityFloat, 'f', int(quantityPrecision), 64)
	}

	fmt.Println(quantityFloat, quantity, quantityPrecision)
	currentOrder.Quantity, err = strconv.ParseFloat(quantity, 64)
	if nil != err {
		fmt.Println(err)
		return
	}

	// 请求下单
	binanceOrder, err = requestBinanceOrder(order.Coin, side, orderType, positionSide, quantity, user.ApiKey, user.ApiSecret)
	if nil != err {
		fmt.Println(err)
		return
	}

	if 0 >= binanceOrder.OrderId {
		fmt.Println(binanceOrder)
		return
	}
	currentOrder.OrderId = strconv.FormatInt(binanceOrder.OrderId, 10)

	currentOrder.CumQuote, err = strconv.ParseFloat(binanceOrder.CumQuote, 64)
	if nil != err {
		fmt.Println(err)
		return
	}

	currentOrder.ExecutedQty, err = strconv.ParseFloat(binanceOrder.ExecutedQty, 64)
	if nil != err {
		fmt.Println(err)
		return
	}

	currentOrder.AvgPrice, err = strconv.ParseFloat(binanceOrder.AvgPrice, 64)
	if nil != err {
		fmt.Println(err)
		return
	}

	// 计算收益 todo 使用数据库量太大以后
	if ("SELL" == order.Side && "LONG" == order.Type) || ("BUY" == order.Side && "SHORT" == order.Type) {
		historyOrders, err = b.binanceUserRepo.GetUserOrderByUserIdAndSymbolAndPositionSide(userBindTrader.UserId, order.Coin, positionSide)
		if nil != err {
			fmt.Println(err)
			return
		}

		var (
			historyAvgPrice float64
			historyCoin     float64 // 币数量
		)

		// 按下单顺序遍历，理论上在任何切面上开单币数永远大于等于关单币数
		for _, vHistoryOrders := range historyOrders {
			if ("SELL" == vHistoryOrders.Side && "SHORT" == vHistoryOrders.PositionSide) || ("BUY" == vHistoryOrders.Side && "LONG" == vHistoryOrders.PositionSide) {
				historyAvgPrice = (historyAvgPrice*historyCoin + vHistoryOrders.ExecutedQty*vHistoryOrders.AvgPrice) / (historyCoin + vHistoryOrders.ExecutedQty)
				historyCoin += vHistoryOrders.ExecutedQty
			} else if ("SELL" == vHistoryOrders.Side && "LONG" == vHistoryOrders.PositionSide) || ("BUY" == vHistoryOrders.Side && "SHORT" == vHistoryOrders.PositionSide) {
				historyCoin -= vHistoryOrders.ExecutedQty
				if 0 > historyCoin {
					fmt.Println("historyCoin err", userBindTrader.UserId, userBindTrader.TraderId, historyCoin, vHistoryOrders.ExecutedQty)
					return
				}
			}
		}

		// 放大精度
		income = int64((currentOrder.AvgPrice - historyAvgPrice) * currentOrder.ExecutedQty * 100000)
	}

	// 写入
	if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
		insertOrder, err = b.binanceUserRepo.InsertUserOrder(ctx, currentOrder)
		if nil != err {
			return err
		}

		// 平
		if ("SELL" == order.Side && "LONG" == order.Type) || ("BUY" == order.Side && "SHORT" == order.Type) {
			_, err = b.binanceUserRepo.UpdatesUserAmount(ctx, currentOrder.UserId, income)
			if nil != err {
				return err
			}

			_, err = b.binanceUserRepo.InsertUserAmountRecord(ctx, &UserAmountRecord{
				UserId:    currentOrder.UserId,
				OrderId:   insertOrder.ID,
				Amount:    income,
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			})
			if nil != err {
				return err
			}

		}

		return nil
	}); err != nil {
		fmt.Println(err)
		return
	}

	return
}

type BinanceOrder struct {
	OrderId       int64
	ExecutedQty   string
	ClientOrderId string
	Symbol        string
	AvgPrice      string
	CumQuote      string
	Side          string
	PositionSide  string
	ClosePosition bool
	Type          string
}

func requestBinanceOrder(symbol string, side string, orderType string, positionSide string, quantity string, apiKey string, secretKey string) (*BinanceOrder, error) {
	var (
		client *http.Client
		req    *http.Request
		resp   *http.Response
		res    *BinanceOrder
		data   string
		b      []byte
		err    error
		apiUrl = "https://fapi.binance.com/fapi/v1/order"
	)

	// 时间
	now := strconv.FormatInt(time.Now().UTC().UnixMilli(), 10)
	// 拼请求数据
	data = "symbol=" + symbol + "&side=" + side + "&type=" + orderType + "&positionSide=" + positionSide + "&newOrderRespType=" + "RESULT" + "&quantity=" + quantity + "&timestamp=" + now
	fmt.Println(data)
	// 加密
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	signature := hex.EncodeToString(h.Sum(nil))
	// 构造请求

	req, err = http.NewRequest("POST", apiUrl, strings.NewReader(data+"&signature="+signature))
	if err != nil {
		return nil, err
	}
	// 添加头信息
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-MBX-APIKEY", apiKey)

	// 请求执行
	client = &http.Client{Timeout: 3 * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	// 结果
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(resp.Header, string(b))

	var o BinanceOrder
	err = json.Unmarshal(b, &o)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res = &BinanceOrder{
		OrderId:       o.OrderId,
		ExecutedQty:   o.ExecutedQty,
		ClientOrderId: o.ClientOrderId,
		Symbol:        o.Symbol,
		AvgPrice:      o.AvgPrice,
		CumQuote:      o.CumQuote,
		Side:          o.Side,
		PositionSide:  o.PositionSide,
		ClosePosition: o.ClosePosition,
		Type:          o.Type,
	}

	return res, nil
}

type BinanceLeverAge struct {
	LeverAge int64
	symbol   string
}

func requestBinanceLeverAge(symbol string, leverAge int64, apiKey string, secretKey string) (*BinanceLeverAge, error) {
	var (
		client *http.Client
		req    *http.Request
		resp   *http.Response
		res    *BinanceLeverAge
		data   string
		b      []byte
		err    error
		apiUrl = "https://fapi.binance.com/fapi/v1/leverage"
	)

	// 时间
	now := strconv.FormatInt(time.Now().UTC().UnixMilli(), 10)
	// 拼请求数据
	data = "symbol=" + symbol + "&leverage=" + strconv.FormatInt(leverAge, 10) + "&timestamp=" + now
	// 加密
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	signature := hex.EncodeToString(h.Sum(nil))
	// 构造请求

	req, err = http.NewRequest("POST", apiUrl, strings.NewReader(data+"&signature="+signature))
	if err != nil {
		return nil, err
	}
	// 添加头信息
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-MBX-APIKEY", apiKey)

	// 请求执行
	client = &http.Client{Timeout: 3 * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	// 结果
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var l BinanceLeverAge
	err = json.Unmarshal(b, &l)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res = &BinanceLeverAge{
		LeverAge: l.LeverAge,
		symbol:   l.symbol,
	}

	return res, nil
}

func requestBinanceOrderInfo(symbol string, orderId int64, apiKey string, secretKey string) (*BinanceOrder, error) {
	var (
		client *http.Client
		req    *http.Request
		resp   *http.Response
		res    *BinanceOrder
		data   string
		b      []byte
		err    error
		apiUrl = "https://fapi.binance.com/fapi/v1/order"
	)

	// 时间
	now := strconv.FormatInt(time.Now().UTC().UnixMilli(), 10)
	// 拼请求数据
	data = "symbol=" + symbol + "&orderId=" + strconv.FormatInt(orderId, 10) + "&timestamp=" + now
	// 加密
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	signature := hex.EncodeToString(h.Sum(nil))
	// 构造请求

	req, err = http.NewRequest("GET", apiUrl, strings.NewReader(data+"&signature="+signature))
	if err != nil {
		return nil, err
	}
	// 添加头信息
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-MBX-APIKEY", apiKey)

	// 请求执行
	client = &http.Client{Timeout: 3 * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	// 结果
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var o BinanceOrder
	err = json.Unmarshal(b, &o)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res = &BinanceOrder{
		OrderId:       o.OrderId,
		ExecutedQty:   o.ExecutedQty,
		ClientOrderId: o.ClientOrderId,
		Symbol:        o.Symbol,
		AvgPrice:      o.AvgPrice,
		CumQuote:      o.CumQuote,
		Side:          o.Side,
		PositionSide:  o.PositionSide,
		ClosePosition: o.ClosePosition,
		Type:          o.Type,
	}

	return res, nil
}

type OrderHistory struct {
	OrderId      int64
	Qty          string
	Symbol       string
	Price        string
	Side         string
	RealizedPnl  string
	QuoteQty     string
	PositionSide string
	Time         int64
}

func requestBinanceOrderHistory(apiKey string, secretKey string, startTime string, endTime string) ([]*OrderHistory, error) {
	var (
		client *http.Client
		req    *http.Request
		resp   *http.Response
		res    []*OrderHistory
		data   string
		b      []byte
		err    error
		apiUrl = "https://fapi.binance.com/fapi/v1/userTrades"
	)

	// 时间
	now := strconv.FormatInt(time.Now().UTC().UnixMilli(), 10)
	// 拼请求数据
	data = "startTime=" + startTime + "&endTime=" + endTime + "&limit=10" + "&timestamp=" + now
	fmt.Println(data)
	// 加密
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	signature := hex.EncodeToString(h.Sum(nil))
	// 构造请求

	req, err = http.NewRequest("GET", apiUrl, strings.NewReader(data+"&signature="+signature))
	if err != nil {
		return nil, err
	}
	// 添加头信息
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-MBX-APIKEY", apiKey)

	// 请求执行
	client = &http.Client{Timeout: 3 * time.Second}
	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}

	// 结果
	defer func(Body io.ReadCloser) {
		err = Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	b, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println(11, string(b))

	var i []*OrderHistory
	err = json.Unmarshal(b, &i)
	if err != nil {
		return nil, err
	}

	res = make([]*OrderHistory, 0)
	for _, v := range i {
		res = append(res, v)

		//res = append(res, &OrderHistory{
		//	OrderId:      v[2].(int64),
		//	Qty:          v[2].(string),
		//	Symbol:       v[2].(string),
		//	Price:        v[2].(string),
		//	Side:         v[2].(string),
		//	PositionSide: v[2].(string),
		//	Time:         v[2].(int64),
		//	RealizedPnl:  v[2].(string),
		//	QuoteQty:     v[2].(string),
		//})
	}

	return res, nil
}

func (b *BinanceUserUsecase) Analyze(ctx context.Context, req *v1.AnalyzeRequest) (*v1.AnalyzeReply, error) {
	var (
		startTime time.Time
		now       = time.Now()
		err       error
	)

	// 指定日期和时间
	dateString := "2024-02-09 04:20:00"

	// 解析日期字符串
	startTime, err = time.Parse("2006-01-02 15:04:05", dateString)
	if err != nil {
		fmt.Println("解析日期出错:", err)
		return nil, nil
	}

	for startTime.Before(now) {

		endTime := startTime.Add(7 * 24 * time.Hour)
		requestBinanceOrderHistory(
			"DhfkUvUqqgQqhB3V7NKkdLXRqOFEcLHvQFzzrnpae2sSjoXogg9vqN4V6Z71i1Sm",
			"77HXUPdPnZiWdbA3qAjQ0eWKA19FHg1shC8qDsTSudcKrZPUMaSnDFSceLwPQhnD",
			strconv.FormatInt(startTime.Add(-8*time.Hour).UnixMilli(), 10),
			strconv.FormatInt(endTime.Add(-8*time.Hour).UnixMilli(), 10),
		)

		startTime = endTime
	}

	return nil, nil
}
