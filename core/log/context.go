package log

import "context"

// constant
const key = "log_fields"

func ToContext(ctx context.Context) context.Context {
	return l.ToContext(ctx)
}

func FromContext(ctx context.Context) Logger {
	return l.FromContext(ctx)
}

func fieldsFromContext(ctx context.Context) Fields {

	var fields Fields

	if ctx == nil {
		return Fields{}
	}

	if param := ctx.Value(key); param != nil {
		fields = ctx.Value(key).(Fields)
	}

	return fields
}
