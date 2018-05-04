package llog

import (
	"testing"
	"fmt"
	"github.com/syyongx/llog/handler"
	"github.com/syyongx/llog/types"
	"github.com/syyongx/llog/formatter"
	"time"
	"os"
)

func Test1(t *testing.T) {
	lr := NewLogger("demo")
	h, err := handler.NewFile("./access.log", types.WARNING, true, false)
	if err != nil {
		fmt.Println(err.Error())
	}
	f := formatter.NewLine("%datetime% [%levelName%] [%channel%] %message%", time.RFC3339)
	h.(*handler.File).SetFormatter(f)
	lr.PushHandler(h)
	lr.Warning("xxx")
	os.Remove("./access.log")
}
