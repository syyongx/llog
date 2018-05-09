package llog

import (
	"testing"
	"fmt"
	"github.com/syyongx/llog/handler"
	"github.com/syyongx/llog/types"
	"github.com/syyongx/llog/formatter"
	"time"
)

func Test1(t *testing.T) {
	logger := NewLogger("demo")
	h, err := handler.NewFile("./llog.log", types.WARNING, true)
	buf := handler.NewBuffer(h, 1, types.WARNING, true, true)
	if err != nil {
		fmt.Println(err.Error())
	}
	f := formatter.NewLine("%datetime% [%levelName%] [%channel%] %message%\n", time.RFC3339)
	h.SetFormatter(f)
	logger.PushHandler(buf)
	logger.Warning("xxx")
	buf.Close()
}
