set CURDIR=`pwd`
set OLDGOPATH=%$GOPATH%
set GOPATH=%cd%

gofmt -w src/server

go install server

set GOPATH=%OLDGOPATH%

echo 'finished'