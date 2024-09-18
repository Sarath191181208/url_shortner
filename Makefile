MAIN_PKG_PATH := ./cmd
BINARY_NAME := main

run/live:
	set -a && source ./.env && set +a
	air \
	--build.cmd "go build -o /tmp/bin/${BINARY_NAME} ${MAIN_PKG_PATH}" \
	--build.bin "/tmp/bin/${BINARY_NAME}" \
	--build.exclude_dir ".mysql-data" \
	--build.exclude_dir "db"

