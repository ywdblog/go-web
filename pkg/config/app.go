package config

var (
	// GlobalAppConfig ...
	GlobalAppConfig AppConfig
)

// AppConfig ...
type AppConfig struct {
	App struct {
		ConsumerArticleCopyrightLogoTopics []string
		ProducerQueueArticleTopic          string
		ConsumerSyncTopics                 []string
		ConsumerArticleQueueTopics         []string
		ArticleTagsTaskTopic               string
		EsArticleIndexName                 string
	}
}
