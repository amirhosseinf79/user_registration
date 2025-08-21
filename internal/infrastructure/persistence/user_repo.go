package persistence

import (
	"github.com/amirhosseinf79/user_registration/internal/domain/model"
	"github.com/amirhosseinf79/user_registration/internal/domain/repository"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
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

func (r *userRepository) GetByMobile(mobile string) (user *model.User, err error) {
	err = r.db.Where("phone_number = ?", mobile).First(&user).Error
	if err != nil {
		return nil, err
	}
	return
}

func (r *userRepository) CheckMobileExists(mobile string) (exists bool, err error) {
	var count int64
	err = r.db.Model(&model.User{}).Where("phone_number = ?", mobile).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

// func (r *userRepository) GetAllByFilter(filter dto.UserFilter) (users []*model.User, total int64, err error) {
// 	query := r.db.Model(&model.User{})
// 	if filter.Email != "" {
// 		query = query.Where("email = ?", filter.Email)
// 	}
// 	err = query.Count(&total).Find(&users).Error
// 	return
// }
