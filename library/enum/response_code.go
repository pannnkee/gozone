package enum

type ResponseCode int

const (
	DefaultSuccess ResponseCode = iota
	DefaultError
)

const (
	TimeSort = 1
	HotSort = 2
)
