package api

import (
	"context"
	apipb "dotTest/pkg/api"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCaseRule struct {
	name    string
	want    *apipb.NewsResponse
	wantErr bool
	err     error
}

func TestImplementation_News(t *testing.T) {
	for _, test := range getTestCaseList() {
		t.Run(test.name, func(t *testing.T) {
			i := createImplementation(t, test.err)

			resp, err := i.News(context.Background(), getListRequest())

			if test.wantErr {
				require.Error(t, err)
				assert.EqualError(t, err, errors.Wrap(test.err, "couldn't get list of news").Error())
				assert.Equal(t, test.want, resp)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.want, resp)
			}
		})
	}
}

func getNewsRequest() *apipb.NewsRequest {
	return &apipb.NewsRequest{
		Filter: "ood",
	}
}

func getTestCaseList() []testCaseList {
	return []testCaseList{
		{
			name: "positive case",
			want: &apipb.NewsResponse{
				OneNews: []*apipb.NewsResponse_OneNews{
					{
						Title: "very good news",
					},
				},
			},
			wantErr: false,
		},
	}
}
