package data

import (
	"binanceexchange_user/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID               uint64    `gorm:"primarykey;type:int"`
	Address          string    `gorm:"type:varchar(100);not null"`
	PlayType         uint64    `gorm:"type:int;not null"`
	ApiStatus        uint64    `gorm:"type:int;not null"`
	ApiKey           string    `gorm:"type:varchar(200);not null"`
	ApiSecret        string    `gorm:"type:varchar(200);not null"`
	BindTraderStatus uint64    `gorm:"type:int;not null"`
	CreatedAt        time.Time `gorm:"type:datetime;not null"`
	UpdatedAt        time.Time `gorm:"type:datetime;not null"`
}

type UserBalance struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null"`
	Balance   string    `gorm:"type:varchar(100);not null"`
	Cost      uint64    `gorm:"type:bigint(20);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserBalanceRecord struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null"`
	Amount    string    `gorm:"type:varchar(100);not null"`
	Balance   string    `gorm:"type:varchar(100);not null"`
	tx        string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserAmount struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null"`
	Amount    int64     `gorm:"type:bigint(20);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserAmountRecord struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null"`
	OrderId   uint64    `gorm:"type:bigint(20);not null"`
	Amount    int64     `gorm:"type:bigint(20);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type Trader struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	IsOpen    uint64    `gorm:"type:int;not null"`
	Amount    uint64    `gorm:"type:bigint(20);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type LhCoinSymbol struct {
	ID                uint64 `gorm:"primarykey;type:int"`
	Symbol            string `gorm:"type:varchar(100);not null"`
	QuantityPrecision int64  `gorm:"type:int;not null"`
}

type UserBindTrader struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null"`
	TraderId  uint64    `gorm:"type:int;not null"`
	Amount    uint64    `gorm:"type:bigint(20);not null"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserOrder struct {
	ID            uint64  `gorm:"primarykey;type:int"`
	UserId        uint64  `gorm:"type:int;not null"`
	TraderId      uint64  `gorm:"type:int;not null"`
	ClientOrderId string  `gorm:"type:varchar(100)"`
	OrderId       string  `gorm:"type:varchar(100);not null"`
	Symbol        string  `gorm:"type:varchar(100);not null"`
	Side          string  `gorm:"type:varchar(100);not null"`
	PositionSide  string  `gorm:"type:varchar(100);not null"`
	Quantity      float64 `gorm:"type:decimal(65,20);not null"`
	Price         float64 `gorm:"type:decimal(65,20);not null"`
	TraderQty     float64 `gorm:"type:decimal(65,20);not null"`
	OrderType     string  `gorm:"type:varchar(100);not null"`
	ClosePosition string  `gorm:"type:varchar(100);not null"`
	CumQuote      float64 `gorm:"type:decimal(65,20);not null"`
	ExecutedQty   float64 `gorm:"type:decimal(65,20);not null"`
	AvgPrice      float64 `gorm:"type:decimal(65,20);not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type BinanceUserRepo struct {
	data *Data
	log  *log.Helper
}

func NewBinanceUserRepo(data *Data, logger log.Logger) biz.BinanceUserRepo {
	return &BinanceUserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// InsertUser .
func (b *BinanceUserRepo) InsertUser(ctx context.Context, user *biz.User) (*biz.User, error) {
	insertUser := &User{
		Address:  user.Address,
		PlayType: user.PlayType,
	}

	res := b.data.DB(ctx).Table("user").Create(&insertUser)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_ERROR", "创建数据失败")
	}

	return &biz.User{
		ID:        insertUser.ID,
		Address:   insertUser.Address,
		ApiKey:    insertUser.ApiKey,
		ApiSecret: insertUser.ApiSecret,
		CreatedAt: insertUser.CreatedAt,
		UpdatedAt: insertUser.UpdatedAt,
	}, nil
}

// UpdateUser .
func (b *BinanceUserRepo) UpdateUser(ctx context.Context, userId uint64, apiKey string, apiSecret string) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{"api_status": 1, "api_key": apiKey, "api_secret": apiSecret, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_ERROR", "UPDATE_USER_ERROR")
	}

	return true, nil
}

// UpdateUserBindTraderStatus .
func (b *BinanceUserRepo) UpdateUserBindTraderStatus(ctx context.Context, userId uint64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("user").Where("id=?", userId).
		Updates(map[string]interface{}{"bind_trader_status": 1, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_ERROR", "UPDATE_USER_ERROR")
	}

	return true, nil
}

// InsertUserBalance .
func (b *BinanceUserRepo) InsertUserBalance(ctx context.Context, userBalance *biz.UserBalance) (bool, error) {
	insertUserBalance := &UserBalance{
		UserId:  userBalance.UserId,
		Balance: userBalance.Balance,
		Cost:    userBalance.Cost,
	}

	res := b.data.DB(ctx).Table("user_balance").Create(&insertUserBalance)
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_BALANCE_ERROR", "创建数据失败")
	}

	return true, nil
}

// UpdatesUserBalance .
func (b *BinanceUserRepo) UpdatesUserBalance(ctx context.Context, userId uint64, balance string, cost uint64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("user_balance").Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance": balance, "cost": cost, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_BALANCE_ERROR", "UPDATE_USER_BALANCE_ERROR")
	}

	return true, nil
}

// InsertUserBalanceRecord .
func (b *BinanceUserRepo) InsertUserBalanceRecord(ctx context.Context, userBalanceRecord *biz.UserBalanceRecord) (bool, error) {
	insertUserBalanceRecord := &UserBalanceRecord{
		UserId:  userBalanceRecord.UserId,
		Amount:  userBalanceRecord.Amount,
		Balance: userBalanceRecord.Balance,
		tx:      userBalanceRecord.Tx,
	}

	res := b.data.DB(ctx).Table("user_balance_record").Create(&insertUserBalanceRecord)
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_BALANCE_RECORD_ERROR", "创建数据失败")
	}

	return true, nil
}

// InsertUserAmount .
func (b *BinanceUserRepo) InsertUserAmount(ctx context.Context, userAmount *biz.UserAmount) (bool, error) {
	insertUserAmount := &UserAmount{
		UserId: userAmount.UserId,
		Amount: 0,
	}

	res := b.data.DB(ctx).Table("user_amount").Create(&insertUserAmount)
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_AMOUNT_ERROR", "创建数据失败")
	}

	return true, nil
}

// UpdatesUserAmount .
func (b *BinanceUserRepo) UpdatesUserAmount(ctx context.Context, userId uint64, amount int64) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("user_amount").Where("user_id=?", userId).
		Updates(map[string]interface{}{"amount": gorm.Expr("amount + ?", amount), "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_AMOUNT_ERROR", "UPDATE_USER_AMOUNT_ERROR")
	}

	return true, nil
}

// InsertUserAmountRecord .
func (b *BinanceUserRepo) InsertUserAmountRecord(ctx context.Context, userAmountRecord *biz.UserAmountRecord) (bool, error) {
	insertUserAmountRecord := &UserAmountRecord{
		UserId:  userAmountRecord.UserId,
		OrderId: userAmountRecord.OrderId,
		Amount:  userAmountRecord.Amount,
	}

	res := b.data.DB(ctx).Table("user_amount_record").Create(&insertUserAmountRecord)
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_AMOUNT_RECORD_ERROR", "创建数据失败")
	}

	return true, nil
}

// InsertUserBindTrader .
func (b *BinanceUserRepo) InsertUserBindTrader(ctx context.Context, userId uint64, traderId uint64, amount uint64) (*biz.UserBindTrader, error) {
	insertUserBindTrader := &UserBindTrader{
		UserId:   userId,
		TraderId: traderId,
		Amount:   amount,
	}

	res := b.data.DB(ctx).Table("user_bind_trader").Create(&insertUserBindTrader)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_BIND_TRADER_ERROR", "创建数据失败")
	}

	return &biz.UserBindTrader{
		ID:        insertUserBindTrader.ID,
		UserId:    insertUserBindTrader.UserId,
		TraderId:  insertUserBindTrader.TraderId,
		Amount:    insertUserBindTrader.Amount,
		CreatedAt: insertUserBindTrader.CreatedAt,
		UpdatedAt: insertUserBindTrader.UpdatedAt,
	}, nil
}

// InsertUserOrder .
func (b *BinanceUserRepo) InsertUserOrder(ctx context.Context, order *biz.UserOrder) (*biz.UserOrder, error) {
	insertUserOrder := &UserOrder{
		UserId:        order.UserId,
		TraderId:      order.TraderId,
		ClientOrderId: order.ClientOrderId,
		OrderId:       order.OrderId,
		Symbol:        order.Symbol,
		Side:          order.Side,
		PositionSide:  order.PositionSide,
		Quantity:      order.Quantity,
		Price:         order.Price,
		TraderQty:     order.TraderQty,
		OrderType:     order.OrderType,
		ClosePosition: order.ClosePosition,
		CumQuote:      order.CumQuote,
		ExecutedQty:   order.ExecutedQty,
		AvgPrice:      order.AvgPrice,
	}

	res := b.data.DB(ctx).Table("user_order").Create(&insertUserOrder)
	if res.Error != nil {
		return nil, errors.New(500, "CREATE_USER_ORDER_ERROR", "创建数据失败")
	}

	return &biz.UserOrder{
		ID:            insertUserOrder.ID,
		UserId:        insertUserOrder.UserId,
		TraderId:      insertUserOrder.TraderId,
		ClientOrderId: insertUserOrder.ClientOrderId,
		OrderId:       insertUserOrder.OrderId,
		Symbol:        insertUserOrder.Symbol,
		Side:          insertUserOrder.Side,
		PositionSide:  insertUserOrder.PositionSide,
		Quantity:      insertUserOrder.Quantity,
		Price:         insertUserOrder.Price,
		TraderQty:     insertUserOrder.TraderQty,
		OrderType:     insertUserOrder.OrderType,
		ClosePosition: insertUserOrder.ClosePosition,
		CumQuote:      insertUserOrder.CumQuote,
		ExecutedQty:   insertUserOrder.ExecutedQty,
		AvgPrice:      insertUserOrder.AvgPrice,
		CreatedAt:     insertUserOrder.CreatedAt,
		UpdatedAt:     insertUserOrder.UpdatedAt,
	}, nil
}

// GetUsers .
func (b *BinanceUserRepo) GetUsers() ([]*biz.User, error) {
	var users []*User
	if err := b.data.db.Table("user").Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ERROR", err.Error())
	}

	res := make([]*biz.User, 0)
	for _, v := range users {
		res = append(res, &biz.User{
			ID:               v.ID,
			PlayType:         v.PlayType,
			Address:          v.Address,
			ApiKey:           v.ApiKey,
			ApiStatus:        v.ApiStatus,
			BindTraderStatus: v.BindTraderStatus,
			ApiSecret:        v.ApiSecret,
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUsersByUserIds .
func (b *BinanceUserRepo) GetUsersByUserIds(userIds []uint64) (map[uint64]*biz.User, error) {
	var users []*User
	if err := b.data.db.Table("user").Where("id in(?)", userIds).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ERROR", err.Error())
	}

	res := make(map[uint64]*biz.User, 0)
	for _, v := range users {
		res[v.ID] = &biz.User{
			ID:               v.ID,
			Address:          v.Address,
			ApiKey:           v.ApiKey,
			ApiStatus:        v.ApiStatus,
			PlayType:         v.PlayType,
			BindTraderStatus: v.BindTraderStatus,
			ApiSecret:        v.ApiSecret,
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
		}
	}

	return res, nil
}

// GetUsersByBindUserStatus .
func (b *BinanceUserRepo) GetUsersByBindUserStatus() ([]*biz.User, error) {
	var users []*User
	if err := b.data.db.Table("user").Where("bind_trader_status=?", 0).Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ERROR", err.Error())
	}

	res := make([]*biz.User, 0)
	for _, v := range users {
		res = append(res, &biz.User{
			ID:               v.ID,
			PlayType:         v.PlayType,
			Address:          v.Address,
			ApiStatus:        v.ApiStatus,
			ApiKey:           v.ApiKey,
			BindTraderStatus: v.BindTraderStatus,
			ApiSecret:        v.ApiSecret,
			CreatedAt:        v.CreatedAt,
			UpdatedAt:        v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserByAddress .
func (b *BinanceUserRepo) GetUserByAddress(ctx context.Context, address string) (*biz.User, error) {
	var user *User
	if err := b.data.DB(ctx).Table("user").Where("address=?", address).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ERROR", err.Error())
	}

	return &biz.User{
		ID:               user.ID,
		ApiStatus:        user.ApiStatus,
		PlayType:         user.PlayType,
		Address:          user.Address,
		ApiKey:           user.ApiKey,
		BindTraderStatus: user.BindTraderStatus,
		ApiSecret:        user.ApiSecret,
		CreatedAt:        user.CreatedAt,
		UpdatedAt:        user.UpdatedAt,
	}, nil
}

// GetUserByApiKeyAndApiSecret .
func (b *BinanceUserRepo) GetUserByApiKeyAndApiSecret(key string, secret string) (*biz.User, error) {
	var user *User
	if err := b.data.db.Table("user").Where("api_key=? or api_secret=?", key, secret).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ERROR", err.Error())
	}

	return &biz.User{
		ID:               user.ID,
		Address:          user.Address,
		ApiStatus:        user.ApiStatus,
		ApiKey:           user.ApiKey,
		BindTraderStatus: user.BindTraderStatus,
		ApiSecret:        user.ApiSecret,
		CreatedAt:        user.CreatedAt,
		UpdatedAt:        user.UpdatedAt,
	}, nil
}

// GetUserBalance .
func (b *BinanceUserRepo) GetUserBalance(ctx context.Context, userId uint64) (*biz.UserBalance, error) {
	var userBalance *UserBalance
	if err := b.data.DB(ctx).Table("user_balance").Where("user_id=?", userId).First(&userBalance).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BALANCE_ERROR", err.Error())
	}

	return &biz.UserBalance{
		ID:        userBalance.ID,
		UserId:    userBalance.UserId,
		Balance:   userBalance.Balance,
		Cost:      userBalance.Cost,
		CreatedAt: userBalance.CreatedAt,
		UpdatedAt: userBalance.UpdatedAt,
	}, nil
}

// GetUserAmount .
func (b *BinanceUserRepo) GetUserAmount(ctx context.Context, userId uint64) (*biz.UserAmount, error) {
	var userAmount *UserAmount
	if err := b.data.DB(ctx).Table("user_amount").Where("user_id=?", userId).First(&userAmount).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BALANCE_ERROR", err.Error())
	}

	return &biz.UserAmount{
		ID:        userAmount.ID,
		UserId:    userAmount.UserId,
		Amount:    userAmount.Amount,
		CreatedAt: userAmount.CreatedAt,
		UpdatedAt: userAmount.UpdatedAt,
	}, nil
}

// GetUserBalanceByUserIds .
func (b *BinanceUserRepo) GetUserBalanceByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*biz.UserBalance, error) {
	var userBalance []*UserBalance
	if err := b.data.DB(ctx).Table("user_balance").Where("user_id in(?)", userIds).Find(&userBalance).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BALANCE_ERROR", err.Error())
	}

	res := make(map[uint64]*biz.UserBalance, 0)
	for _, v := range userBalance {
		res[v.UserId] = &biz.UserBalance{
			ID:        v.ID,
			UserId:    v.UserId,
			Balance:   v.Balance,
			Cost:      v.Cost,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	return res, nil
}

// GetUserAmountByUserIds .
func (b *BinanceUserRepo) GetUserAmountByUserIds(ctx context.Context, userIds []uint64) (map[uint64]*biz.UserAmount, error) {
	var userAmount []*UserAmount
	if err := b.data.DB(ctx).Table("user_amount").Where("user_id in(?)", userIds).Find(&userAmount).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_AMOUNT_ERROR", err.Error())
	}

	res := make(map[uint64]*biz.UserAmount, 0)
	for _, v := range userAmount {
		res[v.UserId] = &biz.UserAmount{
			ID:        v.ID,
			UserId:    v.UserId,
			Amount:    v.Amount,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		}
	}

	return res, nil
}

// GetTradersOrderByAmountDesc .
func (b *BinanceUserRepo) GetTradersOrderByAmountDesc() ([]*biz.Trader, error) {
	var traders []*Trader
	if err := b.data.db.Table("trader").Where("is_open=?", 1).Order("amount desc").Find(&traders).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_TRADER_ERROR", err.Error())
	}

	res := make([]*biz.Trader, 0)
	for _, v := range traders {
		res = append(res, &biz.Trader{
			ID:        v.ID,
			Status:    v.IsOpen,
			Amount:    v.Amount,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

// GetTraders .
func (b *BinanceUserRepo) GetTraders() ([]*biz.Trader, error) {
	var traders []*Trader
	if err := b.data.db.Table("trader").Find(&traders).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_TRADER_ERROR", err.Error())
	}

	res := make([]*biz.Trader, 0)
	for _, v := range traders {
		res = append(res, &biz.Trader{
			ID:        v.ID,
			Status:    v.IsOpen,
			Amount:    v.Amount,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserBindTraderByUserId .
func (b *BinanceUserRepo) GetUserBindTraderByUserId(userId uint64) ([]*biz.UserBindTrader, error) {
	var userBindTrader []*UserBindTrader
	if err := b.data.db.Table("user_bind_trader").Where("user_id=?", userId).Find(&userBindTrader).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BIND_TRADER_ERROR", err.Error())
	}

	res := make([]*biz.UserBindTrader, 0)
	for _, v := range userBindTrader {
		res = append(res, &biz.UserBindTrader{
			ID:        v.ID,
			UserId:    v.UserId,
			TraderId:  v.TraderId,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserBindTraderByTraderIds .
func (b *BinanceUserRepo) GetUserBindTraderByTraderIds(traderIds []uint64) (map[uint64][]*biz.UserBindTrader, error) {
	var userBindTrader []*UserBindTrader
	if err := b.data.db.Table("user_bind_trader").Where("trader_id in(?)", traderIds).Find(&userBindTrader).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_BIND_TRADER_ERROR", err.Error())
	}

	res := make(map[uint64][]*biz.UserBindTrader, 0)
	for _, v := range userBindTrader {
		if _, ok := res[v.TraderId]; !ok {
			res[v.TraderId] = make([]*biz.UserBindTrader, 0)
		}

		res[v.TraderId] = append(res[v.TraderId], &biz.UserBindTrader{
			ID:        v.ID,
			UserId:    v.UserId,
			TraderId:  v.TraderId,
			Amount:    v.Amount,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserOrderByUserTraderIdAndSymbol .
func (b *BinanceUserRepo) GetUserOrderByUserTraderIdAndSymbol(userId uint64, traderId uint64, symbol string) ([]*biz.UserOrder, error) {
	var userOrder []*UserOrder
	if err := b.data.db.Table("user_order").
		Where("user_id=? and trader_id=? and symbol=?", userId, traderId, symbol).
		Find(&userOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ORDER_ERROR", err.Error())
	}

	res := make([]*biz.UserOrder, 0)
	for _, v := range userOrder {
		res = append(res, &biz.UserOrder{
			ID:            v.ID,
			UserId:        v.UserId,
			TraderId:      v.TraderId,
			ClientOrderId: v.ClientOrderId,
			OrderId:       v.OrderId,
			Symbol:        v.Symbol,
			Side:          v.Side,
			PositionSide:  v.PositionSide,
			Quantity:      v.Quantity,
			Price:         v.Price,
			TraderQty:     v.TraderQty,
			OrderType:     v.OrderType,
			ClosePosition: v.ClosePosition,
			CumQuote:      v.CumQuote,
			ExecutedQty:   v.ExecutedQty,
			AvgPrice:      v.AvgPrice,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		})
	}

	return res, nil
}

// GetUserOrderByUserIdAndSymbolAndPositionSide .
func (b *BinanceUserRepo) GetUserOrderByUserIdAndSymbolAndPositionSide(userId uint64, symbol string, positionSide string) ([]*biz.UserOrder, error) {
	var userOrder []*UserOrder
	if err := b.data.db.Table("user_order").
		Where("user_id=? and symbol=? and position_side=?", userId, symbol, positionSide).
		Order("id asc").
		Find(&userOrder).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ORDER_ERROR", err.Error())
	}

	res := make([]*biz.UserOrder, 0)
	for _, v := range userOrder {
		res = append(res, &biz.UserOrder{
			ID:            v.ID,
			UserId:        v.UserId,
			TraderId:      v.TraderId,
			ClientOrderId: v.ClientOrderId,
			OrderId:       v.OrderId,
			Symbol:        v.Symbol,
			Side:          v.Side,
			PositionSide:  v.PositionSide,
			Quantity:      v.Quantity,
			Price:         v.Price,
			TraderQty:     v.TraderQty,
			OrderType:     v.OrderType,
			ClosePosition: v.ClosePosition,
			CumQuote:      v.CumQuote,
			ExecutedQty:   v.ExecutedQty,
			AvgPrice:      v.AvgPrice,
			CreatedAt:     v.CreatedAt,
			UpdatedAt:     v.UpdatedAt,
		})
	}

	return res, nil
}

// GetSymbol .
func (b *BinanceUserRepo) GetSymbol() (map[string]*biz.Symbol, error) {
	var lhCoinSymbol []*LhCoinSymbol
	if err := b.data.db.Table("lh_coin_symbol").Find(&lhCoinSymbol).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_LH_COIN_SYMBOL_ERROR", err.Error())
	}

	res := make(map[string]*biz.Symbol, 0)
	for _, v := range lhCoinSymbol {
		res[v.Symbol+"USDT"] = &biz.Symbol{
			ID:                v.ID,
			Symbol:            v.Symbol,
			QuantityPrecision: v.QuantityPrecision,
		}
	}

	return res, nil
}
