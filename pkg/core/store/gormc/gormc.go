package gormc

import (
	"github.com/spf13/cast"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

// NewEngine 创建db
func NewEngine(c Config, dst ...any) *gorm.DB {
	inner, err := gorm.Open(mysql.New(mysql.Config{
		DSN: c.DSN,
	}), &gorm.Config{
		SkipDefaultTransaction: c.SkipDefaultTransaction,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: c.SingularTable,
		},
	})
	if err != nil {
		log.Fatalf("init db err:%+v", err)
	}

	// 设置默认连接配置
	db, err := inner.DB()
	if err != nil {
		log.Fatalf("set db pool err:%+v", err)
	}

	if c.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(time.Duration(cast.ToInt64(c.ConnMaxLifetime)) * time.Second)
	}

	if c.Debug {
		inner = inner.Debug()
	}

	if len(dst) > 0 {
		WithAutoMigrates(inner, dst...)
	}

	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetMaxOpenConns(c.MaxOpenConns)

	return inner
}

// WithAutoMigrates 数据库迁移
func WithAutoMigrates(inner *gorm.DB, dst ...any) {
	_ = inner.AutoMigrate(dst...)
	inner.Set("gorm:table_options", "CHARSET=utf8mb4")
	inner.Set("gorm:table_options", "collation=utf8mb4_general_ci")

	for _, d := range dst {
		if !inner.Migrator().HasTable(&d) {
			_ = inner.Migrator().CreateTable(&d)
		}
	}
}
