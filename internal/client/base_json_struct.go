package client

type SignResponse struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		Code      string `json:"code"`
		RiskCode  int    `json:"risk_code"`
		Gt        string `json:"gt"`
		Challenge string `json:"challenge"`
		Success   int    `json:"success"`
		IsRisk    bool   `json:"is_risk"`
	} `json:"data"`
}

type SignInfoResponse struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		TotalSignDay  int    `json:"total_sign_day"`
		Today         string `json:"today"`
		IsSign        bool   `json:"is_sign"`
		IsSub         bool   `json:"is_sub"`
		Region        string `json:"region"`
		SignCntMissed int    `json:"sign_cnt_missed"`
		ShortSignDay  int    `json:"short_sign_day"`
	} `json:"data"`
}
