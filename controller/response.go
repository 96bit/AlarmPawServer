package controller

import (
	"AlarmPawServer/modal"
	"fmt"
	"time"
)

// for the fast return success result
func success() modal.CommonResp {
	return modal.CommonResp{
		Code:      200,
		Message:   "success",
		Timestamp: time.Now().Unix(),
	}
}

// for the fast return failed result
func failed(code int, message string, args ...interface{}) modal.CommonResp {
	return modal.CommonResp{
		Code:      code,
		Message:   fmt.Sprintf(message, args...),
		Timestamp: time.Now().Unix(),
	}
}

// for the fast return result with custom data
func data(data interface{}) modal.CommonResp {
	return modal.CommonResp{
		Code:      200,
		Message:   "success",
		Timestamp: time.Now().Unix(),
		Data:      data,
	}
}
