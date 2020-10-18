package load

import (
	"encoding/json"
	"github.com/apex/log"
	"github.com/dhdanie/itinerary-db/db"
	"github.com/dhdanie/itinerary-db/model"
	"io/ioutil"
)

type Loader interface {
	LoadFile(file string) error
}

type venueLoader struct {
	client db.Client
}

func NewVenueLoader(client db.Client) Loader {
	return &venueLoader{
		client: client,
	}
}

func (l *venueLoader) LoadFile(file string) error {
	log.Info("Loading venues")
	fileBytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	log.Debug("Venues read from file")

	var venues []model.Venue
	err = json.Unmarshal(fileBytes, &venues)
	if err != nil {
		return err
	}
	log.Infof("Venues to load - %d", len(venues))

	for _, venue := range venues {
		err := l.client.SaveVenue(&venue)
		if err != nil {
			return err
		}
		log.WithField("venue-id", venue.Id).Info("Venue Saved")
	}
	return nil
}
