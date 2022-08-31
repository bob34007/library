package ptr

import (
	"errors"
	"fmt"
)

func MustNil(v error) {
	if v != nil {
		panic(v)
	}
}

func MustNotZero(args ...interface{}) error {
	for _, arg := range args {
		switch arg.(type) {
		case string:
			if len(arg.(string)) == 0 {
				return errors.New(arg.(string) + "length is zero")
			}
		case int8, int, int16, int32, int64:
			if arg.(int64) == 0 {
				return errors.New(fmt.Sprintf("%v", arg.(int64)) + "is zero")
			}
		case uint8, uint16, uint32, uint64:
			if arg.(uint64) == 0 {
				return errors.New(fmt.Sprintf("%v", arg.(int64)) + "is zero")
			}
		case float32, float64:
			if arg.(float64)-0 < 0.00000001 && arg.(float64)-0 > -0.00000001 {
				return errors.New(fmt.Sprintf("%v", arg.(float64)) + "is zero")
			}
		default:
			if arg == nil {
				return errors.New("arg is nil ")
			}
		}
	}
	return nil
}
