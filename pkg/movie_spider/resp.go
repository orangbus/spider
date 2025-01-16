package movie_spider

type MovieList struct {
	VodID            int         `json:"vod_id"`
	TypeID           int         `json:"type_id"`
	TypeID1          int         `json:"type_id_1"`
	GroupID          int         `json:"group_id"`
	VodName          string      `json:"vod_name"`
	VodSub           string      `json:"vod_sub"`
	VodEn            string      `json:"vod_en"`
	VodStatus        int         `json:"vod_status"`
	VodLetter        string      `json:"vod_letter"`
	VodColor         string      `json:"vod_color"`
	VodTag           string      `json:"vod_tag"`
	VodClass         string      `json:"vod_class"`
	VodPic           string      `json:"vod_pic"`
	VodPicThumb      string      `json:"vod_pic_thumb"`
	VodPicSlide      string      `json:"vod_pic_slide"`
	VodPicScreenshot interface{} `json:"vod_pic_screenshot"`
	VodActor         string      `json:"vod_actor"`
	VodDirector      string      `json:"vod_director"`
	VodWriter        string      `json:"vod_writer"`
	VodBehind        string      `json:"vod_behind"`
	VodBlurb         string      `json:"vod_blurb"`
	VodRemarks       string      `json:"vod_remarks"`
	VodPubdate       string      `json:"vod_pubdate"`
	VodTotal         int         `json:"vod_total"`
	VodSerial        string      `json:"vod_serial"`
	VodTV            string      `json:"vod_tv"`
	VodWeekday       string      `json:"vod_weekday"`
	VodArea          string      `json:"vod_area"`
	VodLang          string      `json:"vod_lang"`
	VodYear          string      `json:"vod_year"`
	VodVersion       string      `json:"vod_version"`
	VodState         string      `json:"vod_state"`
	VodAuthor        string      `json:"vod_author"`
	VodJumpurl       string      `json:"vod_jumpurl"`
	VodTpl           string      `json:"vod_tpl"`
	VodTplPlay       string      `json:"vod_tpl_play"`
	VodTplDown       string      `json:"vod_tpl_down"`
	VodIsend         int         `json:"vod_isend"`
	VodLock          int         `json:"vod_lock"`
	VodLevel         int         `json:"vod_level"`
	VodCopyright     int         `json:"vod_copyright"`
	VodPoints        int         `json:"vod_points"`
	VodPointsPlay    int         `json:"vod_points_play"`
	VodPointsDown    int         `json:"vod_points_down"`
	VodHits          int         `json:"vod_hits"`
	VodHitsDay       int         `json:"vod_hits_day"`
	VodHitsWeek      int         `json:"vod_hits_week"`
	VodHitsMonth     int         `json:"vod_hits_month"`
	VodDuration      string      `json:"vod_duration"`
	VodUp            int         `json:"vod_up"`
	VodDown          int         `json:"vod_down"`
	VodScore         string      `json:"vod_score"`
	VodScoreAll      int         `json:"vod_score_all"`
	VodScoreNum      int         `json:"vod_score_num"`
	VodTime          string      `json:"vod_time"`
	VodTimeAdd       int64       `json:"vod_time_add"`
	VodTimeHits      int         `json:"vod_time_hits"`
	VodTimeMake      int         `json:"vod_time_make"`
	VodTrysee        int         `json:"vod_trysee"`
	VodDoubanID      int         `json:"vod_douban_id"`
	VodDoubanScore   string      `json:"vod_douban_score"`
	VodReurl         string      `json:"vod_reurl"`
	VodRelVod        string      `json:"vod_rel_vod"`
	VodRelArt        string      `json:"vod_rel_art"`
	VodPwd           string      `json:"vod_pwd"`
	VodPwdURL        string      `json:"vod_pwd_url"`
	VodPwdPlay       string      `json:"vod_pwd_play"`
	VodPwdPlayURL    string      `json:"vod_pwd_play_url"`
	VodPwdDown       string      `json:"vod_pwd_down"`
	VodPwdDownURL    string      `json:"vod_pwd_down_url"`
	VodContent       string      `json:"vod_content"`
	VodPlayFrom      string      `json:"vod_play_from"`
	VodPlayServer    string      `json:"vod_play_server"`
	VodPlayNote      string      `json:"vod_play_note"`
	VodPlayURL       string      `json:"vod_play_url"`
	VodDownFrom      string      `json:"vod_down_from"`
	VodDownServer    string      `json:"vod_down_server"`
	VodDownNote      string      `json:"vod_down_note"`
	VodDownURL       string      `json:"vod_down_url"`
	VodPlot          int         `json:"vod_plot"`
	VodPlotName      string      `json:"vod_plot_name"`
	VodPlotDetail    string      `json:"vod_plot_detail"`
	TypeName         string      `json:"type_name"`
}

type ClassList struct {
	TypeID   int    `json:"type_id"`
	TypeName string `json:"type_name"`
	TypePid  int    `json:"type_pid"`
}

type MovieResponse struct {
	Page      string `json:"page"`
	Pagecount int    `json:"pagecount"`
	Total     int64  `json:"total"`
	List      []MovieList
	Class     []ClassList
}
