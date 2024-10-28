package gormc

type Config struct {
	DSN                    string
	Debug                  bool
	ConnMaxLifetime        int  `json:",optional"`
	MaxIdleConns           int  `json:",default=1"`    // 空闲中的最大连接数
	MaxOpenConns           int  `json:",default=10"`   // 打开到数据库的最大连接数
	SkipDefaultTransaction bool `json:",default=true"` //跳过默认事务
	SingularTable          bool `json:",default=true"` //是否使用单数表名
}
