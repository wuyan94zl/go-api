package orm

// 分页返回数
type PageList struct {
	CurrentPage int64       `json:"current_page"`
	FirstPage   int64       `json:"first_page"`
	LastPage    int64       `json:"last_page"`
	PageSize    int64       `json:"page_size"`
	Total       int64       `json:"total"`
	Data        interface{} `json:"data"`
}

func SetPageList(data interface{}, currentPage int64, pageSize ...int64) *PageList {
	pageList := &PageList{CurrentPage: currentPage, FirstPage: 1, Data: data}
	if len(pageSize) > 0 {
		pageList.PageSize = pageSize[0]
	}
	return pageList
}

type Where struct {
	Way   string
	Value interface{}
}
