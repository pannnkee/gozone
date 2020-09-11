package enum

type ResponseCode int

const (
	DefaultSuccess ResponseCode = iota
	DefaultError
)

type SortType int

const (
	TimeSort SortType = iota
	HotSort
)

type ContentType int

const (
	DefaultType ContentType = iota
	ClassType
	TagType
)
