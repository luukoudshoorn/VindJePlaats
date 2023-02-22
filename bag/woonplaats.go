package bag

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type bagStand struct {
	XMLName            xml.Name `xml:"bagStand"`
	Text               string   `xml:",chardata"`
	DatatypenNEN3610   string   `xml:"DatatypenNEN3610,attr"`
	Objecten           string   `xml:"Objecten,attr"`
	Gml                string   `xml:"gml,attr"`
	Historie           string   `xml:"Historie,attr"`
	ObjectenRef        string   `xml:"Objecten-ref,attr"`
	Nen5825            string   `xml:"nen5825,attr"`
	KenmerkInOnderzoek string   `xml:"KenmerkInOnderzoek,attr"`
	SelectiesExtract   string   `xml:"selecties-extract,attr"`
	SlBagExtract       string   `xml:"sl-bag-extract,attr"`
	Sl                 string   `xml:"sl,attr"`
	Xsi                string   `xml:"xsi,attr"`
	Xs                 string   `xml:"xs,attr"`
	SchemaLocation     string   `xml:"schemaLocation,attr"`
	BagInfo            struct {
		Text               string `xml:",chardata"`
		GebiedRegistratief struct {
			Text      string `xml:",chardata"`
			GebiedNLD string `xml:"Gebied-NLD"`
		} `xml:"Gebied-Registratief"`
		LVCExtract struct {
			Text                 string `xml:",chardata"`
			StandTechnischeDatum string `xml:"StandTechnischeDatum"`
		} `xml:"LVC-Extract"`
	} `xml:"bagInfo"`
	StandBestand struct {
		Text    string `xml:",chardata"`
		Dataset string `xml:"dataset"`
		Inhoud  struct {
			Text        string `xml:",chardata"`
			Gebied      string `xml:"gebied"`
			LeveringsId string `xml:"leveringsId"`
			ObjectTypen struct {
				Text       string `xml:",chardata"`
				ObjectType string `xml:"objectType"`
			} `xml:"objectTypen"`
		} `xml:"inhoud"`
		Stand []struct {
			Text      string `xml:",chardata"`
			BagObject struct {
				Text       string `xml:",chardata"`
				Woonplaats struct {
					Text          string `xml:",chardata"`
					Identificatie struct {
						Text   string `xml:",chardata"`
						Domein string `xml:"domein,attr"`
					} `xml:"identificatie"`
					Naam           string    `xml:"naam"`
					Geometrie      geometrie `xml:"geometrie"`
					Status         string    `xml:"status"`
					Geconstateerd  string    `xml:"geconstateerd"`
					Documentdatum  string    `xml:"documentdatum"`
					Documentnummer string    `xml:"documentnummer"`
					Voorkomen      struct {
						Text      string `xml:",chardata"`
						Voorkomen struct {
							Text                   string `xml:",chardata"`
							Voorkomenidentificatie string `xml:"voorkomenidentificatie"`
							BeginGeldigheid        string `xml:"beginGeldigheid"`
							EindGeldigheid         string `xml:"eindGeldigheid"`
							TijdstipRegistratie    string `xml:"tijdstipRegistratie"`
							BeschikbaarLV          struct {
								Text                  string `xml:",chardata"`
								TijdstipRegistratieLV string `xml:"tijdstipRegistratieLV"`
							} `xml:"BeschikbaarLV"`
						} `xml:"Voorkomen"`
					} `xml:"voorkomen"`
				} `xml:"Woonplaats"`
			} `xml:"bagObject"`
		} `xml:"stand"`
	} `xml:"standBestand"`
}

type geometrie struct {
	Text string `xml:",chardata"`
	Vlak struct {
		Text    string     `xml:",chardata"`
		Polygon gmlPolygon `xml:"Polygon"`
	} `xml:"vlak"`
	Multivlak struct {
		Text         string `xml:",chardata"`
		MultiSurface struct {
			Text          string `xml:",chardata"`
			SrsName       string `xml:"srsName,attr"`
			SrsDimension  string `xml:"srsDimension,attr"`
			SurfaceMember []struct {
				Text    string     `xml:",chardata"`
				Polygon gmlPolygon `xml:"Polygon"`
			} `xml:"surfaceMember"`
		} `xml:"MultiSurface"`
	} `xml:"multivlak"`
}

type gmlPolygon struct {
	Text     string `xml:",chardata"`
	Exterior []struct {
		Text       string `xml:",chardata"`
		LinearRing struct {
			Text    string `xml:",chardata"`
			PosList struct {
				Text  string `xml:",chardata"`
				Count string `xml:"count,attr"`
			} `xml:"posList"`
		} `xml:"LinearRing"`
	} `xml:"exterior"`
	Interior []struct {
		Text       string `xml:",chardata"`
		LinearRing struct {
			Text    string `xml:",chardata"`
			PosList struct {
				Text  string `xml:",chardata"`
				Count string `xml:"count,attr"`
			} `xml:"posList"`
		} `xml:"LinearRing"`
	} `xml:"interior"`
}

// type bagStand struct {
// 	BagInfo struct {
// 		Info string `xml:",chardata"`
// 	} `xml:"sl-bag-extract bagInfo"`
// 	StandBestand standBestand `xml:"sl standBestand"`
// }

// type standBestand struct {
// 	Dataset string             `xml:"sl dataset"`
// 	Inhoud  standBestandInhoud `xml:"sl inhoud"`
// 	Stand   []stand            `xml:"sl stand"`
// }

// type standBestandInhoud struct {
// 	Gebied      string      `xml:"sl gebied"`
// 	LeveringsId int         `xml:"sl leveringId"`
// 	ObjectTypen objectTypen `xml:"sl objectTypen"`
// }

// type objectTypen struct {
// 	Typen []string `xml:"sl objectType"`
// }

// type stand struct {
// 	BagObject bagObject `xml:"sl-bag-extract bagObject"`
// }

// type bagObject struct {
// 	Woonplaats bagWoonplaats `xml:"Objecten Woonplaats"`
// }

// type bagWoonplaats struct {
// 	Identificatie int       `xml:"Objecten Identificatie"`
// 	Naam          string    `xml:"Objecten naam"`
// 	Geometrie     geometrie `xml:"Objecten geometrie"`
// }

// type geometrie struct {
// 	Vlak      vlak      `xml:"Objecten vlak"`
// 	Multivlak multivlak `xml:"Objecten multi"`
// }

// type vlak struct {
// 	Polygons []gmlPolygon `xml:"gml Polygon"`
// }

// type multivlak struct {
// 	SurfaceMembers []surfaceMember `xml:"gml surfaceMember"`
// }

// type surfaceMember struct {
// 	Polygons []gmlPolygon `xml:"gml Polygon"`
// }

// type gmlPolygon struct {
// 	Exterior []exterior `xml:"gml exterior"`
// 	Interior []interior `xml:"gml interior"`
// }

// type exterior struct {
// 	Ring linearRing `xml:"gml LinearRing"`
// }

// type interior struct {
// 	Ring linearRing `xml:"gml LinearRing"`
// }

// type linearRing struct {
// 	Coordinates string `xml:"gml posList"`
// }

// Types for converting to
type woonplaats struct {
	Id    int
	Naam  string
	Grens grens
}

type grens struct {
	Exterior []string
	Interior []string
}

func addCoordinatesToGrens(Polygon gmlPolygon, Grens *grens) {
	for _, Exterior := range Polygon.Exterior {
		Grens.Exterior = append(Grens.Exterior, Exterior.LinearRing.PosList.Text)
	}
	for _, Interior := range Polygon.Interior {
		Grens.Interior = append(Grens.Interior, Interior.LinearRing.PosList.Text)
	}
}

func grensFromGeometrie(Geometrie geometrie) grens {
	var Grens grens
	for _, SurfaceMember := range Geometrie.Multivlak.MultiSurface.SurfaceMember {
		addCoordinatesToGrens(SurfaceMember.Polygon, &Grens)
	}
	addCoordinatesToGrens(Geometrie.Vlak.Polygon, &Grens)
	return Grens
}

func ImportWoonplaats(filename string) []woonplaats {
	var Result []woonplaats
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(xmlFile)

	// we initialize our Users array
	var BagStand bagStand
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'bagStand' which we defined above
	if xmlErr := xml.Unmarshal(byteValue, &BagStand); xmlErr != nil {
		fmt.Println(xmlErr)
	}

	for _, Stand := range BagStand.StandBestand.Stand {
		if Stand.BagObject.Woonplaats.Voorkomen.Voorkomen.EindGeldigheid != "" {
			continue
		}
		var Woonplaats woonplaats
		Woonplaats.Id, _ = strconv.Atoi(Stand.BagObject.Woonplaats.Identificatie.Text)
		Woonplaats.Naam = Stand.BagObject.Woonplaats.Naam
		Woonplaats.Grens = grensFromGeometrie(Stand.BagObject.Woonplaats.Geometrie)

		Result = append(Result, Woonplaats)
	}

	return Result
}
