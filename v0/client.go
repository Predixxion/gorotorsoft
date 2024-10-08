package v0

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Predixxion/gosoap"
)

type RotorSoftClient struct {
	HTTPClient *http.Client
	URL        string
	Username   string
	Password   string
}

func (r *RotorSoftClient) GetAllEndPoints() GetAllEndPointsResponse {
	gosoap.SetCustomEnvelope("soap", map[string]string{
		"xmlns:soap": "http://schemas.xmlsoap.org/soap/envelope/",
		"xmlns:v0":   "urn:service:drehpunkt:rotorsoft:soap:v0",
	})
	url := fmt.Sprintf("%s/soap/v0/?WSDL", r.URL)
	client, err := gosoap.SoapClient(url, r.HTTPClient)

	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}

	params := gosoap.ArrayParams{}
	res, err := client.Call("v0:getAllEndPoints", params)
	if err != nil {
		log.Println(string(gosoap.GetPayloadFromError(err)))
		log.Fatalf("Call error: %s", err)
	}

	var response GetAllEndPointsResponse
	err = xml.Unmarshal([]byte(res.Body), &response)
	if err != nil {
		log.Fatalf("xml.Unmarshal error: %s", err)
	}

	return response
}

func (r *RotorSoftClient) GetAllPowerUnits() GetAllPowerUnitsResponse {
	gosoap.SetCustomEnvelope("soap", map[string]string{
		"xmlns:soap": "http://schemas.xmlsoap.org/soap/envelope/",
		"xmlns:v0":   "urn:service:drehpunkt:rotorsoft:soap:v0",
	})
	url := fmt.Sprintf("%s/soap/v0/?WSDL", r.URL)
	client, err := gosoap.SoapClient(url, r.HTTPClient)

	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}

	params := gosoap.ArrayParams{{"userName", r.Username}, {"password", r.Password}}
	res, err := client.Call("v0:getAllPowerUnits", params)
	if err != nil {
		log.Println(string(gosoap.GetPayloadFromError(err)))
		log.Fatalf("Call error: %s", err)
	}

	var response GetAllPowerUnitsResponse
	err = xml.Unmarshal([]byte(res.Body), &response)
	if err != nil {
		log.Fatalf("xml.Unmarshal error: %s", err)
	}

	return response
}

func (r *RotorSoftClient) Ping() PingResponse {
	gosoap.SetCustomEnvelope("soap", map[string]string{
		"xmlns:soap": "http://schemas.xmlsoap.org/soap/envelope/",
		"xmlns:v0":   "urn:service:drehpunkt:rotorsoft:soap:v0",
	})
	url := fmt.Sprintf("%s/soap/v0/?WSDL", r.URL)
	client, err := gosoap.SoapClient(url, r.HTTPClient)

	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}

	params := gosoap.ArrayParams{}
	res, err := client.Call("v0:ping", params)
	if err != nil {
		log.Println(string(gosoap.GetPayloadFromError(err)))
		log.Fatalf("Call error: %s", err)
	}

	var response PingResponse
	err = xml.Unmarshal([]byte(res.Body), &response)
	if err != nil {
		log.Fatalf("xml.Unmarshal error: %s", err)
	}

	return response
}

func (r *RotorSoftClient) GetDataClassesForPowerUnits(powerUnitIdentifier []string) GetDataClassesForPowerUnitsResponse {
	gosoap.SetCustomEnvelope("soap", map[string]string{
		"xmlns:soap": "http://schemas.xmlsoap.org/soap/envelope/",
		"xmlns:v0":   "urn:service:drehpunkt:rotorsoft:soap:v0:rawdata",
	})
	url := fmt.Sprintf("%s/soap/v0/rawdata/?WSDL", r.URL)
	client, err := gosoap.SoapClient(url, r.HTTPClient)

	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	params := gosoap.ArrayParams{{"userName", r.Username}, {"password", r.Password}}

	for _, id := range powerUnitIdentifier {
		params = append(params, [2]interface{}{"powerUnitIdentifier", id})
	}

	res, err := client.Call("v0:getDataClassesForPowerUnits", params)
	if err != nil {
		log.Println(string(gosoap.GetPayloadFromError(err)))
		log.Fatalf("Call error: %s", err)
	}

	var response GetDataClassesForPowerUnitsResponse
	err = xml.Unmarshal([]byte(res.Body), &response)
	if err != nil {
		log.Fatalf("xml.Unmarshal error: %s", err)
	}

	return response
}

func (r *RotorSoftClient) GetDataFieldsForPowerUnits(powerUnitIdentifier []string, dataClassIdentifier []string) GetDataFieldsForPowerUnitsResponse {
	gosoap.SetCustomEnvelope("soap", map[string]string{
		"xmlns:soap": "http://schemas.xmlsoap.org/soap/envelope/",
		"xmlns:v0":   "urn:service:drehpunkt:rotorsoft:soap:v0:rawdata",
	})
	url := fmt.Sprintf("%s/soap/v0/rawdata/?WSDL", r.URL)
	client, err := gosoap.SoapClient(url, r.HTTPClient)

	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	params := gosoap.ArrayParams{{"userName", r.Username}, {"password", r.Password}}

	for _, id := range powerUnitIdentifier {
		params = append(params, [2]interface{}{"powerUnitIdentifier", id})
	}

	for _, id := range dataClassIdentifier {
		params = append(params, [2]interface{}{"dataClassIdentifier", id})
	}

	res, err := client.Call("v0:getDataFieldsForPowerUnits", params)
	if err != nil {
		log.Println(string(gosoap.GetPayloadFromError(err)))
		log.Fatalf("Call error: %s", err)
	}

	var response GetDataFieldsForPowerUnitsResponse
	err = xml.Unmarshal([]byte(res.Body), &response)
	if err != nil {
		log.Fatalf("xml.Unmarshal error: %s", err)
	}

	return response
}

func (r *RotorSoftClient) GetRawDataForPowerUnits(powerUnitIdentifier []string, from time.Time, to time.Time, dataClassIdent string, dataFieldIdents []string) (GetRawDataForPowerUnitsResponse, error) {
	const maxBatchSize = 10

	var combinedResponse GetRawDataForPowerUnitsResponse

	for i := 0; i < len(powerUnitIdentifier); i += maxBatchSize {
		end := i + maxBatchSize
		if end > len(powerUnitIdentifier) {
			end = len(powerUnitIdentifier)
		}

		batch := powerUnitIdentifier[i:end]

		response, err := r.getRawDataForPowerUnitBatch(batch, from, to, dataClassIdent, dataFieldIdents)
		if err != nil {
			return GetRawDataForPowerUnitsResponse{}, err
		}

		combinedResponse = mergeResponses(combinedResponse, response)
	}

	return combinedResponse, nil
}

func (r *RotorSoftClient) getRawDataForPowerUnitBatch(batch []string, from time.Time, to time.Time, dataClassIdent string, dataFieldIdents []string) (GetRawDataForPowerUnitsResponse, error) {
	gosoap.SetCustomEnvelope("soap", map[string]string{
		"xmlns:soap": "http://schemas.xmlsoap.org/soap/envelope/",
		"xmlns:v0":   "urn:service:drehpunkt:rotorsoft:soap:v0:rawdata",
	})
	url := fmt.Sprintf("%s/soap/v0/rawdata/?WSDL", r.URL)
	client, err := gosoap.SoapClient(url, r.HTTPClient)

	if err != nil {
		return GetRawDataForPowerUnitsResponse{}, fmt.Errorf("SoapClient error: %w", err)
	}

	params := gosoap.ArrayParams{{"userName", r.Username}, {"password", r.Password}}

	for _, id := range batch {
		params = append(params, [2]interface{}{"powerUnitIdentifier", id})
	}

	params = append(params, [2]interface{}{"from", from.Format("2006-01-02T15:04:05-07:00")})
	params = append(params, [2]interface{}{"to", to.Format("2006-01-02T15:04:05-07:00")})
	params = append(params, [2]interface{}{"dataClassIdent", dataClassIdent})

	for _, id := range dataFieldIdents {
		params = append(params, [2]interface{}{"dataFieldIdents", id})
	}

	res, err := client.Call("v0:getRawDataForPowerUnits", params)
	if err != nil {
		log.Println(string(gosoap.GetPayloadFromError(err)))
		return GetRawDataForPowerUnitsResponse{}, fmt.Errorf("Call error: %w", err)
	}

	var response GetRawDataForPowerUnitsResponse
	err = xml.Unmarshal([]byte(res.Body), &response)
	if err != nil {
		return GetRawDataForPowerUnitsResponse{}, fmt.Errorf("xml.Unmarshal error: %w", err)
	}

	return response, nil
}
