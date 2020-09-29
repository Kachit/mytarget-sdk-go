package mytarget_sdk

import (
	"fmt"
	"strings"
)

func ArrayToString(a []int, delimiter string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delimiter, -1), "[]")
}
