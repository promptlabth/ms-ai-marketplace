package review

import (
	"context"
)

type ReviewInterface interface {
	NewReview(ctx context.Context, review ReviewRequest) error
}