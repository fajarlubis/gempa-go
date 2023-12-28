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

func LatestEarthQuake() (*EarthQuakeData, error) {
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/autogempa.json", _BMKG), nil)
	if err != nil {
		return nil, err
	}

	var (
		data struct {
			InfoGempa struct {
				Gempa struct {
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
					Shakemap    string    `json:"Shakemap"`
				} `json:"gempa"`
			} `json:"Infogempa"`
		}

		earthQuakeData EarthQuakeData
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

	earthQuakeData.Date = data.InfoGempa.Gempa.Tanggal
	earthQuakeData.Hour = data.InfoGempa.Gempa.Jam
	earthQuakeData.DateTime = &data.InfoGempa.Gempa.DateTime
	earthQuakeData.Coordinates = data.InfoGempa.Gempa.Coordinates

	coords := strings.Split(earthQuakeData.Coordinates, ",")
	earthQuakeData.Latitude, _ = strconv.Atoi(strings.TrimSpace(coords[0]))
	earthQuakeData.Longitude, _ = strconv.Atoi(strings.TrimSpace(coords[1]))

	earthQuakeData.Magnitude, _ = strconv.Atoi(data.InfoGempa.Gempa.Magnitude)

	earthQuakeData.Depth = data.InfoGempa.Gempa.Kedalaman
	earthQuakeData.Region = data.InfoGempa.Gempa.Wilayah

	earthQuakeData.Potential = data.InfoGempa.Gempa.Potensi
	earthQuakeData.AffectedRegion = data.InfoGempa.Gempa.Dirasakan

	earthQuakeData.Shakemap = fmt.Sprintf("%s/%s", _BMKG, data.InfoGempa.Gempa.Shakemap)

	return &earthQuakeData, nil
}
