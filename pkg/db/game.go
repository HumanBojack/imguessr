package db

import (
	"context"
	"imguessr/pkg/domain"
)

func (ms mongoStore) CreateGame(g *domain.Game) error {
	_, err := ms.GameCollection.InsertOne(context.TODO(), g)
	if err != nil {
		return err
	}
	return nil
}
