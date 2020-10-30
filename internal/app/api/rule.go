package api

import (
	"context"
	desc "dotTest/pkg/api"
	"net/http"

	"github.com/pkg/errors"
)

//Rule ...
func (i *Implementation) Rule(ctx context.Context, req *desc.RuleRequest) (*desc.RuleResponse, error) {
	if err := i.api.AddRule(ctx, req.Site, req.Node); err != nil {
		return &desc.RuleResponse{
			Status: http.StatusInternalServerError,
		}, errors.Wrap(err, "couldn't add rule")
	}

	return &desc.RuleResponse{
		Status: http.StatusOK,
	}, nil
}
