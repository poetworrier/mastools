package api

import "github.com/imroc/req/v3"

// Client wraps methods on https://docs.joinmastodon.org/methods/admin/trends/
// TODO: needs to be init'd with OAuth creds
type Trends struct {
	c *req.Client
}

func NewTrends(c *req.Client) *Trends {
	return &Trends{
		c: c,
	}
}

// https://docs.joinmastodon.org/methods/admin/trends/#statuses
func (t *Trends) ListStatus() ([]Status, error) {
	// err is resp.Err
	resp, err := t.c.R().Get("/api/v1/admin/trends/statuses")
	if err != nil {
		// TODO: do something with err
		// if available read code/headers/etc and pass back instructions to caller?
		return nil, err
	}
	var statuses []Status
	err = resp.UnmarshalJson(&statuses)
	if err != nil {
		// TODO: wrap formatting error with details about the request
		return nil, err
	}
	return statuses, nil
}

// https://docs.joinmastodon.org/methods/admin/trends/#statuses
func (t *Trends) ListTags() ([]Tag, error) {
	// err is resp.Err
	resp, err := t.c.R().Get("/api/v1/admin/trends/tags")
	if err != nil {
		// TODO: do something with err
		// if available read code/headers/etc and pass back instructions to caller?
		return nil, err
	}
	var tags []Tag
	err = resp.UnmarshalJson(&tags)
	if err != nil {
		// TODO: wrap formatting error with details about the request
		return nil, err
	}
	return tags, nil
}
