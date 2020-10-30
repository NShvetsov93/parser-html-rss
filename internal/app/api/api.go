package api

import (
	"context"
	"dotTest/internal/db"
)

type Api interface {
	AddRule(ctx context.Context, site string, node string) error
	GetNews(ctx context.Context, filter string) ([]*db.OneNews, error)
}

// Implementation ...
type Implementation struct {
	api Api
}

func NewApi(api Api) *Implementation {
	return &Implementation{
		api: api,
	}
}
