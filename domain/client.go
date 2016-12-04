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
