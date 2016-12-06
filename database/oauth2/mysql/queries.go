package oauth2mysql

import (

)

const ClientInsertQuery = "INSERT INTO `client`(`secret`, `redirect_uri`, `name`, `web_site`, `email`, `enabled`) VALUES (?,?,?,?,?,?)"
