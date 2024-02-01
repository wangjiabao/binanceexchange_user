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
	ID        uint64    `gorm:"primarykey;type:int"`
	Address   string    `gorm:"type:varchar(100)"`
	ApiKey    string    `gorm:"type:varchar(200)"`
	ApiSecret string    `gorm:"type:varchar(200)"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserBalance struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null"`
	Balance   string    `gorm:"type:varchar(100)"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
}

type UserBalanceRecord struct {
	ID        uint64    `gorm:"primarykey;type:int"`
	UserId    uint64    `gorm:"type:int;not null"`
	Amount    string    `gorm:"type:varchar(100)"`
	Balance   string    `gorm:"type:varchar(100)"`
	tx        string    `gorm:"type:varchar(100)"`
	CreatedAt time.Time `gorm:"type:datetime;not null"`
	UpdatedAt time.Time `gorm:"type:datetime;not null"`
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
		Address: user.Address,
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
		Updates(map[string]interface{}{"api_key": apiKey, "api_secret": apiSecret, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_ERROR", "UPDATE_USER_ERROR")
	}

	return true, nil
}

// InsertUserBalance .
func (b *BinanceUserRepo) InsertUserBalance(ctx context.Context, userBalance *biz.UserBalance) (bool, error) {
	insertUserBalance := &UserBalance{
		UserId:  userBalance.UserId,
		Balance: userBalance.Balance,
	}

	res := b.data.DB(ctx).Table("user_balance").Create(&insertUserBalance)
	if res.Error != nil {
		return false, errors.New(500, "CREATE_USER_BALANCE_ERROR", "创建数据失败")
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

// UpdatesUserBalance .
func (b *BinanceUserRepo) UpdatesUserBalance(ctx context.Context, userId uint64, balance string) (bool, error) {
	var (
		err error
		now = time.Now()
	)

	if err = b.data.DB(ctx).Table("user_balance").Where("user_id=?", userId).
		Updates(map[string]interface{}{"balance": balance, "updated_at": now}).Error; nil != err {
		return false, errors.NotFound("UPDATE_USER_BALANCE_ERROR", "UPDATE_USER_BALANCE_ERROR")
	}

	return true, nil
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
			ID:        v.ID,
			Address:   v.Address,
			ApiKey:    v.ApiKey,
			ApiSecret: v.ApiSecret,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
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
		ID:        user.ID,
		Address:   user.Address,
		ApiKey:    user.ApiKey,
		ApiSecret: user.ApiSecret,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
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
		ID:        user.ID,
		Address:   user.Address,
		ApiKey:    user.ApiKey,
		ApiSecret: user.ApiSecret,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
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
		CreatedAt: userBalance.CreatedAt,
		UpdatedAt: userBalance.UpdatedAt,
	}, nil
}
