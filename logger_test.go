package llog

import (
	"github.com/syyongx/llog/formatter"
	"github.com/syyongx/llog/handler"
	"github.com/syyongx/llog/types"
	"testing"
	"time"
)

func TestBasic(t *testing.T) {
	logger := NewLogger("test")
	file := handler.NewFile("/tmp/llog/go.log", types.WARNING, true, 0664)
	buf := handler.NewBuffer(file, 1, types.WARNING, true)
	f := formatter.NewLine("%Datetime% [%LevelName%] [%Channel%] %Message%\n", time.RFC3339)
	file.SetFormatter(f)
	logger.PushHandler(buf)
	logger.Warning("xxx")
	buf.Close()
}

func TestRotatingFile(t *testing.T) {
	logger := NewLogger("test")
	r := handler.NewRotatingFile("/tmp/llog/go.log", 2, types.WARNING, true, 0664)
	buf := handler.NewBuffer(r, 1, types.WARNING, true)
	f := formatter.NewLine("%Datetime% [%LevelName%] [%Channel%] %Message%\n", time.RFC3339)
	r.SetFormatter(f)
	logger.PushHandler(buf)
	logger.Warning("xxx")
	buf.Close()
}

func BenchmarkBasic(b *testing.B) {
	logger := NewLogger("test")
	file := handler.NewFile("/dev/null", types.WARNING, true, 0660)
	f := formatter.NewLine("%Datetime% [%LevelName%] [%Channel%] %Message%\n", time.RFC3339)
	file.SetFormatter(f)
	//buf := handler.NewBuffer(h, 1, types.WARNING, true, true)
	//logger.PushHandler(buf)
	logger.PushHandler(file)
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.Warning("xxx")
			//buf.Close()
		}
	})
}
