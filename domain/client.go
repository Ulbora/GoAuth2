package domain

import (

)

type Client struct{
	clientId uint64
	secret string
	redirectUrl string
	name string
	webSite string
	email string
	enabled bool
}

func (c Client) Name() string{
	return c.name
}

func (c *Client) SetName(name string){
	c.name = name
}
