// Copyright 2023 Innkeeper dairongpeng <dairongpeng@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/rppkg/torchpole.

package user

import (
	"github.com/gin-gonic/gin"

	"github.com/rppkg/torchpole/internal/pkg/core"
	"github.com/rppkg/torchpole/internal/pkg/log"
)

// Get 获取一个用户的详细信息.
func (userController *UserController) Get(c *gin.Context) {
	log.Info(c).Msg("Get user function called")

	user, err := userController.biz.UserBiz().Get(c, c.Param("name"))
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, user)
}
