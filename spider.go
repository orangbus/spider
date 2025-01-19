package spider

import (
	"encoding/json"
	"fmt"
	"github.com/orangbus/spider/pkg/movie_spider"
	"github.com/spf13/cast"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Spider struct {
	debug   bool
	baseUrl string
	page    int
	limit   int
	hour    int
	tp      int
	ac      string
	keyword string
	ids     string
}

func NewSpider() *Spider {
	return &Spider{page: 1, ac: "list"}
}
func (s *Spider) SetAcVideoList() *Spider {
	s.ac = "videolist"
	return s
}

func (s *Spider) BaseUrl(base_url string) *Spider {
	s.baseUrl = base_url
	return s
}

func (s *Spider) SetHour(hour int) *Spider {
	s.hour = hour
	return s
}
func (s *Spider) SetType(t int) *Spider {
	s.tp = t
	return s
}
func (s *Spider) Debug() *Spider {
	s.debug = true
	return s
}

func (s *Spider) get() (movie_spider.MovieResponse, error) {
	param := url.Values{}
	var data movie_spider.MovieResponse

	param.Set("ac", s.ac)
	if s.page > 0 {
		param.Set("pg", cast.ToString(s.page))
	}
	if s.hour > 0 {
		param.Set("h", cast.ToString(s.hour))
	}
	if s.tp > 0 {
		param.Set("t", cast.ToString(s.tp))
	}
	if s.ids != "" {
		param.Set("ids", s.ids)
	}
	if s.keyword != "" {
		param.Set("wd", s.keyword)
	}
	api_url := fmt.Sprintf("%s?%s", s.baseUrl, param.Encode())
	if s.debug {
		log.Printf("请求地址:%s", api_url)
	}
	response, err := http.Get(api_url)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	content, err := io.ReadAll(response.Body)
	if err != nil {
		return data, err
	}

	if err := json.Unmarshal(content, &data); err != nil {
		return data, err
	}
	return data, nil
}

func (s *Spider) GetCateList() ([]movie_spider.ClassList, error) {
	resp, err := s.get()
	if err != nil {
		return []movie_spider.ClassList{}, err
	}
	return resp.Class, nil
}

func (s *Spider) GetList(page int, limit ...int) (movie_spider.MovieResponse, error) {
	if page <= 0 {
		page = 1
	}
	if len(limit) > 0 {
		s.limit = limit[0]
	}
	s.page = page
	return s.get()
}

func (s *Spider) Search(keyword string, page int, limit ...int) (movie_spider.MovieResponse, error) {
	s.keyword = keyword
	s.page = page
	if len(limit) > 0 {
		s.limit = limit[0]
	}
	return s.get()
}

func (s *Spider) Detail(ids string) (movie_spider.MovieResponse, error) {
	s.ids = ids
	s.ac = "videolist"
	return s.get()
}

func (s *Spider) Ping() bool {
	return true
}

func (s *Spider) Parse() *Parse {
	return NewParse()
}

func (s *Spider) Download() *Download {
	return NewDownload()
}
