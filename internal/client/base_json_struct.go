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

type AccountInfoResponse struct {
	Retcode int    `json:"retcode"`
	Message string `json:"message"`
	Data    struct {
		List []struct {
			HasRole         bool   `json:"has_role"`
			GameId          int    `json:"game_id"`
			GameRoleId      string `json:"game_role_id"`
			Nickname        string `json:"nickname"`
			Region          string `json:"region"`
			Level           int    `json:"level"`
			BackgroundImage string `json:"background_image"`
			IsPublic        bool   `json:"is_public"`
			Data            []data `json:"data"`
			RegionName      string `json:"region_name"`
			Url             string `json:"url"`
			DataSwitches    []struct {
				SwitchId   int    `json:"switch_id"`
				IsPublic   bool   `json:"is_public"`
				SwitchName string `json:"switch_name"`
			} `json:"data_switches"`
			H5DataSwitches  []interface{} `json:"h5_data_switches"`
			BackgroundColor string        `json:"background_color"`
		} `json:"list"`
	} `json:"data"`
}

type gameInfo struct {
	GameRoleId string
	Region     string
	Nickname   string
	Level      int
	Data       []data
}

type data struct {
	Name  string `json:"name"`
	Type  int    `json:"type"`
	Value string `json:"value"`
}
