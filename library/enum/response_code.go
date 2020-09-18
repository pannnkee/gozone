package enum

type ResponseCode int

const (
	DefaultSuccess ResponseCode = iota
	DefaultError
)

type SortType int

const (
	TimeSort SortType = iota + 1
	HotSort
)

type ContentType int

const (
	DefaultType ContentType = iota
	ClassType
	TagType

	Mysql ContentType = iota + 98
	Docker
	Istio
	Kubernetes
	Faas
	Redis
)
