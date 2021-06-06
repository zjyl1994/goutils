package libpaging

import (
	"strconv"
	"strings"
)

type PagingContext struct {
	Current int
	NextURL string
	PrevURL string
	PageURL map[int]string
}

var defaultPaging Paging

func init() {
	defaultPaging = Paging{
		PagingBtnNum:        5,
		PageNoPlaceholder:   "$pageNo$",
		PageSizePlaceholder: "$pageSize$",
		DefaultPageSize:     10,
	}
}

func Default() Paging {
	return defaultPaging
}

type Paging struct {
	PagingBtnNum        int
	PageNoPlaceholder   string
	PageSizePlaceholder string
	PagingURL           string
	DefaultPageSize     int
}

func SQLHelper(pageNo, pageSize int) (limit, offset int) {
	return defaultPaging.SQLHelper(pageNo, pageSize)
}

func (p Paging) SQLHelper(pageNo, pageSize int) (limit, offset int) {
	if pageNo <= 0 {
		pageNo = 1
	}
	if pageSize <= 0 {
		pageSize = p.DefaultPageSize
	}
	return pageSize, (pageNo - 1) * pageSize
}

func (p Paging) ButtonHelper(pageNo, pageSize int, totalCount int64) PagingContext {
	totalPage := (int(totalCount) + pageSize - 1) / pageSize
	pages := make(map[int]string)
	var ret PagingContext
	ret.Current = pageNo
	if pageNo > 1 {
		ret.PrevURL = p.buildPagingURL(pageNo-1, pageSize)
	}
	if pageNo < totalPage {
		ret.NextURL = p.buildPagingURL(pageNo+1, pageSize)
	}
	if totalPage > p.PagingBtnNum {
		pagingBtnStartAt := (p.PagingBtnNum+1)/2 - (p.PagingBtnNum / 2)
		for i := pagingBtnStartAt; i <= p.PagingBtnNum; i++ {
			pages[i] = p.buildPagingURL(i, pageSize)
		}
	} else {
		for i := 1; i <= totalPage; i++ {
			pages[i] = p.buildPagingURL(i, pageSize)
		}
	}

	ret.PageURL = pages
	return ret
}

func (p Paging) buildPagingURL(pageNo, pageSize int) string {
	tmpStr := strings.ReplaceAll(p.PagingURL, p.PageNoPlaceholder, strconv.Itoa(pageNo))
	return strings.ReplaceAll(tmpStr, p.PageSizePlaceholder, strconv.Itoa(pageSize))
}
