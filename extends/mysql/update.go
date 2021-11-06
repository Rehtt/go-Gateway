package mysql

type update struct {
	table interface{}
}

func DBUpdate(table interface{}) *update {
	return &update{table}
}

// Save 更新数据库
func (u *update) Save() error {
	return DB.Self.Save(u.table).Error
}

// Updates 更新局部
func (u *update) Updates(where interface{}, data interface{}) error {
	return DB.Self.Model(u.table).Where(where).Updates(data).Error
}
