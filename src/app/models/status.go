package models

// Status struct to hold the status information.
type Status struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
	Version     string `json:"version"`
}
