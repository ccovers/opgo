package dbutil

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"lib/common/clog"
)

type Task func(db *gorm.DB) error

func DoTrans(db *gorm.DB, tasks ...Task) error {
	new_db := db.Begin()
	if err := new_db.Error; err != nil {
		clog.Logger.Error("开启事务失败: %s", err.Error())
		return err
	}

	for _, task := range tasks {
		if err := task(new_db); err != nil {
			if rerr := new_db.Rollback().Error; rerr != nil {
				clog.Logger.Error("回滚事务失败: %s", rerr.Error())
				return rerr
			}
			return err
		}
	}

	if err := new_db.Commit().Error; err != nil {
		new_db.Rollback()
		clog.Logger.Error("提交事务失败: %s", err.Error())
		return err
	}
	return nil
}
