package models

import . "go.mongodb.org/mongo-driver/bson/primitive"

type Beer struct {
	ID      ObjectID `bson:"_id,omitempty"`
	Name    string   `bson:"name,omitempty"`
	Brewery string   `bson:"brewery,omitempty"`
	Country string   `bson:"country,omitempty"`
	Price   []Price  `bson:"price,omitempty"`
}

func NewBeer(name string, brewery string, country string, price []Price) *Beer {
	beer := Beer{Name: name, Brewery: brewery, Country: country, Price: price}
	return &beer
}
