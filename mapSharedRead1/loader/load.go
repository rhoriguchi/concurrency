package loader

import (
	"compress/bzip2"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
)

const defaultFilename = "airports.csv.bz2"
const highestAirfield = 30000

type Record struct {
	Name         string
	Ident        string
	Type         string
	ElevationFt  int
	Country      string
	Municipality string
	GPSCode      string
	IATACode     string
	LocalCode    string
	Longitude    float64 // negative values are west of Greenwich
	Latitude     float64 // negative values are south of equator
}

var regexCoords = regexp.MustCompile(`(-?\d{1,3}(?:\.\d+)?)`)

func parseCoordinates(input string) (longitude float64, latitude float64, err error) {
	m := regexCoords.FindAllString(input, 2)
	if m == nil {
		return 0, 0, fmt.Errorf("coordinates invalid: '%v'", input)
	}

	if long, err := strconv.ParseFloat(m[0], 64); err != nil {
		return 0, 0, fmt.Errorf("longitude not a float: '%v'", m[0])
	} else if lat, err := strconv.ParseFloat(m[1], 64); err != nil {
		return 0, 0, fmt.Errorf("latitude not a float: '%v'", m[1])
	} else {
		return long, lat, nil
	}
}

func load() (map[string]Record, error) {
	f, err := os.Open(defaultFilename)
	if err != nil {
		return nil, fmt.Errorf("open file %v failed: %v", defaultFilename, err)
	}

	z := bzip2.NewReader(f)

	airports := make(map[string]Record)
	cr := csv.NewReader(z)

	_, err = cr.Read()
	if err != nil {
		return nil, fmt.Errorf("failed reading CSV title: %v", err)
	}

	for {
		rec, err := cr.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("failed reading CSV record: %v", err)
		}

		ident := rec[0]

		var elev int64
		if rec[3] != "" {
			elev, err = strconv.ParseInt(rec[3], 10, 32)
			if err != nil {
				return nil, fmt.Errorf("elevation is not a number: '%v'", rec[3])
			}
		}

		if elev > highestAirfield {
			return nil, fmt.Errorf("elevation invalid (too high): %v", elev)
		}

		lon, lat, err := parseCoordinates(rec[9])
		if err != nil {
			return nil, err
		}

		airports[ident] = Record{
			Name:         rec[2],
			Ident:        ident,
			Type:         rec[1],
			ElevationFt:  int(elev),
			Country:      rec[4],
			Municipality: rec[5],
			GPSCode:      rec[6],
			IATACode:     rec[7],
			LocalCode:    rec[8],
			Longitude:    lon,
			Latitude:     lat,
		}
	}

	return airports, nil
}

func GetAll() map[string]Record {
	all, err := load()
	if err != nil {
		panic(err)
	}

	return all
}
