package orm

// 分页返回数
type PageList struct {
	CurrentPage int64
	FirstPage   int64
	LastPage    int64
	PageSize    int64
	Total       int64
	Data        interface{}
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
