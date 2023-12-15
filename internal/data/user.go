package data

import (
	"binanceexchange_user/internal/biz"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type LhBinanceUser struct {
	ID        int64  `gorm:"primarykey;type:int"`
	Address   string `gorm:"type:varchar(200)"`
	ApiKey    string `gorm:"type:varchar(200)"`
	ApiSecret string `gorm:"type:varchar(200)"`
}

type LhBinanceUserStatus struct {
	ID        int64   `gorm:"primarykey;type:int"`
	UserId    int64   `gorm:"type:int;not null"`
	Status    string  `gorm:"type:varchar(200)"`
	BaseMoney float64 `gorm:"type:decimal(65,20);not null"`
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
func (b *BinanceUserRepo) UpdatesUserStatus(ctx context.Context, userId int64, baseMoney float64, status string) (bool, error) {
	var err error

	if err = b.data.DB(ctx).Table("lh_binance_user_status").Where("user_id=?", userId).
		Updates(map[string]interface{}{"status": status, "base_money": baseMoney}).Error; nil != err {
		return false, errors.NotFound("CLOSE_ORDER_POLICY_MACD_COMPARE_INFO_ERROR", "UPDATE_BINANCE_USER_STATUS_ERROR")
	}

	return true, nil
}

// GetUsers .
func (b *BinanceUserRepo) GetUsers() ([]*biz.LhBinanceUser, error) {
	var lhBinanceUser []*LhBinanceUser
	if err := b.data.db.Table("lh_binance_user").Find(&lhBinanceUser).Error; err != nil {
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
func (b *BinanceUserRepo) GetUserStatus(userId int64) (*biz.LhBinanceUserStatus, error) {
	var lhBinanceUserStatus *LhBinanceUserStatus
	if err := b.data.db.Table("lh_binance_user_status").Where("user_id>=?", userId).First(&lhBinanceUserStatus).Error; err != nil {
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
