package logger

import (
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/fatih/color"
)

type VerbosityLevel int

type LogCallback func(message interface{}, verbosityLevel VerbosityLevel, data ...interface{})

var LoggerInstance *Logger

const (
	Debug VerbosityLevel = iota
	Info
	Warning
	Error
	Trace
)

type LoggerOptions struct {
	DateFormat string
}

func NewLoggerOptions(dateFormat string) *LoggerOptions {
	return &LoggerOptions{
		DateFormat: dateFormat,
	}
}

type Logger struct {
	Verbosity VerbosityLevel
	Callback  LogCallback
	Options   *LoggerOptions
}

type LogMessage struct {
	Verbosity VerbosityLevel
	Message   string
}

func Setup(verbosity VerbosityLevel, callback LogCallback, options *LoggerOptions) {
	LoggerInstance = &Logger{
		Verbosity: verbosity,
		Callback:  callback,
		Options:   options,
	}
}

func Log(message interface{}, level VerbosityLevel, data ...interface{}) {
	logger := getLoggerInstance()
	if logger.Callback != nil {
		logger.Callback(message, level, data...)
	}

	var datetime = time.Now()

	colour := color.New(level.Colour()).SprintFunc()
	if int(level) >= int(logger.Verbosity) { // Only print appropriate verbosity messages
		fmt.Printf("%s | %-16s | %s (%s)\n", datetime.Format(logger.Options.DateFormat), colour(level.String()), message, getFrame(2).Function)
	}
}

func (d VerbosityLevel) String() string {
	return []string{"Debug", "Info", "Warning", "Error", "Trace"}[d]
}

func (d VerbosityLevel) Colour() color.Attribute {
	return []color.Attribute{color.FgMagenta, color.FgBlue, color.FgYellow, color.FgRed, color.FgGreen}[d]
}

func getVerbosity(level string) VerbosityLevel {
	verbosityLevels := make(map[string]VerbosityLevel)
	verbosityLevels["debug"] = Debug
	verbosityLevels["info"] = Info
	verbosityLevels["warning"] = Warning
	verbosityLevels["error"] = Error
	verbosityLevels["trace"] = Trace

	var verbosity VerbosityLevel
	if v, ok := verbosityLevels[strings.ToLower(level)]; !ok {
		log.Fatalf("Unknown log verbosity %s", v)
	} else {
		verbosity = v
	}

	return verbosity
}

func getFrame(skipFrames int) runtime.Frame {
	// We need the frame at index skipFrames+2, since we never want runtime.Callers and getFrame
	targetFrameIndex := skipFrames + 2

	// Set size to targetFrameIndex+2 to ensure we have room for one more caller than we need
	programCounters := make([]uintptr, targetFrameIndex+2)
	n := runtime.Callers(0, programCounters)

	frame := runtime.Frame{Function: "unknown"}
	if n > 0 {
		frames := runtime.CallersFrames(programCounters[:n])
		for more, frameIndex := true, 0; more && frameIndex <= targetFrameIndex; frameIndex++ {
			var frameCandidate runtime.Frame
			frameCandidate, more = frames.Next()
			if frameIndex == targetFrameIndex {
				frame = frameCandidate
			}
		}
	}

	return frame
}

func getLoggerInstance() Logger {
	if LoggerInstance == nil {
		newLogger := Logger{Verbosity: getVerbosity("error"), Options: NewLoggerOptions(time.RFC3339), Callback: nil}
		LoggerInstance = &newLogger
		return newLogger
	}
	return *LoggerInstance
}
