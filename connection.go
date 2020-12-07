package freebox

import (
)

type ConnectionStatus struct {
	Type          string `json:"type"`
	RateDown      int    `json:"rate_down"`
	BytesUp       int64  `json:"bytes_up"`
	Ipv4PortRange []int  `json:"ipv4_port_range"`
	RateUp        int    `json:"rate_up"`
	BandwidthUp   int    `json:"bandwidth_up"`
	Ipv6          string `json:"ipv6"`
	BandwidthDown int    `json:"bandwidth_down"`
	Media         string `json:"media"`
	State         string `json:"state"`
	BytesDown     int64  `json:"bytes_down"`
	Ipv4          string `json:"ipv4"`
}

type ConnectionStatusFull struct {
	Type          string `json:"type"`
	RateDown      int    `json:"rate_down"`
	BytesUp       int64  `json:"bytes_up"`
	Ipv4PortRange []int  `json:"ipv4_port_range"`
	RateUp        int    `json:"rate_up"`
	BandwidthUp   int    `json:"bandwidth_up"`
	Ipv6          string `json:"ipv6"`
	BandwidthDown int    `json:"bandwidth_down"`
	Media         string `json:"media"`
	State         string `json:"state"`
	BytesDown     int64  `json:"bytes_down"`
	Ipv4          string `json:"ipv4"`
}

