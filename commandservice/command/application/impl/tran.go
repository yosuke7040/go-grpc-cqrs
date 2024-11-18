package impl

import (
	"context"
	"database/sql"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/yosuke7040/commandservice/infra/sqlboiler/handler"
	"log"
)

type transaction struct{}

func (inc *transaction) begin(ctx context.Context) (*sql.Tx, error) {
	// トランザクションを開始する
	tran, err := boil.BeginTx(ctx, nil)
	if err != nil {
		return nil, handler.DBErrHandler(err)
	}
	return tran, nil
}

func (ins *transaction) complete(tran *sql.Tx, err error) error {
	if err != nil {
		if e := tran.Rollback(); e != nil {
			return handler.DBErrHandler(err)
		} else {
			log.Println("トランザクションをロールバックしました。")
		}
	} else {
		if e := tran.Commit(); e != nil {
			return handler.DBErrHandler(err)
		} else {
			log.Println("トランザクションをコミットしました。")
		}
	}
	return nil
}
