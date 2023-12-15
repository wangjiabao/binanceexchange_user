package biz

import (
	v1 "binanceexchange_user/api/binanceexchange_user/v1"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type BinanceUserRepo interface {
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

func (b *BinanceUserUsecase) SetUser(ctx context.Context, req *v1.SetUserRequest) (*v1.SetUserReply, error) {
	return &v1.SetUserReply{}, nil
}
