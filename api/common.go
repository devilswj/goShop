/*
 * @Author: devilswj
 * @Date: 2021-12-22 19:38:57
 * @LastEditors: devilswj
 * @LastEditTime: 2021-12-22 19:50:22
 * @FilePath: \shop\api\common.go
 */
package api

import (
	"fmt"
	"shop/conf"
	"shop/serializer"

	"github.com/go-playground/validator/v10"
)

func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			field := conf.T(fmt.Sprintf("field.%s", e.Field))
			tag := conf.T(fmt.Sprintf("Tag.Valid.%s", e.Tag))
			return serializer.Response{
				Status: 400001,
				Msg:    fmt.Sprintf("%s%s", field, tag),
				Error:  fmt.Sprint(err),
			}
		}
	}
}
