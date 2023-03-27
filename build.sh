# config
. ./build_config.sh

if [ -z "$GO_MODE" ]; then
    export GO_MODE="release"
fi
BIN_NAME="`echo $APP_NAME`_$GO_MODE"

# Print workspace
echo "[BUILD] Building GO Project with Workspace: $GOPATH"

# clean out dir
rm -rfv ./out/$APP_NAME*

# build to out directory
FLAVOUR=$GO_MODE GOPATH=$GOPATH go build -v -ldflags "-X main.Flavour=$GO_MODE" -o ./out/$BIN_NAME ./src
chmod +x ./out/$BIN_NAME

# copy env file
cp -rfv ./src/conf/env.json ./out/env.json
