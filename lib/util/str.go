package util

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

// 将一个字符串md5加密
// @param str 要加密的字符串
func Md5(str string) string {
	m := md5.New()
	io.WriteString(m, str)
	return strings.ToLower(fmt.Sprintf("%x", m.Sum(nil)))
}


