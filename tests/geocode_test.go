package main

import (
	"context"
	"os"
	"testing"

	"googlemaps.github.io/maps"
)

func TestGeocode(t *testing.T) {
	key := os.Getenv("API_KEY")
	if key == "" {
		t.Logf("API Key not in environment")
		t.Fail()
	}
	c, err := maps.NewClient(maps.WithAPIKey(key))
	if err != nil {
		t.Logf("fatal error: %s", err)
		t.Fail()
	}
	r := &maps.GeocodingRequest{
		Address: "147 Grandview Rd, Ardmore, PA 19003",
	}
	result, err := c.Geocode(context.Background(), r)
	if err != nil {
		t.Logf("fatal error: %s", err)
		t.Fail()
	}

	t.Log(result[0])
}
