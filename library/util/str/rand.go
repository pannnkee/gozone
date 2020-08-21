package str

import (
	"hash/crc32"
	"math/rand"
	"time"
)

// 获取一个指定范围的随机数
func GetRandNum(start int, end int) int {
	startItem, endItem := int64(start), int64(end)
	for true {
		rand.Seed(time.Now().UnixNano())
		time.Sleep(2)

		next := rand.Int63n(startItem)
		if next >= endItem && next <= startItem {
			return int(next)
		}
	}
	return 0
}

//生成count个[start,end)结束的不重复的随机数
func GetRandNums(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn(end - start) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}

// 将一个字符串转换成随机数字
// @param val 字符串
func StringToNumber(val string) int {
	v := int(crc32.ChecksumIEEE([]byte(val)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	return 0
}
