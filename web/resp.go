package web

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

const (
	ParamsErrCode = 2001
)

var ParamsError = NewCodeErr(errors.New("params error"), 2001)

type Response struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result"`
	ErrMsg string      `json:"err_msg,omitempty"`
}

type CodeErr struct {
	err     error
	ErrCode int
}

func (e *CodeErr) Error() string {
	return e.err.Error()
}

func NewCodeErr(err error, code int) *CodeErr {
	return &CodeErr{
		err:     err,
		ErrCode: code,
	}
}

func Resp(writer http.ResponseWriter, result interface{}, err error) {
	resp := Response{
		Result: result,
	}
	writer.Header().Set("Content-Type", "application/json")
	if err == nil {
		resp.Code = 200
		json.NewEncoder(writer).Encode(resp)

	} else if codeErr, ok := err.(*CodeErr); ok {
		resp.Code = codeErr.ErrCode
		resp.ErrMsg = err.Error()
		json.NewEncoder(writer).Encode(resp)
	} else {
		resp.Code = 500
		resp.ErrMsg = err.Error()
		log.Println("[error] server error", err)
		json.NewEncoder(writer).Encode(resp)
	}

}

func Succ(writer http.ResponseWriter, result interface{}) {
	Resp(writer, result, nil)
}

func Err(write http.ResponseWriter, err error) {
	Resp(write, nil, err)
}
