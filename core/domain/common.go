package domain

type ListResponse[T any] struct {
	Items []T
	Total int64
}
