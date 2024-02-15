package exception

type ApiException struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e *ApiException) Error() string {
	return e.Message
}

func NewApiException(status int, message string) *ApiException {
	ex := &ApiException{
		Status:  status,
		Message: message,
	}

	return ex
}
