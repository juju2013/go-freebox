package freebox

// Airmedia
type AirMediaReceiver struct {
		Capabilities struct {
			Photo  bool `json:"photo"`
			Screen bool `json:"screen"`
			Audio  bool `json:"audio"`
			Video  bool `json:"video"`
		} `json:"capabilities"`
		Name              string `json:"name"`
		PasswordProtected bool   `json:"password_protected"`
	}

func (c *Client) GetAirMediaReceivers() ([]AirMediaReceiver, error) {
  payload := []AirMediaReceiver{}
  err := c.GetResult("airmedia/receivers", &payload)
  if err != nil {
    return nil, err
  }
  return payload, nil
}
