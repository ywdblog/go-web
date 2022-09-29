package dao

import (
	"context"
	"database/sql"
	"fmt"

	skydb "github.com/WiFeng/go-sky/database"
)

const (
	userTableName = "llp_user_%d"
	userTableNums = 10
)

const (
	UserTypeKOL = 4 // KOL
	UserTypeUGC = 5 // 优质UGC
)

type UserTable struct {
	db *sql.DB
}

type UserTableRow struct {
	UID             int64
	Nick            sql.NullString
	Avatar          sql.NullString
	Status          int
	UserUgrade      int
	UserType        int
	CreativeAbility int
}

func NewUserTable(ctx context.Context) (*UserTable, error) {
	db, err := skydb.GetInstance(ctx, databaseName)
	if err != nil {
		return nil, err
	}
	table := &UserTable{
		db,
	}
	return table, nil
}

// 根据文章id 计算表名编号
func (t *UserTable) getTableId(userId int64) int {
	return int(userId % userTableNums)
}

// 根据表名标号计算完整表名
func (t *UserTable) getTableName(tableId int) string {
	return fmt.Sprintf(userTableName, tableId)
}

// 根据文章Id直接计算表名
func (t *UserTable) getTableNameById(userId int64) string {
	tableId := t.getTableId(userId)
	tableName := t.getTableName(tableId)
	return tableName
}

func (t *UserTable) GetUgradeByUID(ctx context.Context, uid int64) (*UserTableRow, error) {
	tableName := t.getTableNameById(uid)
	var row UserTableRow

	err := t.db.QueryRowContext(ctx, fmt.Sprintf("SELECT uid,ugrade FROM %s WHERE uid = ?",
		tableName), uid).Scan(&row.UID, &row.UserUgrade)
	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return &row, nil
}
