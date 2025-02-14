package service

import (
	"app/model"
	"fmt"
	"sort"
	"sync"

	"golang.org/x/sync/errgroup"
)

const (
	maxGoroutines = 8
)

// ErrInvalidPayload represents an error for invalid payload data
type ErrInvalidPayload struct {
	message string
}

func (e *ErrInvalidPayload) Error() string {
	return e.message
}

// validatePayload checks if the input payload is valid
func validatePayload(payload model.SortRequestPayload) error {
	if payload.Payload == nil {
		return &ErrInvalidPayload{message: "payload cannot be nil"}
	}
	if len(payload.SortKeys) == 0 {
		return &ErrInvalidPayload{message: "sort keys cannot be empty"}
	}
	return nil
}

// SortPayload sorts arrays based on sortKeys either sequentially or in parallel depending on the size
func SortPayload(payload model.SortRequestPayload) (map[string]interface{}, error) {
	if err := validatePayload(payload); err != nil {
		return nil, err
	}

	sortedPayload := deepCopy(payload.Payload)

	workerCount := min(len(payload.SortKeys), maxGoroutines)
	if workerCount < 2 {
		return processSortSequential(sortedPayload, payload.SortKeys)
	}

	return processSortParallel(sortedPayload, payload.SortKeys, workerCount)
}

// processSortParallel handles parallel sorting of multiple arrays
func processSortParallel(payload map[string]interface{}, sortKeys []string, workerCount int) (map[string]interface{}, error) {
	var g errgroup.Group
	mutex := &sync.Mutex{}

	workChan := make(chan string, len(sortKeys))
	for _, key := range sortKeys {
		workChan <- key
	}
	close(workChan)

	for i := 0; i < workerCount; i++ {
		g.Go(func() error {
			for key := range workChan {
				if err := sortArrayConcurrent(payload, key, mutex); err != nil {
					return fmt.Errorf("error sorting key %s: %w", key, err)
				}
			}
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		return nil, err
	}

	return payload, nil
}

// sortArrayConcurrent handles concurrent sorting of a single array
func sortArrayConcurrent(payload map[string]interface{}, key string, mutex *sync.Mutex) error {
	mutex.Lock()
	val, exists := payload[key]
	mutex.Unlock()

	if !exists {
		return nil
	}

	arr, ok := val.([]interface{})
	if !ok || len(arr) == 0 {
		return nil
	}

	var sortedArr []interface{}
	switch arr[0].(type) {
	case string:
		sortedArr = sortStringArray(arr)
	case float64:
		sortedArr = sortFloatArray(arr)
	default:
		return fmt.Errorf("unsupported type for key %s", key)
	}

	mutex.Lock()
	payload[key] = sortedArr
	mutex.Unlock()

	return nil
}

// sortStringArray sorts an array of strings
func sortStringArray(arr []interface{}) []interface{} {
	strs := make([]string, len(arr))
	for i, v := range arr {
		strs[i] = v.(string)
	}
	sort.Strings(strs)

	result := make([]interface{}, len(strs))
	for i, v := range strs {
		result[i] = v
	}
	return result
}

// sortFloatArray sorts an array of float64
func sortFloatArray(arr []interface{}) []interface{} {
	nums := make([]float64, len(arr))
	for i, v := range arr {
		nums[i] = v.(float64)
	}
	sort.Float64s(nums)

	result := make([]interface{}, len(nums))
	for i, v := range nums {
		result[i] = v
	}
	return result
}

// processSortSequential handles sequential processing of all arrays
func processSortSequential(payload map[string]interface{}, sortKeys []string) (map[string]interface{}, error) {
	for _, key := range sortKeys {
		if val, exists := payload[key]; exists {
			if arr, ok := val.([]interface{}); ok && len(arr) > 0 {
				var sortedArr []interface{}
				switch arr[0].(type) {
				case string:
					sortedArr = sortStringArray(arr)
				case float64:
					sortedArr = sortFloatArray(arr)
				default:
					return nil, fmt.Errorf("unsupported type for key %s", key)
				}
				payload[key] = sortedArr
			}
		}
	}
	return payload, nil
}

// deepCopy creates a deep copy of the input map
func deepCopy(m map[string]interface{}) map[string]interface{} {
	cp := make(map[string]interface{})
	for k, v := range m {
		switch vt := v.(type) {
		case []interface{}:
			tmp := make([]interface{}, len(vt))
			copy(tmp, vt)
			cp[k] = tmp
		default:
			cp[k] = v
		}
	}
	return cp
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
