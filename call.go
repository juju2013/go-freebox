package freebox

import (
  "fmt"
  "time"
  "strconv"
)
// Airmedia
type CallEntry struct {
		Number    string `json:"number"`
		Type      string `json:"type"`
		ID        int    `json:"id"`
		Duration  int    `json:"duration"`
		Datetime  Epoch  `json:"datetime"`
		ContactID int    `json:"contact_id"`
		LineID    int    `json:"line_id"`
		Name      string `json:"name"`
		New       bool   `json:"new"`
	}

// Special handler from epoch to time
type Epoch time.Time

func (e *Epoch) UnmarshalJSON(data []byte) error {
  millis, err := strconv.ParseInt(string(data), 10, 64)
  if err != nil {
      return err
  }
  *e = Epoch(time.Unix(0, millis * int64(time.Millisecond)))
  return nil
}

func (e Epoch) String() string {
  return time.Time(e).Format("2006-01-02 15:04:05")
}

// Get call log
func (c *Client) GetCallEntries() ([]CallEntry, error) {
  payload := []CallEntry{}
  err := c.GetResult("call/log/", &payload)
  if err != nil {
    return nil, err
  }
  return payload, nil
}

// Mark all as read
func (c *Client) MarkAllRead() error {
  _, err := c.PostResource("call/log/mark_all_as_read/", "", true)
  return err
}

// Mark as read
func (c *Client) MarkRead(id int) error {
  payload, _ := c.GetCallEntrie(id)
  payload.New=false
  _, err := c.PutResource(fmt.Sprintf("call/log/%v", id), &payload, true)
  return err
}

// Get a call
func (c *Client) GetCallEntrie(id int) (CallEntry, error) {
  payload := CallEntry{}
  err := c.GetResult(fmt.Sprintf("call/log/%v", id), &payload)
  return payload, err
}
