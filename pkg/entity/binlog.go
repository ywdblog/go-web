package entity

const (
	BinlogTypeInsert = "INSERT"
	BinlogTypeUpdate = "UPDATE"
	BinlogTypeDelete = "DELETE"
)

type FlatMessage struct {
	ID        int64               `json:"id"`
	Database  string              `json:"database"`
	Table     string              `json:"table"`
	PkNames   []string            `json:"pkNames"`
	IsDel     bool                `json:"isDel"`
	Type      string              `json:"type"`
	ES        int64               `json:"es"`
	TS        int64               `json:"ts"`
	Sql       string              `json:"sql"`
	SqlType   map[string]int64    `json:"sqlType"`
	MysqlType map[string]string   `json:"mysqlType"`
	Data      []map[string]string `json:"data"`
	Old       []map[string]string `json:"old"`
}
