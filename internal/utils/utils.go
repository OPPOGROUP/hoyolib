package utils

import "fmt"

func GetSignUrl(api, mark string) (string, string) {
	return fmt.Sprintf("%s/event/%s/sign", api, mark), fmt.Sprintf("%s/event/%s/info", api, mark)
}
