package v3

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Predixxion/gosoap"
)

type RotorSoftClient struct {
	HTTPClient *http.Client
	URL        string
	Username   string
	Password   string
	Helper     *Helper
}

func (r *RotorSoftClient) GetChildLogbookEntries(parentLogEntryId int64, batchSize int) (*GetChildLogbookEntriesResponse, error) {
	gosoap.SetCustomEnvelope("soap", map[string]string{
		"xmlns:soap": "http://schemas.xmlsoap.org/soap/envelope/",
		"xmlns:v3":   "urn:service:drehpunkt:rotorsoft:soap:v3:logbook",
	})
	url := fmt.Sprintf("%s/soap/v3/logbook/?WSDL", r.URL)
	client, err := gosoap.SoapClient(url, r.HTTPClient)

	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	params := gosoap.ArrayParams{{"userName", r.Username}, {"password", r.Password}, {"parentLogEntryId", strconv.FormatInt(parentLogEntryId, 10)}, {"batchsize", strconv.Itoa(batchSize)}}

	res, err := client.Call("v3:getChildLogbookEntries", params)
	if err != nil {
		return nil, err
	}

	var response GetChildLogbookEntriesResponse
	err = xml.Unmarshal([]byte(res.Body), &response)
	if err != nil {
		log.Fatalf("xml.Unmarshal error: %s", err)
	}

	return &response, nil
}

func (r *RotorSoftClient) GetChildLogbookEntriesWithOffset(parentLogEntryId int64, offsetLogbookEntryId int64, batchSize int) (*GetChildLogbookEntriesWithOffsetResponse, error) {
	gosoap.SetCustomEnvelope("soap", map[string]string{
		"xmlns:soap": "http://schemas.xmlsoap.org/soap/envelope/",
		"xmlns:v3":   "urn:service:drehpunkt:rotorsoft:soap:v3:logbook",
	})
	url := fmt.Sprintf("%s/soap/v3/logbook/?WSDL", r.URL)
	client, err := gosoap.SoapClient(url, r.HTTPClient)

	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	params := gosoap.ArrayParams{{"userName", r.Username}, {"password", r.Password}, {"parentLogEntryId", strconv.FormatInt(parentLogEntryId, 10)}, {"offsetLogbookEntryId", strconv.FormatInt(offsetLogbookEntryId, 10)}, {"batchsize", strconv.Itoa(batchSize)}}

	res, err := client.Call("v3:getChildLogbookEntriesWithOffset", params)
	if err != nil {
		return nil, err
	}

	var response GetChildLogbookEntriesWithOffsetResponse
	err = xml.Unmarshal([]byte(res.Body), &response)
	if err != nil {
		log.Fatalf("xml.Unmarshal error: %s", err)
	}

	return &response, nil
}

func (r *RotorSoftClient) GetLastLogbookEntriesForPowerUnit(powerUnitIdentifier string, batchSize int) (*GetLastLogbookEntriesForPowerUnitResponse, error) {
	gosoap.SetCustomEnvelope("soap", map[string]string{
		"xmlns:soap": "http://schemas.xmlsoap.org/soap/envelope/",
		"xmlns:v3":   "urn:service:drehpunkt:rotorsoft:soap:v3:logbook",
	})
	url := fmt.Sprintf("%s/soap/v3/logbook/?WSDL", r.URL)
	client, err := gosoap.SoapClient(url, r.HTTPClient)

	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	params := gosoap.ArrayParams{{"userName", r.Username}, {"password", r.Password}, {"powerUnitIdentifier", powerUnitIdentifier}, {"batchsize", strconv.Itoa(batchSize)}}

	res, err := client.Call("v3:getLastLogbookEntriesForPowerUnit", params)
	if err != nil {
		return nil, err
	}

	var response GetLastLogbookEntriesForPowerUnitResponse
	err = xml.Unmarshal([]byte(res.Body), &response)
	if err != nil {
		log.Fatalf("xml.Unmarshal error: %s", err)
	}

	return &response, nil
}

func (r *RotorSoftClient) GetLastLogbookEntriesForPowerUnitWithOffset(powerUnitIdentifier string, offsetLogbookEntryId int64, batchSize int) (*GetLastLogbookEntriesForPowerUnitWithOffsetResponse, error) {
	gosoap.SetCustomEnvelope("soap", map[string]string{
		"xmlns:soap": "http://schemas.xmlsoap.org/soap/envelope/",
		"xmlns:v3":   "urn:service:drehpunkt:rotorsoft:soap:v3:logbook",
	})
	url := fmt.Sprintf("%s/soap/v3/logbook/?WSDL", r.URL)
	client, err := gosoap.SoapClient(url, r.HTTPClient)
	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	if batchSize > 100 {
		log.Fatalf("Batchsize cant be greater than 100!")
	}

	params := gosoap.ArrayParams{{"userName", r.Username}, {"password", r.Password}, {"powerUnitIdentifier", powerUnitIdentifier}, {"offsetLogbookEntryId", strconv.FormatInt(offsetLogbookEntryId, 10)}, {"batchsize", strconv.Itoa(batchSize)}}

	res, err := client.Call("v3:getLastLogbookEntriesForPowerUnitWithOffset", params)
	if err != nil {
		return nil, err
	}

	var response GetLastLogbookEntriesForPowerUnitWithOffsetResponse
	err = xml.Unmarshal([]byte(res.Body), &response)
	if err != nil {
		log.Fatalf("xml.Unmarshal error: %s", err)
	}

	return &response, nil
}

func (r *RotorSoftClient) GetLastLogbookEntryForPowerUnit(powerUnitIdentifier string) (*GetLastLogbookEntryForPowerUnitResponse, error) {
	gosoap.SetCustomEnvelope("soap", map[string]string{
		"xmlns:soap": "http://schemas.xmlsoap.org/soap/envelope/",
		"xmlns:v3":   "urn:service:drehpunkt:rotorsoft:soap:v3:logbook",
	})
	url := fmt.Sprintf("%s/soap/v3/logbook/?WSDL", r.URL)
	client, err := gosoap.SoapClient(url, r.HTTPClient)

	if err != nil {
		log.Fatalf("SoapClient error: %s", err)
	}
	params := gosoap.ArrayParams{{"userName", r.Username}, {"password", r.Password}, {"powerUnitIdentifier", powerUnitIdentifier}}

	res, err := client.Call("v3:getLastLogbookEntryForPowerUnit", params)
	if err != nil {
		return nil, err
	}

	var response GetLastLogbookEntryForPowerUnitResponse
	err = xml.Unmarshal([]byte(res.Body), &response)
	if err != nil {
		log.Fatalf("xml.Unmarshal error: %s", err)
	}

	return &response, nil
}
