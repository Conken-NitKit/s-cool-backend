package cerror

import "fmt"

type (
	Errors []error
)

func TogetherError(errs ...error) error {
	var e []error
	for _, v := range errs {
		if v == nil {
			continue
		}
		e = append(e, v)
	}
	if len(e) == 0 {
		return nil
	}
	return Errors(e)
}

func (a Errors) Error() string {
	var errStrings string
	for i, v := range a {
		if v == nil {
			continue
		}
		errStrings += fmt.Sprintf("%d: %+v \n", i, v)
	}
	return errStrings
}
