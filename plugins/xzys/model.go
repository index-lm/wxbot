package xzys

type apiResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Todo struct {
			Yi string `json:"yi"`
			Ji string `json:"ji"`
		} `json:"todo"`
		Luckynumber string `json:"source"`
		Fortunetext struct {
			All    string `json:"all"`
			Love   string `json:"love"`
			Work   string `json:"work"`
			Money  string `json:"money"`
			Health string `json:"health"`
		} `json:"fortunetext"`
	} `json:"data"`
}
