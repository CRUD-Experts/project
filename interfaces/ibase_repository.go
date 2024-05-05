package interfaces

type IBaseRepository interface {
	GetCollection() interface{}
	Create(ctx interface{}, document interface{}) (interface{}, error)
	FindByID(ctx interface{}, id interface{}, result interface{}) (interface{}, error)
	FindAll(ctx interface{}, filter interface{}, result interface{}) (interface{}, error)
	FindWithPagination(ctx interface{}, filter interface{}, page, pageSize int64) ([]interface{}, error)
	Update(ctx interface{}, filter interface{}, update interface{}) (interface{}, error)
	Delete(ctx interface{}, filter interface{}) (interface{}, error)
	DeleteMany(ctx interface{}, filter interface{}) (interface{}, error)
}