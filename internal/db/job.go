package db

import (
	"context"
	"strings"

	"github.com/pkg/errors"
)

//RequestData ...
type RequestData struct {
	Name   string
	Node   string
	SiteId int32
	NodeId int32
}

//OneNewsForIns ...
type OneNewsForIns struct {
	SiteId int32
	NodeId int32
	Title  string
	Link   string
}

//GetRequestData ...
func (s Storage) GetRequestData(ctx context.Context) ([]*RequestData, error) {
	var data []*RequestData
	rows, err := s.db.Query(ctx,
		"select s.name,n.value,s.id,n.id from sites as s inner join nodes as n on s.id=n.site_id")

	for rows.Next() {
		row := &RequestData{}
		err = rows.Scan(&row.Name, &row.Node, &row.SiteId, &row.NodeId)
		if err == nil {
			row.Node = strings.Trim(row.Node, "\\")
			data = append(data, row)
		}
	}
	return data, nil
}

//InsertNews ...
func (s Storage) InsertNews(ctx context.Context, news []*OneNewsForIns) error {
	for _, item := range news {
		_, insErr := s.db.Query(ctx,
			"insert into news (site_id,node_id,title,link) values ($1,$2,$3,$4)",
			item.SiteId, item.NodeId, item.Title, item.Link)
		if insErr != nil {
			return errors.Wrapf(insErr, "Новость: \"%s\" не добавлена", item.Title)
		}
	}
	return nil
}
