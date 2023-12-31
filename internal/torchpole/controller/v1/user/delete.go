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

// Delete 删除一个用户.
func (userController *UserController) Delete(c *gin.Context) {
	log.Info(c).Msg("Delete user function called")

	username := c.Param("name")

	if err := userController.biz.UserBiz().Delete(c, username); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	if _, err := userController.auth.RemoveNamedPolicy("p", username, "", ""); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
