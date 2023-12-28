package gempago

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func RecentsEarthQuake() ([]*EarthQuakeData, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/gempaterkini.json", _BMKG), nil)
	if err != nil {
		return nil, err
	}

	var (
		data struct {
			InfoGempa struct {
				Gempa []struct {
					Tanggal     string    `json:"Tanggal"`
					Jam         string    `json:"Jam"`
					DateTime    time.Time `json:"DateTime"`
					Coordinates string    `json:"Coordinates"`
					Lintang     string    `json:"Lintang"`
					Bujur       string    `json:"Bujur"`
					Magnitude   string    `json:"Magnitude"`
					Kedalaman   string    `json:"Kedalaman"`
					Wilayah     string    `json:"Wilayah"`
					Potensi     string    `json:"Potensi"`
					Dirasakan   string    `json:"Dirasakan"`
				} `json:"gempa"`
			} `json:"Infogempa"`
		}

		earthQuakeData = make([]*EarthQuakeData, 0)
	)

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("http status code not 200")
	}

	rb, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rb, &data); err != nil {
		return nil, err
	}

	for _, v := range data.InfoGempa.Gempa {
		coords := strings.Split(v.Coordinates, ",")
		lat, _ := strconv.Atoi(strings.TrimSpace(coords[0]))
		long, _ := strconv.Atoi(strings.TrimSpace(coords[1]))

		magnitude, _ := strconv.Atoi(v.Magnitude)

		earthQuakeData = append(earthQuakeData, &EarthQuakeData{
			Date:        v.Tanggal,
			Hour:        v.Jam,
			DateTime:    &v.DateTime,
			Coordinates: v.Coordinates,

			Latitude:  lat,
			Longitude: long,

			Magnitude: magnitude,

			Depth:  v.Kedalaman,
			Region: v.Wilayah,

			Potential:      v.Potensi,
			AffectedRegion: v.Dirasakan,
		})
	}

	return earthQuakeData, nil
}
