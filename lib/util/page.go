package util

type PagerInstance struct {
	HasNext            bool  // 是否有前一页
	NextPageNumber     int64 // 前一页页码数

	HasPrevious        bool  // 是否有下一页
	PreviousPageNumber int64 // 下一页页码数

	CurrentNumber      int64 // 当前页码数
	TotalNumber        int64 // 总页数
	SortType           int64 // 排序方式
	IsShow             bool  // 是否显示分页
}

// @page 当前页
// @nums 文章总数
// @limit 一页显示数量
// @sortType 排序类型
func HtmlPage(page, nums, limit, sortType int64) (data PagerInstance) {

	data.CurrentNumber = page
	data.SortType = sortType
	if nums > 0 && limit > 0 {
		if nums%limit == 0 {
			data.TotalNumber = nums / limit
		} else {
			data.TotalNumber = (nums / limit) + 1
		}
	}

	if data.CurrentNumber == 1 {
		data.HasPrevious = false
		data.PreviousPageNumber = 0

		data.HasNext = true
		data.NextPageNumber = page + 1

	} else if data.CurrentNumber == data.TotalNumber {
		data.HasNext = false
		data.NextPageNumber = 0

		data.HasPrevious = true
		data.PreviousPageNumber = page - 1

	} else {
		data.HasPrevious = true
		data.PreviousPageNumber = page - 1

		data.HasNext = true
		data.NextPageNumber = page + 1
	}
	return
}
