package test

import (
	"github.com/orangbus/spider/facades"
	"sync"
	"testing"
)

func TestDownload(t *testing.T) {
	name := "斗罗大陆"
	m3u8 := "https://hn.bfvvs.com/play/penWV2Ea/index.m3u8"
	msgs, err := facades.Spider().Download().Start(name, m3u8)
	if err != nil {
		t.Log(err)
		return
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for p := range msgs {
			if p.Stop {
				wg.Done()
			}
			t.Logf("总：%d,进度:%d", p.Total, p.Finish)
		}
	}()
	wg.Wait()
	t.Log("over")
}
