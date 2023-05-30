package configs

import "context"

func SetToCtx[T any](ctx context.Context, config T) context.Context {
	return context.WithValue(ctx, "config", config)
}

func GetFromCtx[T any](ctx context.Context) T {
	return ctx.Value("config").(T)
}
