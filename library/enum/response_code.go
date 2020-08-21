package enum

type ResponseCode int

const (
	DefaultSuccess ResponseCode = iota
	DefaultError
)
