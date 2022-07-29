package mapx

import "sync"

type Map[K comparable, V any] struct {
	val map[K]V
	mtx *sync.RWMutex
}

func New[K comparable, V any]() *Map[K, V] {
	return &Map[K, V]{val: make(map[K]V), mtx: &sync.RWMutex{}}
}

func (m *Map[K, V]) Load(k K) (v V, ok bool) {
	m.mtx.RLock()
	defer m.mtx.RUnlock()
	v, ok = m.val[k]
	return
}

func (m *Map[K, V]) Store(k K, v V) {
	m.mtx.Lock()
	defer m.mtx.Unlock()
	m.val[k] = v
}
