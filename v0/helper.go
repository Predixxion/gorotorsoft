package v0

import (
	"fmt"
	"time"
)

type Helper struct {
	Client *RotorSoftClient
}

func mergeResponses(response1, response2 GetRawDataForPowerUnitsResponse) GetRawDataForPowerUnitsResponse {
	response1.RawData.DataRecords = append(response1.RawData.DataRecords, response2.RawData.DataRecords...)
	return response1
}

func (r *Helper) GetRawDataForPowerUnitsExtended(powerUnitIdentifier []string, from time.Time, to time.Time, dataClassIdent string, dataFieldIdents []string) (GetRawDataForPowerUnitsResponse, error) {
	const maxBatchSize = 10
	const maxDaysPerRequest = 7

	var combinedResponse GetRawDataForPowerUnitsResponse
	var allErrors []error
	successfulRequest := false

	if dataClassIdent == "10m" {
		for start := from; start.Before(to); {
			end := start.AddDate(0, 0, maxDaysPerRequest)
			if end.After(to) {
				end = to
			}
			for i := 0; i < len(powerUnitIdentifier); i += maxBatchSize {
				endBatch := i + maxBatchSize
				if endBatch > len(powerUnitIdentifier) {
					endBatch = len(powerUnitIdentifier)
				}

				batch := powerUnitIdentifier[i:endBatch]

				response, err := r.Client.GetRawDataForPowerUnits(batch, start, end, dataClassIdent, dataFieldIdents)
				if err != nil {
					allErrors = append(allErrors, err)
					continue
				}

				successfulRequest = true
				combinedResponse = mergeResponses(combinedResponse, response)
			}

			start = end
		}
	} else {
		for i := 0; i < len(powerUnitIdentifier); i += maxBatchSize {
			endBatch := i + maxBatchSize
			if endBatch > len(powerUnitIdentifier) {
				endBatch = len(powerUnitIdentifier)
			}

			batch := powerUnitIdentifier[i:endBatch]

			response, err := r.Client.GetRawDataForPowerUnits(batch, from, to, dataClassIdent, dataFieldIdents)
			if err != nil {
				allErrors = append(allErrors, err)
				continue
			}

			successfulRequest = true
			combinedResponse = mergeResponses(combinedResponse, response)
		}
	}

	if !successfulRequest {
		return GetRawDataForPowerUnitsResponse{}, fmt.Errorf("all requests failed: %v", allErrors)
	}

	return combinedResponse, nil
}
