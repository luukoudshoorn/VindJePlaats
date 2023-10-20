package bagImporter

import "encoding/xml"

//Struct was first generated at https://blog.kowalczyk.info/tools/xmltogo/, then adapted for reusability
type bagStand struct {
	XMLName            xml.Name `xml:"bagStand"`
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
		GebiedRegistratief struct {
			GebiedNLD string `xml:"Gebied-NLD"`
		} `xml:"Gebied-Registratief"`
		LVCExtract struct {
			StandTechnischeDatum string `xml:"StandTechnischeDatum"`
		} `xml:"LVC-Extract"`
	} `xml:"bagInfo"`
	StandBestand struct {
		Dataset string `xml:"dataset"`
		Inhoud  struct {
			Gebied      string `xml:"gebied"`
			LeveringsId int64  `xml:"leveringsId"`
			ObjectTypen struct {
				ObjectType string `xml:"objectType"`
			} `xml:"objectTypen"`
		} `xml:"inhoud"`
		Stand []struct {
			BagObject struct {
				Woonplaats      woonplaats      `xml:"Woonplaats"`
				Pand            pand            `xml:"Pand"`
				Verblijfsobject verblijfsobject `xml:"Verblijfsobject"`
			} `xml:"bagObject"`
		} `xml:"stand"`
	} `xml:"standBestand"`
}

type woonplaats struct {
	Identificatie struct {
		Id     int    `xml:",chardata"`
		Domein string `xml:"domein,attr"`
	} `xml:"identificatie"`
	Naam           string    `xml:"naam"`
	Geometrie      geometrie `xml:"geometrie"`
	Status         string    `xml:"status"`
	Geconstateerd  string    `xml:"geconstateerd"`
	Documentdatum  string    `xml:"documentdatum"`
	Documentnummer string    `xml:"documentnummer"`
	Voorkomen      voorkomen `xml:"Voorkomen"`
}

type verblijfsobject struct {
	Text               string `xml:",chardata"`
	HeeftAlsHoofdadres struct {
		NummeraanduidingRef struct {
			Text   string `xml:",chardata"`
			Domein string `xml:"domein,attr"`
		} `xml:"NummeraanduidingRef"`
	} `xml:"heeftAlsHoofdadres"`
	Voorkomen     voorkomen `xml:"Voorkomen"`
	Identificatie struct {
		Id     int64  `xml:",chardata"`
		Domein string `xml:"domein,attr"`
	} `xml:"identificatie"`
	Geometrie       geometrie `xml:"geometrie"`
	Gebruiksdoel    string    `xml:"gebruiksdoel"`
	Oppervlakte     string    `xml:"oppervlakte"`
	Status          string    `xml:"status"`
	Geconstateerd   string    `xml:"geconstateerd"`
	Documentdatum   string    `xml:"documentdatum"`
	Documentnummer  string    `xml:"documentnummer"`
	MaaktDeelUitVan struct {
		PandRef struct {
			Id     int64  `xml:",chardata"`
			Domein string `xml:"domein,attr"`
		} `xml:"PandRef"`
	} `xml:"maaktDeelUitVan"`
}

type pand struct {
	Identificatie struct {
		Id     int64  `xml:",chardata"`
		Domein string `xml:"domein,attr"`
	} `xml:"identificatie"`
	Geometrie              geometrie `xml:"geometrie"`
	OorspronkelijkBouwjaar string    `xml:"oorspronkelijkBouwjaar"`
	Status                 string    `xml:"status"`
	Geconstateerd          string    `xml:"geconstateerd"`
	Documentdatum          string    `xml:"documentdatum"`
	Documentnummer         string    `xml:"documentnummer"`
	Voorkomen              voorkomen `xml:"Voorkomen"`
}

type voorkomen struct {
	HistorieVoorkomen struct {
		Voorkomenidentificatie int    `xml:"voorkomenidentificatie"`
		BeginGeldigheid        string `xml:"beginGeldigheid"`
		EindGeldigheid         string `xml:"eindGeldigheid"`
		TijdstipRegistratie    string `xml:"tijdstipRegistratie"`
		BeschikbaarLV          struct {
			TijdstipRegistratieLV     string `xml:"tijdstipRegistratieLV"`
			TijdstipEindRegistratieLV string `xml:"tijdstipEindRegistratieLV"`
		} `xml:"BeschikbaarLV"`
	} `xml:"Voorkomen"`
}

type geometrie struct {
	Vlak struct {
		Polygon gmlPolygon `xml:"Polygon"`
	} `xml:"vlak"`
	Multivlak struct {
		MultiSurface struct {
			SrsName       string `xml:"srsName,attr"`
			SrsDimension  int    `xml:"srsDimension,attr"`
			SurfaceMember []struct {
				Polygon gmlPolygon `xml:"Polygon"`
			} `xml:"surfaceMember"`
		} `xml:"MultiSurface"`
	} `xml:"multivlak"`
	Polygon gmlPolygon `xml:"Polygon"`
	Punt    struct {
		Point struct {
			SrsName      string `xml:"srsName,attr"`
			SrsDimension string `xml:"srsDimension,attr"`
			Pos          string `xml:"pos"`
		} `xml:"Point"`
	} `xml:"punt"`
}

type gmlPolygon struct {
	SrsName      string `xml:"srsName,attr"`
	SrsDimension int    `xml:"srsDimension,attr"`
	Exterior     []struct {
		LinearRing struct {
			PosList struct {
				Coordinates string `xml:",chardata"`
				Count       int    `xml:"count,attr"`
			} `xml:"posList"`
		} `xml:"LinearRing"`
	} `xml:"exterior"`
	Interior []struct {
		LinearRing struct {
			PosList struct {
				Coordinates string `xml:",chardata"`
				Count       int    `xml:"count,attr"`
			} `xml:"posList"`
		} `xml:"LinearRing"`
	} `xml:"interior"`
}
