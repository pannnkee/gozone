package validate

// 数据验证类
type validate struct {
}

// 检查是否为空字符串
// @param str 要检查的字符串
func (this *validate) CheckString(str string) bool {
	if str == "" {
		return false
	}
	return true
}

// 检查是否为整型字符串
// @param str 要检查的字符串
func (this *validate) CheckInt(str string) bool {
	return RegexpInt.MatchString(str)
}

// 检查是否为一组数字
func (this *validate) CheckIntSet(str string) bool {
	return RegexpNumberSet.MatchString(str)
}

// 检查一个字段是否在一个列表中
// @param val 要检查的字符串
// @param slice 全部种类
func (this *validate) CheckArray(val string, slice []string) bool {
	for _, value := range slice {
		if val == value {
			return true
		}
	}
	return false
}

// 检查是否为浮点型字符串
func (this *validate) CheckFloat(str string) bool {
	return RegexpFloat.MatchString(str)
}

// 检查是否为金额字符串
func (this *validate) CheckMoney(str string) bool {
	return RegexpMoney.MatchString(str)
}

// 检查是否为链接
// @param str 要检查的字符串
func (this *validate) CheckUrl(str string) bool {
	return RegexpUrl.MatchString(str)
}

// 检查是否为十六进制颜色代码
func (this *validate) CheckColorCode(str string) bool {
	return RegexpColorCode.MatchString(str)
}

// 检查日期
func (this *validate) CheckDate(str string) bool {
	return RegexpDate.MatchString(str)
}

// 检查是否为md5值
// @param str 要检查的值
func (this *validate) CheckMd5(str string) bool {
	return RegexpMd5.MatchString(str)
}

// 检查布尔值
// @return bool
func (this *validate) CheckBool(val string) bool {
	if val == "yes" || val == "no" {
		return true
	}
	return false
}

// 过滤一个参数为最小整数
// @param crt 当前数
// @param min 最小数
func (this *validate) HavingMinNumber(crt int, min int) int {
	if crt <= min {
		crt = min
	}
	return crt
}

// 检查一个中国手机号是否正确
// @param phone 中国大陆手机号
func (this *validate) CheckZhPhone(phone string) bool {
	if RegexpZhPhone.MatchString(phone) {
		return true
	}
	return false
}
