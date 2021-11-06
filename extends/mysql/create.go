package mysql

// DBCreate 创建一条数据
func DBCreate(tableAndData interface{}) error {
	return DB.Self.Create(tableAndData).Error
}
