package conn

// redis链接库
var RFStruct = new(RedisFactory).init()

// redis库下标
type RedisDatabase int

//项目redis库下标
const RedisZone RedisDatabase = 0

