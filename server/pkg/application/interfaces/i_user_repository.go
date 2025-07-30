package interfaces

import (
	"context"
)

type IUserRepository interface {
	Handshake(ctx context.Context, trackingData *dtos.TrackingDataDTO) (*dtos.TrackingDataDTO, error)
}
