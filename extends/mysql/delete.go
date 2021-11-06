package mysql

type Delete struct {
	table interface{}
	where interface{}
}

func DBDelete(table, where interface{}) *Delete {
	return &Delete{table, where}
}

func (d *Delete) Delete() error {

	return DB.Self.Where(d.where).Delete(d.table).Error
}

// 永久删除
func (d *Delete) DeleteU() error {
	return DB.Self.Unscoped().Where(d.where).Delete(d.table).Error
}

// 恢复软删除
func (d *Delete) RecoverSoftDelete() error {
	return DB.Self.Unscoped().Model(d.table).Where(d.where).Update("deleted_at", nil).Error
}
