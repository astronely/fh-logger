package logger

import (
	"context"
	"encoding/json"
	"github.com/fatih/color"
	"io"
	"log"
	"log/slog"
)

type PrettyHandlerOptions struct {
	SlogOpts *slog.HandlerOptions
}

type PrettyHandler struct {
	opts PrettyHandlerOptions
	slog.Handler
	l     *log.Logger
	attrs []slog.Attr
}

func (opts PrettyHandlerOptions) NewPrettyHandler(out io.Writer) *PrettyHandler {
	h := &PrettyHandler{
		Handler: slog.NewJSONHandler(out, opts.SlogOpts),
		l:       log.New(out, "", 0),
	}

	return h
}

func (h *PrettyHandler) Handle(_ context.Context, r slog.Record) error {
	level := r.Level.String() + ":"

	switch r.Level {
	case slog.LevelDebug:
		level = color.GreenString(level)
	case slog.LevelInfo:
		level = color.BlueString(level)
	case slog.LevelWarn:
		level = color.YellowString(level)
	case slog.LevelError:
		level = color.RedString(level)
	}

	fields := make(map[string]interface{}, r.NumAttrs())

	r.Attrs(func(a slog.Attr) bool {
		fields[a.Key] = a.Value.Any()

		return true
	})

	for _, a := range h.attrs {
		fields[a.Key] = a.Value.Any()
	}

	var b []byte
	var err error
	if len(fields) > 0 {
		b, err = json.MarshalIndent(fields, "", "  ")
		if err != nil {
			return err
		}
	}

	time := r.Time.Format("[02-01-2006 15:05:05.000]")
	msg := color.CyanString(r.Message)

	h.l.Println(
		time,
		level,
		msg,
		color.WhiteString(string(b)),
	)

	return nil
}

//func (h *PrettyHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
//	return &PrettyHandler{
//		Handler: h.Handler,
//		l:       h.l,
//		attrs:   attrs,
//	}
//}
//
//func (h *PrettyHandler) WithGroup(name string) slog.Handler {
//	return &PrettyHandler{
//		Handler: h.Handler.WithGroup(name),
//		l:       h.l,
//	}
//}
