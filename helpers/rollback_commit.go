package helpers

import (
	"database/sql"
)

func RollbackCommit(tx *sql.Tx) {
	defer func() {
		err := recover()
		if err != nil {
			errRollback := tx.Rollback()
			if errRollback != nil {
				return
			}
		} else {
			errCommit := tx.Commit()
			if errCommit != nil {
				return
			}
		}
	}()
}
