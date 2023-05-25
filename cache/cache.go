package cache

import (
	"context"
	"time"

	"github.com/cloudweops/phoenix/http/request"
)

var (
	cache Cache
)

// C 全局缓存对象, 默认使用.
func C() Cache {
	if cache == nil {
		panic("global cache instance is nil")
	}
	return cache
}

// SetGlobal 设置全局缓存.
func SetGlobal(c Cache) {
	cache = c
}

// Cache 缓存接口.
type Cache interface {
	ListKey(*ListKeyRequest) (*ListKeyResponse, error)
	// SetDefaultTTL 设置默认过期时间.
	SetDefaultTTL(ttl time.Duration)
	// Put 添加缓存, 并设置key的默认过期时间.
	Put(key string, val interface{}) error
	// PutWithTTL 添加缓存, 并指定过期时间.
	PutWithTTL(key string, val interface{}, ttl time.Duration) error
	// Get 通过key获取缓存.
	Get(key string, val interface{}) error
	// Delete 通过key删除缓存.
	Delete(key string) error
	// IsExist 检查名为key的缓存是否存在.
	IsExist(key string) bool
	// ClearAll 清理所有缓存.
	ClearAll() error
	// Incr 增加名为key的缓存的值, 和count类似.
	Incr(key string) error
	// Decr 减少名为key的缓存的值, 和count类似.
	Decr(key string) error
	// Close 关闭缓存.
	Close() error
	// 携带上下文
	WithContext(ctx context.Context) Cache
}

// NewListKeyRequest todo
func NewListKeyRequest(pattern string, ps uint, pn uint) *ListKeyRequest {
	return &ListKeyRequest{
		pattern:     pattern,
		PageRequest: request.NewPageRequest(ps, pn),
	}
}

// ListKeyRequest todo
type ListKeyRequest struct {
	pattern string
	*request.PageRequest
}

// Pattern tood
func (req *ListKeyRequest) Pattern() string {
	return req.pattern
}

// NewListKeyResponse todo
func NewListKeyResponse(keys []string, total uint64) *ListKeyResponse {
	return &ListKeyResponse{
		Keys:  keys,
		Total: total,
	}
}

// ListKeyResponse todo
type ListKeyResponse struct {
	Keys  []string
	Total uint64
}
