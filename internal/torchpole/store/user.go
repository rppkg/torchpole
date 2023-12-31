// Copyright 2023 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rppkg/torchpole.

package store

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"github.com/rppkg/torchpole/internal/pkg/model"
)

// IUserStore 定义了 user 模块在 store 层所实现的方法.
type IUserStore interface {
	Create(ctx context.Context, user *model.User) error
	Get(ctx context.Context, username string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	List(ctx context.Context, offset, limit int) (int64, []*model.User, error)
	Delete(ctx context.Context, username string) error
}

// UserStore 接口的实现.
type userStore struct {
	db *gorm.DB
}

// 确保 users 实现了 UserStore 接口.
var _ IUserStore = (*userStore)(nil)

func newUsers(db *gorm.DB) *userStore {
	return &userStore{db}
}

// Create 插入一条 user 记录.
func (u *userStore) Create(ctx context.Context, user *model.User) error {
	return u.db.Create(&user).Error
}

// Get 根据用户名查询指定 user 的数据库记录.
func (u *userStore) Get(ctx context.Context, username string) (*model.User, error) {
	var user model.User
	if err := u.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

// Update 更新一条 user 数据库记录.
func (u *userStore) Update(ctx context.Context, user *model.User) error {
	return u.db.Save(user).Error
}

// List 根据 offset 和 limit 返回 user 列表.
func (u *userStore) List(ctx context.Context, offset, limit int) (count int64, ret []*model.User, err error) {
	err = u.db.Offset(offset).Limit(defaultLimit(limit)).Order("id desc").Find(&ret).
		Offset(-1).
		Limit(-1).
		Count(&count).
		Error

	return
}

// Delete 根据 username 删除数据库 user 记录.
func (u *userStore) Delete(ctx context.Context, username string) error {
	err := u.db.Where("username = ?", username).Delete(&model.User{}).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	return nil
}
