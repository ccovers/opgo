package errs

import (
	"bytes"
	"fmt"
	"runtime"
	// "path/filepath"
	// "strings"
)

func Stack(skip int) string {
	buf := new(bytes.Buffer)

	callers := make([]uintptr, 32)
	n := runtime.Callers(skip, callers)
	frames := runtime.CallersFrames(callers[:n])
	for {
		if f, ok := frames.Next(); ok {
			fmt.Fprintf(buf, "%s\n\t%s:%d (0x%x)\n", f.Function, f.File, f.Line, f.PC)
		} else {
			break
		}
	}
	return buf.String()

}

func WithStack(err error) string {
	if err == nil {
		return "<nil>"
	}
	if e, ok := err.(interface {
		Stack() string
	}); ok && e.Stack() != "" {
		return err.Error() + "\n" + e.Stack()
	}
	return err.Error()
}

func GetFileInfo() (string, int, string) {
	fileName, line, functionName := "?", 0, "?"
	pc, fileName, line, ok := runtime.Caller(2)
	if ok {
		functionName = runtime.FuncForPC(pc).Name()
		// functionName = filepath.Ext(functionName)
		// functionName = strings.TrimPrefix(functionName, ".")
	}
	return fileName, line, functionName
}
