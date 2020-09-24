module Gozone/src/zone

go 1.13

replace Gozone/library => ../../library

require (
	Gozone/library v0.0.0-00010101000000-000000000000
	github.com/astaxie/beego v1.12.2
	github.com/buger/jsonparser v1.0.0
	github.com/jinzhu/gorm v1.9.16
	github.com/prometheus/common v0.13.0
)
