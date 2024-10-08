package v0

import (
	"encoding/xml"
	"time"
)

type GetAllEndPointsResponse struct {
	XMLName   xml.Name `xml:"getAllEndPointsResponse"`
	EndPoints []string `xml:"return"`
}

type GetAllPowerUnitsResponse struct {
	XMLName xml.Name    `xml:"getAllPowerUnitsResponse"`
	Units   []PowerUnit `xml:"return"`
}

type PingResponse struct {
	XMLName xml.Name `xml:"pingResponse"`
}

type PowerUnit struct {
	PowerUnitName       string `xml:"powerUnitName"`
	PowerUnitIdentifier string `xml:"powerUnitIdentifier"`
}

type GetDataFieldsForPowerUnitsResponse struct {
	XMLName    xml.Name    `xml:"getDataFieldsForPowerUnitsResponse"`
	DataFields []DataField `xml:"DataFields"`
}

type DataField struct {
	PowerUnitIdentifier string `xml:"powerUnitIdentifier"`
	DataFieldName       string `xml:"dataFieldName"`
	DataClassIdentifier string `xml:"dataClassIdentifier"`
	DataFieldIdentifier int    `xml:"dataFieldIdentifier"`
}

type DataClass struct {
	PowerUnitIdentifier string `xml:"powerUnitIdentifier"`
	DataClassName       string `xml:"dataClassName"`
	DataClassIdentifier string `xml:"dataClassIdentifier"`
}

type GetDataClassesForPowerUnitsResponse struct {
	XMLName     xml.Name    `xml:"getDataClassesForPowerUnitsResponse"`
	DataClasses []DataClass `xml:"DataClasses"`
}

type RecordField struct {
	DataFieldIdentifier string `xml:"dataFieldIdentifier"`
	DataFieldValue      string `xml:"dataFieldValue"`
}

type DataRecord struct {
	RecordPowerUnitIdentifier string        `xml:"recordPowerUnitIdentifier"`
	RecordTime                time.Time     `xml:"recordTime"`
	RecordOriginalByScada     bool          `xml:"recordOriginalByScada"`
	RecordFields              []RecordField `xml:"recordFields"`
}

type RawData struct {
	From                time.Time    `xml:"from"`
	To                  time.Time    `xml:"to"`
	DataClassIdentifier string       `xml:"dataClassIdentifier"`
	DataRecords         []DataRecord `xml:"dataRecords"`
}

type GetRawDataForPowerUnitsResponse struct {
	XMLName xml.Name `xml:"getRawDataForPowerUnitsResponse"`
	RawData RawData  `xml:"RawData"`
}

type SensorValueTable struct {
	DataFieldIdentifier string
	Records             map[time.Time]map[string]string // Row (timestamp) -> Column (power unit) -> Value (data field value)
}
