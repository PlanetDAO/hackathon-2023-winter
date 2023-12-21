// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// PromptlikeDao is the data access object for table promptlike.
type PromptlikeDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns PromptlikeColumns // columns contains all the column names of Table for convenient usage.
}

// PromptlikeColumns defines and stores column names for table promptlike.
type PromptlikeColumns struct {
	Id         string // pk
	Uid        string // User ID
	Pid        string // Prompt ID
	Like       string // User like(1) or dislike(-1) prompt
	CreateTime string // Created Time
	UpdateTime string // Updated Time
}

// promptlikeColumns holds the columns for table promptlike.
var promptlikeColumns = PromptlikeColumns{
	Id:         "id",
	Uid:        "uid",
	Pid:        "pid",
	Like:       "like",
	CreateTime: "create_time",
	UpdateTime: "update_time",
}

// NewPromptlikeDao creates and returns a new DAO object for table data access.
func NewPromptlikeDao() *PromptlikeDao {
	return &PromptlikeDao{
		group:   "default",
		table:   "promptlike",
		columns: promptlikeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *PromptlikeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *PromptlikeDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *PromptlikeDao) Columns() PromptlikeColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *PromptlikeDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *PromptlikeDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *PromptlikeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}