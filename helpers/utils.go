package helpers

import (
	"fmt"
	"strconv"
)

func ToInt(val string) int {
	num, err := strconv.Atoi(val)
	if err != nil {
		fmt.Println(err)
		num = 0
	}
	return num
}
