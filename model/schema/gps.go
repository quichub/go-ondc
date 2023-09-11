package schema

import (
	"fmt"
	"strconv"
	"strings"
)

/*
   Gps:
     description: Describes a gps coordinate
     type: string
     pattern: '^[-+]?([1-8]?\d(\.\d+)?|90(\.0+)?),\s*[-+]?(180(\.0+)?|((1[0-7]\d)|([1-9]?\d))(\.\d+)?)$'
*/

type Gps struct {
	Latitude  float64
	Longitude float64
}

func NewGps(latitude, longitude float64) Gps {
	return Gps{
		Latitude:  latitude,
		Longitude: longitude,
	}
}

func (g Gps) Value() (float64, float64) {
	return g.Latitude, g.Longitude
}

func (g Gps) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%f,%f", g.Latitude, g.Longitude)), nil
}

func (g *Gps) UnmarshalJSON(data []byte) error {
	//parse from string to float64
	val := strings.Split(string(data), ",")
	if len(val) == 2 {
		lat, err := strconv.ParseFloat(val[0], 64)
		if err != nil {
			return fmt.Errorf("invalid gps coordinate: %s", string(data))
		}
		lon, err := strconv.ParseFloat(val[1], 64)
		if err != nil {
			return fmt.Errorf("invalid gps coordinate: %s", string(data))
		}
		g.Latitude = lat
		g.Longitude = lon

		return nil
	}

	return fmt.Errorf("invalid gps coordinate: %s", string(data))
}
