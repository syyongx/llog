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
	logger := NewLogger("demo")
	h, err := handler.NewFile("./access.log", types.WARNING, true, false)
	if err != nil {
		fmt.Println(err.Error())
	}
	f := formatter.NewLine("%datetime% [%levelName%] [%channel%] %message%\n", time.RFC3339)
	h.SetFormatter(f)
	logger.PushHandler(h)
	logger.Warning("xxx")
	os.Remove("./access.log")
}
