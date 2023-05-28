package model

import "strings"

func ModelTem(svr string) string {
	t := `
package model

type Sample struct{

}
`
	t = strings.ReplaceAll(t, "Sample", strings.Title(svr[:1])+svr[1:])
	return t
}
