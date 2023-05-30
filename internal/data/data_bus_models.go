package data

func NewBusModelMigrate() []interface{} {
	return []interface{}{
		&BusDemo{},
	}
}

// func (dm *DomainModel) BeforeUpdate(tx *gorm.DB) (err error) {
// 	// tx.Statement.Context
// 	// if dm.Updater == "admin" {
// 	// 	return errors.New("admin user not allowed to update")
// 	// }
// 	fmt.Println("更新前")
// 	return
// }

// func (dm *DomainModel) BeforeCreate(tx *gorm.DB) (err error) {
// 	// if dm.Creator == "admin" {
// 	// 	return errors.New("invalid role")
// 	// }
// 	fmt.Println("创建前")
// 	return
// }

// SysDemo 领域
type BusDemo struct {
	DomainModel
	Name    string `gorm:"type:varchar(255);column:name;not null;comment:名称;"`
	Sort    int32  `gorm:"type:int(10);column:sort;not null;default:100;comment:排序"`
	State   int32  `gorm:"type:tinyint(1);column:state;not null;default:1;index;comment:状态 0 未指定  1 启用 2 停用;"`
	Remarks string `gorm:"type:varchar(255);column:remarks;not null;default:'';comment:备注;"`
}
