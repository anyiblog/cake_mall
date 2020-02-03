package conf

import (
	"fmt"
	"os"
)

// Init 初始化配置项
func Init() {
	// 连接数据库
	MysqlDsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=%s&timeout=%s",
		os.Getenv("MYSQL_USER_NAME"),
		os.Getenv("MYSQL_USER_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_HOST_PORT"),
		os.Getenv("MYSQL_DB"),
		os.Getenv("MYSQL_Loc"),
		os.Getenv("MYSQL_TIMEOUT"),
	)
	Database(MysqlDsn)
}
