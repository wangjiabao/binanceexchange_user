package biz

import (
	v1 "binanceexchange_user/api/binanceexchange_user/v1"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type LhBinanceUser struct {
	ID        int64
	Address   string
	ApiKey    string
	ApiSecret string
}

type LhBinanceUserStatus struct {
	ID        int64
	UserId    int64
	Status    string
	BaseMoney float64
}

type BinanceUserRepo interface {
	InsertUser(ctx context.Context, lhBinanceUser *LhBinanceUser) (bool, error)
	InsertUserStatus(ctx context.Context, lhBinanceUserStatus *LhBinanceUserStatus) (bool, error)
	UpdatesUserStatus(ctx context.Context, userId int64, baseMoney float64, status string) (bool, error)
	GetUsers() ([]*LhBinanceUser, error)
	GetUserStatus(userId int64) (*LhBinanceUserStatus, error)
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
	var (
		err error
	)

	if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
		_, err = b.binanceUserRepo.InsertUser(ctx, &LhBinanceUser{
			Address:   req.SendBody.Address,
			ApiKey:    req.SendBody.Apikey,
			ApiSecret: req.SendBody.Apisecret,
		})

		if nil != err {
			return err
		}

		return nil
	}); err != nil {
		b.log.Error(err)
	}

	return &v1.SetUserReply{}, nil
}

func (b *BinanceUserUsecase) PullUserStatus(ctx context.Context, req *v1.PullUserStatusRequest) (*v1.PullUserStatusReply, error) {

	return &v1.PullUserStatusReply{}, nil
}

func (b *BinanceUserUsecase) GetUserStatus(userId int64) (*LhBinanceUserStatus, error) {
	return b.binanceUserRepo.GetUserStatus(userId)
}

func (b *BinanceUserUsecase) InsertUserStatus(ctx context.Context, userId int64, baseMoney float64) (bool, error) {
	return b.binanceUserRepo.InsertUserStatus(ctx, &LhBinanceUserStatus{
		UserId:    userId,
		BaseMoney: baseMoney,
	})
}

func (b *BinanceUserUsecase) UpdateUserStatusOpen(ctx context.Context, userId int64, baseMoney float64) (bool, error) {
	return b.binanceUserRepo.UpdatesUserStatus(ctx, userId, baseMoney, "open")
}

func (b *BinanceUserUsecase) UpdateUserStatusClose(ctx context.Context, userId int64) (bool, error) {
	return b.binanceUserRepo.UpdatesUserStatus(ctx, userId, 0, "close")
}

func (b *BinanceUserUsecase) GetUsers() ([]*LhBinanceUser, error) {
	return b.binanceUserRepo.GetUsers()
}
