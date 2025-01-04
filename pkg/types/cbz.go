package types

import "encoding/xml"

type ComicInfo struct {
	Title       string `xml:"Title,omitempty"`
	Series      string `xml:"Series,omitempty"`
	Number      int    `xml:"Number,omitempty"`
	Summary     string `xml:"Summary,omitempty"`
	Year        int    `xml:"Year,omitempty"`
	Month       int    `xml:"Month,omitempty"`
	Day         int    `xml:"Day,omitempty"`
	Writer      string `xml:"Writer,omitempty"`
	Penciller   string `xml:"Penciller,omitempty"`
	Publisher   string `xml:"Publisher,omitempty"`
	Genre       string `xml:"Genre,omitempty"`
	PageCount   int    `xml:"PageCount,omitempty"`
	LanguageISO string `xml:"LanguageISO,omitempty"`
	Format      string `xml:"Format,omitempty"`
	AgeRating   int    `xml:"AgeRating,omitempty"`
	GTIN        string `xml:"GTIN,omitempty"`
}

type ComicInfoWrapper struct {
	XMLName xml.Name `xml:"ComicInfo"`
	XSI     string   `xml:"xmlns:xsi,attr"`
	XSD     string   `xml:"xmlns:xsd,attr"`
	ComicInfo
}
