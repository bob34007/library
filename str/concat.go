package str

import "strings"

func SliceConcat(s []string, lineTerminator string) (str string) {
	for _, v := range s {
		str += v + lineTerminator
	}
	return strings.Trim(str, lineTerminator)
}
