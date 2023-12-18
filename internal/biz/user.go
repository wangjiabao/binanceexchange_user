package biz

import (
	v1 "binanceexchange_user/api/binanceexchange_user/v1"
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type LhBinanceUser struct {
	ID        uint64
	Address   string
	ApiKey    string
	ApiSecret string
}

type LhBinanceUserStatus struct {
	ID        uint64
	UserId    uint64
	Status    string
	BaseMoney float64
}

type BinanceUserRepo interface {
	InsertUser(ctx context.Context, lhBinanceUser *LhBinanceUser) (bool, error)
	UpdateUser(ctx context.Context, userId uint64, apiKey string, apiSecret string) (bool, error)
	InsertUserStatus(ctx context.Context, lhBinanceUserStatus *LhBinanceUserStatus) (bool, error)
	UpdatesUserStatus(ctx context.Context, userId uint64, baseMoney float64, status string) (bool, error)
	GetUsers() ([]*LhBinanceUser, error)
	GetUserStatus(userId uint64) (*LhBinanceUserStatus, error)
	GetUserByAddress(address string) (*LhBinanceUser, error)
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

func (b *BinanceUserUsecase) SetUser(ctx context.Context, address string, apiKey string, apiSecret string) error {
	var (
		lhBinanceUser *LhBinanceUser
		err           error
	)

	lhBinanceUser, err = b.binanceUserRepo.GetUserByAddress(address)
	if nil != err {
		return err
	}

	if nil == lhBinanceUser {
		if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
			_, err = b.binanceUserRepo.InsertUser(ctx, &LhBinanceUser{
				Address:   address,
				ApiKey:    apiKey,
				ApiSecret: apiSecret,
			})

			if nil != err {
				return err
			}

			return nil
		}); err != nil {
			b.log.Error(err)
		}
	} else if apiKey != lhBinanceUser.ApiKey || apiSecret != lhBinanceUser.ApiSecret {
		if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
			_, err = b.binanceUserRepo.UpdateUser(ctx, lhBinanceUser.ID, apiKey, apiSecret)

			if nil != err {
				return err
			}

			return nil
		}); err != nil {
			b.log.Error(err)
		}
	}

	return nil
}

func (b *BinanceUserUsecase) PullUserStatus(ctx context.Context, req *v1.PullUserStatusRequest) (*v1.PullUserStatusReply, error) {

	return &v1.PullUserStatusReply{}, nil
}

func (b *BinanceUserUsecase) GetUserStatus(userId uint64) (*LhBinanceUserStatus, error) {
	return b.binanceUserRepo.GetUserStatus(userId)
}

func (b *BinanceUserUsecase) InsertUserStatus(ctx context.Context, userId uint64, baseMoney float64) (bool, error) {
	return b.binanceUserRepo.InsertUserStatus(ctx, &LhBinanceUserStatus{
		UserId:    userId,
		BaseMoney: baseMoney,
	})
}

func (b *BinanceUserUsecase) UpdateUserStatusOpen(ctx context.Context, userId uint64, baseMoney float64) (bool, error) {
	return b.binanceUserRepo.UpdatesUserStatus(ctx, userId, baseMoney, "open")
}

func (b *BinanceUserUsecase) UpdateUserStatusClose(ctx context.Context, userId uint64) (bool, error) {
	return b.binanceUserRepo.UpdatesUserStatus(ctx, userId, 0, "close")
}

func (b *BinanceUserUsecase) GetUsers() ([]*LhBinanceUser, error) {
	return b.binanceUserRepo.GetUsers()
}
