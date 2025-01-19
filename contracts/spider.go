package contracts

import (
	"github.com/orangbus/spider"
	"github.com/orangbus/spider/pkg/movie_spider"
)

type Spider interface {
	Debug() *spider.Spider
	BaseUrl(base_url string) *spider.Spider
	SetHour(hour int) *spider.Spider
	SetType(type_id int) *spider.Spider

	GetCateList() ([]movie_spider.ClassList, error)
	GetList(page int, limit ...int) (movie_spider.MovieResponse, error)
	Search(keyword string, page int, limit ...int) (movie_spider.MovieResponse, error)
	Detail(ids string) (movie_spider.MovieResponse, error)
	Ping() bool

	// 解析
	Parse() *spider.Parse

	// 下载
	Download() *spider.Download
}
