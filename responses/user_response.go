package responses

type FactureResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}
