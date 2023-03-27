# config
. ./build_config.sh

# build
. ./build.sh

# run
BIN_NAME="`echo $APP_NAME`_$GO_MODE"
./out/$BIN_NAME

