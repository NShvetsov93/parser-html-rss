package api

import (
	"context"
	apipb "dotTest/pkg/api"
	"net/http"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type testCaseRule struct {
	name    string
	want    *apipb.RuleResponse
	wantErr bool
	err     error
}

func TestImplementation_Rule(t *testing.T) {
	for _, test := range getTestCaseRule() {
		t.Run(test.name, func(t *testing.T) {
			i := createImplementation(t, test.err)

			resp, err := i.Rule(context.Background(), getRuleRequest())

			if test.wantErr {
				require.Error(t, err)
				assert.EqualError(t, err, errors.Wrap(test.err, "couldn't add rule").Error())
				assert.Equal(t, test.want, resp)
			} else {
				require.NoError(t, err)
				assert.Equal(t, test.want, resp)
			}
		})
	}
}

func getRuleRequest() *apipb.RuleRequest {
	return &apipb.RuleRequest{
		Site: "https://yandex.ru",
		Node: `//*[@id="news_panel_news"]/ol[1]/li`,
	}
}

func getTestCaseRule() []testCaseRule {
	return []testCaseRule{
		{
			name:    "positive case",
			want:    &apipb.RuleResponse{Status: http.StatusOK},
			wantErr: false,
		},
		{
			name:    "negative case",
			want:    &apipb.RuleResponse{Status: http.StatusInternalServerError},
			wantErr: true,
			err:     assert.AnError,
		},
	}
}
