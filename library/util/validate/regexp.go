package validate

import "regexp"

var (
	RegexpInt, _       = regexp.Compile("^-?\\d+$")                                                 // 检查整形
	RegexpFloat, _     = regexp.Compile("^-?\\d+(\\.\\d{0,3})?$")                                   // 检查浮点型
	RegexpMoney, _     = regexp.Compile("^\\d+(\\.\\d{0,3})?$")                                     // 检查金额
	RegexpUrl, _       = regexp.Compile("^https?://[\\w]+[(.\\w+){1,4}]/?[\\w=?&-_:*#&%\\[\\].]*$") // 检查链接
	RegexpDate, _      = regexp.Compile("^\\d{4}-\\d{2}-\\d{2}$")                                   // 检查日期
	RegexpNumberSet, _ = regexp.Compile("^(\\d+)(,\\d+)*$")                                         // 检查一组连续的数字
	RegexpColorCode, _ = regexp.Compile("^[a-fA-F0-9]{6}$")                                         // 检查一组颜色代码
	RegexpMd5, _       = regexp.Compile("^[a-f0-9]{32}$")                                           // 检查是否是md5字符串
	RegexpZhPhone, _   = regexp.Compile("^1\\d{10}$")                                               // 检查中国大陆手机号
)
