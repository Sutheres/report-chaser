package models

import "encoding/xml"

type XBRLParser struct {
	XMLName xml.Name        `xml:"xbrli\:xbrl"`
	Ctx     []ReportContext `xml:"context"`
	Origin  []string        `xml:"origin"`
}

type ReportContext struct {
	XMLName xml.Name `xml:"xbrli\:context,omitempty"`
	ID      string   `xml:"id,attr"`
	Entity  Entity   `xml:"entity,omitempty"`
	Period  Period   `xml:"period,omitempty"`
}

type Entity struct {
	XMLName    xml.Name   `xml:"xbrli\:entity,omitempty"`
	Identifier Identifier `xml:"identifier,omitempty"`
	Segment    Segment    `xml:"segment,omitempty"`
}

type Period struct {
	XMLName   xml.Name `xml:"xbrli\:period,omitempty"`
	StartDate string   `xml:"startDate,omitempty"`
	EndDate   string   `xml:"endDate,omitempty"`
	Instant string `xml:"instant,omitempty"`
}

type Identifier struct {
	XMLName xml.Name `xml:"xbrli\:identifier,omitempty"`
	Scheme  string   `xml:"scheme,attr,omitempty"`
	Value   string   `xml:",chardata"`
}

type Segment struct {
	XMLName        xml.Name       `xml:"xbrli\:segment"`
	ExplicitMember ExplicitMember `xml:"explicitMember"`
}

type ExplicitMember struct {
	XMLName   xml.Name `xml:"xbrldi\:explicitMember"`
	Dimension string   `xml:"dimension,attr"`
	Value     string   `xml:",chardata"`
}

type Instant struct {
	XMLName xml.Name `xml:"xbrli\:instant"`
	Value string `xml:",chardata"`
}