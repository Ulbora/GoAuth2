package oauth2mysql

import (

)

const ClientInsertQuery = "INSERT INTO `client`(`secret`, `redirect_uri`, `name`, `web_site`, `email`, `enabled`) VALUES (?,?,?,?,?,?)"
const ClientReadRecordQuery = "SELECT `client_id`, `secret`, `redirect_uri`, `name`, `web_site`, `email`, `enabled`" +
							  "FROM `client` WHERE client_id = ?"