package task

const (
	tableArticlePrefix  = "llp_article_" // 文章
	tableCommentPrefix  = "llp_comment_" // 评论
	tableUserPrefix     = "llp_user_"    // 用户
	articleNumTableName = "llp_article_num"
)

const (
	binlogTypeInsert = "INSERT"
	binlogTypeUpdate = "UPDATE"
	binlogTypeDelete = "DELETE"
)

type Binlog struct {
	Data []map[string]string `json:"data"`
	Old  []map[string]string `json:"old"`

	ID int64 `json:"id"`
	ES int64 `json:"es"`
	TS int64 `json:"ts"`

	IsDel bool   `json:"isDel"`
	Type  string `json:"type"`

	Table    string `json:"table"`
	Database string `json:"database"`
}
