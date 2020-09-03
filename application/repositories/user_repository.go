package repositories

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/rodrigorrch/challengs_platform/domain"
)

type UserRepository interface {
	Insert(user *domain.User) (*domain.User, error)
	Find(email string) (*domain.User, error)
}

type UserRepositoryDb struct {
	Db *gorm.DB
}

func (repo UserRepositoryDb) Insert(user *domain.User) (*domain.User, error) {

	err := user.Prepare()

	if err != nil {
		return user, err
	}

	err = repo.Db.Create(user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (repo UserRepositoryDb) Find(email string) (*domain.User, error) {

	var user domain.User
	repo.Db.First(&user, "email = ?", email)

	if user.Token == "" {
		return nil, fmt.Errorf("User does not exist")
	}

	return &user, nil
}
