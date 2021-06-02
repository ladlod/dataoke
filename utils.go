package dataoke

import "fmt"

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
