package clients

import (
	"github.com/covidlocaldev/common/keys"
	"github.com/covidlocaldev/common/logging"
	"googlemaps.github.io/maps"
)

func NewGoogleMapsClient() *maps.Client {
	client, err := maps.NewClient(maps.WithAPIKey(keys.GoogleApiKeys.GoogleApiServerKey))
	if err != nil {
		logging.Error("could not create google maps api client: %s", err)
		return nil
	}

	return client
}
