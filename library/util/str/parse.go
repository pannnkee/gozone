package str

import (
	"Gozone/library/util/validate"
	"fmt"
	"strconv"
	"strings"
)

// 将一个字符串直接转换成整形
// @param val 要转换的字符串
func ParseInt(val string) int {

	// 判断是否为整形
	if validate.RegexpInt.MatchString(val) == false {
		return 0
	}

	ret, err := strconv.Atoi(val)
	if err != nil {
		panic(err)
	}
	return ret
}

// 将一个字符串转换成int64形式
// @param val 要转换的字符串
func ParseInt64(val string) int64 {

	// 判断是否为整形
	if validate.RegexpInt.MatchString(val) == false {
		return 0
	}

	ret, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		panic(err)
	}
	return ret
}

// 将一个整形转换成浮点型
// @param val 要转换的整形
func ParseStringByInt(val int) string {
	return strconv.Itoa(val)
}

// 将一个整形64转换成浮点型
// @param val 要转换的整形
func ParseStringByInt64(val int64) string {
	return strconv.FormatInt(val, 10)
}

// 将一个浮点型转换成字符串
// @param val 要转换的浮点型
// @param len 保留浮点数长度
func ParseStringByFloat64(val float64, len int) string {
	return strconv.FormatFloat(val, 'f', len, 32)
}

// 将一个字符串转换成浮点型
func ParseFloat(val string) float64 {
	ret, err := strconv.ParseFloat(val, 64)
	if err != nil {
		panic(err)
	}
	return ret
}

// 格式化金额信息
// @param money 要格式化的金额
func FormatPrice(money float64, length ...int) float64 {

	var index = 3
	if len(length) >= 1 {
		index = length[0]
	}

	sMoney := fmt.Sprintf("%.10f", money)
	offset := strings.IndexByte(sMoney, '.')
	if offset != 0 {
		sMoney = sMoney[:offset+index]
	}

	float, _ := strconv.ParseFloat(sMoney, 64)
	return float
}
