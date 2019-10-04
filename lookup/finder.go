package lookup

import (
	"context"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"googlemaps.github.io/maps"
)

type Config struct {
	GoogleApiKey string `envconfig:"GOOGLE_API_KEY" required:"true"`
}

type Opts struct {
	lgr    *logrus.Entry
	client *maps.Client
}

func defaultConf() (*Config, error) {
	c := new(Config)
	if err := envconfig.Process("", c); err != nil {
		return nil, err
	}
	return c, nil
}

type Finder struct {
	config *Config
	opts   *Options
}

func New(opts ...Option) (*Finder, error) {
	c, err := defaultConf()
	if err != nil {
		return nil, err
	}

	options := &Options{
		Logger: baseLogger(),
		Client: nil,
	}

	for _, opt := range opts {
		opt(options)
	}
	if options.Client != nil {
		return &Finder{
			config: c,
			opts:   options,
		}, nil
	}
	// base maps client
	options.Client, err = baseClient(c)
	if err != nil {
		return nil, err
	}
	return &Finder{
		config: c,
		opts:   options,
	}, nil
}

func (s Finder) Do(ctx context.Context, req *Request) (*NearbyResult, error) {
	res, err := s.opts.Client.NearbySearch(ctx, &maps.NearbySearchRequest{
		Location: &maps.LatLng{
			Lat: req.Location.Lat,
			Lng: req.Location.Lng,
		},
		Radius:  req.Radius,
		Keyword: req.Keyword,
		OpenNow: req.OpenNow,
	})

	if err != nil {
		return nil, err
	}

	return s.buildResponse(10, &res)
}

func (s Finder) buildResponse(limit int, input *maps.PlacesSearchResponse) (*NearbyResult, error) {
	if len(input.Results) == 0 {
		return nil, NoAvailablePlaceError
	}
	result := &NearbyResult{List: []*Place{}}
	for i, p := range input.Results {
		if i >= limit {
			break
		}
		result.List = append(result.List, &Place{
			Name:   p.Name,
			Rating: p.Rating,
			Types:  p.Types,
		})
	}
	return result, nil
}
