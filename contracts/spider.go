package contracts

import (
	"github.com/orangbus/spider"
	"github.com/orangbus/spider/pkg/movie_spider"
)

type Spider interface {
	Debug() *spider.Spider
	BaseUrl(base_url string, proxy_url ...string) *spider.Spider
	SetHour(hour int) *spider.Spider
	SetType(type_id int) *spider.Spider
	SetKeyword(keyword string) *spider.Spider

	GetCateList() ([]movie_spider.ClassList, error)
	GetList(page int, limit ...int) (movie_spider.MovieResponse, error)
	Get(api_url string) (movie_spider.MovieResponse, error)
	Search(page int, limit ...int) (movie_spider.MovieResponse, error)
	Detail(ids string) (movie_spider.MovieItem, error)
	Ping() bool                 // ping一下地址是否正常访问
	Parse() *spider.Parse       // 解析
	Download() *spider.Download // 下载
	Live() *spider.Live         // 直播
}
