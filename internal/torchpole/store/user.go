package store

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/combizent/torchpole/internal/pkg/model"
)

// UserStore 定义了 user 模块在 store 层所实现的方法.
type UserStore interface {
	Create(ctx context.Context, user *model.User) error
	Get(ctx context.Context, username string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	List(ctx context.Context, offset, limit int) (int64, []*model.User, error)
	Delete(ctx context.Context, username string) error
}

// UserStore 接口的实现.
type users struct {
	db *gorm.DB
}

// 确保 users 实现了 UserStore 接口.
var _ UserStore = (*users)(nil)

func newUsers(db *gorm.DB) *users {
	return &users{db}
}

// Create 插入一条 user 记录.
func (u *users) Create(ctx context.Context, user *model.User) error {
	return u.db.Create(&user).Error
}

// Get 根据用户名查询指定 user 的数据库记录.
func (u *users) Get(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Update 更新一条 user 数据库记录.
func (u *users) Update(ctx context.Context, user *model.User) error {
	return u.db.Save(user).Error
}

// List 根据 offset 和 limit 返回 user 列表.
func (u *users) List(ctx context.Context, offset, limit int) (count int64, ret []*model.User, err error) {
	err = u.db.Offset(offset).Limit(defaultLimit(limit)).Order("id desc").Find(&ret).
		Offset(-1).
		Limit(-1).
		Count(&count).
		Error

	return
}

// Delete 根据 username 删除数据库 user 记录.
func (u *users) Delete(ctx context.Context, username string) error {
	err := u.db.Where("username = ?", username).Delete(&model.User{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return nil
}
