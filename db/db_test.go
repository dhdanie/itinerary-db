package db

import (
	"fmt"
	"testing"
)

func TestClient_GetVenues(t *testing.T) {
	client := NewClient("mongodb://mongoadmin:booger@localhost:27017", 10, "itinerary", "itineraries", "venues")
	err := client.Connect()
	if err != nil {
		t.Fatal(err)
	}
	defer client.Disconnect()

	venues := client.GetVenues()
	for _, venue := range venues {
		fmt.Printf("%d - %s\n", venue.Id, venue.Name)
	}
}
