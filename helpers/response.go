package helpers

type MetaFormat struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type ResponseFormat struct {
	Meta MetaFormat  `json:"meta"`
	Data interface{} `json:"data"`
}

func Response(code int, status string, message string, data interface{}) ResponseFormat {
	return ResponseFormat{
		Meta: MetaFormat{
			Code:    code,
			Status:  status,
			Message: message,
		},
		Data: data,
	}
}
