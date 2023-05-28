package Err

import "strings"

func ErrTem(svr string) string {
	t := `
package Err

import "sample/infrastructure/err"

var ParameterError = err.NewDefaultErr("parameter error")
var SampleErr = err.NewErr(996, "sample err")
`
	t = strings.ReplaceAll(t, "sample", svr)
	return t
}
