package service

import (
	"app/model"
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestSortPayload(t *testing.T) {
	tests := []struct {
		name      string
		input     model.SortRequestPayload
		expected  map[string]interface{}
		expectErr bool
	}{
		{
			name: "Valid string and number arrays",
			input: model.SortRequestPayload{
				SortKeys: []string{"fruits", "numbers"},
				Payload: map[string]interface{}{
					"fruits":  []interface{}{"banana", "apple", "orange"},
					"numbers": []interface{}{1333.0, 4.0, 2431.0, 7.0}, // The json package converts numbers to float64 when decoding
					"colors":  []interface{}{"red", "blue", "green"},
				},
			},
			expected: map[string]interface{}{
				"fruits":  []interface{}{"apple", "banana", "orange"},
				"numbers": []interface{}{4.0, 7.0, 1333.0, 2431.0},
				"colors":  []interface{}{"red", "blue", "green"},
			},
			expectErr: false,
		},
		{
			name: "Empty payload",
			input: model.SortRequestPayload{
				SortKeys: []string{"fruits"},
				Payload:  nil,
			},
			expected:  nil,
			expectErr: true,
		},
		{
			name: "Empty sort keys",
			input: model.SortRequestPayload{
				SortKeys: []string{},
				Payload: map[string]interface{}{
					"fruits": []interface{}{"banana", "apple"},
				},
			},
			expected:  nil,
			expectErr: true,
		},
		{
			name: "Unsupported data type",
			input: model.SortRequestPayload{
				SortKeys: []string{"fruits"},
				Payload: map[string]interface{}{
					"fruits": []interface{}{true, false},
				},
			},
			expected:  nil,
			expectErr: true,
		},
		{
			name: "Parallel processing with large number arrays",
			input: model.SortRequestPayload{
				SortKeys: []string{"numbers1", "numbers2"},
				Payload: map[string]interface{}{
					"numbers1": generateLargeNumberArray(10000),
					"numbers2": generateLargeNumberArray(10000),
				},
			},
			expected:  nil,
			expectErr: false,
		},
		{
			name: "Parallel processing with large string arrays",
			input: model.SortRequestPayload{
				SortKeys: []string{"strings1", "strings2"},
				Payload: map[string]interface{}{
					"strings1": generateLargeStringArray(10000),
					"strings2": generateLargeStringArray(10000),
				},
			},
			expected:  nil,
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SortPayload(tt.input)
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}

			if strings.Contains(tt.name, "Parallel processing with large") && !tt.expectErr {
				for key, value := range result {
					arr := value.([]interface{})
					if !isSorted(arr) {
						t.Errorf("array %s is not sorted", key)
					}
				}
			} else if !tt.expectErr && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected result: %v, got: %v", tt.expected, result)
			}
		})
	}
}

func generateLargeNumberArray(size int) []interface{} {
	arr := make([]interface{}, size)
	for i := 0; i < size; i++ {
		arr[i] = float64(size - i)
	}
	return arr
}

func generateLargeStringArray(size int) []interface{} {
	arr := make([]interface{}, size)
	for i := 0; i < size; i++ {
		arr[i] = fmt.Sprintf("string_%d", size-i)
	}
	return arr
}

func isSorted(arr []interface{}) bool {
	for i := 1; i < len(arr); i++ {
		switch arr[i-1].(type) {
		case float64:
			if arr[i-1].(float64) > arr[i].(float64) {
				return false
			}
		case string:
			if arr[i-1].(string) > arr[i].(string) {
				return false
			}
		}
	}
	return true
}
