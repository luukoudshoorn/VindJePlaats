package bagImporter

import (
	"io"
	"math"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"oudshoorn.dev/mijnwoonplaats/bag"
)

func TestTransform(t *testing.T) {
	coord := transform(tr, bag.Coordinate{Lon: 190927.9890527073, Lat: 443886.1709509538})
	if math.Abs(5.910181-coord.Lon) > 0.000001 || math.Abs(51.982213-coord.Lat) > 0.000001 {
		t.Fail()
	}
}

func TestParseWoonplaatsen(t *testing.T) {
	f, err := os.Open("parseWoonplaatsTest.xml")
	if err != nil {
		t.Fail()
	}
	bytes, readErr := io.ReadAll(f)
	if readErr != nil {
		t.Fail()
	}
	assert.Equal(t, ParseWoonplaatsen(bytes), 1)
}
