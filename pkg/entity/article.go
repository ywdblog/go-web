package entity

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	// 发布平台
	publishPlatformMap = map[int8]int64{
		1: 1 << 0, // 西五街
		2: 1 << 1, // 微博
		3: 1 << 2, // B站
		4: 1 << 3, // 小红书
		5: 1 << 4, // 抖音
		6: 1 << 5, // 绿洲
		7: 1 << 6, // 快手
		8: 1 << 7, // 测评徽章
	}

	// 淘口令格式
	taoCodeRe = regexp.MustCompile(`\[\#\#(\w+)\$\$(.+?)\$\$(.+?)\#\#\]`)
)

const (
	ArticleTagsTypeTips    = "rec_tips"  // 情感标签
	ArticleTagsTypeTopic   = "rec_topic" // 话题标签
	ArticleTagsTypeComment = "rec_cmts"  // 热评标签
	ArticleTagsTypeJiequ   = "jiequ"     // 街主推荐
)

type Picture struct {
	Source string     `json:"src"`
	Width  int        `json:"w"`
	Height int        `json:"h"`
	Type   string     `json:"t"`
	Tag    []struct{} `json:"tag"`
	RGB    string     `json:"rgb"`
}

type Video struct {
	// Filter   string  `json:"filter"`
	Type     string  `json:"t"`
	Src      string  `json:"src"`
	Duration float64 `json:"duration"`
	Cover    string  `json:"cover"` // 封面图，静态
	Height   int     `json:"h"`
	Width    int     `json:"w"`
}

type Route struct {
	ActionName string `json:"action_name"`
	Page       string `json:"page"`
	Url        string `json:"url,omitempty"`
	Title      string `json:"title,omitempty"`
	Tab        string `json:"tab,omitempty"`
	Index      int64  `json:"index,omitempty"`
	StrId      string `json:"strId,omitempty"`
	NumId      int64  `json:"numId,omitempty"`
}

// 秒针监测参数
type MZTrackParams struct {
	AdId         int64  `json:"id"`
	IOSExp       string `json:"ios_exp"`
	IOSClick     string `json:"ios_click"`
	AndroidExp   string `json:"android_exp"`
	AndroidClick string `json:"android_click"`
	H5Exp        string `json:"h5_exp"`
	H5Click      string `json:"h5_click"`
}

// 神策推荐埋点字段
type LogPoints struct {
	SectionId        string `json:"section_id"`
	ExpId            string `json:"exp_id"`
	ItemType         string `json:"item_type"`
	ItemId           string `json:"item_id"`
	IsRecommendation bool   `json:"is_recommendation"`
	StrategyId       string `json:"strategy_id"`
	RetrieveId       string `json:"retrieve_id"`
	Weight           int    `json:"weight"`
	LogId            string `json:"log_id"`

	CardId    string `json:"card_id"` // 管理后台card id
	CardType  string `json:"card_type"`
	CardTitle string `json:"card_title"`
}

// 不喜欢类型字段
type DislikeOpts struct {
	Type      string `json:"type"`
	TypeId    int64  `json:"type_id"`
	TypeTitle string `json:"type_title"`
}

// 文章标签
type ArticleTags struct {
	TagType  string        `json:"type"`
	TipsData []interface{} `json:"tipsData"`
}

// 文章标签数据
type ArticleTagsData struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// 数字转换为易于阅读的格式
func NumberFormat(number int64) string {
	var vol string
	if number >= 1000000 {
		vol = fmt.Sprintf("%dw", number/10000)
	} else if number >= 10000 {
		vol = fmt.Sprintf("%.1fw", float64(number)/10000)
	} else {
		vol = strconv.FormatInt(number, 10)
	}
	return vol
}

// 日期转换为易于阅读的格式
func DateStrFormat(dateStr string) (string, error) {

	date, err := time.ParseInLocation(DatetimeFormat, dateStr, time.Local)
	if err != nil {
		return dateStr, err
	}

	now := time.Now()
	if now.Unix()-date.Unix() < 60 {
		return "刚刚", nil
	}

	dayDiff := now.Day() - date.Day()
	monthDiff := now.Month() - date.Month()
	yearDiff := now.Year() - date.Year()

	simpleLayout := date.Format("15:04")
	simpleLayout2 := date.Format("2006-01-02 15:04")
	if dayDiff == 0 && monthDiff == 0 && yearDiff == 0 {
		return fmt.Sprintf("今天 %s", simpleLayout), nil
	} else if dayDiff == 1 && monthDiff == 0 && yearDiff == 0 {
		return fmt.Sprintf("昨天 %s", simpleLayout), nil
	}

	return simpleLayout2, nil
}

func BuildPicture(URL string) Picture {
	picture := Picture{
		Source: URL,
		Width:  800,
		Height: 800,
		Type:   "image/jpeg",
		Tag:    []struct{}{},
		RGB:    "#000000",
	}

	slices := strings.Split(URL, "|")
	if len(slices) < 2 {
		return picture
	}

	picture.Source = slices[0]
	attrStr := slices[1]
	attrSlices := strings.Split(attrStr, "x")
	if len(attrSlices) >= 2 {
		width, _ := strconv.Atoi(attrSlices[0])
		height, _ := strconv.Atoi(attrSlices[1])
		picture.Width = width
		picture.Height = height
	}

	if len(attrSlices) >= 4 {
		imgType := attrSlices[2]
		picture.Type = fmt.Sprintf("image/%s", imgType)
	}

	return picture
}

// 转换发布平台
func TransPublishPlatform(p int64) []int8 {
	var plats []int8
	for k, v := range publishPlatformMap {
		if (p & v) == v {
			plats = append(plats, k)
		}
	}
	return plats
}

type TopicItem struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	SubjectType int    `json:"subject_type"`
	IsGood      int    `json:"is_good"`
	Pid         int    `json:"p_id"`
	Logo        string `json:"logo"`
	ArticleNums int    `json:"article_nums"`
}

type ArticleInfo struct {
	ArticleId   int64         `json:"aid"`
	Uid         int64         `json:"uid"`
	Type        int           `json:"type"`
	Title       string        `json:"title"`
	Desc        string        `json:"desc"`
	AddTime     string        `json:"add_time"`
	PublishTime string        `json:"publish_time"`
	UpdateTime  string        `json:"update_time"`
	Topic       []TopicItem   `json:"topic"`
	Banner      []Picture     `json:"banner"`
	Video       *Video        `json:"video"`
	Content     string        `json:"content"`
	PictURL     *Picture      `json:"pict_url"`
	IsRecommend int           `json:"is_recommend"`
	IsRec       bool          `json:"isRec"`
	ArticleTags []ArticleTags `json:"artTags"`
	IsKOL       int           `json:"is_kol"`
}

type ArticleInfoMGetRequest struct {
	BaseRequest

	ArticleIds   []int64 `json:"aids"`
	ForceNoCache bool    `json:"force_no_cache"`
}

type ArticleInfoMGetRespData struct {
	Infos []ArticleInfo `json:"infos"`
}

type ArticleInfoMGetResponse struct {
	BaseResponse

	Data ArticleInfoMGetRespData `json:"data"`
}

// 替换特殊内容
func replaceSpecialText(text string) string {
	replaceText := taoCodeRe.ReplaceAllString(text, "[##${1}$$...$$${3}##]")
	return replaceText
}
