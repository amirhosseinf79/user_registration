package persistence

import (
	"errors"
	"strings"

	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
	"github.com/amirhosseinf79/user_registration/internal/dto/shared"
	"github.com/amirhosseinf79/user_registration/internal/dto/user"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *model.User) error {
	uniqueErr := "UNIQUE constraint failed: "
	err := r.db.Create(user).Error
	if strings.Contains(err.Error(), uniqueErr) {
		keyName := strings.ReplaceAll(err.Error(), uniqueErr, "")
		errKeyName := errors.New(keyName)
		return errors.Join(errKeyName, shared.ErrAlreadyExists)
	}
	return err
}

func (r *userRepository) Update(user *model.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&model.User{}, id).Error
}

func (r *userRepository) GetByID(id uint) (user *model.User, err error) {
	err = r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return
}

func (r *userRepository) GetByEmail(email string) (user *model.User, err error) {
	err = r.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return
}

func (r *userRepository) GetByMobile(mobile string) (user *model.User, err error) {
	err = r.db.Where("phone_number = ?", mobile).First(&user).Error
	if err != nil {
		return nil, err
	}
	return
}

func (r *userRepository) CheckEmailExists(email string) (exists bool, err error) {
	var count int64
	err = r.db.Model(&model.User{}).Where("email = ?", email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *userRepository) CheckMobileExists(mobile string) (exists bool, err error) {
	var count int64
	err = r.db.Model(&model.User{}).Where("phone_number = ?", mobile).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (r *userRepository) GetAllByFilter(filter user.FilterUser) (users []*model.User, total int64, err error) {
	query := r.db.Model(&model.User{})
	if filter.PhoneNumber != "" {
		query = query.Where("phone_number LIKE ?", "%"+filter.PhoneNumber+"%")
	}
	if filter.Email != "" {
		query = query.Where("email LIKE ?", "%"+filter.Email+"%")
	}
	if filter.Name != "" {
		query = query.Where("concat(first_name, ' ', last_name) like ?", "%"+filter.Name+"%")
	}
	err = query.Count(&total).Error
	if err != nil {
		return
	}

	if filter.Page < 1 {
		filter.Page = 1
	}
	if filter.PageSize <= 0 || filter.PageSize > 100 {
		filter.PageSize = 10
	}

	err = query.Offset((filter.Page - 1) * filter.PageSize).Limit(filter.PageSize).Find(&users).Error
	return
}
