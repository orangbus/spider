package test

import (
	"sync"
	"testing"

	"github.com/orangbus/spider/facades"
)

func TestDownload(t *testing.T) {
	name := "1"
	m3u8 := "https://play.gayzyv.com/play/NbW6Jvva/index.m3u8"
	msgs, err := facades.Spider().Download().Start(name, m3u8)
	if err != nil {
		t.Log(err)
		return
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for p := range msgs {
			t.Logf("总：%d,进度:%d", p.Total, p.Finish)
		}
	}()
	wg.Wait()
	t.Log("下载成功")
}

func TestGenerate(t *testing.T) {
	type item struct {
		Name string
		Url  string
	}
	var list = []item{}
	list = append(list, item{
		Name: "gayapi_21.txt",
		Url:  "https://91jpzyw.com/api.php/provide/vod?ac=videolist&t=80",
	})

	for _, v := range list {
		path, err := facades.Spider().Download().SetPrefixUrl("https://spider.orangbus.cn?url=").GenerateFile(v.Name, v.Url)
		if err != nil {
			t.Log(err)
			return
		}
		t.Logf("保存位置:%s", path)
	}
}
