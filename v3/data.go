package v3

import (
	"encoding/xml"
	"time"
)

type GetLastLogbookEntryForPowerUnitResponse struct {
	XMLName          xml.Name     `xml:"getLastLogbookEntryForPowerUnitResponse"`
	LastLogbookEntry LogbookEntry `xml:"LastLogbookEntry"`
}

type GetLastLogbookEntriesForPowerUnitResponse struct {
	XMLName            xml.Name       `xml:"getLastLogbookEntriesForPowerUnitResponse"`
	LastLogbookEntries []LogbookEntry `xml:"LastLogbookEntries"`
}

type GetLastLogbookEntriesForPowerUnitWithOffsetResponse struct {
	XMLName            xml.Name       `xml:"getLastLogbookEntriesForPowerUnitWithOffsetResponse"`
	LastLogbookEntries []LogbookEntry `xml:"LastLogbookEntries"`
}

type GetChildLogbookEntriesResponse struct {
	XMLName             xml.Name            `xml:"getChildLogbookEntriesResponse"`
	ChildLogbookEntries []ChildLogbookEntry `xml:"ChildLogbookEntries"`
}

type GetChildLogbookEntriesWithOffsetResponse struct {
	XMLName             xml.Name            `xml:"getChildLogbookEntriesWithOffsetResponse"`
	ChildLogbookEntries []ChildLogbookEntry `xml:"ChildLogbookEntries"`
}

type MultiLangText struct {
	XMLName xml.Name `xml:"multiLangText"`
	Items   []Item   `xml:"item"`
}

type Item struct {
	Key   string `xml:"key,attr"`
	Value string `xml:"value,attr"`
}

type ErrorState struct {
	XMLName xml.Name      `xml:"errorStateName"`
	Text    MultiLangText `xml:"multiLangText"`
}

type RotorsoftError struct {
	XMLName xml.Name      `xml:"rotorsoftErrorName"`
	Text    MultiLangText `xml:"multiLangText"`
}

type OriginalError struct {
	XMLName xml.Name      `xml:"originalErrorName"`
	Text    MultiLangText `xml:"multiLangText"`
}

type LogbookEntry struct {
	RSID                    int64          `xml:"RSID"`
	PowerUnitName           string         `xml:"powerUnitName"`
	PowerUnitIdentifier     string         `xml:"powerUnitIdentifier"`
	StartDate               time.Time      `xml:"startDate"`
	Duration                int64          `xml:"duration"`
	EndDate                 time.Time      `xml:"endDate"`
	ErrorStateName          ErrorState     `xml:"errorStateName"`
	Color                   string         `xml:"color"`
	RotorsoftErrorName      RotorsoftError `xml:"rotorsoftErrorName"`
	OriginalErrorIdentifier string         `xml:"originalErrorIdentifier"`
	OriginalErrorname       OriginalError  `xml:"originalErrorName"`
	ProductionLoss          float64        `xml:"productionLoss"`
	Comment                 string         `xml:"string"`
	ConfirmNeeded           bool           `xml:"confirmNeeded"`
	ConfirmUserId           int64          `xml:"confirmUserId"`
	ConfirmTime             time.Time      `xml:"confirmTime"`
	ParentRSID              int64          `xml:"Parent_RSID"`
	NoChildren              int64          `xml:"noChildren"`
}

type ChildLogbookEntry LogbookEntry
