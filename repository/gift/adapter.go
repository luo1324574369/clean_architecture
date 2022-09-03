package gift

import (
	"context"

	"github.com/luo1324574369/clean_architecture/enity/gift"
)

type Adapter struct {
}

func (s *Adapter) Create(ctx context.Context, entity *gift.Entity) error {
	return nil
}
