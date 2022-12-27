package rest

type ListResponse struct {
	Data []map[string]any
	Meta struct {
		Pagination struct {
			Page      int
			PageSize  int `json:"pageSize"`
			PageCount int `json:"pageCount"`
			Total     int
		}
	}
}
