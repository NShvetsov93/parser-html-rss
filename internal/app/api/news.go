package api

import (
	"context"
	"dotTest/internal/db"
	desc "dotTest/pkg/api"

	"github.com/pkg/errors"
)

func (i *Implementation) News(ctx context.Context, req *desc.NewsRequest) (*desc.NewsResponse, error) {
	news, err := i.api.GetNews(ctx, req.Filter)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't get list of products")
	}
	return convertToListResponse(news), nil
}

func convertToListResponse(news []*db.OneNews) *desc.NewsResponse {
	listNews := make([]*desc.NewsResponse_OneNews, 0, len(news))

	for _, n := range news {
		ln := &desc.NewsResponse_OneNews{
			Title: n.Title,
		}

		listNews = append(listNews, ln)
	}

	return &desc.NewsResponse{Onenews: listNews}
}
