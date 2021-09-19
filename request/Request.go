package request

// 抖音直播数据
type LogRequest struct {
	// binding:"required"修饰的字段，若接收为空值，则报错，是必须字段
	LiveId   string `form:"live_id"  binding:"required"`
	Data     string `form:"data"  binding:"required"`
	Msg_Type string `form:"msg_type"  binding:"required"`
}
