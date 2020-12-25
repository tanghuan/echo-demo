package vos

// PageVo 分页查询 分页数据封装
type PageVo struct {
	// 当前页
	Page uint

	// 每页大小
	Size uint

	// 总记录数
	Total uint

	Data interface{}
}
