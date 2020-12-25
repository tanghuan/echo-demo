package dtos

// PageDto 分页查询 分页数据封装
type PageDto struct {

	// 查询页 从 0 开始
	Page uint `json:"page" form:"page" query:"page"`

	// 每页大小
	Size uint `json:"size" form:"size" query:"size"`

	// 搜索关键字
	SearchKey string `json:"search_key" form:"search_key" query:"search_key"`
}

// Offset mysql 分页的 offset
func (dto *PageDto) Offset() uint {
	return (dto.Page - 1) * dto.Size
}
