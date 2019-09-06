package models

// 返回信息格式
type Response struct {
	RespCode string
	RespDesc string
	RespData interface{}
}
