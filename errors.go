package bugsnag_stacktracer

import (
	bugsnagerrors "github.com/bugsnag/bugsnag-go/errors"
	"github.com/pkg/errors"
)

type Error struct {
	err         error
	stackFrames []bugsnagerrors.StackFrame
}

func (Error) TypeName() string {
	return "error"
}

func (e Error) Error() string {
	return e.err.Error()
}

func (e Error) StackFrames() []bugsnagerrors.StackFrame {
	return e.stackFrames
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

var skipFrames = 0

// SetSkipFrames globally sets how many lines of the stacktrace are overhead by our logging/tracing solution
func SetSkipFrames(sf int) {
  skipFrames = 0
}

// FromStackTracer wraps an error that implements the github.com/pkg/errors stackTracer interface with a Bugsnag compatible error
func FromError(err error) Error {
	// extract the stacktrace
	var st errors.StackTrace
	errWithStackTrace, ok := err.(stackTracer)
	if ok {
		st = errWithStackTrace.StackTrace()
	}

	// convert github.com/pkg/errors []StackFrame to Bugsnag stackframe
	stack := make([]bugsnagerrors.StackFrame, 0)
	if st != nil {
		for i, f := range st {
			if i < skipFrames {
				continue
			}

			stack = append(stack, bugsnagerrors.NewStackFrame(uintptr(f)))
		}
	}

	return Error{
		err:         err,
		stackFrames: stack,
	}
}
