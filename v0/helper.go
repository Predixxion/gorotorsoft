package v0

func mergeResponses(response1, response2 GetRawDataForPowerUnitsResponse) GetRawDataForPowerUnitsResponse {
	response1.RawData.DataRecords = append(response1.RawData.DataRecords, response2.RawData.DataRecords...)
	return response1
}
