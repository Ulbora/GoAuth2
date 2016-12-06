package domain

import (

)

type Client struct{
	ClientId uint64
	Secret string
	RedirectUrl string
	Name string
	WebSite string
	Email string
	Enabled bool
}
/*
func (c Client) ClientId() uint64{
	return c.clientId
}

func (c *Client) SetClientId(clientId uint64){
	c.clientId = clientId
}

func (c Client) Secret() string{
	return c.secret
}

func (c *Client) SetSecret(secret string){
	c.secret = secret
}

func (c Client) RedirectUrl() string{
	return c.redirectUrl
}

func (c *Client) SetRedirectUrl(redirectUrl string){
	c.redirectUrl = redirectUrl
}

func (c Client) Name() string{
	return c.name
}

func (c *Client) SetName(name string){
	c.name = name
}

func (c Client) WebSite() string{
	return c.webSite
}

func (c *Client) SetWebSite(webSite string){
	c.webSite = webSite
}

func (c Client) Email() string{
	return c.email
}

func (c *Client) SetEmail(email string){
	c.email = email
}

func (c Client) Enabled() bool{
	return c.enabled
}

func (c *Client) SetEnabled(enabled bool){
	c.enabled = enabled
}
*/
