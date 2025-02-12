package interfaces

import (
	"context"
	"ecommerce-order/external"
)

type IExternal interface {
	GetProfile(ctx context.Context, token string) (external.Profile, error)
	ProductKafkaMessage(ctx context.Context, data []byte) error
}
