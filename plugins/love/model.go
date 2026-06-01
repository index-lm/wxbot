package sjyjyy

type apiResponse struct {
	Success bool   `json:"success"`
	Type    string `json:"type"`
	Data    Data   `json:"data"`
}

type Data struct {
	Content string `json:"content"`
	Id      int    `json:"id"`
}
