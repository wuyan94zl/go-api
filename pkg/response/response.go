package response

type Response struct {
	Code    int
	Data    interface{}
	Message interface{}
}

func Success(data interface{}) {
	panic(&Response{Code: 200, Message: "请求成功", Data: data})
}

func Error(code int, message interface{}) {
	panic(&Response{Code: code, Message: message})
}
