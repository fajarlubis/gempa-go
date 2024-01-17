package gempago

import "time"

type EarthQuakeData struct {
	Date           string     `json:"date"`
	Hour           string     `json:"hour"`
	DateTime       *time.Time `json:"datetime"`
	Coordinates    string     `json:"coordinates"`
	Latitude       int        `json:"latitude"`
	Longitude      int        `json:"longitude"`
	Magnitude      int        `json:"magnitude"`
	Depth          string     `json:"depth"`
	Region         string     `json:"region"`
	Potential      string     `json:"potential"`
	AffectedRegion string     `json:"affected_region"`
	Shakemap       string     `json:"shakemap"`
	GoogleMapsUrl  string     `json:"google_maps_url"`
}
