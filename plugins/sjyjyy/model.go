package sjyjyy

type apiResponse struct {
	Success bool   `json:"success"`
	Type    string `json:"type"`
	Data    struct {
		Id      int64  `json:"id"`
		Vhan    string `json:"vhan"`
		Source  string `json:"source"`
		Creator string `json:"creator"`
	} `json:"data"`
}
type apiResponse2 struct {
	Success bool   `json:"success"`
	Ishan   string `json:"ishan"`
}
type apiResponse3 struct {
	Success bool   `json:"success"`
	Joke    string `json:"joke"`
}
