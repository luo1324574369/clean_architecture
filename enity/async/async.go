package async

import (
	"context"
	"time"
)

func Go(ctx context.Context, timeout time.Duration, handler func(context.Context)) error {
	return nil
}
