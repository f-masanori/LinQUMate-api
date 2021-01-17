package usecase

import (
	"linqumate/domain/repositoryinterface"
)

// UserUsecase : Userのユースケースで使われるrepositoryを定義
// domain/repositoryinterfaceに依存させる
type UserUsecase struct {
	UserRepository *repositoryinterface.UserRepository
}

func NewUserUsecase(UserRepo *repositoryinterface.UserRepository) *UserUsecase {
	return &UserUsecase{
		UserRepository: UserRepo,
	}
}

// Create :uid,emailを指定してユーザーを作成する
func (u *UserUsecase) Create(uid string, email string) (int, error) {
	return 0, nil
}
