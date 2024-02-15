package exception

type ApiException struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewApiException(status int, message string) *ApiException {
	ex := &ApiException{
		Status:  status,
		Message: message,
	}

	return ex
}
