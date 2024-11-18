package handler

import (
	"database/sql"
	"fmt"
	"github.com/BurntSushi/toml"
	_ "github.com/go-sql-driver/mysql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"os"
	"time"
)

type DBConfig struct {
	Dbname string `toml:"dbname"` //	データベース名
	Host   string `toml:"host"`   //	ホスト名
	Port   int64  `toml:"port"`   //	ポート番号
	User   string `toml:"user"`   //	ユーザー名
	Pass   string `toml:"pass"`   //	パスワード
}

func tomlRead() (*DBConfig, error) {
	path := os.Getenv("DATABASE_TOML_PATH")
	if path == "" {
		path = "infra/sqlboiler/handler/database.toml"
	}

	m := map[string]DBConfig{}
	_, err := toml.DecodeFile(path, &m)
	if err != nil {
		return nil, err
	}

	config := m["mysql"]
	return &config, nil
}

func DBConnect() error {
	config, err := tomlRead()
	if err != nil {
		return DBErrHandler(err)
	}

	rdbms := "mysql"
	connect_str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.User, config.Pass, config.Host, config.Port, config.Dbname)

	conn, err := sql.Open(rdbms, connect_str)
	if err != nil {
		return DBErrHandler(err)
	}

	if err = conn.Ping(); err != nil {
		return DBErrHandler(err)
	}

	MAX_IDLE_CONNS := 10                  // 初期接続数
	MAX_OPEN_CONNS := 100                 // 最大接続数
	CONN_MAX_LIFTIME := 300 * time.Second // 最大生存期間

	conn.SetMaxIdleConns(MAX_IDLE_CONNS)
	conn.SetMaxOpenConns(MAX_OPEN_CONNS)
	conn.SetConnMaxLifetime(CONN_MAX_LIFTIME)

	boil.SetDB(conn)
	boil.DebugMode = true
	return nil
}
