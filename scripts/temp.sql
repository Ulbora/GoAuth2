SELECT * 
FROM client_role cr inner join 
uri_role ur on cr.id = ur.client_role_id
left join client_allowed_uri cau on cau.id = ur.client_allowed_uri_id
WHERE cr.client_id = 403
order by ur.client_role_id

SELECT cr.id as role_id, cr.role,  cau.id as uri_id, cau.uri, cr.client_id
FROM client_role cr inner join 
uri_role ur on cr.id = ur.client_role_id
left join client_allowed_uri cau on cau.id = ur.client_allowed_uri_id
WHERE cr.client_id = 403
order by ur.client_role_id



SELECT a.authorization_code, a.client_id, s.scope
FROM authorization_code a inner join auth_code_scope s
on a.authorization_code = s.authorization_code
WHERE a.client_id = ? and a.user_id = ? and s.scope = ?