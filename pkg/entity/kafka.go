package entity

const (
	KafkaEventNameArticleTagsMachine = "label_article_tags_machine" // 文章标签机器标识
	KafkaEventNameIsTopicSelected    = "is_topic_selected"          // 文章标签话题优选
)

type KafkaBaseMsg struct {
	Time  string `json:"time"`
	Event string `json:"event"`
}

type KafkaArticleTagsMachineMsgData struct {
	ArticleId   int64 `json:"aid"`
	ChannelList []int `json:"channel"`
	TagList     []int `json:"tags"`
}

type KafkaArticleTagsMachineMsg struct {
	KafkaBaseMsg
	Data KafkaArticleTagsMachineMsgData `json:"data"`
}

type KafkaArticleTagsTaskMsgData struct {
	ArticleId int64 `json:"aid"`
}

type KafkaArticleTagsTaskMsg struct {
	KafkaBaseMsg
	Data KafkaArticleTagsMachineMsgData `json:"data"`
}
