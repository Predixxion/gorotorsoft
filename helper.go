package gorotorsoft

import (
	v0 "github.com/Predixxion/gorotorsoft/v0"
	"time"
)

func ConvertRotorSoftRawDataToTable(rawData v0.GetRawDataForPowerUnitsResponse, powerUnitMap map[string]string, dataFieldMap map[string]string) []v0.SensorValueTable {
	sensorValueTables := make(map[string]*v0.SensorValueTable)

	for _, dataRecord := range rawData.RawData.DataRecords {
		for _, recordField := range dataRecord.RecordFields {
			if _, ok := sensorValueTables[dataFieldMap[recordField.DataFieldIdentifier]]; !ok {
				sensorValueTables[dataFieldMap[recordField.DataFieldIdentifier]] = &v0.SensorValueTable{
					DataFieldIdentifier: dataFieldMap[recordField.DataFieldIdentifier],
					Records:             make(map[time.Time]map[string]string),
				}
			}

			table := sensorValueTables[dataFieldMap[recordField.DataFieldIdentifier]]

			if _, ok := table.Records[dataRecord.RecordTime]; !ok {
				table.Records[dataRecord.RecordTime] = make(map[string]string)
			}

			table.Records[dataRecord.RecordTime][powerUnitMap[dataRecord.RecordPowerUnitIdentifier]] = recordField.DataFieldValue
		}
	}

	var result []v0.SensorValueTable
	for _, table := range sensorValueTables {
		result = append(result, *table)
	}

	return result
}
