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

type LhBinanceUserApiError struct {
	ID     uint64
	UserId uint64
	Msg    string
}

type BinanceUserRepo interface {
	InsertUser(ctx context.Context, lhBinanceUser *LhBinanceUser) (bool, error)
	UpdateUser(ctx context.Context, userId uint64, apiKey string, apiSecret string) (bool, error)
	InsertUserStatus(ctx context.Context, lhBinanceUserStatus *LhBinanceUserStatus) (bool, error)
	UpdatesUserStatus(ctx context.Context, userId uint64, baseMoney float64, status string) (bool, error)
	GetUsers() ([]*LhBinanceUser, error)
	GetUserStatus(userId uint64) (*LhBinanceUserStatus, error)
	GetUserByAddress(address string) (*LhBinanceUser, error)
	GetUserByApiKeyAndApiSecret(key string, secret string) (*LhBinanceUser, error)
	GetUserApiErrByUserId(userId uint64) (*LhBinanceUserApiError, error)
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
		lhBinanceUser  *LhBinanceUser
		lhBinanceUser2 *LhBinanceUser
		err            error
	)

	lhBinanceUser, err = b.binanceUserRepo.GetUserByAddress(address)
	if nil != err {
		return err
	}

	lhBinanceUser2, err = b.binanceUserRepo.GetUserByApiKeyAndApiSecret(apiKey, apiSecret)
	if nil != err {
		return err
	}

	if nil == lhBinanceUser && nil == lhBinanceUser2 { // 地址和api信息都不存在
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
	} else if nil != lhBinanceUser { // 地址存在
		if nil == lhBinanceUser2 || (lhBinanceUser2.ID == lhBinanceUser.ID) { // api不存在 或 地址和api指向相同ID(同一条记录)
			if apiKey != lhBinanceUser.ApiKey || apiSecret != lhBinanceUser.ApiSecret { // api_key和api_secret发生了变化
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

func (b *BinanceUserUsecase) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	var (
		lhBinanceUser       *LhBinanceUser
		lhBinanceUserApiErr *LhBinanceUserApiError
		status              string
		err                 error
	)

	if 0 >= len(req.Address) || 300 < len(req.Address) {
		return &v1.GetUserReply{Status: status}, err
	}

	lhBinanceUser, err = b.binanceUserRepo.GetUserByAddress(req.Address)
	if nil != err {
		return nil, err
	}

	lhBinanceUserApiErr, err = b.binanceUserRepo.GetUserApiErrByUserId(lhBinanceUser.ID)
	if nil != err {
		return nil, err
	}

	if nil == lhBinanceUserApiErr {
		status = "ok"
	}

	return &v1.GetUserReply{Status: status}, err
}
