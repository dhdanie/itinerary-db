package db

import (
	"context"
	"fmt"
	"github.com/apex/log"
	"github.com/dhdanie/itinerary-db/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Client interface {
	Connect() error
	Disconnect()
	SaveVenue(venue *model.Venue) error
	GetVenue(id int) *model.Venue
	GetVenues() []*model.Venue
	DeleteVenue(id int) error
}

type client struct {
	connected                 bool
	uri                       string
	timeoutSecs               int
	dbName                    string
	itinerariesCollectionName string
	venuesCollectionName      string
	itinerariesCollection     *mongo.Collection
	venuesCollection          *mongo.Collection
	mongoCli                  *mongo.Client
}

type VenueId struct {
	Id int `json:"id"`
}

var mongoClient Client

func NewClient(dbConn string, timeoutSecs int, dbName string, itinerariesCollection string, venuesCollection string) Client {
	if mongoClient != nil {
		return mongoClient
	}
	mongoClient = &client{
		connected:                 false,
		uri:                       dbConn,
		timeoutSecs:               timeoutSecs,
		dbName:                    dbName,
		itinerariesCollectionName: itinerariesCollection,
		venuesCollectionName:      venuesCollection,
	}
	return mongoClient
}

func GetClient() Client {
	return mongoClient
}

func (c *client) Connect() error {
	if c.connected {
		return nil
	}
	var err error

	c.mongoCli, err = mongo.NewClient(options.Client().ApplyURI(c.uri))
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.timeoutSecs))
	err = c.mongoCli.Connect(ctx)
	cancel()

	c.itinerariesCollection = c.mongoCli.Database(c.dbName).Collection(c.itinerariesCollectionName)
	c.venuesCollection = c.mongoCli.Database(c.dbName).Collection(c.venuesCollectionName)

	if err != nil {
		return err
	}
	c.connected = true
	return nil
}

func (c *client) Disconnect() {
	if !c.connected {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.timeoutSecs))
	err := c.mongoCli.Disconnect(ctx)
	cancel()
	if err != nil {
		log.WithError(err).Warn("Error received while disconnecting mongo client")
	}

	c.venuesCollection = nil
	c.itinerariesCollection = nil

	c.connected = false
}

func (c *client) SaveVenue(venue *model.Venue) error {
	if !c.connected {
		return fmt.Errorf("db not connected")
	}

	filterVenue := &VenueId{Id: venue.Id}
	upsertVenue := bson.M{
		"$set": bson.M{
			"Name":       venue.Name,
			"Abbr":       venue.Abbr,
			"Address":    venue.Address,
			"PhoneNo":    venue.PhoneNo,
			"AvatarUrl":  venue.AvatarUrl,
			"CoverPhoto": venue.CoverPhoto,
			"InfoUrl":    venue.InfoUrl,
			"Notes":      venue.Notes,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.timeoutSecs))
	updateResult, err := c.venuesCollection.UpdateOne(ctx, filterVenue, upsertVenue, options.Update().SetUpsert(true))
	cancel()

	if err != nil {
		return err
	}
	log.Debugf("Upserted Count: %d", updateResult.UpsertedCount)
	log.WithField("insert-id", updateResult.UpsertedID).Debug("Venue updated")
	return nil
}

func (c *client) GetVenue(id int) *model.Venue {
	if !c.connected {
		return nil
	}
	venueId := &VenueId{Id: id}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.timeoutSecs))
	result := c.venuesCollection.FindOne(ctx, venueId)
	cancel()

	if result.Err() != nil {
		return nil
	}

	venue := &model.Venue{}
	err := result.Decode(venue)
	if err != nil {
		return nil
	}
	return venue
}

func (c *client) GetVenues() []*model.Venue {
	if !c.connected {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.timeoutSecs))
	defer cancel()

	cursor, err := c.venuesCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil
	}

	var venues []*model.Venue
	for cursor.Next(ctx) {
		var venue model.Venue
		err := cursor.Decode(&venue)
		if err != nil {
			log.WithError(err).Error("Error decoding venue")
		}
		venues = append(venues, &venue)
	}
	return venues
}

func (c *client) DeleteVenue(id int) error {
	if !c.connected {
		return fmt.Errorf("not connected")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.timeoutSecs))
	defer cancel()

	venueId := &VenueId{Id: id}

	result, err := c.venuesCollection.DeleteOne(ctx, venueId)
	if err != nil {
		return err
	}

	if result.DeletedCount < 1 {
		log.Warnf("Delete called for nonexistent venue ID %d", id)
	}

	return nil
}
