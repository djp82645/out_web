package pojo

// api_info 实体，对应表api_info
type ApiInfo struct {
	ID
	Sn         string    `json:"sn" gorm:"type:varchar(100);not null;index:index_name,class:FULLTEXT,option:WITH PARSER ngram;comment:编号"`
	Api    string    `json:"api" gorm:"type:text;not null;comment:接口地址"`
	Status    int64     `json:"status" gorm:"type:smallint(5);not null;default:-1;comment:状态 -1，初始状态 0，已取号 1，成功 2，失败 3，异常"`
	CreatedAt    LocalTime `json:"created_at"`
	UpdatedAt    LocalTime `json:"updated_at"`
	SoftDeletes
}

// TableName 指定表名
func (ApiInfo) TableName() string {
	return "api_info"
}
