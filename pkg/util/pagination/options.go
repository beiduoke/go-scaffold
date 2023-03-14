package pagination

import (
	"encoding/json"
	"log"
)

type Option func(*Pagination)

type Pagination struct {
	// 当前页
	Page int32
	// 每一页的行数
	PageSize int32
	// 查询参数
	Query map[string]string
	// 排序
	OrderBy map[string]bool
	// 是否不分页
	Nopaging bool
}

func (p *Pagination) queryFormat() {
	if query, ok := p.Query["query"]; ok {
		err := json.Unmarshal([]byte(query), &p.Query)
		if err != nil {
			log.Println(err)
		}
	}
}

func (p *Pagination) orderByFormat() {
	// if order, ok := p.OrderBy["orderBy"]; ok {
	// 	err := json.Unmarshal([]byte(order), &p.OrderBy)
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// }
}

func NewPagination(opts ...Option) *Pagination {
	p := Pagination{
		Page:     PAGE,
		PageSize: PAGE_SIZE,
	}
	for _, o := range opts {
		o(&p)
	}
	p.orderByFormat()
	p.queryFormat()
	return &p
}

func WithNopaging() Option {
	return func(p *Pagination) {
		p.Nopaging = true
	}
}

func WithPage(page int32) Option {
	return func(p *Pagination) {
		if page == 0 {
			page = PAGE
		}
		p.Page = page
	}
}

func WithPageSize(size int32) Option {
	return func(p *Pagination) {
		if size == 0 {
			size = PAGE_SIZE
		}
		p.PageSize = size
	}
}
