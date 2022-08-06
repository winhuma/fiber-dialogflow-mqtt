package models

type MyResponse struct {
	Message    interface{}
	Success    interface{}
	StatusCode interface{}
	Data       interface{}
}

func SetMyResponse(msg string, code int, success bool, data interface{}) interface{} {
	return MyResponse{
		Message:    msg,
		StatusCode: code,
		Data:       data,
		Success:    success,
	}
}

func ResponseSuccess(msg string, data ...interface{}) interface{} {
	var setData interface{}
	if len(data) == 0 {
		setData = nil
	} else {
		setData = data[0]
	}
	return MyResponse{
		Message:    msg,
		Data:       setData,
		StatusCode: 200,
		Success:    true,
	}
}

func ResponseFail(msg string) interface{} {
	return MyResponse{
		Message:    msg,
		StatusCode: 400,
	}
}
