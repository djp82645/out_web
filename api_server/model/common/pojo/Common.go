package pojo

import (
	"gorm.io/gorm"
)

// 自增ID主键
type ID struct {
	ID int64 `json:"id" gorm:"primaryKey"`
}

// // 创建、更新时间
// type Timestamps struct {
//     CreatedAt LocalTime
//     UpdatedAt LocalTime
// }

// 软删除
type SoftDeletes struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
