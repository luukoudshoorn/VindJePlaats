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

func readAllForTest(t *testing.T, filename string) []byte {
	f, err := os.Open(filename)
	if err != nil {
		t.Error(err)
	}
	bytes, readErr := io.ReadAll(f)
	if readErr != nil {
		t.Error(readErr)
	}
	return bytes
}

func TestParseWoonplaatsen(t *testing.T) {
	bytes := readAllForTest(t, "9999WPL08022023-000001.xml")
	ParseWoonplaatsen(bytes)
	assert.Equal(t, 239, len(alleWoonplaatsen))
}

func TestParsePanden(t *testing.T) {
	bytes := readAllForTest(t, "9999PND08022023-000001.xml")
	ParsePanden(bytes)
	assert.Equal(t, 10000, len(allePanden))
}

func TestParseVerblijfsobjecten(t *testing.T) {
	bytes := readAllForTest(t, "9999VBO08022023-000001.xml")
	ParseVerblijfsobjecten(bytes)
	assert.Equal(t, 10000, len(alleVerblijfsobjecten))
}
