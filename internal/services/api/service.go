package api

import (
	"context"
	"dotTest/internal/db"
	"net/url"
	"unicode"

	"github.com/pkg/errors"
)

const (
	rssExt string = ".rss"
)

type Storage interface {
	InsertRule(ctx context.Context, site string, node string) error
	SelectNews(ctx context.Context, filter string) ([]*db.OneNews, error)
}

type Service struct {
	storage Storage
}

func NewService(storage Storage) *Service {
	return &Service{
		storage: storage,
	}
}

func (s *Service) AddRule(ctx context.Context, site string, node string) error {
	_, err := url.ParseRequestURI(site)
	if err != nil {
		return errors.Wrapf(err, "couldn't connect to url %s", site)
	}

	if err := s.storage.InsertRule(ctx, site, node); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetNews(ctx context.Context, filter string) ([]*db.OneNews, error) {
	if !isLetter(filter) {
		var res []*db.OneNews
		err := errors.New("filter should be alphabet")
		return res, errors.Wrapf(err, "received : %v", filter)
	}
	return s.storage.SelectNews(ctx, filter)
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
