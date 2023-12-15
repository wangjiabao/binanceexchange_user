package service

import (
	v1 "binanceexchange_user/api/binanceexchange_user/v1"
	"binanceexchange_user/internal/biz"
	"context"
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

func (b *BinanceUserService) SetUser(ctx context.Context, req *v1.SetUserRequest) (*v1.SetUserReply, error) {
	return b.buc.SetUser(ctx, req)
}
