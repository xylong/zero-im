package dao

import (
	"gorm.io/gorm"
	"zero-im/models"
)

type UserDao interface {
	// CreateUser 创建用户
	CreateUser(user *models.User) error

	// GetByID 根据id查找
	GetByID(id string) (*models.User, error)

	// FindByPhone 根据手机号查询用户
	FindByPhone(phone string) (*models.User, error)

	ListByName(name string) ([]*models.User, error)
	ListByIds(ids []string) ([]*models.User, error)
}

type userDao struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &userDao{db: db}
}

func (d *userDao) CreateUser(user *models.User) error {
	return d.db.Create(user).Error
}

func (d *userDao) GetByID(id string) (*models.User, error) {
	var (
		err  error
		user models.User
	)

	if err = d.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *userDao) FindByPhone(phone string) (*models.User, error) {
	var (
		err  error
		user models.User
	)

	if err = d.db.Where("phone = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (d *userDao) ListByName(name string) ([]*models.User, error) {
	var (
		err   error
		users []*models.User
	)

	if err = d.db.Where("nickname like ?", "%"+name+"%").Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}

func (d *userDao) ListByIds(ids []string) ([]*models.User, error) {
	var (
		err   error
		users []*models.User
	)

	if err = d.db.Where("id in ?", ids).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
