package cte

import (
	"fmt"
)

var (
	userAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 14_0_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) %s"
)

func GetHeaders(oversea bool) map[string]string {
	var headers = map[string]string{}
	if oversea {
		headers["User-Agent"] = fmt.Sprintf(userAgent, "miHoYoBBSOversea/1.5.0")
	} else {
		headers["User-Agent"] = fmt.Sprintf(userAgent, "miHoYoBBS/2.36.1")
	}
	return headers
}
