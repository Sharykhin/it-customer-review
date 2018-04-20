package analyzer

import (
	"testing"

	"sync"

	"encoding/json"

	"github.com/stretchr/testify/assert"
)

func TestAnalyzer_basicAuth(t *testing.T) {
	res := basicAuth("foo", "bar")
	assert.Equal(t, "Zm9vOmJhcg==", res)
}

func TestAnalyzer_calculateScore(t *testing.T) {
	tt := []struct {
		name          string
		jsonStr       []byte
		expectedValue int64
	}{
		{
			name: "a bit negative",
			jsonStr: []byte(`
					{
					  "document_tone": {
						"tones": [
						  {
							"score": 0.520417,
							"tone_id": "sadness",
							"tone_name": "Sadness"
						  },
						  {
							"score": 0.822231,
							"tone_id": "tentative",
							"tone_name": "Tentative"
						  }
						]
					  }
					}`),
			expectedValue: 40,
		},
		{
			name: "joyable",
			jsonStr: []byte(`
					{
					  "document_tone": {
						"tones": [
						  {
							"score": 0.880435,
							"tone_id": "joy",
							"tone_name": "Joy"
						  }
						]
					  }
					}`),
			expectedValue: 88,
		},
		{
			name: "fully anger",
			jsonStr: []byte(`
					{
					  "document_tone": {
						"tones": [
						  {
							"score": 1,
							"tone_id": "anger",
							"tone_name": "Anger"
						  },
						  {
							"score": 0.916667,
							"tone_id": "sadness",
							"tone_name": "Sadness"
						  },
						  {
							"score": 0.931034,
							"tone_id": "fear",
							"tone_name": "Fear"
						  }
						]
					  }
					}`),
			expectedValue: 0,
		},
		{
			name: "fully sad",
			jsonStr: []byte(`
					{
					  "document_tone": {
						"tones": [
						  {
							"score": 1,
							"tone_id": "sadness",
							"tone_name": "Sadness"
						  }
						]
					  }
					}`),
			expectedValue: 0,
		},
		{
			name: "more positive",
			jsonStr: []byte(`
					{
					  "document_tone": {
						"tones": [
						  {
							"score": 0.873263,
							"tone_id": "tentative",
							"tone_name": "Tentative"
						  }
						]
					  }
					}`),
			expectedValue: 69,
		},
	}

	var wg sync.WaitGroup
	for _, tc := range tt {
		wg.Add(1)
		go t.Run(tc.name, func(t *testing.T) {
			defer wg.Done()
			var d documentAnalysis
			err := json.Unmarshal(tc.jsonStr, &d)
			if err != nil {
				t.Fatal(err)
			}
			v := calculateScore(d)
			assert.Equal(t, tc.expectedValue, v)
		})
	}
	wg.Wait()
}
