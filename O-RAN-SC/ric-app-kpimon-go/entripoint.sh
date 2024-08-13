#!/bin/bash
./f1apServer &  # Start the API server in the background
# ./kpimon   # Run the client application
go-fuzz-run --file=/go/src/gerrit.o-ran-sc.org/r/scp/ric-app/kpimon/gofuzzdep-fuzz.zip \
    --trackdir=gopath/src/example.com --wkdir=/go/src/gerrit.o-ran-sc.org/r/scp/ric-app/kpimon --coverfreq=10
