package lookup

import (
	"github.com/sirupsen/logrus"
	"googlemaps.github.io/maps"
)

func WithLogger(lgr *logrus.Entry) Option {
	return func(options *Options) {
		options.Logger = lgr
	}
}

func WithClient(client *maps.Client) Option {
	return func(options *Options) {
		options.Client = client
	}
}

func baseLogger() *logrus.Entry {
	return logrus.WithField("service", "google_finder")
}

func baseClient(c *Config) (*maps.Client, error) {
	return maps.NewClient(maps.WithAPIKey(c.GoogleApiKey))
}

type Options struct {
	Logger *logrus.Entry
	Client *maps.Client
}

type Option func(*Options)
