package health

import "context"

type Checker interface {
	Check(ctx context.Context) error
}
