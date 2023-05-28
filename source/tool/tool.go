package tool

func ToolTem() string {
	t := `
package tool

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"os"
	"strings"
	"time"
)

type Tool struct {
}

var tool *Tool

func init() {
	tool = new(Tool)
}

func GetTool() *Tool {
	return tool
}
func (t *Tool) GetEnv() string {
	env := os.Getenv("environment")
	return env
}

//func NowString() string {
//	return time.Now().Format("2006-01-02 15:04:05")
//}

func (t *Tool) ConObjectIDToString(obj primitive.ObjectID) string {
	i := strings.TrimLeft(obj.Hex(), "0")
	return i
}

func (t *Tool) ConStringToObjectID(s string) primitive.ObjectID {
	obj, _ := primitive.ObjectIDFromHex(s)
	return obj
}

func (t *Tool) Milliseconds() int64 {
	return time.Now().UnixMilli()
}

`
	return t
}
