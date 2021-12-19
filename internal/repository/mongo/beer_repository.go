package repository

import (
	. "github.com/v8tix/bender/internal/models"
	. "github.com/v8tix/kit/app"
	. "github.com/v8tix/kit/utils"
	"go.mongodb.org/mongo-driver/bson"
	. "go.mongodb.org/mongo-driver/bson/primitive"
)

const BeerCollection = "beer"

type BeerRepositoryI struct {
	App
}

func NewBeerRepository(app App) BeerRepositoryI {
	beerRepository := BeerRepositoryI{
		app,
	}
	return beerRepository
}

func (b BeerRepositoryI) FindBeerById(id string) (*Beer, error) {
	objectID, err := ToObjectID(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}
	var beer Beer
	collection := GetCollection(b.MongoServer.Db, BeerCollection, b.MongoServer.Client)
	if err := collection.FindOne(b.MongoServer.Ctx, filter).Decode(&beer); err != nil {

		return nil, err

	}
	return &beer, nil

}

func (b BeerRepositoryI) InsertBeer(beer *Beer) (string, error) {
	collection := GetCollection(b.MongoServer.Db, BeerCollection, b.MongoServer.Client)
	result, err := collection.InsertOne(b.MongoServer.Ctx, beer)
	if err != nil {

		return "", err

	}

	if oid, ok := result.InsertedID.(ObjectID); ok {

		return oid.Hex(), nil

	}

	return "", err
}

func (b BeerRepositoryI) UpdateBeer(beer *Beer) (string, error) {
	collection := GetCollection(b.MongoServer.Db, BeerCollection, b.MongoServer.Client)
	filter := bson.M{"_id": beer.ID}
	result, err := collection.ReplaceOne(b.MongoServer.Ctx, filter, beer)
	if err != nil || result.ModifiedCount <= 0 {

		return "", err

	}
	return beer.ID.Hex(), nil
}
