package sicgolib

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadFromRequestBody(t *testing.T) {	

	tests := []struct{
		name string
		request io.Reader
		expected error		
	}{
		{
			name: "Full filled field",
			request: strings.NewReader(`
				{
					"data_string": "Siapa namaku?",
					"data_number": 1202,
					"data_boolean": true,
					"data_array_string": [
						"shopee.com",
						"google.com",
						"tiket.com",
						"dana.com"
					],
					"data_array_number": [12, 2, 2022],
					"data_nested_map": {
						"kucing": "meong"
					},
					"data_null": null
				}
			`),
			expected: nil,
		},
	}

	for _, test := range tests {
		var data interface{}
		t.Run(test.name, func(t *testing.T) {
			err := ReadFromRequestBody(test.request, &data)
			fmt.Printf("\nTest name: \n%s\n\nTest result: \n%v\n\n", test.name, data)
			assert.Equal(t, test.expected, err)
		})
	}
}