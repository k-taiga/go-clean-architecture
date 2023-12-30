package usecase

import (
	"context"
	"sluck/model"
	"sluck/repository"
	"sluck/transaction"
)

type UserUsecase interface {
	GetByID(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, user *model.User) (string, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
}

type userUsecase struct {
	r           repository.UserRepository
	mr          repository.MessageRepository
	transaction transaction.Transaction
}

func NewUserUsecase(r repository.UserRepository, mr repository.MessageRepository, transaction transaction.Transaction) UserUsecase {
	return &userUsecase{r, mr, transaction}
}

func (u *userUsecase) GetByID(ctx context.Context, id string) (*model.User, error) {
	return nil, nil
}

func (u *userUsecase) Create(ctx context.Context, user *model.User) (string, error) {
	id, err := u.r.Create(ctx, user)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (u *userUsecase) Update(ctx context.Context, user *model.User) error {
	err := u.r.Update(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) Delete(ctx context.Context, id string) error {
	// トランザクション内で実行する処理を関数として渡す
	u.transaction.DoInTx(ctx, func(ctx context.Context) (any, error) {
		// funcで渡す関数
		err := u.r.Delete(ctx, id)
		if err != nil {
			return nil, err
		}

		err = u.mr.Delete(ctx, id)
		if err != nil {
			return nil, err
		}

		return nil, nil
	})

	return nil
}
