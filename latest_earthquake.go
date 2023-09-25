package gempago

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type AutoGempa struct {
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

func LatestEarthQuake() (*AutoGempa, error) {
	req, err := http.NewRequest(http.MethodGet, "https://data.bmkg.go.id/DataMKG/TEWS/autogempa.json", nil)
	if err != nil {
		return nil, err
	}

	var data AutoGempa

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("status code not 200")
	}

	rb, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(rb, &data); err != nil {
		return nil, err
	}

	return &data, nil
}
