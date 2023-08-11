package entity

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// 创建审核实体
type CreationAuditedEntity struct {
	Entity
	//创建时间
	CreationTime time.Time `json:"creationTime" bson:"creationTime" gorm:"column:CreationTime"`
	//创建人
	CreatorId *uuid.UUID `json:"creatorId" bson:"creatorId" gorm:"column:CreatorId"`
}

//创建时设置对象的基本信息
func (entity *CreationAuditedEntity) BeforeCreate() {
	//调用基类的函数
	entity.Entity.BeforeCreate()
	if entity.CreationTime.IsZero() {
		entity.CreationTime = time.Now()
	}
}
