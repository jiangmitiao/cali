package models

type Api struct {
	Result     bool        `json:"result"`
	StatusCode int         `json:"statusCode"`
	Message    string      `json:"message"`
	Info       interface{} `json:"info"`
}

func NewOKApi() Api {
	return Api{Result: true, StatusCode: 200, Message: "ok"}
}

func NewOKApiWithInfo(info interface{}) Api {
	return Api{Result: true, StatusCode: 200, Message: "ok", Info: info}
}

func NewOKApiWithMessageAndInfo(message string, info interface{}) Api {
	return Api{Result: true, StatusCode: 200, Message: message, Info: info}
}

func NewErrorApi() Api {
	return Api{Result: false, StatusCode: 500, Message: "error"}
}

func NewErrorApiWithInfo(info interface{}) Api {
	return Api{Result: false, StatusCode: 500, Message: "error", Info: info}
}

func NewErrorApiWithMessageAndInfo(message string, info interface{}) Api {
	return Api{Result: false, StatusCode: 500, Message: message, Info: info}
}
