set -euxo pipefail

mkdir -p "$(pws)/functions"
GOBIN=$(pwd)/functions go install ./...
chmod +x "$(pwd)"/functions/*
go env
