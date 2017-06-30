package infrastructure

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

var _mysqlHandlerLst []*MysqlHandler = []*MysqlHandler{nil}

//实现DbHandler接口
type MysqlHandler struct {
	Conn *sql.DB
}

//类似于单例模式，全局只有一个实例
func NewMysqlHandler() *MysqlHandler {
	if _mysqlHandlerLst[0] != nil {
		return _mysqlHandlerLst[0]
	}

	conn, _ := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/Test?charset=utf8")
	handler := new(MysqlHandler)
	handler.Conn = conn
	_mysqlHandlerLst[0] = handler
	return handler
}

func (handler *MysqlHandler) Execute(statement string) {
	handler.Conn.Exec(statement)
}

func (handler *MysqlHandler) Query(statement string) *sql.Rows {
	rows, err := handler.Conn.Query(statement)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return rows
}
