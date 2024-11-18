package service

import (
	"context"
	"github.com/yosuke7040/commandservice/domain/models/categories"
)

type CategoryService interface {
	Add(ctx context.Context, category *categories.Category) error
	Update(ctx context.Context, category *categories.Category) error
	Delete(ctx context.Context, categoryId *categories.Category) error
}
