package jobs

import (
	"bytes"
	"context"
	"dotTest/internal/db"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/antchfx/xmlquery"
	"golang.org/x/net/html"
)

const extRss string = ".rss"

//Storage ...
type Storage interface {
	GetRequestData(ctx context.Context) ([]*db.RequestData, error)
	InsertNews(ctx context.Context, news []*db.OneNewsForIns) error
}

//Run ...
func Run(ctx context.Context, storage Storage) error {
	for {
		data, err := storage.GetRequestData(ctx)
		// fmt.Printf("%v %T", data, data)
		if err != nil || len(data) == 0 {
			fmt.Println("Не удалось получить критерии поиска")
			return nil
		}
		var insData []*db.OneNewsForIns
		for _, item := range data {
			if filepath.Ext(item.Name) != extRss {
				insData, err = getItemsForHtml(item)
			} else {
				insData, err = getItemsForRss(item)
			}
			if err != nil {
				fmt.Println(err)
				continue
			}
			errIns := storage.InsertNews(ctx, insData)
			if errIns != nil {

			}
		}
		timeSleep, err := strconv.Atoi(os.Getenv("TIME_RELOAD"))
		if err != nil {
			break
		}
		time.Sleep(time.Duration(timeSleep) * time.Second)
	}
	return nil
}

func getHtmlDataLink(n *html.Node, title *bytes.Buffer) {
	if n.Type == html.TextNode {
		title.WriteString(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		getHtmlDataLink(c, title)
	}
}

func getItemsForHtml(item *db.RequestData) ([]*db.OneNewsForIns, error) {
	doc, errLoad := htmlquery.LoadURL(item.Name)
	if errLoad != nil {
		return nil, errLoad
	}
	nodes := htmlquery.Find(doc, item.Node)
	var res []*db.OneNewsForIns

	title := &bytes.Buffer{}
	for _, node := range nodes {
		node = htmlquery.FindOne(node, "//a")
		getHtmlDataLink(node, title)
		one := &db.OneNewsForIns{
			SiteId: item.SiteId,
			NodeId: item.NodeId,
			Title:  title.String(),
			Link:   htmlquery.SelectAttr(node, "href"),
		}
		res = append(res, one)
	}
	return res, nil
}

func getItemsForRss(item *db.RequestData) ([]*db.OneNewsForIns, error) {
	doc, errLoad := xmlquery.LoadURL(item.Name)
	if errLoad != nil {
		return nil, errLoad
	}
	nodes := xmlquery.Find(doc, item.Node)
	var res []*db.OneNewsForIns

	for _, node := range nodes {
		one := &db.OneNewsForIns{
			SiteId: item.SiteId,
			NodeId: item.NodeId,
			Title:  xmlquery.FindOne(node, "//title").InnerText(),
			Link:   xmlquery.FindOne(node, "//link").InnerText(),
		}

		res = append(res, one)
	}

	return res, nil
}
