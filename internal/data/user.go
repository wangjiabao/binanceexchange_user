package data

import (
	"binanceexchange_user/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type LhBinanceUser struct {
	ID        uint64 `gorm:"primarykey;type:int"`
	Address   string `gorm:"type:varchar(200)"`
	ApiKey    string `gorm:"type:varchar(200)"`
	ApiSecret string `gorm:"type:varchar(200)"`
}

type LhBinanceUserStatus struct {
	ID        uint64  `gorm:"primarykey;type:int"`
	UserId    uint64  `gorm:"type:int;not null"`
	Status    string  `gorm:"type:varchar(200)"`
	BaseMoney float64 `gorm:"type:decimal(65,20);not null"`
}

type LhBinanceUserApiError struct {
	ID     uint64 `gorm:"primarykey;type:int"`
	UserId uint64 `gorm:"type:int;not null"`
	Msg    string `gorm:"type:varchar(200)"`
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
func (b *BinanceUserRepo) InsertUser(ctx context.Context, lhBinanceUser *biz.LhBinanceUser) (bool, error) {
	insertLhBinanceUser := &LhBinanceUser{
		Address:   lhBinanceUser.Address,
		ApiKey:    lhBinanceUser.ApiKey,
		ApiSecret: lhBinanceUser.ApiSecret,
	}

	res := b.data.DB(ctx).Table("lh_binance_user").Create(&insertLhBinanceUser)
	if res.Error != nil {
		return false, errors.New(500, "CREATE_BINANCE_USER_ERROR", "创建数据失败")
	}

	return true, nil
}

// UpdateUser .
func (b *BinanceUserRepo) UpdateUser(ctx context.Context, userId uint64, apiKey string, apiSecret string) (bool, error) {
	var err error

	if err = b.data.DB(ctx).Table("lh_binance_user").Where("id=?", userId).
		Updates(map[string]interface{}{"api_key": apiKey, "api_secret": apiSecret}).Error; nil != err {
		return false, errors.NotFound("", "UPDATE_BINANCE_USER_STATUS_ERROR")
	}

	return true, nil
}

// InsertUserStatus .
func (b *BinanceUserRepo) InsertUserStatus(ctx context.Context, lhBinanceUserStatus *biz.LhBinanceUserStatus) (bool, error) {
	insertLhBinanceUserStatus := &LhBinanceUserStatus{
		UserId:    lhBinanceUserStatus.UserId,
		Status:    "open",
		BaseMoney: lhBinanceUserStatus.BaseMoney,
	}

	res := b.data.DB(ctx).Table("lh_binance_user_status").Create(&insertLhBinanceUserStatus)
	if res.Error != nil {
		return false, errors.New(500, "CREATE_BINANCE_USER_STATUS_ERROR", "创建数据失败")
	}

	return true, nil
}

// UpdatesUserStatus .
func (b *BinanceUserRepo) UpdatesUserStatus(ctx context.Context, userId uint64, baseMoney float64, status string) (bool, error) {
	var err error

	if err = b.data.DB(ctx).Table("lh_binance_user_status").Where("user_id=?", userId).
		Updates(map[string]interface{}{"status": status, "base_money": baseMoney}).Error; nil != err {
		return false, errors.NotFound("UPDATE_BINANCE_USER_STATUS_ERROR", "UPDATE_BINANCE_USER_STATUS_ERROR")
	}

	return true, nil
}

// GetUsers .
func (b *BinanceUserRepo) GetUsers() ([]*biz.LhBinanceUser, error) {
	var lhBinanceUser []*LhBinanceUser
	if err := b.data.db.Table("lh_binance_user").Where("status<=?", 1).Find(&lhBinanceUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ERROR", err.Error())
	}

	res := make([]*biz.LhBinanceUser, 0)
	for _, v := range lhBinanceUser {
		res = append(res, &biz.LhBinanceUser{
			ID:        v.ID,
			Address:   v.Address,
			ApiKey:    v.ApiKey,
			ApiSecret: v.ApiSecret,
		})
	}

	return res, nil
}

// GetUserStatus .
func (b *BinanceUserRepo) GetUserStatus(userId uint64) (*biz.LhBinanceUserStatus, error) {
	var lhBinanceUserStatus *LhBinanceUserStatus
	if err := b.data.db.Table("lh_binance_user_status").Where("user_id=?", userId).First(&lhBinanceUserStatus).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_STATUS_ERROR", err.Error())
	}

	return &biz.LhBinanceUserStatus{
		ID:        lhBinanceUserStatus.ID,
		UserId:    lhBinanceUserStatus.UserId,
		Status:    lhBinanceUserStatus.Status,
		BaseMoney: lhBinanceUserStatus.BaseMoney,
	}, nil
}

// GetUserByAddress .
func (b *BinanceUserRepo) GetUserByAddress(address string) (*biz.LhBinanceUser, error) {
	var lhBinanceUser *LhBinanceUser
	if err := b.data.db.Table("lh_binance_user").Where("address=?", address).First(&lhBinanceUser).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_ERROR", err.Error())
	}

	return &biz.LhBinanceUser{
		ID:        lhBinanceUser.ID,
		Address:   lhBinanceUser.Address,
		ApiKey:    lhBinanceUser.ApiKey,
		ApiSecret: lhBinanceUser.ApiSecret,
	}, nil
}

// GetUserApiErrByUserId .
func (b *BinanceUserRepo) GetUserApiErrByUserId(userId uint64) (*biz.LhBinanceUserApiError, error) {
	var lhBinanceUserApiError *LhBinanceUserApiError
	if err := b.data.db.Table("lh_binance_user_api_err").Where("user_id=?", userId).First(&lhBinanceUserApiError).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, errors.New(500, "FIND_USER_STATUS_ERROR", err.Error())
	}

	return &biz.LhBinanceUserApiError{
		ID:     lhBinanceUserApiError.ID,
		UserId: lhBinanceUserApiError.UserId,
		Msg:    lhBinanceUserApiError.Msg,
	}, nil
}
