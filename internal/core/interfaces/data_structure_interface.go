package interfaces

// CacheType represents the type of cache being used.
type CacheType string

const (
	FixedQACache     CacheType = "fixed_qa"
	StepByStepCache  CacheType = "step_by_step"
	CodeCache        CacheType = "code"
	APICache         CacheType = "api"
	FewShotCache     CacheType = "few_shot"
	EntityBasedCache CacheType = "entity_based"
)

// CacheData is the interface that all cache data types must implement.
type CacheData interface {
	GetType() CacheType
	Store(data interface{}) error
	Retrieve(key string) (interface{}, error)
	Delete(key string) error
	Update(key string, data interface{}) error
	Clear() error
}
