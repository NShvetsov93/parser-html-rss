package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/pkg/errors"
)

//Storage ...
type Storage struct {
	db *pgxpool.Pool
}

//OneNews ...
type OneNews struct {
	Title string
}

//NewStorage ...
func NewStorage(db *pgxpool.Pool) *Storage {
	return &Storage{
		db: db,
	}
}

//InsertRule ...
func (s Storage) InsertRule(ctx context.Context, site string, node string) error {
	_, insErr := s.db.Query(ctx, "insert into sites (name) values ($1)", site)
	if insErr != nil {
		return errors.Wrap(insErr, "Сайт не добавлено")
	}
	id, errS := getSiteIDByName(ctx, s, site)
	if errS != nil {
		return errors.Wrap(errS, "Что то пошло не так")
	}
	_, errN := s.db.Query(ctx, "insert into nodes (site_id,value) values ($1,$2)", id, node)
	if errN != nil {
		return errors.Wrapf(errN, "Правило не добавлено,данные: id - %v, node - %v", id, node)
	}
	return nil
}

//SelectNews ...
func (s Storage) SelectNews(ctx context.Context, filter string) ([]*OneNews, error) {
	var news []*OneNews
	sel := "select n.title from news as n"
	if filter != "" {
		sel += " where n.title like '%" + filter + "%'"
	}
	rows, err := s.db.Query(ctx, sel)

	if err != nil {
		return nil, err
	}
	for rows.Next() {
		row := &OneNews{}
		err = rows.Scan(&row.Title)
		if err == nil {
			news = append(news, row)
		} else {
			fmt.Println(err)
		}
	}
	return news, nil
}

func getSiteIDByName(ctx context.Context, s Storage, name string) (int32, error) {
	var id int32
	errS := s.db.QueryRow(ctx, "select id from sites where name=$1", name).Scan(&id)
	if errS != nil {
		return id, errors.Wrap(errS, "Не найден идентификатор сайта")
	}
	return id, nil
}

func getNodeIDByName(ctx context.Context, s Storage, name string) (int32, error) {
	var id int32
	errN := s.db.QueryRow(ctx, "select id from nodes where value=$1", name).Scan(&id)
	if errN != nil {
		return id, errors.Wrap(errN, "Не найден идентификатор правила")
	}
	return id, nil
}
