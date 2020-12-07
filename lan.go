package freebox

import (
	"fmt"
)

// browsable LAN interfaces
type LanBrowserInterfaces struct {
	Name      string `json:"name"`
	HostCount int    `json:"host_count"`
}

// list of hosts on a given interface
type LanBrowserInterface struct {
	L2Ident struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"l2ident"`
	Active            bool   `json:"active"`
	ID                string `json:"id"`
	LastTimeReachable int    `json:"last_time_reachable"`
	Persistent        bool   `json:"persistent"`
	Names             []struct {
		Name   string `json:"name"`
		Source string `json:"source"`
	} `json:"names"`
	VendorName       string `json:"vendor_name"`
	L3Connectivities []struct {
		Addr              string `json:"addr"`
		Active            bool   `json:"active"`
		Af                string `json:"af"`
		Reachable         bool   `json:"reachable"`
		LastActivity      int    `json:"last_activity"`
		LastTimeReachable int    `json:"last_time_reachable"`
	} `json:"l3connectivities"`
	Reachable         bool   `json:"reachable"`
	LastActivity      int    `json:"last_activity"`
	PrimaryNameManual bool   `json:"primary_name_manual"`
	PrimaryName       string `json:"primary_name"`
}

// Lan configuration
type LanConfig struct {
	NameDNS     string `json:"name_dns"`
	NameMdns    string `json:"name_mdns"`
	Name        string `json:"name"`
	Mode        string `json:"mode"`
	NameNetbios string `json:"name_netbios"`
	IP          string `json:"ip"`
}

// Get LanConfig
func (c *Client) GetLanConfig() (*LanConfig, error) {
	payload := LanConfig{}
	err := c.GetResult("lan/config/", &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

func (c *Client) GetLanBrowserInterfaces() ([]LanBrowserInterfaces, error) {
	payload := []LanBrowserInterfaces{}
	err := c.GetResult("lan/browser/interfaces/", &payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func (c *Client) GetLanBrowserInterface(nic string) ([]LanBrowserInterface, error) {
	payload := []LanBrowserInterface{}
	err := c.GetResult(fmt.Sprintf("lan/browser/%s/", nic), &payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
