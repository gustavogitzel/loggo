package loggo

import "context"

type LogContextTransformer func(context.Context) context.Context

func NewContext(ctx context.Context, transformers ...LogContextTransformer) context.Context {
	for _, t := range transformers {
		ctx = t(ctx)
	}
	return ctx
}

func getRequestId(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(requestIdContextKey).(string)
	return id, ok
}

func WithSessionId(id string) LogContextTransformer {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, sessionIdContextKey, id)
	}
}

func getSessionId(ctx context.Context) (string, bool) {
	id, ok := ctx.Value(sessionIdContextKey).(string)
	return id, ok
}

func WithAdditionalFields(fields Fields) LogContextTransformer {
	return func(ctx context.Context) context.Context {
		// If there are any additional fields registered already, merge them with the new fields (without modifying
		// the previously existing map, which would affect the upstream context). New fields take priority.
		if existing, ok := getAdditionalFields(ctx); ok {
			for k, v := range existing {
				if _, ok := fields[k]; !ok {
					fields[k] = v
				}
			}
		}

		return context.WithValue(ctx, additionalFieldsContextKey, fields)
	}
}

func WithAdditionalField(name string, value interface{}) LogContextTransformer {
	return WithAdditionalFields(Fields{name: value})
}

func getAdditionalFields(ctx context.Context) (Fields, bool) {
	// For this method in particular, we might want to call it on a nil context.
	if ctx == nil {
		return nil, false
	}

	fields, ok := ctx.Value(additionalFieldsContextKey).(Fields)
	return fields, ok
}

type requestIdContextKeyType struct{}

var requestIdContextKey requestIdContextKeyType = requestIdContextKeyType(struct{}{})

type sessionIdContextKeyType struct{}

var sessionIdContextKey sessionIdContextKeyType = sessionIdContextKeyType(struct{}{})

type additionalFieldsContextKeyType struct{}

var additionalFieldsContextKey additionalFieldsContextKeyType = additionalFieldsContextKeyType(struct{}{})
