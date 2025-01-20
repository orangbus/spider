package spider

import (
	"encoding/json"
	"fmt"
	"github.com/goravel/framework/facades"
	"io"
	"net/http"
)

type Live struct {
	api_url string // 接口地址
}

type ZhuboResponse struct {
	Zhubo []ZhuboItem `json:"zhubo"`
}

type ZhuboItem struct {
	Title   string `json:"title"`
	Img     string `json:"img"`
	Address string `json:"address"`
}

type PintaiResponse struct {
	PingTai []PintaiItem `json:"pingtai"`
}

type PintaiItem struct {
	Title   string `json:"title"`
	Xinimg  string `json:"img"`
	Number  any    `json:"number"`
	Address string `json:"address"`
}

func NewLive() *Live {
	return &Live{
		api_url: facades.Config().GetString("spider.live_api"),
	}
}

func get(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, err
	}
	readAll, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return readAll, nil
}
func (l *Live) GetPinTai() ([]PintaiItem, error) {
	res, err := get(fmt.Sprintf("%s/mf/json.txt", l.api_url))
	if err != nil {
		return nil, err
	}
	var data PintaiResponse
	if err := json.Unmarshal(res, &data); err != nil {
		return nil, err
	}
	return data.PingTai, nil
}

func (l *Live) GetZhubo(pintai_address string) ([]ZhuboItem, error) {
	res, err := get(fmt.Sprintf("%s/mf/%s", l.api_url, pintai_address))
	if err != nil {
		return nil, err
	}
	var data ZhuboResponse
	if err := json.Unmarshal(res, &data); err != nil {
		return nil, err
	}
	return data.Zhubo, nil
}
