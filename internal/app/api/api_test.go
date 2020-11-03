package api

import (
	"context"
	"dotTest/internal/app/api/mocks"
	"dotTest/internal/db"
	"testing"
)

func createImplementation(t *testing.T, errIn error) *Implementation {
	api := mocks.NewApiMock(t)
	api.AddRuleMock.Set(func(ctx context.Context, site string, node string) (err error) {
		return errIn
	})
	retList := []*db.OneNews{
		{
			Title: "very good news",
		},
	}

	api.GetNewsMock.Set(func(ctx context.Context, filter string) ([]*db.OneNews, error) {
		return retList, errIn
	})

	return NewApi(api)
}
