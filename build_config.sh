GOPATH=`pwd`/src
GO_MODE="release"
APP_NAME="rara-ms-notification"
IMAGE_NAME=$APP_NAME:`cat VERSION.txt`

# implement go modules and remove these
#export GOFLAGS=""
#export GO111MODULE=on