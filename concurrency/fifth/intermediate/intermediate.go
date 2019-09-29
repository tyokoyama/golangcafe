package intermediate

import (
	"fmt"
	"os/exec"
	"runtime/debug"
	"github.com/tyokoyama/golangcafe/concurrency/fifth/lowlevel"
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

type IntermediateErr struct {
	error
}

func RunJob(id string) error {
	const jobBinPath = "/bad/job/binary"
	isExecuable, err := lowlevel.IsGloballyExec(jobBinPath)
	if err != nil {
		// きちんとエラーを内包すると、予期したエラーとしてエラーメッセージが表示できる。P.156
		// return err
		return IntermediateErr{wrapError(
			err,
			"cannot run job %q: requisite binaries not available",
			id,
		)}
	} else if isExecuable == false {
		return wrapError(nil, "job binary is not executable")
	}

	return exec.Command(jobBinPath, "--id="+id).Run()
}