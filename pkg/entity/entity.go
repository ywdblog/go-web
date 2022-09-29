package entity

const (
	DayFormat      = "20060102" // 每日日期格式
	DatetimeFormat = "2006-01-02 15:04:05"
)

const (
	KafkaNameSync = "article_kafka_sync"
	KafkaNameLogo = "article_kafka_logo"
)

const (
	ArticleInfoMgetURI = "/article_service/article/mget" // 文章信息获取

)

const (
	CodeSucess = 0
	MsgSuccess = ""
)

type BaseRequest struct {
	AppId  string `json:"app_id"`
	AppKey string `json:"app_key"`
}

type BaseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
