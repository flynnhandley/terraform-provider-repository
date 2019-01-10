package tpr

import (
	"net/http"
)

// Provider returns the binary
type Provider struct {
	Name     string `json:"name,omitempty"`
	Provider []byte
}

// GetProvider returns the provider
func GetProvider(w http.ResponseWriter, r *http.Request) {
	//
}
