package freebox

import (
	"fmt"
)

// Airmedia
type ContactEntry struct {
	LastName string `json:"last_name"`
	Company  string `json:"company"`
	PhotoURL string `json:"photo_url"`
	ID       int    `json:"id"`
	Birthday string `json:"birthday"`
	Numbers  []struct {
		Number    string `json:"number"`
		Type      string `json:"type"`
		ID        int    `json:"id"`
		ContactID int    `json:"contact_id"`
		IsDefault bool   `json:"is_default"`
		IsOwn     bool   `json:"is_own"`
	} `json:"numbers"`
	LastUpdate  Epoch  `json:"last_update"`
	DisplayName string `json:"display_name"`
	Notes       string `json:"notes"`
	FirstName   string `json:"first_name"`
}

// Get all contacts
func (c *Client) GetContacts() ([]ContactEntry, error) {
	payload := []ContactEntry{}
	err := c.GetResult("contact/", &payload)
	return payload, err
}

// Get a contact
func (c *Client) GetContact(id int) (CallEntry, error) {
	payload := CallEntry{}
	err := c.GetResult(fmt.Sprintf("contact/%v", id), &payload)
	return payload, err
}
