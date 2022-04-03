package sampling

type Entity struct {
	Id          int    `json:"id"`
	DisplayName string `json:"displayName"`
}

type Sample struct {
	Time      int64
	EntityId  int
	Latitude  float64
	Longitude float64
}
