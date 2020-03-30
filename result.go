package main

type Result struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Page  int         `json:"page;omitempty"`
	Limit int         `json:"limit;omitempty"`
	Count int         `json:"count;omitempty"`
	Data  interface{} `json:"data;omitempty"`
}

func NewResultError(code int, err error) *Result {
	return &Result{Code: code, Msg: err.Error()}
}
func NewResultSuccess(code int, msg string, data interface{}) *Result {
	return &Result{Code: code, Msg: msg, Data: data}
}
func NewResultsSuccess(code, page, limit, count int, msg string, data interface{}) *Result {
	return &Result{Code: code, Msg: msg, Page: page, Limit: limit, Count: count, Data: data}
}
