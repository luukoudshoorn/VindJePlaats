package bag

type Grens struct {
	Exterior []Polygon
	Interior []Polygon
}

type Polygon struct {
	Vertices []Coordinate
}

type Coordinate struct {
	Lon float64
	Lat float64
}

// As per https://en.wikipedia.org/wiki/Centroid
func PolygonCenter(polygon Polygon) Coordinate {
	var Result Coordinate
	var Area float64
	for i := 0; i < len(polygon.Vertices)-1; i++ {
		var Shoelace = (polygon.Vertices[i].Lon*polygon.Vertices[i+1].Lat - polygon.Vertices[i+1].Lon*polygon.Vertices[i].Lat)
		Result.Lon += (polygon.Vertices[i].Lon * polygon.Vertices[i+1].Lon) * Shoelace
		Result.Lat += (polygon.Vertices[i].Lat * polygon.Vertices[i+1].Lat) * Shoelace
		Area += Shoelace
	}
	var Shoelace = (polygon.Vertices[len(polygon.Vertices)-1].Lon*polygon.Vertices[0].Lat - polygon.Vertices[0].Lon*polygon.Vertices[len(polygon.Vertices)-1].Lat)
	Result.Lon += (polygon.Vertices[len(polygon.Vertices)-1].Lon * polygon.Vertices[0].Lon) * Shoelace
	Result.Lat += (polygon.Vertices[len(polygon.Vertices)-1].Lat * polygon.Vertices[0].Lat) * Shoelace
	Area = (Area + Shoelace) / 2
	Result.Lon = Result.Lon / (6 * Area)
	Result.Lat = Result.Lat / (6 * Area)
	return Result
}
