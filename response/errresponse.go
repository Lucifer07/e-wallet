package response

type ResponseMsgErr struct {
	StatusCode int `json:"-"`
	Message    string `json:"message"`
}
