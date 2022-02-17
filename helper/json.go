package helper

type meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

type response struct {
	Meta meta        `json:"meta"`
	Data interface{} `json:"data"`
}

func APIResponse(message string, code int, data interface{}) response {
	status := "error"

	if code == 200 || code == 201 {
		status = "success"
	}

	metaRes := meta{Message: message, Code: code, Status: status}
	res := response{Meta: metaRes, Data: data}

	return res
}
