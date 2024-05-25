package response

type ResponseMsg struct {
	Message string `json:"message"`
	Data    any `json:"data,omitempty"`
}
