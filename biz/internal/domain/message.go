package domain

type IMMessageEntity struct {
	Event string `json:"event"`
	Type  string `json:"type"`
	Data  string `json:"data"`
}
