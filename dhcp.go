package freebox

import ()

// DhcpConfig
type DhcpConfig struct {
	Enabled         bool     `json:"enabled"`
	Gateway         string   `json:"gateway"`
	StickyAssign    bool     `json:"sticky_assign"`
	IPRangeEnd      string   `json:"ip_range_end"`
	Netmask         string   `json:"netmask"`
	DNS             []string `json:"dns"`
	AlwaysBroadcast bool     `json:"always_broadcast"`
	IPRangeStart    string   `json:"ip_range_start"`
}

type DhcpHost struct {
	L2Ident struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"l2ident"`
	Active     bool `json:"active"`
	Persistent bool `json:"persistent"`
	Names      []struct {
		Name   string `json:"name"`
		Source string `json:"source"`
	} `json:"names"`
	VendorName        string `json:"vendor_name"`
	HostType          string `json:"host_type"`
	Interface         string `json:"interface"`
	ID                string `json:"id"`
	LastTimeReachable int    `json:"last_time_reachable"`
	PrimaryNameManual bool   `json:"primary_name_manual"`
	DefaultName       string `json:"default_name"`
	L3Connectivities  []struct {
		Addr              string `json:"addr"`
		Active            bool   `json:"active"`
		Reachable         bool   `json:"reachable"`
		LastActivity      int    `json:"last_activity"`
		Af                string `json:"af"`
		LastTimeReachable int    `json:"last_time_reachable"`
	} `json:"l3connectivities"`
	Reachable    bool `json:"reachable"`
	LastActivity int  `json:"last_activity"`
	AccessPoint  struct {
		Mac                 string `json:"mac"`
		Type                string `json:"type"`
		ConnectivityType    string `json:"connectivity_type"`
		UID                 string `json:"uid"`
		EthernetInformation struct {
			Duplex string `json:"duplex"`
			Speed  string `json:"speed"`
			Link   string `json:"link"`
		} `json:"ethernet_information"`
	} `json:"access_point"`
	PrimaryName string `json:"primary_name"`
}

type DhcpStaticLease struct {
	Mac      string   `json:"mac"`
	Comment  string   `json:"comment"`
	Hostname string   `json:"hostname"`
	ID       string   `json:"id"`
	Host     DhcpHost `json:"host"`
	IP       string   `json:"ip"`
}

type DhcpDynamicLease struct {
	Mac  string `json:"mac"`
	Host     DhcpHost `json:"host"`
	RefreshTime    int    `json:"refresh_time"`
	Hostname       string `json:"hostname"`
	AssignTime     int    `json:"assign_time"`
	LeaseRemaining int    `json:"lease_remaining"`
	IsStatic       bool   `json:"is_static"`
	IP             string `json:"ip"`
}

// Get the current DHCP configuration
func (c *Client) GetDhcpConfig() (*DhcpConfig, error) {
	payload := DhcpConfig{}
	err := c.GetResult("dhcp/config/", &payload)
	if err != nil {
		return nil, err
	}
	return &payload, nil
}

//Get the list of DHCP static leases
func (c *Client) GetDhcpStaticLease() ([]DhcpStaticLease, error) {
	payload := []DhcpStaticLease{}
	err := c.GetResult("dhcp/static_lease/", &payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

//Get the list of DHCP dynamic leases
func (c *Client) GetDhcpDynamicLease() ([]DhcpDynamicLease, error) {
	payload := []DhcpDynamicLease{}
	err := c.GetResult("dhcp/static_lease/", &payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}
