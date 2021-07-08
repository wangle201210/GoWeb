// @BeeOverwrite YES
// @BeeGenerateTime 20200911_202853
package models

type Pagination struct {
	Current  int `json:"current"`
	Total    int `json:"total"`
	PageSize int `json:"pageSize"`
}

type ListData struct {
	List []interface{} `json:"data"`
	Pagination
}

func NewPagination(current int, pageSize int) *Pagination {
	p := &Pagination{}
	p.Current = current
	p.PageSize = pageSize
	if p.Current == 0 {
		p.Current = 1
	}
	if p.PageSize == 0 {
		p.PageSize = 20
	}
	return p
}
