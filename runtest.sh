cd mysqldb
go test -coverprofile=coverage.out
sleep 15
cd ..
cd managers
go test -coverprofile=coverage.out
cd ..
cd oauthclient
go test -coverprofile=coverage.out