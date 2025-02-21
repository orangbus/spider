package spider

import (
	"fmt"
	"github.com/goravel/framework/errors"
	"github.com/goravel/framework/facades"
	"github.com/spf13/cast"
	"sync"

	//facades2 "github.com/orangbus/spider/facades"
	"github.com/orangbus/spider/pkg/downloader/dl"
	"log"
	"net/url"
	"path/filepath"
)

type Download struct {
	thread    int
	save_path string
}

func NewDownload() *Download {
	return &Download{
		thread:    facades.Config().GetInt("spider.thread", 30),
		save_path: facades.Config().GetString("spider.path", "./download"),
	}
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
func (d *Download) GenerateFile(fileName, api_url string) (string, error) {
	file_path := facades.Storage().Path(fileName)
	if facades.Storage().Exists(file_path) {
		if err := facades.Storage().Delete(file_path); err != nil {
			return "", err
		}
	}

	reqUrl, err := nextPageUrl(api_url, 1)
	if err != nil {
		return "", err
	}

	content, pageCount, err := getUrlData(reqUrl)
	if err != nil {
		return "", err
	}

	var wg sync.WaitGroup
	ch := make(chan string, 2)
	if pageCount > 1 {
		for i := 1; i < pageCount; i++ {
			wg.Add(1)
			go func(page int) {
				defer wg.Done()
				q, _ := nextPageUrl(api_url, i+1)

				content2, _, err2 := getUrlData(q)
				if err2 != nil {
					ch <- ""
				} else {
					ch <- content2
				}
			}(i)
		}
	}

	// 接受channel消息
	go func() {
		wg.Wait() // 确保所有goroutine完成后再关闭通道
		close(ch)
	}()

	for i := 1; i <= pageCount; i++ {
		data := <-ch
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

func getUrlData(api_url string) (string, int, error) {
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
			content += fmt.Sprintf("%s %s.mp4\n", urlItem.Url, item.VodName)
			if total > 1 {
				content += fmt.Sprintf("%s %s-%s.mp4\n", urlItem.Url, item.VodName, urlItem.Name)
			}
		}
	}
	return content, movieResponse.PageCount, err
}
