package durable

import (
	"context"
)

func AddOne(ctx context.Context, num int) (int, error) {
	return num + 1, nil
}
