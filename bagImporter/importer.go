package bagImporter

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/ctessum/geom"
	"github.com/ctessum/geom/proj"
	"github.com/pbnjay/memory"
	"oudshoorn.dev/mijnwoonplaats/bag"
)

var importFunctions = map[string]func([]byte){
	"[0-9]+WPL[0-9]+.zip": ParseWoonplaatsen,
	"[0-9]+PND[0-9]+.zip": parsePanden,
}

func innerZipIterator(file *zip.File, action func([]byte)) {
	fmt.Printf("File: %s\n", file.Name)

	var zipReaderAt io.ReaderAt
	var size int64
	//For some reason, uncompressing the file takes up to six times the file size worth of memory
	if file.UncompressedSize64*6 > memory.FreeMemory() {
		//File plus overhead might not fit in memory, use a slower but more memory efficient method
		fmt.Println("using memory efficient method")
		zipReaderAt = newInefficentReaderAt(file.Open)
		size = int64(file.UncompressedSize64)
	} else {
		//File plus overhead fit entirely in memory, use fast method
		fmt.Println("using fast method")
		in, openErr := file.Open()
		if openErr != nil {
			fmt.Printf("Could not open inner zip %s: %s\n", file.Name, openErr)
			return
		}
		buffer, err := ioutil.ReadAll(in)
		if err != nil {
			fmt.Printf("Could not read inner zip %s into buffer: %s\n", file.Name, openErr)
			return
		}
		zipReaderAt = bytes.NewReader(buffer)
		size = int64(len(buffer))
	}

	zipReader, err := zip.NewReader(zipReaderAt, size)
	if err != nil {
		fmt.Printf("Could not open inner zip %s: %s\n", file.Name, err)
		return
	}

	for _, f := range zipReader.File {
		reader, err := f.Open()
		if err != nil {
			fmt.Printf("Could not read inner file %s in inner zip %s\n", f.Name, file.Name)
			return
		}
		defer reader.Close()
		bytesValue, _ := ioutil.ReadAll(reader)
		go action(bytesValue)
	}
}

func ImportBag(filename string) {
	fileInfo, err := os.Stat(filename)

	if err != nil || fileInfo.IsDir() {
		log.Fatalf("Cannot open file '%s'\n", filename)
	}

	read, zipErr := zip.OpenReader(filename)
	if zipErr != nil {
		log.Fatalf("Error while reading zip file '%s'\n", filename)
	}
	defer read.Close()

	//The zip should contain multiple zips, of which we're interested in a few
	for _, f := range read.File {
		for k, v := range importFunctions {
			//Todo: Match map using regex
			if match, _ := regexp.MatchString(k, f.Name); match {
				innerZipIterator(f, v)
			}
		}
	}
}

func createTransformerToWgs84(projDef string) proj.Transformer {
	wgs, srcErr := proj.Parse("+proj=longlat +datum=WGS84 +no_defs")
	if srcErr != nil {
		log.Fatal(srcErr)
	}

	sr, dstErr := proj.Parse(projDef)
	if dstErr != nil {
		log.Fatal(dstErr)
	}

	transform, transformError := sr.NewTransform(wgs)
	if transformError != nil {
		log.Fatal(transformError)
	}

	return transform
}

var tr = createTransformerToWgs84("+proj=sterea +lat_0=52.1561605555556 +lon_0=5.38763888888889 +k=0.9999079 +x_0=155000 +y_0=463000 +ellps=bessel +towgs84=565.4171,50.3319,465.5524,1.9342,-1.6677,9.1019,4.0725 +units=m +no_defs")

func transform(transformer proj.Transformer, input bag.Coordinate) bag.Coordinate {
	pt := geom.NewPoint(input.Lon, input.Lat)
	Geom, _ := pt.Transform(tr)
	pt2 := Geom.Points()()
	return bag.Coordinate{Lon: pt2.X, Lat: pt2.Y}
}

func addCoordinatesToGrens(GmlPolygon gmlPolygon, Grens *bag.Grens) {
	for _, Exterior := range GmlPolygon.Exterior {
		coordinatesText := strings.Split(Exterior.LinearRing.PosList.Coordinates, " ")
		Grens.Exterior = append(Grens.Exterior, bag.Polygon{})
		for i := 0; i < Exterior.LinearRing.PosList.Count*GmlPolygon.SrsDimension; i += GmlPolygon.SrsDimension {
			lon, _ := strconv.ParseFloat(coordinatesText[i], 64)
			lat, _ := strconv.ParseFloat(coordinatesText[i+1], 64)
			Grens.Exterior[len(Grens.Exterior)-1].Vertices = append(Grens.Exterior[len(Grens.Exterior)-1].Vertices, transform(tr, bag.Coordinate{Lon: lon, Lat: lat}))
		}
	}
	for _, Interior := range GmlPolygon.Interior {
		coordinatesText := strings.Split(Interior.LinearRing.PosList.Coordinates, " ")
		Grens.Interior = append(Grens.Interior, bag.Polygon{})
		for i := 0; i < Interior.LinearRing.PosList.Count*GmlPolygon.SrsDimension; i += GmlPolygon.SrsDimension {
			lon, _ := strconv.ParseFloat(coordinatesText[i], 64)
			lat, _ := strconv.ParseFloat(coordinatesText[i+1], 64)
			Grens.Interior[len(Grens.Interior)-1].Vertices = append(Grens.Interior[len(Grens.Interior)-1].Vertices, transform(tr, bag.Coordinate{Lon: lon, Lat: lat}))
		}
	}
}

func grensFromGeometrie(Geometrie geometrie) bag.Grens {
	var Grens bag.Grens
	for _, SurfaceMember := range Geometrie.Multivlak.MultiSurface.SurfaceMember {
		addCoordinatesToGrens(SurfaceMember.Polygon, &Grens)
	}
	addCoordinatesToGrens(Geometrie.Vlak.Polygon, &Grens)
	addCoordinatesToGrens(Geometrie.Polygon, &Grens)
	return Grens
}

func ParseWoonplaatsen(bytesValue []byte) {
	var Woonplaatsen []bag.Woonplaats

	var BagStand bagStand
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'bagStand' which we defined above
	xmlErr := xml.Unmarshal(bytesValue, &BagStand)
	if xmlErr != nil {
		fmt.Println(xmlErr)
		return
	}

	for _, Stand := range BagStand.StandBestand.Stand {
		if Stand.BagObject.Woonplaats.Voorkomen.HistorieVoorkomen.EindGeldigheid != "" {
			continue
		}
		var Woonplaats bag.Woonplaats
		Woonplaats.Id = Stand.BagObject.Woonplaats.Identificatie.Id
		Woonplaats.Naam = Stand.BagObject.Woonplaats.Naam
		Woonplaats.Grens = grensFromGeometrie(Stand.BagObject.Woonplaats.Geometrie)

		Woonplaatsen = append(Woonplaatsen, Woonplaats)
	}

	fmt.Printf("Aantal woonplaatsen geïmporteerd: %d\n", len(Woonplaatsen))
}

var allePanden []bag.Pand

func parsePanden(bytesValue []byte) {
	var Panden []bag.Pand

	var BagStand bagStand
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'bagStand' which we defined above
	xmlErr := xml.Unmarshal(bytesValue, &BagStand)
	if xmlErr != nil {
		fmt.Println(xmlErr)
		return
	}

	for _, Stand := range BagStand.StandBestand.Stand {
		var Pand bag.Pand
		Pand.Id = Stand.BagObject.Pand.Identificatie.Id
		Pand.Omtrek = grensFromGeometrie(Stand.BagObject.Pand.Geometrie)

		Panden = append(Panden, Pand)
	}
	allePanden = append(allePanden, Panden...)

	fmt.Printf("Aantal Panden geïmporteerd: %d\n", len(allePanden))
}
