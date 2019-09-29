package lowlevel

import (
	"fmt"
	"runtime/debug"
	"os"
)

type MyError struct {
	Inner error
	Message string
	StackTrace string
	Misc map[string]interface{}
}

func wrapError(err error, messagef string, msgArgs ...interface{}) MyError {
	return MyError{
		Inner: err,
		Message: fmt.Sprintf(messagef, msgArgs...),
		StackTrace: string(debug.Stack()),
		Misc: make(map[string]interface{}),
	}
}

func (err MyError) Error() string {
	return err.Message
}

type LowLevelError struct {
	error
}

func IsGloballyExec(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, LowLevelError{(wrapError(err, err.Error()))}
	}

	return info.Mode().Perm()&0100 == 0100, nil
}