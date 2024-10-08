package v0

import "time"

func mergeResponses(response1, response2 GetRawDataForPowerUnitsResponse) GetRawDataForPowerUnitsResponse {
	response1.RawData.DataRecords = append(response1.RawData.DataRecords, response2.RawData.DataRecords...)
	return response1
}

func ConvertRotorSoftRawDataToTable(rawData GetRawDataForPowerUnitsResponse, powerUnitMap map[string]string, dataFieldMap map[string]string) []SensorValueTable {
	sensorValueTables := make(map[string]*SensorValueTable)

	for _, dataRecord := range rawData.RawData.DataRecords {
		for _, recordField := range dataRecord.RecordFields {
			if _, ok := sensorValueTables[dataFieldMap[recordField.DataFieldIdentifier]]; !ok {
				sensorValueTables[dataFieldMap[recordField.DataFieldIdentifier]] = &SensorValueTable{
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

	var result []SensorValueTable
	for _, table := range sensorValueTables {
		result = append(result, *table)
	}

	return result
}
