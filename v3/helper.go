package v3

import (
	"log"
	"sync"
	"time"
)

type Helper struct {
	Client *RotorSoftClient
}

func (r *Helper) GetLastLogbookEntries(powerUnitIdentifiers []string) (*map[string]LogbookEntry, error) {
	lastLookbookEntries := make(map[string]LogbookEntry)
	for _, id := range powerUnitIdentifiers {
		res, err := r.Client.GetLastLogbookEntryForPowerUnit(id)
		if err != nil {
			return nil, err
		}
		lastLookbookEntries[id] = res.LastLogbookEntry
	}
	return &lastLookbookEntries, nil
}

func (r *Helper) GetLogbookEntriesByDateRangeForPowerUnits(powerUnitIdentifiers []string, from time.Time, to time.Time) (*map[string]*[]LogbookEntry, error) {
	lookbookEntries := make(map[string]*[]LogbookEntry)

	var wg sync.WaitGroup
	var mu sync.Mutex
	errCh := make(chan error, len(powerUnitIdentifiers))

	for _, id := range powerUnitIdentifiers {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()

			res, err := r.Client.Helper.GetLogbookEntriesByDateRangeForPowerUnit(id, from, to)
			if err != nil {
				errCh <- err
				return
			}

			mu.Lock()
			lookbookEntries[id] = res
			mu.Unlock()
		}(id)
	}

	wg.Wait()
	close(errCh)

	for err := range errCh {
		if err != nil {
			return nil, err
		}
	}

	return &lookbookEntries, nil
}

// Filtering the Logbook after start date in range
func (r *Helper) GetLogbookEntriesByDateRangeForPowerUnit(powerUnitIdentifier string, from time.Time, to time.Time) (*[]LogbookEntry, error) {
	lastLogbookEntry, err := r.Client.GetLastLogbookEntryForPowerUnit(powerUnitIdentifier)
	if err != nil {
		return nil, err
	}
	var logbookEntries []LogbookEntry
	logbookEntries = append(logbookEntries, lastLogbookEntry.LastLogbookEntry)
	fromDateIndex := -1
	toDateIndex := -1
	var prevLogbookEntry LogbookEntry
	for (fromDateIndex == -1) || (toDateIndex == -1) {
		if prevLogbookEntry.RSID == logbookEntries[len(logbookEntries)-1].RSID {
			break
		}
		prevLogbookEntry = logbookEntries[len(logbookEntries)-1]
		response, err := r.Client.GetLastLogbookEntriesForPowerUnitWithOffset(powerUnitIdentifier, logbookEntries[len(logbookEntries)-1].RSID, 100)
		if err != nil {
			return nil, err
		}
		logbookEntries = append(logbookEntries, response.LastLogbookEntries...)
		if toDateIndex == -1 {
			for i, entry := range logbookEntries {
				if entry.StartDate.Before(to) {
					log.Println(logbookEntries[i])
					toDateIndex = i
					break
				}

			}
		}

		if fromDateIndex == -1 {
			for i, entry := range logbookEntries {
				if entry.StartDate.Before(from) {
					log.Println(logbookEntries[i-1])
					fromDateIndex = i - 1
					break
				}

			}
		}
	}
	if fromDateIndex == -1 {
		fromDateIndex = len(logbookEntries)
	}
	slicedLogbookEntries := logbookEntries[toDateIndex : fromDateIndex+1]
	return &slicedLogbookEntries, nil
}
