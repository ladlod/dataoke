package dataoke

import (
	"fmt"
	"strconv"
)

func getErrorsError(errors []error) (err error) {
	if len(errors) == 0 {
		return nil
	}
	for _, e := range errors {
		if e != nil {
			err = fmt.Errorf("%s | %s", err, e)
		}
	}
	return
}

func toString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return v
	case []byte:
		return string(v)
	case int64:
		return strconv.FormatInt(v, 10)
	case int32:
		return strconv.FormatInt(int64(v), 10)
	case int:
		return strconv.Itoa(v)
	case float64:
		return strconv.FormatFloat(v, 'g', 10, 64)
	case float32:
		return strconv.FormatFloat(float64(v), 'g', 10, 32)
	case bool:
		return strconv.FormatBool(v)
	default:
		bytes, _ := dJson.Marshal(v)
		return string(bytes)
	}
}
