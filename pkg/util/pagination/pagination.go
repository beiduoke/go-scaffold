package pagination

import "github.com/iancoleman/strcase"

const (
	PAGE_NUM      = 1
	PAGE_SIZE     = 10
	MAX_PAGE_SIZE = 1000
)

func GetPageOffset(pageNum, pageSize int32) int {
	return int((pageNum - 1) * pageSize)
}

type Condition struct {
	Query string
	Args  []interface{}
}

type Order struct {
	Column string
	Desc   bool
}

type PaginationHandler interface {
	GetPageNum() int32
	GetPageSize() int32
	GetPageOffset() int
	GetConditions() []Condition
	GetOrders() []Order
	GetNopaging() bool
	SetCondition(...Condition)
	SetOrder(...Order)
	Reset()
}

type Pagination struct {
	pageNum, pageSize int32
	nopaging          bool
	conditions        []Condition
	orders            []Order
}

type Option func(*Pagination)

func WithNopaging() Option {
	return func(p *Pagination) {
		p.nopaging = true
	}
}

func WithPageNum(num int32) Option {
	return func(p *Pagination) {
		if num == 0 {
			num = PAGE_NUM
		}
		p.pageNum = num
	}
}

func WithPageSize(size int32) Option {
	return func(p *Pagination) {
		if size == 0 {
			size = PAGE_SIZE
		}
		p.pageSize = size
	}
}

func WithCondition(query string, args ...interface{}) Option {
	return func(p *Pagination) {
		p.SetCondition(Condition{Query: query, Args: args})
	}
}

func WithConditions(condition ...Condition) Option {
	return func(p *Pagination) {
		p.SetCondition(condition...)
	}
}

func WithOrder(column string, desc bool) Option {
	return func(p *Pagination) {
		p.SetOrder(Order{Column: column, Desc: desc})
	}
}

func WithOrders(order ...Order) Option {
	return func(p *Pagination) {
		p.SetOrder(order...)
	}
}

func NewPagination(opts ...Option) PaginationHandler {
	p := Pagination{
		nopaging: false,
		pageNum:  PAGE_NUM,
		pageSize: PAGE_SIZE,
	}
	for _, o := range opts {
		o(&p)
	}
	return &p
}

func (p *Pagination) String() string {
	return "Pagination"
}

func (p *Pagination) Reset() {
	p.nopaging = false
	p.pageNum = PAGE_NUM
	p.pageSize = PAGE_SIZE
	p.conditions = []Condition{}
	p.orders = []Order{}
}

func (p *Pagination) GetPageOffset() int {
	if p.nopaging {
		return 0
	}
	return GetPageOffset(p.pageNum, p.pageSize)
}

func (p *Pagination) GetPageNum() int32 {
	return p.pageNum
}

func (p *Pagination) GetPageSize() int32 {
	if p.nopaging {
		return MAX_PAGE_SIZE
	}
	return p.pageSize
}

func (p *Pagination) GetNopaging() bool {
	return p.nopaging
}

func (p *Pagination) GetConditions() []Condition {
	return p.conditions
}

func (p *Pagination) GetOrders() []Order {
	return p.orders
}

func (p *Pagination) SetCondition(condition ...Condition) {
	for _, v := range condition {
		v.Query = strcase.ToSnakeWithIgnore(v.Query, " ")
		for k, vc := range p.conditions {
			if v.Query == vc.Query {
				p.conditions[k].Args = v.Args
				break
			}
		}
		p.conditions = append(p.conditions, v)
	}
}

func (p *Pagination) SetOrder(order ...Order) {
	for _, v := range order {
		v.Column = strcase.ToSnakeWithIgnore(v.Column, " ")
		for k, vc := range p.orders {
			if v.Column == vc.Column {
				p.orders[k].Desc = v.Desc
				break
			}
		}
		p.orders = append(p.orders, v)
	}
}
