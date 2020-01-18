# GoAuth2
A complete standalone Oauth2 Server RFC 6749 implementation written in Golang and licensed under the GPL V3 license.

### [Documentation](https://github.com/Ulbora/GoAuth2/wiki)

[![Build Status](https://travis-ci.org/Ulbora/GoAuth2.svg?branch=master)](https://travis-ci.org/Ulbora/GoAuth2)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=Ulbora_GoAuth2&metric=alert_status)](https://sonarcloud.io/dashboard?id=Ulbora_GoAuth2)
[![Go Report Card](https://goreportcard.com/badge/github.com/Ulbora/GoAuth2)](https://goreportcard.com/report/github.com/Ulbora/GoAuth2)




GoAuth2 is an Oauth2 server implementation written in Golang. Currently authorization code, 
implicit, and client credentials grant, and password grant types are supported.

GoAuth2 issues a compressed enhanced JWT token that can be used to secure individual REST endpoints for users using roles. REST endpoints can be coded to validate the JWT token using the user's role. There is also a access token REST service that can validate a compressed token. Token compression can be turned off at startup if desired.

### GoAuth2 will provide the security infrastructure for the new Ulbora Labs eCommerce Platform project, 66GoCart.
#### (66GoCart is an eCommerce Platform server solution written in golang that provides REST endpoints for backend operations. 66GoCart frontend soltutions can be written in golang or any other language.)

## Authorization Code Grant Type

 * Authorize

```
   Example
   GET:

http://localhost:3000/oauth/authorize?response_type=code&client_id=403&redirect_uri=CALLBACK_URL&scope=read&state=xyz
  
```

```  
   Test on localhost
   GET:

http://localhost:3000/oauth/authorize?response_type=code&client_id=403&redirect_uri=http://www.google.com&scope=read&state=xyz

```

 * Access Token 

```
   Example
   POST:

http://localhost:3000/oauth/token?client_id=403&client_secret=554444vfg55ggfff22454sw2fff2dsfd&grant_type=authorization_code&code=i76y13e340akRn6Ipkdbii&redirect_uri=http://www.google.com
 
```

```  
   Test on localhost
   POST:

http://localhost:3000/oauth/token?client_id=403&client_secret=554444vfg55ggfff22454sw2fff2dsfd&grant_type=authorization_code&code=i76y13e340akRn6Ipkdbii&redirect_uri=http://www.google.com

```

 * Refresh Token

```
   Example
   POST:

http://localhost:3000/oauth/token?grant_type=refresh_token&client_id=CLIENT_ID&client_secret=CLIENT_SECRET&refresh_token=REFRESH_TOKEN
   
```

``` 
   Test on localhost
   POST:

http://localhost:3000/oauth/token?grant_type=refresh_token&client_id=403&client_secret=554444vfg55ggfff22454sw2fff2dsfd&refresh_token=efssffffnnlf

```

## Implicit Grant Type

* Authorize

```
   Example
   GET:

http://localhost:3000/oauth/authorize?response_type=token&client_id=403&redirect_uri=CALLBACK_URL&scope=read&state=xyz
  
```

```  
   Test on localhost
   GET:

http://localhost:3000/oauth/authorize?response_type=token&client_id=403&redirect_uri=http://www.google.com&scope=read&state=xyz

```


## Client Credentials Grant Type

 * Access Token    

```
   Example
   POST:

http://localhost:3000/oauth/token?client_id=403&client_secret=554444vfg55ggfff22454sw2fff2dsfd&grant_type=client_credentials
 
```

```  
   Test on localhost
   POST:

http://localhost:3000/oauth/token?client_id=403&client_secret=554444vfg55ggfff22454sw2fff2dsfd&grant_type=client_credentials

```

## Password Grant Type

* Access Token 

```
   Example
   POST:
   
http://localhost:3000/oauth/token

grant_type=password&client_id=403&username=someUser&password=somePw

```




## Access Token Validation


```
   Example
   POST:

http://localhost:3000/rs/token/validate
 
```

```  
   Request:

{
   "userId":null,
   "clientId": 403,
   "role":"admin",
   "url":"http:localhost:3000/rs/updateClient",
   "scope":null,
   "accessToken":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhY2Nlc3MiLCJncmFudCI6ImNsaWVudF9jcmVkZW50aWFscyIsImNsaWVudElkIjo0MDMsInJvbGVVcmlzIjpbeyJjbGllbnRSb2xlSWQiOjEsInJvbGUiOiJhZG1pbiIsInVyaUlkIjo2MywidXJpIjoiaHR0cDpsb2NhbGhvc3Q6MzAwMC9ycy91cGRhdGVDbGllbnQiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6MSwicm9sZSI6ImFkbWluIiwidXJpSWQiOjc3LCJ1cmkiOiJodHRwOmxvY2FsaG9zdDozMDAwL3JzL2FkZENsaWVudFNjb3BlIiwiY2xpZW50SWQiOjQwM30seyJjbGllbnRSb2xlSWQiOjIsInJvbGUiOiJ1c2VyIiwidXJpSWQiOjY4LCJ1cmkiOiJodHRwOmxvY2FsaG9zdDozMDAwL3JzL2RlbGV0ZUNsaWVudEFsbG93ZWRVcmkiLCJjbGllbnRJZCI6NDAzfSx7ImNsaWVudFJvbGVJZCI6Miwicm9sZSI6InVzZXIiLCJ1cmlJZCI6ODAsInVyaSI6Imh0dHA6bG9jYWxob3N0OjMwMDAvcnMvYWRkQ2xpZW50Um9sZVVyaSIsImNsaWVudElkIjo0MDN9XSwiZXhwaXJlc0luIjozNjAwMCwiaWF0IjoxNDg3NTUwNTcxLCJ0b2tlblR5cGUiOiJhY2Nlc3MiLCJleHAiOjE0ODc1ODY1NzEsImlzcyI6IlVsYm9yYSBPYXV0aDIgU2VydmVyIn0.1Isnysob52ujgYOu9Oi"
}



```


```  
   Response:

{
  "valid": true
}


```

# Client Micro Service


 Oauth2 Client Micro Service


## Headers
Content-Type: application/json (for POST and PUT)
Authorization: Bearer atToken
clientId: clientId (example 33477)


## Add Client

```
POST:
URL: http://localhost:3000/rs/client/add

Example Request
{
   "name":"ulbora",
   "webSite":"www.ulboralabs.com",
   "email":"ulbora@ulbora.com",
   "enabled":true,
   "redirectUrls":[
      {
         "uri":"http://www.google.com",
         "clientId":null
      },
      {
         "uri":"http://www.ulboralabs.com",
         "clientId":null
      }
   ]
}
  
```

```
Example Response   

{
  "success": true, 
  "message": ""
}

```



## Update Client

```
PUT:
URL: http://localhost:3000/rs/client/update

Example Request
{
  "clientId": 510,
  "name": "ulbora",
  "webSite": "www.ulboralabs.com",
  "email": "ulbora@ulbora.com",
  "enabled": false
}
  
```

```
Example Response   

{
  "success": true, 
  "message": ""
}

```



## Get Client

```
GET:
URL: http://localhost:3000/rs/client/get/510
  
```

```
Example Response   

{
    "id": 94,    
    "clientId": 510,
    "name": "ulbora",
    "webSite": "www.ulboralabs.com",
    "email": "ulbora@ulbora.com",
    "enabled": false

}

```



## Get Client List

```
GET:
URL: http://localhost:3000/rs/client/list

  
```

```
Example Response   

[
    {
    "id": 94,    
    "clientId": 510,
    "name": "ulbora",
    "webSite": "www.ulboralabs.com",
    "email": "ulbora@ulbora.com",
    "enabled": false

    }
]

```



## Delete Client

```
DELETE:
URL: http://localhost:3000/rs/client/delete/509
  
```

```
Example Response   

{
  "success": true,
  "message": ""
}

```



## Add Client Grant Type

```
POST:
URL: http://localhost:3000/rs/clientGrantType/add

Example Request
{
   "grantType":"code",
   "clientId":581
}
  
```

```
Example Response   

{
  "success": true, 
  "message": ""
}

```



## Get Client Grant Type

```
GET:
URL: http://localhost:3000/rs/clientGrantType/list/581
  
```

```
Example Response   

{
   "grantType":"code",
   "clientId":581
}

```



## Delete Client Grant Type

```
DELETE:
URL: http://localhost:3000/rs/clientGrantType/delete/221
  
```

```
Example Response   

{
  "success": true,
  "message": ""
}

```




## Add Client Allowed URI

```
POST:
URL: http://localhost:3000/rs/clientAllowedUri/add

Example Request
{
   "uri":"www.ulboralabs.com",
   "clientId":616
   
}
  
```

```
Example Response   

{
  "success": true, 
  "message": ""
}

```



## Update Client Allowed URI

```
PUT:
URL: http://localhost:3000/rs/clientAllowedUri/update

Example Request
{
   "uri":"www.ulbora.com",
   "id":139
   
}
  
```

```
Example Response   

{
  "success": true, 
  "message": ""
}

```



## Get Client Allowed URI

```
GET:
URL: http://localhost:3000/rs/clientAllowedUri/get/139
  
```

```
Example Response   

{
   "uri":"www.ulbora.com",
   "id":139
   
}

```



## Get Client Allowed URI List

```
GET:
URL: http://localhost:3000/rs/clientAllowedUri/list/616

  
```

```
Example Response   

[
    {
        "uri":"www.ulbora.com",
        "id":139   
    }
]

```



## Delete Client Allowed URI

```
DELETE:
URL: http://localhost:3000/rs/clientAllowedUri/delete/139
  
```

```
Example Response   

{
  "success": true,
  "message": ""
}

```




## Add Client Redirect URI

```
POST:
URL: http://localhost:3000/rs/clientRedirectUri/add

Example Request
{
   "uri":"www.ulboralabs.com",
   "clientId":616
   
}
  
```

```
Example Response   

{
  "success": true, 
  "message": ""
}

```




## Get Client Redirect URI List

```
GET:
URL: http://localhost:3000/rs/clientRedirectUri/list/616

  
```

```
Example Response   

[
    {
        "uri":"www.ulbora.com",
        "id":139   
    }
]

```



## Delete Client Redirect URI

```
DELETE:

URL: http://localhost:3000/rs/clientRedirectUri/delete/681
  
```

```
Example Response   

{
  "success": true,
  "message": ""
}

```





## Add Client Role

```
POST:
URL: http://localhost:3000/rs/clientRole/add

Example Request
{
   "role":"tester2",
   "clientId":616
   
}
  
```

```
Example Response   

{
  "success": true, 
  "message": ""
}

```




## Get Client Role List

```
GET:
URL: http://localhost:3000/rs/clientRole/list/616

  
```

```
Example Response   

[
    {
        "role":"tester2",
        "clientId":616   
    }
]

```



## Delete Client Role

```
DELETE:

URL: http://localhost:3000/rs/clientRole/delete/25
  
```

```
Example Response   

{
  "success": true,
  "message": ""
}

```




## Add Client Role URI

```
POST:
URL: http://localhost:3000/rs/clientRoleUri/add

Example Request
{  
   "clientRoleId":24,
   "clientAllowedUriId":167
}
  
```

```
Example Response   

{
  "success": true, 
  "message": ""
}

```




## Get Client Role URI List

```
GET:
URL: http://localhost:3000/rs/clientRoleUri/list/24

  
```

```
Example Response   

[
    {  
        "clientRoleId":24,
        "clientAllowedUriId":167
    }
]

```



## Delete Client Role URI

```
DELETE:

URL: http://localhost:3000/rs/clientRoleUri/delete/{clientRoleId}/{clientAllowedUriId}

URL: http://localhost:3000/rs/clientRoleUri/delete/24/167
  
```

```
Example Response   

{
  "success": true,
  "message": ""
}

```


## Delete Client Role URI with POST

```
POST:

URL: http://localhost:3000/rs/clientRoleUri/delete

Example Request

{  
   "clientRoleId":24,
   "clientAllowedUriId":167
}
  
```

```
Example Response   

{
  "success": true,
  "message": ""
}

```




This server should run behind nginx and nginx should handle certs.

If you would **like to contribute** to this project, **send a pull request**.


Contributors:
Ken Williamson

