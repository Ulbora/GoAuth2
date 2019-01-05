use ulbora_oauth2_server;
insert into client(email, enabled, name, secret, web_site) 
values('kw@ulboralabs.com', 1, 'Ulbora Labs', '554444vfg55ggfff22454sw2fff2dsfd', 'www.ulboralabs.com');


insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/addClient');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/updateClient');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/getClient');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/getClientList');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/deleteClient');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/addClientAllowedUri');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/getClientAllowedUriList');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/deleteClientAllowedUri');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/addClientRedirectUri');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/getClientRedirectUriList');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/deleteClientRedirectUri');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/deleteAllClientRedirectUri');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/addClientRole');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/getClientRoleList');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/deleteClientRole');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/addClientScope');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/getClientScopeList');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/deleteClientScope');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/addClientRoleUri');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/getClientRoleAllowedUriList');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/deleteClientRoleUri');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/addClientGrantType');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/getClientGrantTypeList');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/deleteClientGrantType');

insert into client_allowed_uri(client_id, uri) 
values((select client_id from client where name = 'Ulbora Labs'), 'http:localhost:3000/rs/deleteClientGrantType');


