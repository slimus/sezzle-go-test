package storage

import "context"

type MemoryStorage struct {
	maxCount int
	data     []string
}

func NewMemoryStorage(maxCount int) *MemoryStorage {
	return &MemoryStorage{
		data:     []string{},
		maxCount: maxCount,
	}
}

func (ms *MemoryStorage) Store(ctx context.Context, s string) error {
	ms.data = append(ms.data, s)
	if len(ms.data) > ms.maxCount {
		ms.data = ms.data[len(ms.data)-ms.maxCount:]
	}
	return nil
}

func (ms *MemoryStorage) List(ctx context.Context) ([]string, error) {
	return ms.data, nil
}
