package data

import (
	"binanceexchange_user/internal/biz"
	"github.com/go-kratos/kratos/v2/log"
)

type lhBinanceUser struct {
	ID int64 `gorm:"primarykey;type:int"`
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
