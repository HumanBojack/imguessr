package db

import (
	"context"
	"imguessr/pkg/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetAllGames retrieves all games from the MongoDB database.
func (ms mongoStore) GetAllGames() ([]*domain.Game, error) {
	// Find all documents in the GameCollection
	cur, err := ms.GameCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	gameList, err := decodeGame(cur)

	// Return the gameList slice
	return gameList, err
}

// GetAllGamesByUserID retrieves all games from the MongoDB database that have the specified userID.
func (ms mongoStore) GetAllGamesByUserID(userType string, userID string) ([]*domain.Game, error) {
	// Choose the correct filter
	var filter bson.M
	switch userType {
	case "owner":
		filter = bson.M{"user._id": userID}
	case "player":
		filter = bson.M{"gameparameters.users_ids": bson.M{"$in": []string{userID}}}
	default:
		filter = bson.M{
			"$or": []bson.M{
				{"user._id": userID},
				{"gameparameters.users_ids": bson.M{"$in": []string{userID}}},
			},
		}
	}

	// Find games that match the filter
	cur, err := ms.GameCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	// Close the cursor after the function returns
	defer cur.Close(context.TODO())

	gameList, err := decodeGame(cur)

	return gameList, err
}

// decodeGame decodes the documents returned by the MongoDB cursor into Game structs.
func decodeGame(cur *mongo.Cursor) ([]*domain.Game, error) {
	// Create a slice to hold the games
	var gameList = []*domain.Game{}

	// Iterate over the documents and decode them into Game structs
	for cur.Next(context.TODO()) {
		var elem domain.Game
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		// Append the decoded Game struct to the gameList slice
		gameList = append(gameList, &elem)
	}

	// Return the gameList slice
	return gameList, nil
}

func (ms mongoStore) CreateGame(g *domain.Game) error {
	_, err := ms.GameCollection.InsertOne(context.TODO(), g)
	if err != nil {
		return err
	}
	return nil
}
