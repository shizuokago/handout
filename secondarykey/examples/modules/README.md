# on

export GO111MODULE=on
go run cmd/main.go

go.modが存在しないのでエラーになります。

go mod init mypkgs

で動作します。

# off

GO111MODULE=off
export GOPATH=`pwd`

go run cmd/main.go

src/mypkgs/で動作します。

# auto

GO111MODULE=auto

go run cmd/main.go

$GOPATH/go.mod exists but should not
//GOPATH直下にgo.modがあるので動作できません

export GOPATH=

go run cmd/main.go

on と同様の動作になります。

export GOPATH=`pwd`
rm go.mod

