package test

import (
	"github.com/orangbus/spider/facades"
	"sync"
	"testing"
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
