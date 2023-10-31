package access

type Paginator[T any] struct {
	FromId T      `json:"fromId"` // 分页的起始ID（不含）
	Limit  uint16 `json:"limit"`  // 限制返回的结果数
}
