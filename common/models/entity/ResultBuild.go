package entity

import (
	"encoding/json"
	"github.com/kynchen/rachel_golang/common/constants"
)

type ResultBuild struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
}

// ResultSuccess 统一返回类
func ResultSuccess() []byte {
	result := ResultBuild{
		Code: constants.SUCCESS,
		Msg:  constants.SuccessMsg,
		Data: nil,
	}
	marshal, _ := json.Marshal(result)
	return marshal
}
// ResultSuccessData  统一返回类
func ResultSuccessData(data interface{}) []byte {
	result :=  &ResultBuild{
		Code: constants.SUCCESS,
		Msg:  constants.SuccessMsg,
		Data: data,
	}
	marshal, _ := json.Marshal(result)
	return marshal
}

// ResultError 返回系统异常
func ResultError() []byte {
	result := ResultBuild{
		Code: constants.SystemError,
		Msg:  constants.SystemErrorMsg,
		Data: nil,
	}
	marshal, _ := json.Marshal(result)
	return marshal
}

// ResultErrorMsg
// code 状态码，msg 消息
///*
func ResultErrorMsg(code int,msg string) []byte {
	result := ResultBuild{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	marshal, _ := json.Marshal(result)
	return marshal
}