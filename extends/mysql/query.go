package mysql

import (
	"math"
)

type Query struct {
	limit int64
	page  int64
	table interface{}
	where interface{}
	from  int64
	to    int64
}

func DBQuery(table, where interface{}) *Query {
	return &Query{
		table: table,
		where: where,
	}
}

// GetCount 获取总数
func (p *Query) GetCount() (count int) {
	DB.Self.Model(p.table).Where(p.where).Count(&count)
	return
}

// getPageCount 获取页数
func (p *Query) getPageCount() int {
	sum := p.GetCount()
	return int(math.Ceil(float64(sum) / float64(p.limit)))
}

// getContentFromRange 获取指定范围内的数据
func (p *Query) getContentFromRange() {
	DB.Self.Where(p.where).Offset(p.from).Limit(p.to).Find(p.table)
}

// GetContentFromPage  获取指定页面内的数据
// page 为负数时获取全部内容
func (p *Query) GetContentFromPage(page int64) {
	p.page = page
	if page > 0 {
		p.page--
	}
	if p.page < 0 {
		p.from = -1
		p.to = -1
	} else {
		p.from = p.page * p.limit
		p.to = p.limit
	}
	p.getContentFromRange()
}

func (p *Query) SetLimit(limit int64) *Query {
	p.limit = limit
	return p
}

// GetContent 查找（不包括软删除）
func (p *Query) GetContent() error {
	return DB.Self.Where(p.where).Find(p.table).Error
}

// GetALL 查找所有，包括软删除
func (p *Query) GetALL() error {
	return DB.Self.Unscoped().Where(p.where).Find(p.table).Error
}
