package spider

import "strings"

type Parse struct {
}

type MovieUrlItem struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func NewParse() *Parse {
	return &Parse{}
}

func (p *Parse) Url(vodPlayNote, vodPlayFrom, vodPlayURL string) []MovieUrlItem {
	// 是否有分隔符
	var cate []string       // 分类列表
	var cateList []string   // 分类视频列表
	var list []MovieUrlItem // 返回视频链接地址
	if vodPlayNote != "" {
		cate = strings.Split(vodPlayFrom, vodPlayNote)
		cateList = strings.Split(vodPlayURL, vodPlayNote)

		for i := range cate {
			listItem := parseUrl(cateList[i])
			if len(listItem) > 0 {
				for _, item := range listItem {
					list = append(list, item)
				}
			}
		}
		return list
	}
	return parseUrl(vodPlayURL)
}

/*
*
vodPlayURL: 视频地址
cateList: 分类列表
*/
func parseUrl(vodPlayURL string) []MovieUrlItem {
	var list []MovieUrlItem // 返回视频链接地址
	if strings.Contains(vodPlayURL, "#") {
		cateList := strings.Split(vodPlayURL, "#")
		for _, urlStr := range cateList {
			urlList := strings.Split(urlStr, "$")
			var movieUrlItem MovieUrlItem
			if len(urlList) >= 2 {
				movieUrlItem.Name = urlList[0]
				movieUrlItem.Url = urlList[1]
			}
			list = append(list, movieUrlItem)
		}
		return list
	}
	// 只有单个地址的情况
	if strings.Contains(vodPlayURL, "$") {
		item := strings.Split(vodPlayURL, "$")
		if len(item) >= 2 {
			list = append(list, MovieUrlItem{Name: item[0], Url: item[1]})
		}
		return list
	}
	// 如果直接就是播放地址
	list = append(list, MovieUrlItem{Name: "1", Url: vodPlayURL})
	return list
}
