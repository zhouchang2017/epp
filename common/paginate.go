package common

type Paginate struct {
	Data        interface{} `json:"data"`
	CurrentPage uint        `json:"current_page"`
	PerPage     uint        `json:"per_page"`
	Count       uint        `json:"count"`
	Total       uint        `json:"total"`
	LastPage    uint        `json:"last_page"`
	NextPage    uint        `json:"next_page"`
}

func NewPaginate(data interface{}, currentPage uint, perPage uint, count uint) *Paginate {
	p := &Paginate{Data: data, CurrentPage: currentPage, PerPage: perPage, Count: count}

	p.setTotal()
	p.setLastPage()
	p.setNextPage()

	return p
}

func (this *Paginate) setLastPage() {

	if this.Total-1 <= 0 {
		this.LastPage = 1
	} else {
		this.LastPage = this.Total
	}
}

func (this *Paginate) setNextPage() {

	if this.CurrentPage+1 >= this.Total {
		this.NextPage = this.Total
	} else {
		this.NextPage = this.CurrentPage + 1
	}
}

func (this *Paginate) setTotal() {
	this.Total = this.Count / this.PerPage
	if this.Count%this.PerPage > 1 {
		this.Total += 1
	}
}
