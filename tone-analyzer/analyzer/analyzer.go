package analyzer

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"net/http"
	"os"
	"time"

	"github.com/Sharykhin/it-customer-review/tone-analyzer/util"
)

const (
	url        = "https://gateway.watsonplatform.net/tone-analyzer/api/v3/tone"
	sadness    = 0.3
	fear       = 0.3
	joy        = 1.0
	anger      = 0.2
	analytical = 0.8
	confident  = 0.8
	tentative  = 0.8
)

var (
	// Analyzer keeps a reference to a private struct that implements methods for getting
	// tone description and return final score
	Analyzer = analyzer{client: newCustomClient()}
	username = os.Getenv("TONE_ANALYZER_USERNAME")
	password = os.Getenv("TONE_ANALYZER_PASSWORD")
	tones    = map[string]float64{
		"sadness":    sadness,
		"fear":       fear,
		"joy":        joy,
		"anger":      anger,
		"analytical": analytical,
		"confident":  confident,
		"tentative":  tentative,
	}
	toneEmotion = map[string]bool{
		"sadness":    true,
		"fear":       true,
		"joy":        true,
		"anger":      true,
		"analytical": false,
		"confident":  false,
		"tentative":  false,
	}
	toneIDs = map[string]string{
		"sadness":    "sadness",
		"fear":       "fear",
		"joy":        "joy",
		"anger":      "anger",
		"analytical": "analytical",
		"confident":  "confident",
		"tentative":  "tentative",
	}
)

type (
	// HTTPProvider is a specific interface that describes necessary method for making remote requests
	HTTPProvider interface {
		NewRequest(method, url string, body io.Reader) (*http.Request, error)
		Do(req *http.Request) (*http.Response, error)
	}
	analyzer struct {
		client HTTPProvider
	}

	customClient struct {
		*http.Client
	}

	payload struct {
		Text string `json:"text"`
	}

	documentAnalysis struct {
		DocumentTone tonesList `json:"document_tone"`
	}

	tonesList struct {
		Tones []toneScore `json:"tones"`
	}

	toneScore struct {
		Score    float64 `json:"score"`
		ToneID   string  `json:"tone_id"`
		ToneName string  `json:"tone_name"`
	}
)

func newCustomClient() *customClient {
	return &customClient{Client: &http.Client{}}
}

func (c customClient) NewRequest(method, url string, body io.Reader) (*http.Request, error) {
	return http.NewRequest(
		method,
		url,
		body,
	)
}

func (c customClient) Do(req *http.Request) (*http.Response, error) {
	return c.Client.Do(req)
}

func (a analyzer) Analyze(content string) (int64, error) {
	p := payload{Text: content}

	jsonStr, err := json.Marshal(p)
	if err != nil {
		return 0, fmt.Errorf("could not marshal struct: %v, error: %v", p, err)
	}

	req, err := a.client.NewRequest(
		"POST",
		fmt.Sprintf(url+"?version=%s&sentences=false", time.Now().UTC().Format("2006-01-02")),
		bytes.NewBuffer(jsonStr),
	)
	if err != nil {
		return 0, fmt.Errorf("could not create a request instance: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Basic "+basicAuth(username, password))

	resp, err := a.client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("could not make a request to a tone analyzer api: %v", err)
	}
	defer util.Check(resp.Body.Close)
	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("could not make a success request to a tone analyzer api, status code is %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))
	if err != nil {
		return 0, fmt.Errorf("could not read body from tone analyzer api response: %v", err)
	}

	var d documentAnalysis
	err = json.Unmarshal(body, &d)
	if err != nil {
		return 0, fmt.Errorf("could not get a proper response from tone analyzer api: %v", err)
	}

	score := calculateScore(d)
	return score, nil

}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func calculateScore(d documentAnalysis) int64 {
	var finalScore float64
	var numTones = len(d.DocumentTone.Tones)
	if numTones == 0 {
		return 50
	}

	for _, v := range d.DocumentTone.Tones {
		if v.Score == 1 {
			if isEmotion, ok := toneEmotion[v.ToneID]; ok && isEmotion {
				toneValue := v.Score * 100
				if name, ok := toneIDs[v.ToneID]; ok && name == "joy" {
					return int64(toneValue)
				}
				return int64(toneValue - 100)
			}
		}
		finalScore = finalScore + tones[v.ToneID]*v.Score
	}

	finalScore = (finalScore * 100) / float64(numTones)
	return int64(math.Abs(finalScore))
}
