package models

type Endpoint struct {
	Name       string                 `json:"name"`
	Type       string                 `json:"type"`
	StatusCode int                    `json:"status_code"`
	Response   map[string]interface{} `json:"response"`
}

type Service struct {
	Name      string     `json:"name"`
	Endpoints []Endpoint `json:"endpoints"`
}
