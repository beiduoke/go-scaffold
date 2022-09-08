package pagination

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
	SetCondition(condition Condition)
	SetOrder(order Order)
	Reset()
}

type Pagination struct {
	pageNum, pageSize int32
	nopaging          bool
	conditions        []*Condition
	orders            []*Order
}

type Option func(*Pagination)

func WithNopaging() Option {
	return func(p *Pagination) {
		p.nopaging = true
	}
}

func WithPageNum(num int32) Option {
	return func(p *Pagination) {
		p.pageNum = num
	}
}

func WithPageSize(size int32) Option {
	return func(p *Pagination) {
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
		for _, v := range condition {
			p.conditions = append(p.conditions, &v)
		}
	}
}

func WithOrder(column string, desc bool) Option {
	return func(p *Pagination) {
		p.SetOrder(Order{Column: column, Desc: desc})
	}
}

func WithOrders(order ...Order) Option {
	return func(p *Pagination) {
		for _, v := range order {
			p.orders = append(p.orders, &v)
		}
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
	p.conditions = []*Condition{}
	p.orders = []*Order{}
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
	conditions := make([]Condition, 0, len(p.conditions))
	for _, v := range p.conditions {
		conditions = append(conditions, *v)
	}
	return conditions
}

func (p *Pagination) GetOrders() []Order {
	orders := make([]Order, 0, len(p.orders))
	for _, v := range p.orders {
		orders = append(orders, *v)
	}
	return orders
}

func (p *Pagination) SetCondition(condition Condition) {
	isAdd := true
	for _, v := range p.conditions {
		if v.Query == condition.Query {
			isAdd = false
			v.Args = condition.Args
		}
	}
	if isAdd {
		p.conditions = append(p.conditions, &condition)
	}
}

func (p *Pagination) SetOrder(order Order) {
	isAdd := true
	for _, v := range p.orders {
		if v.Column == order.Column {
			isAdd = false
			v.Desc = order.Desc
		}
	}
	if isAdd {
		p.orders = append(p.orders, &order)
	}
}
