// Copyright 2023 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rppkg/torchpole.

package biz

import (
	"github.com/rppkg/torchpole/internal/torchpole/biz/user"
	"github.com/rppkg/torchpole/internal/torchpole/store"
)

// IBiz 定义了 Biz 层需要实现的方法.
type IBiz interface {
	UserBiz() user.IUserBiz
}

// 确保 biz 实现了 IBiz 接口.
var _ IBiz = (*biz)(nil)

// biz 是 IBiz 的一个具体实现.
type biz struct {
	ds store.IStore
}

// NewBiz 创建一个 IBiz 类型的实例.
func NewBiz(s store.IStore) *biz {
	return &biz{ds: s}
}

// UserBiz 返回一个实现了 UserBiz 接口的实例.
func (b *biz) UserBiz() user.IUserBiz {
	return user.New(b.ds)
}
