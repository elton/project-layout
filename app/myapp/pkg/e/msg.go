package e

// MsgFlags is the map of error code and error message
var MsgFlags = map[int]string{
	SUCCESS:                "ok",
	ERROR:                  "fail",
	INVALID_PARAMS:         "请求参数错误",
	ERROR_RECORD_NOT_FOUND: "记录不存在",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
