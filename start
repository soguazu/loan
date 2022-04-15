#!/bin/sh
awhile=3
rm -rf docs && swag init --parseDependency --parseInternal --parseDepth 2 -g cmd/main.go
sleep $awhile && open http://localhost:8085/swagger/index.html &
make
