package db

import (
	"context"
	"fmt"
	"github.com/apex/log"
	"github.com/dhdanie/itinerary-db/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type ItineraryId struct {
	Id string `json:"id"`
}

func (c *client) GetItinerary(id string) *model.Itinerary {
	if !c.connected {
		return nil
	}

	itineraryId := &ItineraryId{Id: id}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.timeoutSecs))
	result := c.itinerariesCollection.FindOne(ctx, itineraryId)
	cancel()

	if result.Err() != nil {
		return nil
	}

	itinerary := &model.Itinerary{}
	err := result.Decode(itinerary)
	if err != nil {
		return nil
	}
	return itinerary
}

func (c *client) GetItineraries() []*model.Itinerary {
	if !c.connected {
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.timeoutSecs))
	defer cancel()

	cursor, err := c.itinerariesCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil
	}

	var itineraries []*model.Itinerary
	for cursor.Next(ctx) {
		var itinerary model.Itinerary
		err := cursor.Decode(&itinerary)
		if err != nil {
			log.WithError(err).Error("Error decoding itinerary")
		}
		itineraries = append(itineraries, &itinerary)
	}
	return itineraries
}

func (c *client) SaveItinerary(itinerary *model.Itinerary) error {
	if !c.connected {
		return fmt.Errorf("db not connected")
	}

	filterItinerary := &ItineraryId{Id: itinerary.Id}
	upsertItinerary := bson.M{
		"$set": bson.M{
			"StartDate":           itinerary.StartDate,
			"EndDate":             itinerary.EndDate,
			"SalutationName":      itinerary.SalutationName,
			"SalutationText":      itinerary.SalutationText,
			"SalutationSignature": itinerary.SalutationSignature,
			"Infos":               itinerary.Infos,
			"Attractions":         itinerary.Attractions,
			"HotelStays":          itinerary.HotelStays,
			"DiningReservations":  itinerary.DiningReservations,
			"Flights":             itinerary.Flights,
		},
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.timeoutSecs))
	updateResult, err := c.itinerariesCollection.UpdateOne(ctx, filterItinerary, upsertItinerary, options.Update().SetUpsert(true))
	cancel()

	if err != nil {
		return err
	}
	log.Debugf("Upserted Count: %d", updateResult.UpsertedCount)
	log.WithField("insert-id", updateResult.UpsertedID).Debug("Itinerary updated")
	return nil
}

func (c *client) DeleteItinerary(id string) error {
	if !c.connected {
		return fmt.Errorf("not connected")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(c.timeoutSecs))
	defer cancel()

	itineraryId := &ItineraryId{Id: id}

	result, err := c.itinerariesCollection.DeleteOne(ctx, itineraryId)
	if err != nil {
		return err
	}

	if result.DeletedCount < 1 {
		log.Warnf("Delete called for nonexistent itinerary ID %d", id)
	}

	return nil
}
