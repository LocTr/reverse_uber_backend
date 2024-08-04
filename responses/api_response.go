package responses

type ApiResponse struct {
	Status   int                    `json:status`
	Messaage string                 `json:"message"`
	Data     map[string]interface{} `json:"data"`
}
