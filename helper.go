package gorotorsoft

import (
	v0 "github.com/Predixxion/gorotorsoft/v0"
	"time"
)

func ConvertRotorSoftRawDataToTable(rawData v0.GetRawDataForPowerUnitsResponse, powerUnitMap map[string]string, dataFieldMap map[string]string) []v0.SensorValueTable {
	sensorValueTables := make(map[string]*v0.SensorValueTable)

	allTimestamps := make(map[time.Time]struct{})
	for _, dataRecord := range rawData.RawData.DataRecords {
		allTimestamps[dataRecord.RecordTime] = struct{}{}
	}

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

	for _, table := range sensorValueTables {
		for timestamp := range allTimestamps {
			if _, ok := table.Records[timestamp]; !ok {
				table.Records[timestamp] = make(map[string]string)
			}
			for _, turbineID := range powerUnitMap {
				if _, exists := table.Records[timestamp][turbineID]; !exists {
					table.Records[timestamp][turbineID] = ""
				}
			}
		}
	}

	var result []v0.SensorValueTable
	for _, table := range sensorValueTables {
		result = append(result, *table)
	}

	return result
}
