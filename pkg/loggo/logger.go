package loggo

import (
	"context"
	"encoding/json"
	"fmt"
)

type Fields = map[string]interface{}

type LogLevel = string

const (
	FatalLogLevel LogLevel = "FATAL"
	ErrorLogLevel          = "ERROR"
	WarnLogLevel           = "WARN"
	InfoLogLevel           = "INFO"
	DebugLogLevel          = "DEBUG"
	TraceLogLevel          = "TRACE"
)

type Backend interface {
	Log(fields Fields, level LogLevel, params ...interface{})
}

type Logger interface {
	Fatal(context.Context, ...interface{})
	Error(context.Context, ...interface{})
	Warn(context.Context, ...interface{})
	Info(context.Context, ...interface{})
	Debug(context.Context, ...interface{})
	Trace(context.Context, ...interface{})

	Fallback(context.Context, ...interface{})
	CircuitBreakerOpen(context.Context, ...interface{})
}

type Loggo struct {
	ServiceName string
	Backend     Backend
}

func NewLogger(serviceName string, backend Backend) Logger {
	return &Loggo{serviceName, backend}
}

func formatFieldValue(value interface{}) string {
	if str, ok := value.(string); ok {
		return str
	}

	data, err := json.Marshal(value)
	if err != nil {
		return string(data)
	}

	return fmt.Sprintf("%+v", value)
}

const (
	requestIdFieldName string = "request_id"
	sessionIdFieldName        = "session_id"
)

func (l *Loggo) data(ctx context.Context, params []interface{}) (Fields, []interface{}) {
	fields := make(Fields)

	fields["service_name"] = l.ServiceName
	if ctx != nil {
		if id, ok := getRequestId(ctx); ok {
			fields[requestIdFieldName] = id
		}
		if id, ok := getSessionId(ctx); ok {
			fields[sessionIdFieldName] = id
		}
	}

	var c Fields
	if additional, ok := getAdditionalFields(ctx); ok {
		c = additional
	} else {
		c = make(Fields)
	}
	if f, ok := params[len(params)-1].(Fields); ok {
		params = params[:len(params)-1]
		for k, v := range f {
			c[k] = formatFieldValue(v)
		}
	}
	if len(c) != 0 {
		fields["context"] = c
	}

	return fields, params
}

func (l Loggo) Fatal(ctx context.Context, params ...interface{}) {
	if l.Backend == nil {
		return
	}
	fields, params := l.data(ctx, params)
	l.Backend.Log(fields, FatalLogLevel, params...)
}

func (l Loggo) Error(ctx context.Context, params ...interface{}) {
	if l.Backend == nil {
		return
	}
	fields, params := l.data(ctx, params)
	l.Backend.Log(fields, ErrorLogLevel, params...)
}

func (l Loggo) Warn(ctx context.Context, params ...interface{}) {
	if l.Backend == nil {
		return
	}
	fields, params := l.data(ctx, params)
	l.Backend.Log(fields, WarnLogLevel, params...)
}

func (l Loggo) Info(ctx context.Context, params ...interface{}) {
	if l.Backend == nil {
		return
	}
	fields, params := l.data(ctx, params)
	l.Backend.Log(fields, InfoLogLevel, params...)
}

func (l Loggo) Debug(ctx context.Context, params ...interface{}) {
	if l.Backend == nil {
		return
	}
	fields, params := l.data(ctx, params)
	l.Backend.Log(fields, DebugLogLevel, params...)
}

func (l Loggo) Trace(ctx context.Context, params ...interface{}) {
	if l.Backend == nil {
		return
	}
	fields, params := l.data(ctx, params)
	l.Backend.Log(fields, TraceLogLevel, params...)
}

func (l Loggo) Fallback(ctx context.Context, params ...interface{}) {
	l.Warn(ctx, append([]interface{}{"[FALLBACK] "}, params...)...)
}

func (l *Loggo) CircuitBreakerOpen(ctx context.Context, params ...interface{}) {
	l.Warn(ctx, params...)
}
