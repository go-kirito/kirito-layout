/**
 * @Author : nopsky
 * @Email : cnnopsky@gmail.com
 * @Date : 2021/8/21 10:05
 */
package ecode

import "github.com/go-kirito/pkg/errors"

var (
	ErrNotFound = errors.NotFound("100001", "用户不存在")
)
