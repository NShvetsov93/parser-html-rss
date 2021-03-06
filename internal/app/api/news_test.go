package api

import (
	"context"
	apipb "dotTest/pkg/api"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCaseNews struct {
	name    string
	want    *apipb.NewsResponse
	wantErr bool
	err     error
}

func TestImplementation_News(t *testing.T) {
	for _, test := range getTestCaseNews() {
		t.Run(test.name, func(t *testing.T) {
			i := createImplementation(t, test.err)

			resp, err := i.News(context.Background(), getNewsRequest())

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

func getTestCaseNews() []testCaseNews {
	return []testCaseNews{
		{
			name: "positive case",
			want: &apipb.NewsResponse{
				Onenews: []*apipb.NewsResponse_OneNews{
					{
						Title: "very good news",
					},
				},
			},
			wantErr: false,
		},
	}
}
