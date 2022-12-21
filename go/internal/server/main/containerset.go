package servermain

import (
	"context"
	"sync"

	"zenhack.net/go/sandstorm-next/go/internal/database"
	"zenhack.net/go/sandstorm-next/go/internal/server/container"
)

type ContainerSet struct {
	mu                  sync.Mutex
	db                  database.DB
	containersByGrainId map[string]*container.Container
}

func (cset *ContainerSet) Get(ctx context.Context, grainId string) (*container.Container, error) {
	cset.mu.Lock()
	defer cset.mu.Unlock()
	c, ok := cset.containersByGrainId[grainId]
	if ok {
		return c, nil
	}
	c, err := container.Start(ctx, cset.db, grainId)
	if err == nil {
		cset.containersByGrainId[grainId] = c
	}
	return c, err
}
