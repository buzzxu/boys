package types

type (
	KV[K comparable, V any] struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	Pair[K comparable, V any] struct {
		KV[K, V]
	}

	Pager[T any] struct {
		PageNumber int64 `json:"pageNumber"`
		PageSize   int64 `json:"pageSize"`
		TotalPage  int64 `json:"totalPage"`
		TotalRow   int64 `json:"totalRow"`
		Count      bool  `json:"count"`
		HasNext    bool  `json:"hasNext"`
		Data       []T   `json:"data"`
	}

	Tuple3[T1, T2, T3 any] struct {
		T1 T1 `json:"t1"`
		T2 T2 `json:"t2"`
		T3 T3 `json:"t3"`
	}
	Tuple4[T1, T2, T3, T4 any] struct {
		Tuple3[T1, T2, T3]
		T4 T4 `json:"t4"`
	}
	Tuple5[T1, T2, T3, T4, T5 any] struct {
		Tuple4[T1, T2, T3, T4]
		T5 T5 `json:"t5"`
	}
)

func (p *Pager[T]) FirstRow() int64 {
	return (p.PageNumber - 1) * p.PageSize
}

// Transform takes a transformation function and applies it to each element in the Data slice.
func (p Pager[T]) Transform(transform func(T) interface{}) *Pager[interface{}] {
	newData := make([]interface{}, len(p.Data))
	for i, v := range p.Data {
		newData[i] = transform(v)
	}
	return &Pager[interface{}]{
		PageNumber: p.PageNumber,
		PageSize:   p.PageSize,
		TotalPage:  p.TotalPage,
		TotalRow:   p.TotalRow,
		Count:      p.Count,
		HasNext:    p.HasNext,
		Data:       newData,
	}
}

// PageConvert  takes a transformation function and applies it to each element in the Data slice.
// The function is also generic, accepting a slice of any type T and returning a slice of any type
func PageConvert[T any, U any](pager *Pager[T], transform func(T) *U) *Pager[U] {
	newData := make([]U, len(pager.Data))
	for i, v := range pager.Data {
		newData[i] = *transform(v)
	}
	return &Pager[U]{
		PageNumber: pager.PageNumber,
		PageSize:   pager.PageSize,
		TotalPage:  pager.TotalPage,
		TotalRow:   pager.TotalRow,
		Count:      pager.Count,
		HasNext:    pager.HasNext,
		Data:       newData,
	}
}
