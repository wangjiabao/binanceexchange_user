package biz

import (
	v1 "binanceexchange_user/api/binanceexchange_user/v1"
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"math/big"
	"time"
)

type User struct {
	ID        uint64
	Address   string
	ApiKey    string
	ApiSecret string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserBalance struct {
	ID        uint64
	UserId    uint64
	Balance   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserBalanceRecord struct {
	ID        uint64
	UserId    uint64
	Amount    string
	Balance   string
	Tx        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type BinanceUserRepo interface {
	InsertUser(ctx context.Context, User *User) (*User, error)
	UpdateUser(ctx context.Context, userId uint64, apiKey string, apiSecret string) (bool, error)
	InsertUserBalance(ctx context.Context, userBalance *UserBalance) (bool, error)
	InsertUserBalanceRecord(ctx context.Context, userBalance *UserBalanceRecord) (bool, error)
	UpdatesUserBalance(ctx context.Context, userId uint64, balance string) (bool, error)
	GetUsers() ([]*User, error)
	GetUserByAddress(ctx context.Context, address string) (*User, error)
	GetUserByApiKeyAndApiSecret(key string, secret string) (*User, error)
	GetUserBalance(ctx context.Context, userId uint64) (*UserBalance, error)
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

func (b *BinanceUserUsecase) SetUserBalanceAndUser(ctx context.Context, address string, balance string) error {
	var (
		user            *User
		userBalance     *UserBalance
		err             error
		lastUserBalance string // 上次用户余额
	)

	user, err = b.binanceUserRepo.GetUserByAddress(ctx, address)
	if nil != err {
		return err
	}

	// 初始化 用户和余额
	if nil == user {
		if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
			user, err = b.binanceUserRepo.InsertUser(ctx, &User{
				Address: address,
			})
			if nil != err {
				return err
			}

			if nil == user {
				return errors.New(500, "CREATE_USER_ERROR", "未发现创建的用户")
			}

			_, err = b.binanceUserRepo.InsertUserBalance(ctx, &UserBalance{
				UserId:  user.ID,
				Balance: balance,
			})
			if nil != err {
				return err
			}

			_, err = b.binanceUserRepo.InsertUserBalanceRecord(ctx, &UserBalanceRecord{
				UserId:  user.ID,
				Amount:  balance,
				Balance: lastUserBalance,
			})
			if nil != err {
				return err
			}

			return nil
		}); err != nil {
			b.log.Error(err)
		}

	} else {
		userBalance, err = b.binanceUserRepo.GetUserBalance(ctx, user.ID)
		if nil != err {
			return err
		}

		// 余额未变化
		if userBalance.Balance == balance {
			return nil
		}

		// 上次余额
		tmpLastUserBalance := new(big.Int)
		tmpLastUserBalance.SetString(userBalance.Balance, 10)
		// 本次余额
		tmpCurrentUserBalance := new(big.Int)
		tmpCurrentUserBalance.SetString(balance, 10)
		// 增长
		tmpAmount := new(big.Int)
		tmpAmount.Sub(tmpCurrentUserBalance, tmpCurrentUserBalance)

		// 更新余额
		if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
			_, err = b.binanceUserRepo.UpdatesUserBalance(ctx, user.ID, balance)
			if nil != err {
				return err
			}

			_, err = b.binanceUserRepo.InsertUserBalanceRecord(ctx, &UserBalanceRecord{
				UserId:  user.ID,
				Amount:  tmpAmount.String(),
				Balance: userBalance.Balance,
			})
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

func (b *BinanceUserUsecase) UpdateUser(ctx context.Context, user *User, apiKey string, apiSecret string) error {
	var (
		user2 *User
		err   error
	)

	user2, err = b.binanceUserRepo.GetUserByApiKeyAndApiSecret(apiKey, apiSecret)
	if nil != err {
		return err
	}

	if nil == user2 || (user2.ID == user.ID) { // api不存在 或 地址和api指向相同ID(同一条记录)
		if apiKey != user.ApiKey || apiSecret != user.ApiSecret { // api_key或api_secret发生了变化
			if err = b.tx.ExecTx(ctx, func(ctx context.Context) error {
				_, err = b.binanceUserRepo.UpdateUser(ctx, user.ID, apiKey, apiSecret)
				if nil != err {
					return err
				}

				return nil
			}); err != nil {
				b.log.Error(err)
			}
		}
	}

	return nil
}

func (b *BinanceUserUsecase) GetUsers() ([]*User, error) {
	return b.binanceUserRepo.GetUsers()
}

func (b *BinanceUserUsecase) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	var (
		//user   *User
		status string
		err    error
	)

	if 0 >= len(req.Address) || 300 < len(req.Address) {
		return &v1.GetUserReply{Status: status}, err
	}

	//user, err = b.binanceUserRepo.GetUserByAddress(ctx, req.Address)
	//if nil != err {
	//	return nil, err
	//}

	//lhBinanceUserApiErr, err = b.binanceUserRepo.GetUserApiErrByUserId(lhBinanceUser.ID)
	//if nil != err {
	//	return nil, err
	//}
	//
	//if nil == lhBinanceUserApiErr {
	//	status = "ok"
	//}

	return &v1.GetUserReply{Status: status}, err
}
