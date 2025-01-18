package spider

import (
	"github.com/goravel/framework/facades"
	"github.com/orangbus/spider/pkg/downloader/dl"
	"log"
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
		defer close(ch)

		err = task.Start(d.thread, ch)
		if err != nil {
			log.Printf("下载错误：%s", err.Error())
		}
	}()
	return ch, err
}
