package str

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/url"
	"strings"
)

// 将一个字符串md5加密
// @param str 要加密的字符串
func Md5(str string) string {
	m := md5.New()
	io.WriteString(m, str)
	return strings.ToLower(fmt.Sprintf("%x", m.Sum(nil)))
}

// 将一个字符串basa64解码
// @param str 要加密的字符串
func Base64Decode(str string) string {
	bytes, _ := base64.StdEncoding.DecodeString(str)
	return string(bytes)
}

// 将一个字符串basa64编码
// @param str 要加密的字符串
func Base64Encode(str string) string {
	bytes := base64.StdEncoding.EncodeToString([]byte(str))
	return bytes
}

// 将一个字符串url解码
// @param str 需要解码的字符串
func UrlDecode(str string) string {
	val, _ := url.QueryUnescape(str)
	return val
}

// 将一个字符串url编码
// @param str 需要解码的字符串
func UrlEncode(str string) string {
	val := url.QueryEscape(str)
	return val
}

// json字符串编码
// @param str 需要解码的字符串
func JsonEncode(str interface{}) string {
	bytes, _ := json.Marshal(str)
	return string(bytes)
}
