package logger

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
	"time"
)

var (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
)

func colorForLevel(level slog.Level) string {
	switch level {
	case slog.LevelDebug:
		return green
	case slog.LevelInfo:
		return blue
	case slog.LevelWarn:
		return yellow
	case slog.LevelError:
		return red
	default:
		return reset
	}
}

type colorWriter struct {
	writer io.Writer
	level  slog.Level
}

func (cw *colorWriter) Write(p []byte) (int, error) {
	s := string(p)
	color := reset

	switch {
	case strings.Contains(s, "level=DEBUG"):
		color = green
	case strings.Contains(s, "level=INFO"):
		color = blue
	case strings.Contains(s, "level=WARN"):
		color = yellow
	case strings.Contains(s, "level=ERROR"):
		color = red
	}

	colored := fmt.Sprintf("%s%s%s", color, s, reset)
	return cw.writer.Write([]byte(colored))
}


func New(level slog.Level) *slog.Logger {
	cw := &colorWriter{
		writer: os.Stdout,
		level:  level,
	}

	handler := slog.NewTextHandler(cw, &slog.HandlerOptions{
		Level:     level,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				a.Value = slog.StringValue(strings.ToUpper(a.Value.String()))
			}
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(time.Now().Format("15:04:05"))
			}
			return a
		},
	})

	return slog.New(handler)
}

var Log *slog.Logger

func Init(level slog.Level) {
	Log = New(level)
}
