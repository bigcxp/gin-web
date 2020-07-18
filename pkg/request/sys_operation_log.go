package request

import "time"

// 创建操作日志结构体
type CreateOperationLogRequestStruct struct {
	ApiDesc    string        `json:"apiDesc"`
	Path       string        `json:"path"`
	Method     string        `json:"method"`
	Params     string        `json:"params"`
	Body       string        `json:"body"`
	Data       string        `json:"data"`
	Status     int           `json:"status"`
	Username   string        `json:"username"`
	RoleName   string        `json:"roleName"`
	Ip         string        `json:"ip"`
	IpLocation string        `json:"ipLocation"`
	Latency    time.Duration `json:"latency"`
	UserAgent  string        `json:"userAgent"`
}
