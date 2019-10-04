package lookup

type Request struct {
	Location *Location
	Radius   uint
	Keyword  string
	OpenNow  bool
}

type Location struct {
	Lat float64
	Lng float64
}

type NearbyResult struct {
	List []*Place
}

type Place struct {
	Name   string
	Rating float32
	Types  []string
}
