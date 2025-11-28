package spider

import (
	"fmt"
	"log"
	"net/url"
	"path/filepath"
	"sync"

	"github.com/goravel/framework/errors"
	"github.com/goravel/framework/facades"
	"github.com/orangbus/spider/pkg/downloader/dl"
	"github.com/spf13/cast"
)

type Download struct {
	prefix_url string
	proxy_url  string
	thread     int
	save_path  string
}

func NewDownload() *Download {
	return &Download{
		thread:    facades.Config().GetInt("spider.thread", 30),
		save_path: facades.Config().GetString("spider.path", "./download"),
	}
}

//	func (d *Download) SetProxyUrl(proxy_url string) *Download {
//		d.proxy_url = proxy_url
//		return d
//	}
func (d *Download) SetPrefixUrl(prefix_url string) *Download {
	d.prefix_url = prefix_url
	return d
}

func (d *Download) Start(name, m3u8_url string) (<-chan dl.Progress, error) {
	abs, err := filepath.Abs(d.save_path)
	if err != nil {
		return nil, err
	}
	task, err := dl.NewTask(abs, m3u8_url)
	if err != nil {
		return nil, err
	}
	ch := make(chan dl.Progress, 30)
	go func() {
		err = task.Start(name, d.thread, ch)
		if err != nil {
			log.Printf("下载错误：%s", err.Error())
		}
	}()
	return ch, err
}
func (d *Download) GenerateFile(fileName, api_url string, sensoryList ...[]string) (string, error) {
	file_path := facades.Storage().Path(fileName)
	if facades.Storage().Exists(file_path) {
		if err := facades.Storage().Delete(file_path); err != nil {
			return "", err
		}
	}
	log.Print("file_path:", file_path)

	sensorys := []string{}
	if len(sensoryList) > 0 {
		sensorys = sensoryList[0]
	}

	reqUrl, err := nextPageUrl(api_url, 1)
	if err != nil {
		return "", err
	}

	content, pageCount, err := getUrlData(fmt.Sprintf("%s%s", d.prefix_url, reqUrl))
	if err != nil {
		return "", err
	}

	ch := make(chan string, 2)
	go func() {
		if pageCount > 1 {
			var wg sync.WaitGroup
			for i := 1; i < pageCount; i++ {
				wg.Add(1)
				go func(page int) {
					defer wg.Done()
					q, _ := nextPageUrl(api_url, i+1)

					content2, _, err2 := getUrlData(fmt.Sprintf("%s%s", d.prefix_url, q), sensorys)
					if err2 != nil {
						ch <- ""
					} else {
						ch <- content2
					}
				}(i)
			}
			wg.Wait() // 确保所有goroutine完成后再关闭通道
		}
		close(ch)
	}()

	for data := range ch {
		content += data
	}

	// 创建文件
	if err := facades.Storage().Put(fileName, content); err != nil {
		return "", err
	}
	return file_path, nil
}

func nextPageUrl(api_url string, page int) (string, error) {
	u, err := url.Parse(api_url)
	if err != nil {
		return "", err
	}
	param := u.Query()
	param.Set("pg", cast.ToString(page))
	api_url = fmt.Sprintf("%s://%s%s?%s", u.Scheme, u.Host, u.Path, param.Encode())
	return api_url, err
}

func getUrlData(api_url string, sensoryList ...[]string) (string, int, error) {
	sensorys := []string{}
	if len(sensoryList) > 0 {
		sensorys = sensoryList[0]
	}
	content := ""
	var spider = NewSpider()
	movieResponse, err := spider.Get(api_url)
	if err != nil {
		return "", 0, errors.New(fmt.Sprintf("接口请求错误：%s", err.Error()))
	}
	for _, item := range movieResponse.List {
		urlItems := spider.Parse().Url(item.VodPlayNote, item.VodPlayFrom, item.VodPlayURL)
		total := len(urlItems)
		for _, urlItem := range urlItems {
			if urlItem.Url == "" {
				continue
			}
			if checkHasSensory(sensorys, item.VodName) {
				continue
			}
			if total > 1 {
				content += fmt.Sprintf("%s %s-%s.mp4\n", urlItem.Url, item.VodName, urlItem.Name)
			} else {
				content += fmt.Sprintf("%s %s.mp4\n", urlItem.Url, item.VodName)
			}
		}
	}
	return content, movieResponse.PageCount, err
}

/*
*
是否包含敏感词
*/
func checkHasSensory(list []string, name string) bool {
	for _, item := range list {
		if item == name {
			return true
		}
	}
	return false

}
