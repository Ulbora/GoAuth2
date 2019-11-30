cd mysqldb
go test -coverprofile=coverage.out
sleep 15
cd ..
cd managers
go test -coverprofile=coverage.out
sleep 15
cd ..
cd oauthclient
go test -coverprofile=coverage.out
sleep 15
cd ..
cd rolecontrol
go test -coverprofile=coverage.out
sleep 15
cd ..
cd handlers
go test -coverprofile=coverage.out
sleep 15