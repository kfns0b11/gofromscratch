package cuslog

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

// log level
type Level uint8

const (
	FmtEmptySeparate = ""
)

// const log level
const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
	PanicLevel
	FatalLevel
)

// mapping log level with string name
var LevelNameMapping = map[Level]string{
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
	PanicLevel: "PANIC",
	FatalLevel: "FALTAL",
}

// log options
type options struct {
	output        io.Writer
	level         Level
	stdlevel      Level
	formatter     Formatter
	disableCaller bool
}

type Option func(*options)

func initOptions(ops ...Option) (o *options) {
	o = &options{}
	for _, opt := range ops {
		opt(o)
	}

	if o.output == nil {
		o.output = os.Stderr
	}

	if o.formatter == nil {
		o.formatter = &TextFormatter{}
	}

	return
}

var errUnmarshalNilLevel = errors.New("can't unmarshal a nil *Level")

func (l *Level) unmarshalText(text []byte) bool {
	switch string(text) {
	case "debug", "DEBUG":
		*l = DebugLevel
	case "info", "INFO", "": // make the zero value useful
		*l = InfoLevel
	case "warn", "WARN":
		*l = WarnLevel
	case "error", "ERROR":
		*l = ErrorLevel
	case "panic", "PANIC":
		*l = PanicLevel
	case "fatal", "FATAL":
		*l = FatalLevel
	default:
		return false
	}
	return true
}

// UnmarshalText unmarshals text to a level.
func (l *Level) UnmarshalText(text []byte) error {
	if l == nil {
		return errUnmarshalNilLevel
	}
	if !l.unmarshalText(text) && !l.unmarshalText(bytes.ToLower(text)) {
		return fmt.Errorf("unrecognized level: %q", text)
	}
	return nil
}

func WithOutput(output io.Writer) Option {
	return func(o *options) {
		o.output = output
	}
}

func WithLevel(level Level) Option {
	return func(o *options) {
		o.level = level
	}
}

func WithStdLevel(level Level) Option {
	return func(o *options) {
		o.stdlevel = level
	}
}

func WithFormatter(farmatter Formatter) Option {
	return func(o *options) {
		o.formatter = farmatter
	}
}

func WithDisableCaller(caller bool) Option {
	return func(o *options) {
		o.disableCaller = caller
	}
}
