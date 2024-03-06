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
	ID                  uint64
	Address             string
	PlayType            uint64
	ApiStatus           uint64
	ApiKey              string
	ApiSecret           string
	BindTraderStatus    uint64
	BindTraderStatusTfi uint64
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

type UserBalance struct {
	ID        uint64
	UserId    uint64
	Balance   string
	Cost      uint64
	Open      uint64
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
	Status    uint64
	InitOrder uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserBindAfterUnbind struct {
	ID           uint64
	UserId       uint64
	TraderId     uint64
	Symbol       string
	PositionSide string
	Quantity     float64
	Status       uint64
	Amount       uint64
	CreatedAt    time.Time
	UpdatedAt    time.Time
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
	HandleStatus  int64
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
	UpdateUserBindTraderStatusNo(ctx context.Context, userId uint64) (bool, error)
	UpdateUserBindTraderStatusNoTfi(ctx context.Context, userId uint64) (bool, error)
	InsertUserBalance(ctx context.Context, userBalance *UserBalance) (bool, error)
	InsertUserBalanceTwo(ctx context.Context, userBalance *UserBalance) (bool, error)
	UpdatesUserBalance(ctx context.Context, userId uint64, balance string, cost uint64, open uint64) (bool, error)
	UpdatesUserBalanceTwo(ctx context.Context, userId uint64, balance string, cost uint64, open uint64) (bool, error)
	InsertUserBalanceRecord(ctx context.Context, userBalance *UserBalanceRecord) (bool, error)
	InsertUserBalanceRecordTwo(ctx context.Context, userBalanceRecord *UserBalanceRecord) (bool, error)
	InsertUserAmount(ctx context.Context, userAmount *UserAmount) (bool, error)
	InsertUserAmountTwo(ctx context.Context, userAmount *UserAmount) (bool, error)
	UpdatesUserAmount(ctx context.Context, userId uint64, amount int64) (bool, error)
	UpdatesUserAmountTwo(ctx context.Context, userId uint64, amount int64) (bool, error)
	InsertUserAmountRecord(ctx context.Context, userAmount *UserAmountRecord) (bool, error)
	InsertUserAmountRecordTwo(ctx context.Context, userAmount *UserAmountRecord) (bool, error)
	InsertUserBindTrader(ctx context.Context, userId uint64, traderId uint64, amount uint64) (*UserBindTrader, error)
	UpdatesUserBindTraderStatus(ctx context.Context, userId uint64, status uint64) (bool, error)
	UpdatesUserBindTraderStatusById(ctx context.Context, id uint64, status uint64) (bool, error)
	UpdatesUserBindTraderStatusAndInitOrderById(ctx context.Context, id uint64, status uint64, initOrder uint64) (bool, error)
	UpdatesUserBindTraderInitOrderById(ctx context.Context, id uint64) (bool, error)
	UpdatesUserBindTraderTwoInitOrderById(ctx context.Context, id uint64) (bool, error)
	UpdatesUserBindTraderStatusAndAmountById(ctx context.Context, id uint64, status uint64, amount uint64) (bool, error)
	DeleteUserBindTrader(ctx context.Context, userId uint64) (bool, error)
	InsertUserOrder(ctx context.Context, order *UserOrder) (*UserOrder, error)
	InsertUserOrderTwo(ctx context.Context, order *UserOrder) (*UserOrder, error)
	UpdatesUserOrderHandleStatus(ctx context.Context, id uint64) (bool, error)
	UpdatesUserOrderTwoHandleStatus(ctx context.Context, id uint64) (bool, error)
	InsertUserBindAfterUnbind(ctx context.Context, u *UserBindAfterUnbind) (*UserBindAfterUnbind, error)
	UpdatesUserBindAfterUnbind(ctx context.Context, id uint64, status uint64, quantity float64) (bool, error)
	UpdatesUserBindAfterUnbindTwo(ctx context.Context, id uint64, status uint64, quantity float64) (bool, error)
	GetUsers() ([]*User, error)
	GetUsersByUserIds(userIds []uint64) (map[uint64]*User, error)
	GetUsersByBindUserStatus() ([]*User, error)
	GetUsersByBindUserStatusReBind() ([]*User, error)
	GetUserByAddress(ctx context.Context, address string) (*User, error)
	GetUserById(ctx context.Context, userId uint64) (*User, error)
	GetUserByApiKeyAndApiSecret(key string, secret string) (*User, error)
	GetUserBalance(ctx context.Context, userId uint64) (*UserBalance, error)
	GetUserBalanceTwo(ctx context.Context, userId uint64) (*UserBalance, error)
	GetUserAmount(ctx context.Context, userId uint64) (*UserAmount, error)
	GetUserAmountTwo(ctx context.Context, userId uint64) (*UserAmount, error)
	GetUserBalanceByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*UserBalance, error)
	GetUserBalanceTwoByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*UserBalance, error)
	GetUserAmountByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*UserAmount, error)
	GetUserAmountTwoByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*UserAmount, error)
	GetTradersOrderByAmountDesc() ([]*Trader, error)
	GetTraders() (map[uint64]*Trader, error)
	GetUserBindTraderByUserId(userId uint64) ([]*UserBindTrader, error)
	GetUserBindTrader() (map[uint64][]*UserBindTrader, error)
	GetUserBindTraderByStatus(status uint64) (map[uint64][]*UserBindTrader, error)
	GetUserBindTraderMapByUserIds(userIds []uint64) (map[uint64][]*UserBindTrader, error)
	GetUserBindTraderByTraderIds(traderIds []uint64) (map[uint64][]*UserBindTrader, error)
	GetUserBindTraderTwoByTraderIds(traderIds []uint64) (map[uint64][]*UserBindTrader, error)
	GetUserBindAfterUnbindByTraderIds(traderIds []uint64) (map[uint64][]*UserBindAfterUnbind, error)
	GetUserBindAfterUnbindByUserIdAndTraderIdAndSymbol(ctx context.Context, userId uint64, traderId uint64, symbol string, positionSide string) (*UserBindAfterUnbind, error)
	GetSymbol() (map[string]*Symbol, error)
	GetUserOrderByUserTraderIdAndSymbol(userId uint64, traderId uint64, symbol string) ([]*UserOrder, error)
	GetUserOrderTwoByUserTraderIdAndSymbol(userId uint64, traderId uint64, symbol string) ([]*UserOrder, error)
	GetUserOrderByUserTraderId(userId uint64, traderId uint64) (map[string][]*UserOrder, error)
	GetUserOrderByUserIdMapId(userId uint64) (map[string]*UserOrder, error)
	GetUserOrderByUserIdAndSymbolAndPositionSide(userId uint64, symbol string, positionSide string) ([]*UserOrder, error)
	GetUserOrderByHandleStatus() ([]*UserOrder, error)
	GetUserOrderByUserIdGroupBySymbol(userId uint64) ([]*UserOrder, error)
	GetUserOrderByIds(ids []int64) ([]*UserOrder, error)
	GetUserOrderTwoByIds(ids []int64) ([]*UserOrder, error)
	GetUserOrderById(orderId int64) (*UserOrder, error)

	SAddOrderSetSellLongOrBuyShort(ctx context.Context, OrderId int64) error
	SMembersOrderSetSellLongOrBuyShort(ctx context.Context) ([]string, error)
	SRemOrderSetSellLongOrBuyShort(ctx context.Context, OrderId int64) error

	SAddOrderSetSellLongOrBuyShortTwo(ctx context.Context, OrderId int64) error
	SMembersOrderSetSellLongOrBuyShortTwo(ctx context.Context) ([]string, error)
	SRemOrderSetSellLongOrBuyShortTwo(ctx context.Context, OrderId int64) error
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

func (b *BinanceUserUsecase) SetUserBalanceAndUser(ctx context.Context, address string, balance string, cost uint64, open bool, playType uint64) error {
	var (
		user            *User
		userBalance     *UserBalance
		balance1        string
		balance2        string
		cost1           uint64
		cost2           uint64
		open1           uint64
		open2           uint64
		err             error
		lastUserBalance string // 上次用户余额
	)

	if 1 == playType {
		balance1 = balance
		cost1 = cost
		// todo
		//if open {
		open1 = 1
		//}

	} else if 2 == playType {
		balance2 = balance
		cost2 = cost
		//if open {
		open2 = 1
		//}

		var balanceTmp int64
		lengthToKeep := len(balance2) - 17 // todo 当len=19时表示1个代币，截取到10

		if lengthToKeep > 0 {
			balanceTmpStr := balance2[:lengthToKeep]
			balanceTmp, err = strconv.ParseInt(balanceTmpStr, 10, 64)
			if nil != err || 0 >= balanceTmp {
				return nil
			}
			// todo
			if balanceTmp*100 <= int64(cost2) {
				cost2 = uint64(balanceTmp * 100)
			}
		} else {
			// 不足1个tfi
			cost2 = 0
		}

	}
	user, err = b.binanceUserRepo.GetUserByAddress(ctx, address)
	if nil != err {
		return err
	}

	// 初始化 用户和余额
	if nil == user {
		if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
			user, err = b.binanceUserRepo.InsertUser(ctx, &User{
				Address: address, // 初始化玩法，模式
			})
			if nil != err {
				return err
			}

			if nil == user {
				return errors.New(500, "CREATE_USER_ERROR", "未发现创建的用户")
			}

			_, err = b.binanceUserRepo.InsertUserBalance(ctx, &UserBalance{
				UserId:  user.ID,
				Balance: balance1,
				Cost:    cost1,
				Open:    open1,
			})
			if nil != err {
				return err
			}

			_, err = b.binanceUserRepo.InsertUserBalanceTwo(ctx, &UserBalance{
				UserId:  user.ID,
				Balance: balance2,
				Cost:    cost2,
				Open:    open2,
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

			_, err = b.binanceUserRepo.InsertUserAmountTwo(ctx, &UserAmount{
				UserId: user.ID,
			})
			if nil != err {
				return err
			}

			_, err = b.binanceUserRepo.InsertUserBalanceRecord(ctx, &UserBalanceRecord{
				UserId:  user.ID,
				Amount:  balance1,
				Balance: lastUserBalance,
			})
			if nil != err {
				return err
			}

			_, err = b.binanceUserRepo.InsertUserBalanceRecordTwo(ctx, &UserBalanceRecord{
				UserId:  user.ID,
				Amount:  balance2,
				Balance: lastUserBalance,
			})
			if nil != err {
				return err
			}

			return nil
		}); err != nil {
			return err
		}

	} else {
		// 模式对应才给充值，初始化用户时已经决定了，条件很关键
		if 1 == playType {
			userBalance, err = b.binanceUserRepo.GetUserBalance(ctx, user.ID)
			if nil != err {
				return err
			}

			// 无变化
			if userBalance.Balance == balance1 && userBalance.Cost == cost1 && userBalance.Open == open1 {
				return nil
			}

			// 有变化
			tmpAmount := new(big.Int)
			if userBalance.Balance != balance1 {
				// 上次余额
				tmpLastUserBalance := new(big.Int)
				tmpLastUserBalance.SetString(userBalance.Balance, 10)
				// 本次余额
				tmpCurrentUserBalance := new(big.Int)
				tmpCurrentUserBalance.SetString(balance1, 10)
				// 增长
				tmpAmount.Sub(tmpCurrentUserBalance, tmpLastUserBalance)
			}

			tmpCost := userBalance.Cost
			if tmpCost != cost1 {
				tmpCost = cost1
			}

			// 更新余额
			if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
				_, err = b.binanceUserRepo.UpdatesUserBalance(ctx, user.ID, balance1, tmpCost, open1)
				if nil != err {
					return err
				}

				if userBalance.Balance != balance1 {
					_, err = b.binanceUserRepo.InsertUserBalanceRecord(ctx, &UserBalanceRecord{
						UserId:  user.ID,
						Amount:  tmpAmount.String(),
						Balance: userBalance.Balance,
					})
					if nil != err {
						return err
					}
				}

				// 开单额度变更时，更改绑定状态
				if userBalance.Cost != tmpCost {
					_, err = b.binanceUserRepo.UpdateUserBindTraderStatusNo(ctx, user.ID)
					if nil != err {
						return err
					}
				}

				return nil
			}); err != nil {
				return err
			}
		} else if 2 == playType {
			userBalance, err = b.binanceUserRepo.GetUserBalanceTwo(ctx, user.ID)
			if nil != err {
				return err
			}

			// 无变化
			if userBalance.Balance == balance2 && userBalance.Cost == cost2 && userBalance.Open == open2 {
				return nil
			}

			// 有变化
			tmpAmount := new(big.Int)
			if userBalance.Balance != balance2 {
				// 上次余额
				tmpLastUserBalance := new(big.Int)
				tmpLastUserBalance.SetString(userBalance.Balance, 10)
				// 本次余额
				tmpCurrentUserBalance := new(big.Int)
				tmpCurrentUserBalance.SetString(balance2, 10)
				// 增长
				tmpAmount.Sub(tmpCurrentUserBalance, tmpLastUserBalance)
			}

			tmpCost := userBalance.Cost
			if tmpCost != cost2 {
				tmpCost = cost2
			}

			// 更新余额
			if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
				_, err = b.binanceUserRepo.UpdatesUserBalanceTwo(ctx, user.ID, balance2, tmpCost, open2)
				if nil != err {
					return err
				}

				if userBalance.Balance != balance2 {
					_, err = b.binanceUserRepo.InsertUserBalanceRecordTwo(ctx, &UserBalanceRecord{
						UserId:  user.ID,
						Amount:  tmpAmount.String(),
						Balance: userBalance.Balance,
					})
					if nil != err {
						return err
					}
				}

				// 开单额度变更时，更改绑定状态
				if userBalance.Cost != tmpCost {
					_, err = b.binanceUserRepo.UpdateUserBindTraderStatusNoTfi(ctx, user.ID)
					if nil != err {
						return err
					}
				}

				return nil
			}); err != nil {
				return err
			}
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
				return err
			}

			symbol, err = b.binanceUserRepo.GetSymbol()
			if nil != err {
				return err
			}

			_, err = requestBinancePositionSide(apiKey, apiSecret)
			if nil != err {
				fmt.Println("更改持仓模式异常", err, user)
			}

			for k, _ := range symbol {
				_, err = requestBinanceLeverAge(k, int64(20), apiKey, apiSecret)
				if nil != err {
					continue
				}

				_, err = requestBinanceMarginType(k, apiKey, apiSecret)
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
		user           *User
		userBalance    *UserBalance
		userBalanceTwo *UserBalance
		userAmount     *UserAmount
		userAmountTwo  *UserAmount
		err            error
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

	userBalanceTwo, err = b.binanceUserRepo.GetUserBalanceTwo(ctx, user.ID)
	if nil != err {
		return &v1.GetUserReply{}, nil
	}

	userAmount, err = b.binanceUserRepo.GetUserAmount(ctx, user.ID)
	if nil != err {
		return &v1.GetUserReply{}, nil
	}

	userAmountTwo, err = b.binanceUserRepo.GetUserAmountTwo(ctx, user.ID)
	if nil != err {
		return &v1.GetUserReply{}, nil
	}
	return &v1.GetUserReply{Status: int64(user.ApiStatus), Play: int64(user.PlayType), Balance: userBalance.Balance, Amount: userAmount.Amount, BalanceTfi: userBalanceTwo.Balance, AmountTfi: userAmountTwo.Amount}, err
}

// BindTrader
// 用户的绑定交易员，和重新换绑的方法
// 重新换绑会将所有交易员绑定历史更新为不可用
// 重新绑定，且历史开单未全部平仓的交易员的代币数据会加入新数据方便下单逻辑查询，平仓使用
func (b *BinanceUserUsecase) BindTrader(ctx context.Context) (*v1.BindTraderReply, error) {
	var (
		err error
	)

	// 处理绑定
	err = b.UserBindTrader(ctx, nil)
	if nil != err {
		return nil, err
	}

	// 因为额度换绑
	err = b.ReBindTrader(ctx)
	if nil != err {
		return nil, err
	}

	// 处理更换单个trader
	err = b.ChangeBindTrader(ctx)
	if nil != err {
		return nil, err
	}

	return nil, nil
}

func (b *BinanceUserUsecase) UserBindTrader(ctx context.Context, users []*User) error {

	var (
		userBalance map[uint64]*UserBalance
		traders     []*Trader
		err         error
	)

	// 获取未绑定交易员的用户
	if nil == users {
		users, err = b.binanceUserRepo.GetUsersByBindUserStatus()
		if nil != err {
			return err
		}
	}

	if 0 == len(users) {
		return nil
	}

	// 获取用户cost
	userIds := make([]uint64, 0)
	for _, vUsers := range users {
		userIds = append(userIds, vUsers.ID)
	}
	userBalance, err = b.binanceUserRepo.GetUserBalanceByUserIds(ctx, userIds)
	if nil != err {
		return err
	}

	// 按amount排序好的交易员
	traders, err = b.binanceUserRepo.GetTradersOrderByAmountDesc()
	if nil != err {
		return err
	}

	if 0 == len(traders) {
		return nil
	}

	// 1. 遍历，轮询用户
	// 2. 轮询交易员（查询时已经按amount desc排序，用户cost的百分之30>=当前交易员的amount字段，建立绑定关系。
	// 3. 第二轮保留最后的交易员额度限制，将此和用户剩余可用额度对比，按顺序分配给此额度及以下的交易员。
	// 4. 轮询结束时，用户额度剩余不予理会。
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

		// 上述待绑定交易员结果集
		if 0 >= len(bindTrader) {
			continue
		}

		num := 0
		// 新增 和 更新
		insertUserBindTrader := make([]*UserBindTrader, 0)
		for k, v := range bindTrader {
			// todo 分配暂时10人 1人100
			if 9 < num {
				break
			}

			insertUserBindTrader = append(insertUserBindTrader, &UserBindTrader{
				UserId:   vUsers.ID,
				TraderId: k,
				Amount:   v.Amount,
			})

			num++
		}

		// 写入
		if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
			_, err = b.binanceUserRepo.UpdateUserBindTraderStatus(ctx, vUsers.ID)
			if nil != err {
				return err
			}

			for _, vInsertUserBindTrader := range insertUserBindTrader {
				_, err = b.binanceUserRepo.InsertUserBindTrader(ctx, vInsertUserBindTrader.UserId, vInsertUserBindTrader.TraderId, vInsertUserBindTrader.Amount)
				if nil != err {
					return err
				}
			}

			return nil
		}); err != nil {
			fmt.Println("换绑交易员写入异常", vUsers, err)
			continue
		}
	}

	return nil
}

func (b *BinanceUserUsecase) ReBindTrader(ctx context.Context) error {

	var (
		userBalance       map[uint64]*UserBalance
		userBindTraderMap map[uint64][]*UserBindTrader
		users             []*User
		traders           []*Trader
		symbols           map[string]*Symbol
		err               error
	)

	// 获取需要更新绑定交易员的用户
	users, err = b.binanceUserRepo.GetUsersByBindUserStatusReBind()
	if nil != err {
		return err
	}

	if 0 == len(users) {
		return nil
	}

	// 获取用户cost
	userIds := make([]uint64, 0)
	for _, vUsers := range users {
		userIds = append(userIds, vUsers.ID)
	}
	userBalance, err = b.binanceUserRepo.GetUserBalanceByUserIds(ctx, userIds)
	if nil != err {
		return err
	}

	// 获取用户绑定记录
	userBindTraderMap, err = b.binanceUserRepo.GetUserBindTraderMapByUserIds(userIds)
	if nil != err {
		return err
	}

	// 按amount排序好的交易员
	traders, err = b.binanceUserRepo.GetTradersOrderByAmountDesc()
	if nil != err {
		return err
	}

	if 0 == len(traders) {
		return nil
	}

	// 代币
	symbols, err = b.binanceUserRepo.GetSymbol()
	if nil != err {
		return err
	}

	// 1. 遍历，轮询用户
	// 2. 轮询交易员（查询时已经按amount desc排序，用户cost的百分之30>=当前交易员的amount字段，建立绑定关系。
	// 3. 第二轮保留最后的交易员额度限制，将此和用户剩余可用额度对比，按顺序分配给此额度及以下的交易员。
	// 4. 轮询结束时，用户额度剩余不予理会。
	for _, vUsers := range users {
		if _, exists := userBalance[vUsers.ID]; !exists {
			continue
		}

		tmpCost := userBalance[vUsers.ID].Cost
		// 不存在绑定记录, 新增
		if _, exists := userBindTraderMap[vUsers.ID]; !exists {
			// 新增
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

			// 上述待绑定交易员结果集
			if 0 >= len(bindTrader) {
				continue
			}

			// 新增 和 更新
			num := 0
			insertUserBindTrader := make([]*UserBindTrader, 0)
			for k, v := range bindTrader {
				// todo 分配暂时10人 1人100
				if 9 < num {
					break
				}

				insertUserBindTrader = append(insertUserBindTrader, &UserBindTrader{
					UserId:   vUsers.ID,
					TraderId: k,
					Amount:   v.Amount,
				})

				num++
			}

			// 写入
			if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
				_, err = b.binanceUserRepo.UpdateUserBindTraderStatus(ctx, vUsers.ID)
				if nil != err {
					return err
				}

				for _, vInsertUserBindTrader := range insertUserBindTrader {
					_, err = b.binanceUserRepo.InsertUserBindTrader(ctx, vInsertUserBindTrader.UserId, vInsertUserBindTrader.TraderId, vInsertUserBindTrader.Amount)
					if nil != err {
						return err
					}
				}

				return nil
			}); err != nil {
				fmt.Println("换绑交易员写入异常", vUsers, err)
			}

			continue
		}

		// 存在正在绑定记录，检查
		var (
			bindCost               uint64
			alreadyBindTraderCount uint64
		)

		alreadyBindTrader := make(map[uint64]uint64, 0)
		selfUpdateBindTrader := make(map[uint64]uint64, 0)
		for _, vUserBindTraderMap := range userBindTraderMap[vUsers.ID] {
			if 0 == vUserBindTraderMap.Status {
				bindCost += vUserBindTraderMap.Amount
				alreadyBindTraderCount++
			}

			// 因为额度调整换绑的不算
			if 3 == vUserBindTraderMap.Status {
				selfUpdateBindTrader[vUserBindTraderMap.TraderId] = vUserBindTraderMap.ID
				continue
			}

			alreadyBindTrader[vUserBindTraderMap.TraderId] = vUserBindTraderMap.TraderId
		}

		// 需要新增
		if bindCost < tmpCost {
			tmpCost -= bindCost

			var limitAmount uint64
			bindTrader := make(map[uint64]*Trader, 0)

			// 第一轮
			for _, vTraders := range traders {
				// 已经绑定过的交易员
				if _, ok := alreadyBindTrader[vTraders.ID]; ok {
					continue
				}

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
				// 已经绑定过的交易员
				if _, ok := alreadyBindTrader[vTraders.ID]; ok {
					continue
				}

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

			// 上述待绑定交易员结果集
			if 0 >= len(bindTrader) {
				continue
			}

			// 新增 和 更新
			// todo 分配暂时10人 1人100
			if alreadyBindTraderCount >= 10 {
				continue
			}

			num := uint64(0)
			insertUserBindTrader := make([]*UserBindTrader, 0)
			for k, v := range bindTrader {
				if 10-alreadyBindTraderCount <= num {
					break
				}

				insertUserBindTrader = append(insertUserBindTrader, &UserBindTrader{
					UserId:   vUsers.ID,
					TraderId: k,
					Amount:   v.Amount,
				})

				num++
			}

			// 写入
			if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
				_, err = b.binanceUserRepo.UpdateUserBindTraderStatus(ctx, vUsers.ID)
				if nil != err {
					return err
				}

				for _, vInsertUserBindTrader := range insertUserBindTrader {
					// 存在的写入，不存在更新
					if _, ok := selfUpdateBindTrader[vInsertUserBindTrader.TraderId]; ok {
						_, err = b.binanceUserRepo.UpdatesUserBindTraderStatusAndInitOrderById(ctx, selfUpdateBindTrader[vInsertUserBindTrader.TraderId], 0, 0)
					} else {
						_, err = b.binanceUserRepo.InsertUserBindTrader(ctx, vInsertUserBindTrader.UserId, vInsertUserBindTrader.TraderId, vInsertUserBindTrader.Amount)
					}

					if nil != err {
						return err
					}
				}

				return nil
			}); err != nil {
				fmt.Println("换绑交易员写入异常", vUsers, err)
			}

			continue
		} else if bindCost > tmpCost {
			// 需要减掉
			tmpMoreCost := tmpCost - bindCost
			updateBindTrader := make([]*UserBindTrader, 0)
			// 查询的结果已经排序好按amount
			for _, vUserBindTraderMap := range userBindTraderMap[vUsers.ID] {
				if 0 == vUserBindTraderMap.Status && tmpMoreCost >= vUserBindTraderMap.Amount {
					if 0 >= tmpMoreCost {
						break
					} else if tmpMoreCost >= vUserBindTraderMap.Amount {
						tmpMoreCost -= vUserBindTraderMap.Amount
						updateBindTrader = append(updateBindTrader, vUserBindTraderMap)
					}
				}
			}

			// 更新未status=1不再绑定的状态
			for _, vUpdateBindTrader := range updateBindTrader {
				// 解绑后的还会平单的数据
				userBindAfterUnbind := make([]*UserBindAfterUnbind, 0)
				if 1 == vUpdateBindTrader.InitOrder {
					// 添加未关单数据到后续表
					var (
						currentOrders map[string][]*UserOrder
					)
					currentOrders, err = b.binanceUserRepo.GetUserOrderByUserTraderId(vUpdateBindTrader.UserId, vUpdateBindTrader.TraderId)
					if nil != err {
						continue
					}
					if 0 < len(currentOrders) {
						// 用户下遍历对应的交易员
						for symbol, vCurrentOrders := range currentOrders {
							// 精度
							if _, ok := symbols[symbol]; !ok {
								continue
							}

							var (
								historyQuantityFloatMore  float64
								historyQuantityFloatEmpty float64
								quantity                  float64
							)

							if 0 >= symbols[symbol].QuantityPrecision {
								quantity = 1
							} else {
								quantity = 1 / float64(symbols[symbol].QuantityPrecision)
							}

							for _, vVCurrentOrders := range vCurrentOrders {

								// 历史开多和平多
								if "BUY" == vVCurrentOrders.Side && "LONG" == vVCurrentOrders.PositionSide {
									historyQuantityFloatMore += vVCurrentOrders.ExecutedQty
								} else if "SELL" == vVCurrentOrders.Side && "LONG" == vVCurrentOrders.PositionSide {
									historyQuantityFloatMore -= vVCurrentOrders.ExecutedQty
								}

								// 历史开空和平空
								if "SELL" == vVCurrentOrders.Side && "SHORT" == vVCurrentOrders.PositionSide {
									historyQuantityFloatEmpty += vVCurrentOrders.ExecutedQty
								} else if "BUY" == vVCurrentOrders.Side && "SHORT" == vVCurrentOrders.PositionSide {
									historyQuantityFloatEmpty -= vVCurrentOrders.ExecutedQty
								}
							}

							// 当前代币的多单还剩多少, 有的记录
							if 0 < historyQuantityFloatMore && quantity < historyQuantityFloatMore {
								userBindAfterUnbind = append(userBindAfterUnbind, &UserBindAfterUnbind{
									UserId:       vUpdateBindTrader.UserId,
									TraderId:     vUpdateBindTrader.TraderId,
									Amount:       vUpdateBindTrader.Amount,
									Symbol:       symbol,
									PositionSide: "LONG",
									Quantity:     historyQuantityFloatMore,
								})
							}

							// 当前代币的空单还剩多少, 有的记录
							if 0 < historyQuantityFloatEmpty && quantity < historyQuantityFloatEmpty {
								userBindAfterUnbind = append(userBindAfterUnbind, &UserBindAfterUnbind{
									UserId:       vUpdateBindTrader.UserId,
									TraderId:     vUpdateBindTrader.TraderId,
									Amount:       vUpdateBindTrader.Amount,
									Symbol:       symbol,
									PositionSide: "SHORT",
									Quantity:     historyQuantityFloatEmpty,
								})
							}
						}
					}
				}

				// 写入
				if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
					_, err = b.binanceUserRepo.UpdatesUserBindTraderStatusById(ctx, vUpdateBindTrader.ID, 3)
					if nil != err {
						return err
					}

					// 放置平仓表
					for _, vUserBindAfterUnbind := range userBindAfterUnbind {
						_, err = b.binanceUserRepo.InsertUserBindAfterUnbind(ctx, vUserBindAfterUnbind)
						if nil != err {
							return err
						}
					}

					return nil
				}); err != nil {
					fmt.Println("用户额度变更换绑失败", err, vUpdateBindTrader)
				}
			}

			continue
		} else {
			continue
		}

	}

	return nil
}

func (b *BinanceUserUsecase) ChangeBindTrader(ctx context.Context) error {
	var (
		userBindTrader map[uint64][]*UserBindTrader
		traders        map[uint64]*Trader
		symbols        map[string]*Symbol
		err            error
	)

	// 获取待更换的绑定关系
	userBindTrader, err = b.binanceUserRepo.GetUserBindTraderByStatus(2)
	if nil != err {
		return err
	}

	if 0 == len(userBindTrader) {
		return nil
	}

	// 代币
	symbols, err = b.binanceUserRepo.GetSymbol()
	if nil != err {
		return err
	}

	// 全部交易员
	traders, err = b.binanceUserRepo.GetTraders()
	if nil != err {
		return err
	}

	// 每个换绑的用户，换一个未绑定的同等amount的可用的交易员
	for traderId, vUserBindTrader := range userBindTrader {
		// 获取amount相同的traders
		if _, ok := traders[traderId]; !ok {
			fmt.Println("无法换绑，不存在交易员", traderId)
			continue
		}

		// 可替换的，去除当前的
		okTraders := make([]*Trader, 0)
		for _, vTraders := range traders {
			if traderId != vTraders.ID && vTraders.Amount == traders[traderId].Amount && 1 == vTraders.Status {
				okTraders = append(okTraders, vTraders)
			}
		}

		// 遍历需要换绑的用户
		for _, vVUserBindTrader := range vUserBindTrader {
			var (
				userBindTradersAll []*UserBindTrader
				insertTrader       *Trader
			)

			// 已经绑定的
			userBindTradersAll, err = b.binanceUserRepo.GetUserBindTraderByUserId(vVUserBindTrader.UserId)
			if nil != err {
				continue
			}
			userBindTradersAllMap := make(map[uint64]*UserBindTrader, 0)
			selfUpdateBindTrader := make(map[uint64]uint64, 0)
			for _, vUserBindTradersAll := range userBindTradersAll {
				// 因为额度调整，自己换绑的不算
				if 3 == userBindTradersAllMap[vUserBindTradersAll.TraderId].Status {
					selfUpdateBindTrader[vUserBindTradersAll.TraderId] = vUserBindTradersAll.ID
					continue
				}
				userBindTradersAllMap[vUserBindTradersAll.TraderId] = vUserBindTradersAll
			}

			// 差集，未绑定过的
			for _, vOkTraders := range okTraders {
				if _, ok := userBindTradersAllMap[vOkTraders.ID]; !ok {
					insertTrader = vOkTraders // 新增
					break
				}
			}

			if nil == insertTrader {
				continue
			}

			// 解绑后的还会平单的数据
			userBindAfterUnbind := make([]*UserBindAfterUnbind, 0)
			// 添加未关单数据到后续表
			if 1 == vVUserBindTrader.InitOrder {
				var (
					currentOrders map[string][]*UserOrder
				)
				currentOrders, err = b.binanceUserRepo.GetUserOrderByUserTraderId(vVUserBindTrader.UserId, vVUserBindTrader.TraderId)
				if nil != err {
					continue
				}
				if 0 < len(currentOrders) {
					// 用户下遍历对应的交易员
					for symbol, vCurrentOrders := range currentOrders {
						// 精度
						if _, ok := symbols[symbol]; !ok {
							continue
						}

						var (
							historyQuantityFloatMore  float64
							historyQuantityFloatEmpty float64
							quantity                  float64
						)

						if 0 >= symbols[symbol].QuantityPrecision {
							quantity = 1
						} else {
							quantity = 1 / float64(symbols[symbol].QuantityPrecision)
						}

						for _, vVCurrentOrders := range vCurrentOrders {

							// 历史开多和平多
							if "BUY" == vVCurrentOrders.Side && "LONG" == vVCurrentOrders.PositionSide {
								historyQuantityFloatMore += vVCurrentOrders.ExecutedQty
							} else if "SELL" == vVCurrentOrders.Side && "LONG" == vVCurrentOrders.PositionSide {
								historyQuantityFloatMore -= vVCurrentOrders.ExecutedQty
							}

							// 历史开空和平空
							if "SELL" == vVCurrentOrders.Side && "SHORT" == vVCurrentOrders.PositionSide {
								historyQuantityFloatEmpty += vVCurrentOrders.ExecutedQty
							} else if "BUY" == vVCurrentOrders.Side && "SHORT" == vVCurrentOrders.PositionSide {
								historyQuantityFloatEmpty -= vVCurrentOrders.ExecutedQty
							}
						}

						// 当前代币的多单还剩多少, 有的记录
						if 0 < historyQuantityFloatMore && quantity < historyQuantityFloatMore {
							userBindAfterUnbind = append(userBindAfterUnbind, &UserBindAfterUnbind{
								UserId:       vVUserBindTrader.UserId,
								TraderId:     vVUserBindTrader.TraderId,
								Amount:       vVUserBindTrader.Amount,
								Symbol:       symbol,
								PositionSide: "LONG",
								Quantity:     historyQuantityFloatMore,
							})
						}

						// 当前代币的空单还剩多少, 有的记录
						if 0 < historyQuantityFloatEmpty && quantity < historyQuantityFloatEmpty {
							userBindAfterUnbind = append(userBindAfterUnbind, &UserBindAfterUnbind{
								UserId:       vVUserBindTrader.UserId,
								TraderId:     vVUserBindTrader.TraderId,
								Amount:       vVUserBindTrader.Amount,
								Symbol:       symbol,
								PositionSide: "SHORT",
								Quantity:     historyQuantityFloatEmpty,
							})
						}

					}
				}
			}

			// 写入
			if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
				_, err = b.binanceUserRepo.UpdatesUserBindTraderStatusById(ctx, vVUserBindTrader.ID, 1)
				if nil != err {
					return err
				}

				// 存在则更新
				if _, ok := selfUpdateBindTrader[insertTrader.ID]; ok {
					_, err = b.binanceUserRepo.UpdatesUserBindTraderStatusAndInitOrderById(ctx, selfUpdateBindTrader[insertTrader.ID], 0, 0)
				} else {
					_, err = b.binanceUserRepo.InsertUserBindTrader(ctx, vVUserBindTrader.UserId, insertTrader.ID, vVUserBindTrader.Amount)
				}

				if nil != err {
					return err
				}

				// 放置平仓表
				for _, vUserBindAfterUnbind := range userBindAfterUnbind {
					_, err = b.binanceUserRepo.InsertUserBindAfterUnbind(ctx, vUserBindAfterUnbind)
					if nil != err {
						return err
					}
				}

				return nil
			}); err != nil {
				return err
			}

		}

	}

	return nil
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
		wg sync.WaitGroup
	)

	wg.Add(1) // 启动一个goroutine就登记+1
	go b.ListenTradersHandle(ctx, req, &wg)

	wg.Add(1) // 启动一个goroutine就登记+1
	go b.ListenTradersHandleTwo(ctx, req, &wg)

	wg.Wait() // 等待所有登记的goroutine都结束

	return &v1.ListenTraderAndUserOrderReply{
		Status: "ok",
	}, nil
}

func (b *BinanceUserUsecase) ListenTradersHandle(ctx context.Context, req *v1.ListenTraderAndUserOrderRequest, wg *sync.WaitGroup) {
	defer wg.Done()

	var (
		initOrderReq   = req.SendBody.InitOrder
		traderIds      []uint64
		userBindTrader map[uint64][]*UserBindTrader
		userIds        []uint64
		userIdsMap     map[uint64]uint64
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
		return
	}

	userIds = make([]uint64, 0)
	userIdsMap = make(map[uint64]uint64, 0)
	for _, vUserBindTrader := range userBindTrader {
		for _, vVUserBindTrader := range vUserBindTrader {
			if 1 == initOrderReq && 0 != vVUserBindTrader.InitOrder {
				continue
			}

			if _, ok := userIdsMap[vVUserBindTrader.UserId]; ok {
				continue
			}

			userIdsMap[vVUserBindTrader.UserId] = vVUserBindTrader.UserId
			userIds = append(userIds, vVUserBindTrader.UserId)
		}
	}

	if 0 >= len(userIds) {
		return
	}

	// 获取用户信息，余额信息，收益信息
	users, err = b.binanceUserRepo.GetUsersByUserIds(userIds)
	if nil != err {
		return
	}

	userBalance, err = b.binanceUserRepo.GetUserBalanceByUserIds(ctx, userIds)
	if nil != err {
		return
	}

	userAmount, err = b.binanceUserRepo.GetUserAmountByUserIds(ctx, userIds)
	if nil != err {
		return
	}

	symbol, err = b.binanceUserRepo.GetSymbol()
	if nil != err {
		return
	}

	for _, vOrders := range req.SendBody.Orders {
		for _, vOrdersData := range vOrders.Data {
			if _, exists := userBindTrader[vOrders.Uid]; !exists {
				continue
			}

			for _, vUserBindTrader := range userBindTrader[vOrders.Uid] {
				if 0 == vUserBindTrader.Status { // 绑定
					// 初始化仓位
					if 1 == initOrderReq && 0 != vUserBindTrader.InitOrder {
						continue
					}

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

						// 余额不足，收益大于余额的10倍
						if userAmount[vUserBindTrader.UserId].Amount > balanceTmp*10 {
							continue
						}

					} else {
						continue
					}

					// 精度
					if _, ok := symbol[vOrdersData.Symbol]; !ok {
						continue
					}

					fmt.Println(time.Now(), vOrdersData, vUserBindTrader)
					// 发送订单
					wg.Add(1) // 启动一个goroutine就登记+1
					go b.userOrderGoroutine(ctx, wg, &OrderData{
						Coin:  vOrdersData.Symbol,
						Type:  vOrdersData.Type,
						Price: vOrdersData.Price,
						Side:  vOrdersData.Side,
						Qty:   vOrdersData.Qty,
					}, vOrders.BaseMoney, users[vUserBindTrader.UserId], vUserBindTrader, symbol[vOrdersData.Symbol].QuantityPrecision, initOrderReq, vOrdersData.Proportion)
				}
			}
		}
	}
}

func (b *BinanceUserUsecase) ListenTradersHandleTwo(ctx context.Context, req *v1.ListenTraderAndUserOrderRequest, wg *sync.WaitGroup) {
	defer wg.Done()

	var (
		initOrderReq   = req.SendBody.InitOrder
		traderIds      []uint64
		userBindTrader map[uint64][]*UserBindTrader
		userIds        []uint64
		userIdsMap     map[uint64]uint64
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

	userBindTrader, err = b.binanceUserRepo.GetUserBindTraderTwoByTraderIds(traderIds)
	if nil != err {
		return
	}

	userIds = make([]uint64, 0)
	userIdsMap = make(map[uint64]uint64, 0)
	for _, vUserBindTrader := range userBindTrader {
		for _, vVUserBindTrader := range vUserBindTrader {
			if 1 == initOrderReq && 0 != vVUserBindTrader.InitOrder {
				continue
			}

			if _, ok := userIdsMap[vVUserBindTrader.UserId]; ok {
				continue
			}

			userIdsMap[vVUserBindTrader.UserId] = vVUserBindTrader.UserId
			userIds = append(userIds, vVUserBindTrader.UserId)
		}
	}

	if 0 >= len(userIds) {
		return
	}

	// 获取用户信息，余额信息，收益信息
	users, err = b.binanceUserRepo.GetUsersByUserIds(userIds)
	if nil != err {
		return
	}

	userBalance, err = b.binanceUserRepo.GetUserBalanceTwoByUserIds(ctx, userIds)
	if nil != err {
		return
	}

	userAmount, err = b.binanceUserRepo.GetUserAmountTwoByUserIds(ctx, userIds)
	if nil != err {
		return
	}

	symbol, err = b.binanceUserRepo.GetSymbol()
	if nil != err {
		return
	}

	for _, vOrders := range req.SendBody.Orders {
		for _, vOrdersData := range vOrders.Data {
			if _, exists := userBindTrader[vOrders.Uid]; !exists {
				continue
			}

			for _, vUserBindTrader := range userBindTrader[vOrders.Uid] {
				if 0 == vUserBindTrader.Status { // 绑定
					// 初始化仓位
					if 1 == initOrderReq && 0 != vUserBindTrader.InitOrder {
						continue
					}

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
						//// 余额不足，10u的收益，要1u的余额
						//// 精度按代币18位，截取小数点后到5位计算
						//var balanceTmp int64
						//lengthToKeep := len(userBalance[vUserBindTrader.UserId].Balance) - 13
						//
						//if lengthToKeep > 0 {
						//	balanceTmpStr := userBalance[vUserBindTrader.UserId].Balance[:lengthToKeep]
						//	balanceTmp, err = strconv.ParseInt(balanceTmpStr, 10, 64)
						//	if nil != err || 0 >= balanceTmp {
						//		continue
						//	}
						//} else {
						//	continue
						//}
						//
						//// 余额不足，收益大于余额的1000倍
						//if userAmount[vUserBindTrader.UserId].Amount > balanceTmp*1000 {
						//	continue
						//}
					} else {
						continue
					}

					// 精度
					if _, ok := symbol[vOrdersData.Symbol]; !ok {
						continue
					}
					fmt.Println(time.Now(), vOrdersData, vUserBindTrader)
					// 发送订单
					wg.Add(1) // 启动一个goroutine就登记+1
					go b.userOrderGoroutineTwo(ctx, wg, &OrderData{
						Coin:  vOrdersData.Symbol,
						Type:  vOrdersData.Type,
						Price: vOrdersData.Price,
						Side:  vOrdersData.Side,
						Qty:   vOrdersData.Qty,
					}, vOrders.BaseMoney, users[vUserBindTrader.UserId], vUserBindTrader, symbol[vOrdersData.Symbol].QuantityPrecision, initOrderReq, vOrdersData.Proportion)
				}
			}
		}
	}
}

// 用户下单
func (b *BinanceUserUsecase) userOrderGoroutine(ctx context.Context, wg *sync.WaitGroup, order *OrderData, amount string, user *User, userBindTrader *UserBindTrader, quantityPrecision int64, initOrderReq uint64, proportion string) {
	defer wg.Done() // goroutine结束就登记-1

	var (
		binanceOrder    *BinanceOrder
		orderType       = "MARKET"
		positionSide    string
		quantity        string
		qty             float64
		price           float64
		traderAmount    float64
		proportionFloat float64
		currentOrders   []*UserOrder
		currentOrder    *UserOrder
		insertOrder     *UserOrder
		err             error
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
		fmt.Println(err, order, userBindTrader)
		return
	}
	currentOrder.TraderQty = qty

	price, err = strconv.ParseFloat(order.Price, 64)
	if nil != err {
		fmt.Println(err, order, userBindTrader)
		return
	}
	currentOrder.Price = price

	traderAmount, err = strconv.ParseFloat(amount, 64)
	if nil != err {
		fmt.Println(err, order, userBindTrader)
		return
	}

	if 0 < len(proportion) {
		proportionFloat, err = strconv.ParseFloat(proportion, 64)
		if nil != err {
			fmt.Println(err, order, userBindTrader)
			return
		}
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

	var (
		quantityFloat float64
	)
	if 1 == initOrderReq {
		if 0 > currentOrder.Price {
			fmt.Println("err price", currentOrder.Price, order, userBindTrader)
			return
		}

		quantityFloat = float64(userBindTrader.Amount) * proportionFloat / currentOrder.Price
	} else {
		quantityFloat = float64(userBindTrader.Amount) * qty / traderAmount // 本次开单数量
	}

	var historyQuantityFloat float64
	// 本次平仓
	if ("SELL" == order.Side && "LONG" == order.Type) || ("BUY" == order.Side && "SHORT" == order.Type) {
		// 订单统计
		currentOrders, err = b.binanceUserRepo.GetUserOrderByUserTraderIdAndSymbol(userBindTrader.UserId, userBindTrader.TraderId, order.Coin)
		if nil != err {
			fmt.Println(err, order)
			return
		}
		// 查出用户的BUY单的币的数量，在对应的trader下，超过了BUY不能SELL todo 使用数据库量太大以后
		if 0 >= len(currentOrders) {
			return
		}

		// 多的部分不管，按剩余的数量关 todo 交易员全部平仓，少的部分另一个程序解决
		for _, vCurrentOrders := range currentOrders {
			// 本次平多
			if "SELL" == order.Side && "LONG" == order.Type {
				// 历史开多和平多
				if "BUY" == vCurrentOrders.Side && "LONG" == vCurrentOrders.PositionSide {
					historyQuantityFloat += vCurrentOrders.ExecutedQty
				} else if "SELL" == vCurrentOrders.Side && "LONG" == vCurrentOrders.PositionSide {
					historyQuantityFloat -= vCurrentOrders.ExecutedQty
				}
			} else if "BUY" == order.Side && "SHORT" == order.Type {
				// 历史开空和平空
				if "SELL" == vCurrentOrders.Side && "SHORT" == vCurrentOrders.PositionSide {
					historyQuantityFloat += vCurrentOrders.ExecutedQty
				} else if "BUY" == vCurrentOrders.Side && "SHORT" == vCurrentOrders.PositionSide {
					historyQuantityFloat -= vCurrentOrders.ExecutedQty
				}
			}
		}

		// 开单历史数量不足了
		if 0 > historyQuantityFloat {
			fmt.Println("trader的开单数量小于关单数量了，可能是精度问题", order.Coin, userBindTrader.UserId, userBindTrader.TraderId, historyQuantityFloat)
			return
		} else if 0 == historyQuantityFloat {
			fmt.Println("trader的开单数量等于关单数量了", order.Coin, userBindTrader.UserId, userBindTrader.TraderId, historyQuantityFloat)
			return
		}

		// 超过了净开单数量
		// todo 并发操作时数据不一致的可能性，会导致数量对不上，例如下单程序和更换绑定程序同时执行时，是否程序中统计的总数字会漏计算这里的新增的订单的数字
		if quantityFloat > historyQuantityFloat {
			quantityFloat = historyQuantityFloat
		}

	} else if ("SELL" == order.Side && "SHORT" == order.Type) || ("BUY" == order.Side && "LONG" == order.Type) {
		// 开仓 todo
		if 5 == user.ID {
			return
		}
	} else {
		fmt.Println("err order side")
		return
	}

	currentOrder.Side = order.Side

	// 精度调整
	if 0 >= quantityPrecision {
		quantity = fmt.Sprintf("%d", int64(quantityFloat))
	} else {
		quantity = strconv.FormatFloat(quantityFloat, 'f', int(quantityPrecision), 64)
	}

	currentOrder.Quantity, err = strconv.ParseFloat(quantity, 64)
	if nil != err {
		fmt.Println(err)
		return
	}

	// 请求下单
	binanceOrder, err = requestBinanceOrder(order.Coin, order.Side, orderType, positionSide, quantity, user.ApiKey, user.ApiSecret)
	if nil != err {
		fmt.Println(err)
		return
	}

	if 0 >= binanceOrder.OrderId {
		fmt.Println("下单异常:数量,", quantity, "下单信息:", currentOrder, "是否初始化订单:", initOrderReq, proportion)
		return
	}

	currentOrder.OrderId = strconv.FormatInt(binanceOrder.OrderId, 10)

	currentOrder.CumQuote, err = strconv.ParseFloat(binanceOrder.CumQuote, 64)
	if nil != err {
		fmt.Println(err, binanceOrder)
		return
	}

	currentOrder.ExecutedQty, err = strconv.ParseFloat(binanceOrder.ExecutedQty, 64)
	if nil != err {
		fmt.Println(err, binanceOrder)
		return
	}

	currentOrder.AvgPrice, err = strconv.ParseFloat(binanceOrder.AvgPrice, 64)
	if nil != err {
		fmt.Println(err, binanceOrder)
		return
	}

	// 写入
	if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
		insertOrder, err = b.binanceUserRepo.InsertUserOrder(ctx, currentOrder)
		if nil != err {
			return err
		}

		return nil
	}); err != nil {
		fmt.Println(err, currentOrder)
		return
	}

	// 计算收益 卖单去binance查询本单收益
	if ("SELL" == order.Side && "LONG" == order.Type) || ("BUY" == order.Side && "SHORT" == order.Type) {
		if 0 >= insertOrder.ID {
			fmt.Println("错误的数据写入", insertOrder)
			return
		}

		err = b.binanceUserRepo.SAddOrderSetSellLongOrBuyShort(ctx, int64(insertOrder.ID))
		if nil != err {
			fmt.Println("错误的数据写入, redis", insertOrder, err)
			return
		}
	}

	// 新事务 写入防止锁等待可能影响订单的insert
	if 1 == initOrderReq {
		if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
			_, err = b.binanceUserRepo.UpdatesUserBindTraderInitOrderById(ctx, userBindTrader.ID)
			if nil != err {
				return err
			}

			return nil
		}); err != nil {
			fmt.Println(err, currentOrder)
			return
		}
	}

	return
}

func (b *BinanceUserUsecase) userOrderGoroutineTwo(ctx context.Context, wg *sync.WaitGroup, order *OrderData, amount string, user *User, userBindTrader *UserBindTrader, quantityPrecision int64, initOrderReq uint64, proportion string) {
	defer wg.Done() // goroutine结束就登记-1

	var (
		binanceOrder    *BinanceOrder
		orderType       = "MARKET"
		positionSide    string
		quantity        string
		qty             float64
		price           float64
		traderAmount    float64
		proportionFloat float64
		currentOrders   []*UserOrder
		currentOrder    *UserOrder
		insertOrder     *UserOrder
		err             error
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
		fmt.Println(err, order, userBindTrader)
		return
	}
	currentOrder.TraderQty = qty

	price, err = strconv.ParseFloat(order.Price, 64)
	if nil != err {
		fmt.Println(err, order, userBindTrader)
		return
	}
	currentOrder.Price = price

	traderAmount, err = strconv.ParseFloat(amount, 64)
	if nil != err {
		fmt.Println(err, order, userBindTrader)
		return
	}

	if 0 < len(proportion) {
		proportionFloat, err = strconv.ParseFloat(proportion, 64)
		if nil != err {
			fmt.Println(err, order, userBindTrader)
			return
		}
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

	var (
		quantityFloat float64
	)
	if 1 == initOrderReq {
		if 0 > currentOrder.Price {
			fmt.Println("err price", currentOrder.Price, order, userBindTrader)
			return
		}

		quantityFloat = float64(userBindTrader.Amount) * proportionFloat / currentOrder.Price
	} else {
		quantityFloat = float64(userBindTrader.Amount) * qty / traderAmount // 本次开单数量
	}

	var historyQuantityFloat float64
	// 本次平仓
	if ("SELL" == order.Side && "LONG" == order.Type) || ("BUY" == order.Side && "SHORT" == order.Type) {
		// 订单统计
		currentOrders, err = b.binanceUserRepo.GetUserOrderTwoByUserTraderIdAndSymbol(userBindTrader.UserId, userBindTrader.TraderId, order.Coin)
		if nil != err {
			fmt.Println(err)
			return
		}
		// 查出用户的BUY单的币的数量，在对应的trader下，超过了BUY不能SELL todo 使用数据库量太大以后
		if 0 >= len(currentOrders) {
			return
		}

		// 多的部分不管，按剩余的数量关 todo 交易员全部平仓，少的部分另一个程序解决
		for _, vCurrentOrders := range currentOrders {
			// 本次平多
			if "SELL" == order.Side && "LONG" == order.Type {
				// 历史开多和平多
				if "BUY" == vCurrentOrders.Side && "LONG" == vCurrentOrders.PositionSide {
					historyQuantityFloat += vCurrentOrders.ExecutedQty
				} else if "SELL" == vCurrentOrders.Side && "LONG" == vCurrentOrders.PositionSide {
					historyQuantityFloat -= vCurrentOrders.ExecutedQty
				}
			} else if "BUY" == order.Side && "SHORT" == order.Type {
				// 历史开空和平空
				if "SELL" == vCurrentOrders.Side && "SHORT" == vCurrentOrders.PositionSide {
					historyQuantityFloat += vCurrentOrders.ExecutedQty
				} else if "BUY" == vCurrentOrders.Side && "SHORT" == vCurrentOrders.PositionSide {
					historyQuantityFloat -= vCurrentOrders.ExecutedQty
				}
			}
		}

		// 开单历史数量不足了
		if 0 > historyQuantityFloat {
			fmt.Println("trader的开单数量小于关单数量了，可能是精度问题", order.Coin, userBindTrader.UserId, userBindTrader.TraderId, historyQuantityFloat)
			return
		} else if 0 == historyQuantityFloat {
			fmt.Println("trader的开单数量等于关单数量了", order.Coin, userBindTrader.UserId, userBindTrader.TraderId, historyQuantityFloat)
			return
		}

		// 超过了净开单数量
		// todo 并发操作时数据不一致的可能性，会导致数量对不上，例如下单程序和更换绑定程序同时执行时，是否程序中统计的总数字会漏计算这里的新增的订单的数字
		if quantityFloat > historyQuantityFloat {
			quantityFloat = historyQuantityFloat
		}

	} else if ("SELL" == order.Side && "SHORT" == order.Type) || ("BUY" == order.Side && "LONG" == order.Type) {
		// 开仓 todo
		if 3 == user.ID || 6 == user.ID {
			return
		}
	} else {
		fmt.Println("err order side")
		return
	}

	currentOrder.Side = order.Side

	// 精度调整
	if 0 >= quantityPrecision {
		quantity = fmt.Sprintf("%d", int64(quantityFloat))
	} else {
		quantity = strconv.FormatFloat(quantityFloat, 'f', int(quantityPrecision), 64)
	}

	currentOrder.Quantity, err = strconv.ParseFloat(quantity, 64)
	if nil != err {
		fmt.Println(err)
		return
	}

	// 请求下单
	binanceOrder, err = requestBinanceOrder(order.Coin, order.Side, orderType, positionSide, quantity, user.ApiKey, user.ApiSecret)
	if nil != err {
		fmt.Println(err)
		return
	}

	if 0 >= binanceOrder.OrderId {
		fmt.Println("下单异常:数量,", quantity, "下单信息:", currentOrder, "是否初始化订单:", initOrderReq, proportion)
		return
	}

	currentOrder.OrderId = strconv.FormatInt(binanceOrder.OrderId, 10)

	currentOrder.CumQuote, err = strconv.ParseFloat(binanceOrder.CumQuote, 64)
	if nil != err {
		fmt.Println(err, binanceOrder)
		return
	}

	currentOrder.ExecutedQty, err = strconv.ParseFloat(binanceOrder.ExecutedQty, 64)
	if nil != err {
		fmt.Println(err, binanceOrder)
		return
	}

	currentOrder.AvgPrice, err = strconv.ParseFloat(binanceOrder.AvgPrice, 64)
	if nil != err {
		fmt.Println(err, binanceOrder)
		return
	}

	// 写入
	if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
		insertOrder, err = b.binanceUserRepo.InsertUserOrderTwo(ctx, currentOrder)
		if nil != err {
			return err
		}

		return nil
	}); err != nil {
		fmt.Println(err, insertOrder)
		return
	}

	// 计算收益 卖单去binance查询本单收益
	if ("SELL" == order.Side && "LONG" == order.Type) || ("BUY" == order.Side && "SHORT" == order.Type) {
		if 0 >= insertOrder.ID {
			fmt.Println("错误的数据写入", insertOrder)
			return
		}

		err = b.binanceUserRepo.SAddOrderSetSellLongOrBuyShortTwo(ctx, int64(insertOrder.ID))
		if nil != err {
			fmt.Println("错误的数据写入, redis", insertOrder, err)
			return
		}
	}

	// 新事务 写入防止锁等待可能影响订单的insert
	if 1 == initOrderReq {
		if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
			_, err = b.binanceUserRepo.UpdatesUserBindTraderTwoInitOrderById(ctx, userBindTrader.ID)
			if nil != err {
				return err
			}

			return nil
		}); err != nil {
			fmt.Println(err, currentOrder)
			return
		}
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
		fmt.Println(string(b))
		fmt.Println(err)
		return nil, err
	}

	var o BinanceOrder
	err = json.Unmarshal(b, &o)
	if err != nil {
		fmt.Println(string(b))
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

	if 0 >= res.OrderId {
		fmt.Println(string(b))
	}

	return res, nil
}

type BinanceLeverAge struct {
	LeverAge int64
	symbol   string
}

type BinanceMarginType struct {
	Code int64
	Msg  string
}

type BinancePositionSide struct {
	Code int64
	Msg  string
}

// 更改杠杆倍率
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

// 更改逐全仓模式
func requestBinanceMarginType(symbol string, apiKey string, secretKey string) (*BinanceMarginType, error) {
	var (
		client *http.Client
		req    *http.Request
		resp   *http.Response
		res    *BinanceMarginType
		data   string
		b      []byte
		err    error
		apiUrl = "https://fapi.binance.com/fapi/v1/marginType"
	)

	// 时间
	now := strconv.FormatInt(time.Now().UTC().UnixMilli(), 10)
	// 拼请求数据
	data = "symbol=" + symbol + "&marginType=CROSSED" + "&timestamp=" + now
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

	var l BinanceMarginType
	err = json.Unmarshal(b, &l)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res = &BinanceMarginType{
		Code: l.Code,
		Msg:  l.Msg,
	}

	return res, nil
}

// 更改持仓模式
func requestBinancePositionSide(apiKey string, secretKey string) (*BinancePositionSide, error) {
	var (
		client *http.Client
		req    *http.Request
		resp   *http.Response
		res    *BinancePositionSide
		data   string
		b      []byte
		err    error
		apiUrl = "https://fapi.binance.com/fapi/v1/positionSide/dual"
	)

	// 时间
	now := strconv.FormatInt(time.Now().UTC().UnixMilli(), 10)
	// 拼请求数据
	data = "dualSidePosition=true&timestamp=" + now
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

	var l BinancePositionSide
	err = json.Unmarshal(b, &l)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res = &BinancePositionSide{
		Code: l.Code,
		Msg:  l.Msg,
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

func (b *BinanceUserUsecase) Analyze(ctx context.Context, req *v1.AnalyzeRequest) (*v1.AnalyzeReply, error) {
	//var (
	//	orders     []*UserOrder
	//	ordersUser map[string]*UserOrder
	//	order1     map[string]*UserOrder
	//	total      int64
	//	u          float64
	//	totalU     int64
	//	u2         float64
	//	totalU2    int64
	//	err        error
	//)
	//
	//orders, err = b.binanceUserRepo.GetUserOrderByUserIdGroupBySymbol(4)
	//if nil != err {
	//	return nil, nil
	//}
	//
	//ordersUser, err = b.binanceUserRepo.GetUserOrderByUserIdMapId(4)
	//if nil != err {
	//	return nil, nil
	//}
	//
	//order1 = make(map[string]*UserOrder, 0)
	//
	//for _, v := range orders {
	//	var (
	//		binanceOrder []*OrderHistory
	//		start        = int64(1707346800000)
	//	)
	//
	//	for start <= 1708337023116 {
	//
	//		var (
	//			tmpBinanceOrder []*OrderHistory
	//			tmpStart        string
	//			end             string
	//		)
	//
	//		tmpStart = strconv.FormatInt(start, 10)
	//		if start+432000000 >= 1708337023116 {
	//			end = strconv.FormatInt(1708337023116, 10)
	//		} else {
	//			end = strconv.FormatInt(start+432000000, 10)
	//		}
	//
	//		//fmt.Println(v.Symbol, tmpStart, end)
	//		tmpBinanceOrder, err = requestBinanceOrderHistory(
	//			"Bc8Id20hm7sBIrKewsRAvA4RX2MPmy546X6q1wOGlsBQsSoRdn6Uze6lHAim8YY1",
	//			"DAzqoU9KbrQfxrWNMKdbNYgtWSxZpFVPSRPxeN1k08ed7cNGTACc7dy9oK8eZwrG",
	//			v.Symbol,
	//			0,
	//			tmpStart, end)
	//		if nil != err {
	//			fmt.Println("错误查询", err, v.Symbol)
	//		}
	//
	//		for _, vTmpBinanceOrder := range tmpBinanceOrder {
	//			binanceOrder = append(binanceOrder, vTmpBinanceOrder)
	//		}
	//
	//		start += 432000000
	//
	//	}
	//
	//	for _, vBinanceOrder := range binanceOrder {
	//		tmpOrderId := strconv.FormatInt(vBinanceOrder.OrderId, 10)
	//		//if _, ok := ordersUser[tmpOrderId]; !ok {
	//		//	fmt.Println("不存在系统下单", tmpOrderId)
	//		//}
	//
	//		total++
	//		// 平多
	//		if "SELL" == vBinanceOrder.Side && "LONG" == vBinanceOrder.PositionSide {
	//			var tmp float64
	//			tmp, err = strconv.ParseFloat(vBinanceOrder.RealizedPnl, 64)
	//			if nil != err {
	//				return nil, err
	//			}
	//
	//			fmt.Println("收益此时：", u, vBinanceOrder.RealizedPnl, tmp, vBinanceOrder)
	//
	//			u += tmp
	//			totalU++
	//
	//			if _, ok := ordersUser[tmpOrderId]; ok {
	//				u2 += tmp
	//
	//				if _, ok2 := order1[tmpOrderId]; !ok2 {
	//					totalU2++
	//					order1[tmpOrderId] = ordersUser[tmpOrderId]
	//				}
	//			}
	//		}
	//
	//		// 平空
	//		if "BUY" == vBinanceOrder.Side && "SHORT" == vBinanceOrder.PositionSide {
	//			var tmp float64
	//			tmp, err = strconv.ParseFloat(vBinanceOrder.RealizedPnl, 64)
	//			if nil != err {
	//				return nil, err
	//			}
	//
	//			fmt.Println("收益此时：", u, vBinanceOrder.RealizedPnl, tmp, vBinanceOrder)
	//
	//			u += tmp
	//			totalU++
	//
	//			if _, ok := ordersUser[tmpOrderId]; ok {
	//				fmt.Println("系统收益此时：", u2)
	//
	//				u2 += tmp
	//
	//				if _, ok2 := order1[tmpOrderId]; !ok2 {
	//
	//					totalU2++
	//
	//					order1[tmpOrderId] = ordersUser[tmpOrderId]
	//				}
	//			}
	//		}
	//
	//		// 开多
	//		if "BUY" == vBinanceOrder.Side && "LONG" == vBinanceOrder.PositionSide {
	//
	//		}
	//
	//		// 开空
	//		if "SELL" == vBinanceOrder.Side && "SHORT" == vBinanceOrder.PositionSide {
	//
	//		}
	//	}
	//
	//}
	//
	//fmt.Println("共单数：", total, "收益：", u, "平仓单共：", totalU, "系统订单：", len(ordersUser), "系统收益：", u2, "系统收益共", totalU2, "a", len(order1))
	var (
		userOrders []*UserOrder
		err        error
	)
	userOrders, err = b.binanceUserRepo.GetUserOrderByHandleStatus()
	if nil != err {
		return nil, err
	}

	for _, v := range userOrders {
		err = b.binanceUserRepo.SAddOrderSetSellLongOrBuyShort(ctx, int64(v.ID))
		if nil != err {
			fmt.Println("错误的数据写入, redis", v, err)
		}
	}

	return nil, nil
}

func (b *BinanceUserUsecase) OrderHandle(ctx context.Context, req *v1.OrderHandleRequest) (*v1.OrderHandleReply, error) {
	var (
		wg          sync.WaitGroup
		orderIdsStr []string
		orderIds    []int64
		userOrders  []*UserOrder
		userIdsMap  map[uint64]uint64
		userIds     []uint64
		users       map[uint64]*User
		err         error
	)

	orderIdsStr, err = b.binanceUserRepo.SMembersOrderSetSellLongOrBuyShort(ctx)
	if nil != err {
		fmt.Println(err)
	}

	if 0 >= len(orderIdsStr) {
		return &v1.OrderHandleReply{}, nil
	}

	// 订单信息
	orderIds = make([]int64, 0)
	for _, vOrderIdsStr := range orderIdsStr {
		var tmp int64
		tmp, err = strconv.ParseInt(vOrderIdsStr, 10, 64)
		if nil != err {
			fmt.Println(err)
		}

		orderIds = append(orderIds, tmp)
	}
	userOrders, err = b.binanceUserRepo.GetUserOrderByIds(orderIds)
	if nil != err {
		fmt.Println("更新平单后收益数据失败, 空查询, 订单不存在", err, orderIds)
		return &v1.OrderHandleReply{}, nil
	}

	// 用户信息
	userIdsMap = make(map[uint64]uint64, 0)
	for _, vUserOrders := range userOrders {
		userIdsMap[vUserOrders.UserId] = vUserOrders.UserId
	}
	userIds = make([]uint64, 0)
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}
	users, err = b.binanceUserRepo.GetUsersByUserIds(userIds)
	if nil != err {
		fmt.Println("更新平单后收益数据失败, 空查询，用户查询不存在", err, userIds)
		return &v1.OrderHandleReply{}, nil
	}

	for k, vUserOrders := range userOrders {
		if 100 <= k {
			break
		}

		if _, ok := users[vUserOrders.UserId]; !ok {
			fmt.Println("更新平单收益，不存在用户", vUserOrders)
			continue
		}

		wg.Add(1) // 启动一个goroutine就登记+1
		go b.userOrderHandleGoroutine(ctx, &wg, vUserOrders, users[vUserOrders.UserId])
	}

	wg.Wait() // 等待所有登记的goroutine都结束

	return &v1.OrderHandleReply{}, nil
}

func (b *BinanceUserUsecase) userOrderHandleGoroutine(ctx context.Context, wg *sync.WaitGroup, userOrder *UserOrder, user *User) {
	defer wg.Done() // goroutine结束就登记-1

	var (
		binanceOrder []*OrderHistory
		income       float64
		orderIdInt   int64
		err          error
	)

	// 平
	if !(("SELL" == userOrder.Side && "LONG" == userOrder.PositionSide) || ("BUY" == userOrder.Side && "SHORT" == userOrder.PositionSide)) {
		fmt.Println("非平仓", userOrder.ID)
		return
	}

	if 0 < userOrder.HandleStatus {
		//继续执行删除
	} else {

		// 查询
		orderIdInt, err = strconv.ParseInt(userOrder.OrderId, 10, 64)

		binanceOrder, err = requestBinanceOrderHistory(user.ApiKey, user.ApiSecret, userOrder.Symbol, orderIdInt, "", "")
		if nil != err {
			fmt.Println("更新平单后收益数据失败, 错误查询", err, userOrder.ID)
			return
		}

		if 0 >= len(binanceOrder) {
			fmt.Println("更新平单后收益数据失败, 空数据", binanceOrder, userOrder.ID)
			return
		}

		// 收益
		for _, vBinanceOrder := range binanceOrder {
			var (
				tmpIncome float64
			)
			tmpIncome, err = strconv.ParseFloat(vBinanceOrder.RealizedPnl, 64)
			if nil != err {
				fmt.Println("更新平单后收益数据失败", err, vBinanceOrder)
				return
			}

			income += tmpIncome
		}

		// 写入
		if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
			_, err = b.binanceUserRepo.UpdatesUserOrderHandleStatus(ctx, userOrder.ID)
			if nil != err {
				return err
			}

			_, err = b.binanceUserRepo.UpdatesUserAmount(ctx, userOrder.UserId, int64(income*100000))
			if nil != err {
				return err
			}

			_, err = b.binanceUserRepo.InsertUserAmountRecord(ctx, &UserAmountRecord{
				UserId:    userOrder.UserId,
				OrderId:   userOrder.ID,
				Amount:    int64(income * 100000),
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			})
			if nil != err {
				return err
			}

			return nil
		}); err != nil {
			fmt.Println("更新平单后收益数据失败", err, userOrder.ID)
			return
		}
	}

	err = b.binanceUserRepo.SRemOrderSetSellLongOrBuyShort(ctx, int64(userOrder.ID))
	if nil != err {
		fmt.Println("更新平单后收益数据失败, 错误的数据删除, redis，order_id:", int64(userOrder.ID), userOrder.ID, err)
		return
	}

	return
}

func (b *BinanceUserUsecase) OrderHandleTwo(ctx context.Context, req *v1.OrderHandleRequest) (*v1.OrderHandleReply, error) {
	var (
		wg          sync.WaitGroup
		orderIdsStr []string
		orderIds    []int64
		userOrders  []*UserOrder
		userIdsMap  map[uint64]uint64
		userIds     []uint64
		users       map[uint64]*User
		err         error
	)

	orderIdsStr, err = b.binanceUserRepo.SMembersOrderSetSellLongOrBuyShortTwo(ctx)
	if nil != err {
		fmt.Println(err)
	}

	if 0 >= len(orderIdsStr) {
		return &v1.OrderHandleReply{}, nil
	}

	// 订单信息
	orderIds = make([]int64, 0)
	for _, vOrderIdsStr := range orderIdsStr {
		var tmp int64
		tmp, err = strconv.ParseInt(vOrderIdsStr, 10, 64)
		if nil != err {
			fmt.Println(err)
		}

		orderIds = append(orderIds, tmp)
	}
	userOrders, err = b.binanceUserRepo.GetUserOrderTwoByIds(orderIds)
	if nil != err {
		fmt.Println("更新平单后收益数据失败, 空查询, 订单不存在", err, orderIds)
		return &v1.OrderHandleReply{}, nil
	}

	// 用户信息
	userIdsMap = make(map[uint64]uint64, 0)
	for _, vUserOrders := range userOrders {
		userIdsMap[vUserOrders.UserId] = vUserOrders.UserId
	}
	userIds = make([]uint64, 0)
	for _, v := range userIdsMap {
		userIds = append(userIds, v)
	}
	users, err = b.binanceUserRepo.GetUsersByUserIds(userIds)
	if nil != err {
		fmt.Println("更新平单后收益数据失败, 空查询，用户查询不存在", err, userIds)
		return &v1.OrderHandleReply{}, nil
	}

	for k, vUserOrders := range userOrders {
		if 100 <= k {
			break
		}

		if _, ok := users[vUserOrders.UserId]; !ok {
			fmt.Println("更新平单收益，不存在用户", vUserOrders)
			continue
		}

		wg.Add(1) // 启动一个goroutine就登记+1
		go b.userOrderHandleGoroutineTwo(ctx, &wg, vUserOrders, users[vUserOrders.UserId])
	}

	wg.Wait() // 等待所有登记的goroutine都结束

	return &v1.OrderHandleReply{}, nil
}

func (b *BinanceUserUsecase) userOrderHandleGoroutineTwo(ctx context.Context, wg *sync.WaitGroup, userOrder *UserOrder, user *User) {
	defer wg.Done() // goroutine结束就登记-1

	var (
		binanceOrder []*OrderHistory
		income       float64
		orderIdInt   int64
		err          error
	)

	// 平
	if !(("SELL" == userOrder.Side && "LONG" == userOrder.PositionSide) || ("BUY" == userOrder.Side && "SHORT" == userOrder.PositionSide)) {
		fmt.Println("非平仓", userOrder.ID)
		return
	}

	if 0 < userOrder.HandleStatus {
		//继续执行删除
	} else {

		// 查询
		orderIdInt, err = strconv.ParseInt(userOrder.OrderId, 10, 64)

		binanceOrder, err = requestBinanceOrderHistory(user.ApiKey, user.ApiSecret, userOrder.Symbol, orderIdInt, "", "")
		if nil != err {
			fmt.Println("更新平单后收益数据失败, 错误查询", err, userOrder.ID)
			return
		}

		if 0 >= len(binanceOrder) {
			fmt.Println("更新平单后收益数据失败, 空数据", binanceOrder, userOrder.ID)
			return
		}

		// 收益
		for _, vBinanceOrder := range binanceOrder {
			var (
				tmpIncome float64
			)
			tmpIncome, err = strconv.ParseFloat(vBinanceOrder.RealizedPnl, 64)
			if nil != err {
				fmt.Println("更新平单后收益数据失败", err, vBinanceOrder)
				return
			}

			income += tmpIncome
		}

		// 写入
		if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
			_, err = b.binanceUserRepo.UpdatesUserOrderTwoHandleStatus(ctx, userOrder.ID)
			if nil != err {
				return err
			}

			_, err = b.binanceUserRepo.UpdatesUserAmountTwo(ctx, userOrder.UserId, int64(income*100000))
			if nil != err {
				return err
			}

			_, err = b.binanceUserRepo.InsertUserAmountRecordTwo(ctx, &UserAmountRecord{
				UserId:    userOrder.UserId,
				OrderId:   userOrder.ID,
				Amount:    int64(income * 100000),
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			})
			if nil != err {
				return err
			}

			return nil
		}); err != nil {
			fmt.Println("更新平单后收益数据失败", err, userOrder.ID)
			return
		}
	}

	err = b.binanceUserRepo.SRemOrderSetSellLongOrBuyShortTwo(ctx, int64(userOrder.ID))
	if nil != err {
		fmt.Println("更新平单后收益数据失败, 错误的数据删除, redis，order_id:", int64(userOrder.ID), userOrder.ID, err)
		return
	}

	return
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

func requestBinanceOrderHistory(apiKey string, secretKey string, symbol string, orderId int64, startTime string, endTime string) ([]*OrderHistory, error) {
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
	if 0 < len(startTime) && 0 < len(endTime) {
		data = "symbol=" + symbol + "&startTime=" + startTime + "&endTime=" + endTime + "&limit=1000&timestamp=" + now
	} else if 0 < orderId {
		data = "symbol=" + symbol + "&orderId=" + strconv.FormatInt(orderId, 10) + "&timestamp=" + now
	} else {
		return nil, nil
	}

	// 加密
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	signature := hex.EncodeToString(h.Sum(nil))
	// 构造请求

	req, err = http.NewRequest("GET", apiUrl+"?"+data+"&signature="+signature, nil)
	if err != nil {
		return nil, err
	}

	// 添加头信息
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
		fmt.Println(string(b))
		fmt.Println(err)
		return nil, err
	}

	var i []*OrderHistory
	err = json.Unmarshal(b, &i)
	if err != nil {
		fmt.Println(string(b))
		return nil, err
	}

	res = make([]*OrderHistory, 0)
	for _, v := range i {
		res = append(res, v)
	}

	return res, nil
}
