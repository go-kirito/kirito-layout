/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/8/21 10:05
 */
package ecode

import "github.com/go-kirito/pkg/errors"

var (
	ErrInvalidArgument = errors.BadRequest("100001", "参数错误")
)
