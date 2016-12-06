package domain

import ()

type Client struct {
	ClientId    uint64
	Secret      string
	RedirectUrl string
	Name        string
	WebSite     string
	Email       string
	Enabled     bool
}
