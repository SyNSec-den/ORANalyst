build:
	go build

rtmgr: build
	./constraint_collector --path-constraint=true --module=rtmgr dir/ric-plt-rtmgr/pkg/ | tee logs/rtmgr.log 2>&1

onos-e2t: build
	./constraint_collector --path-constraint=true --module=onos-e2t-indication dir/onos-e2t | tee logs/onos-e2t.log 2>&1

onos-e2t-e2setup: build
	./constraint_collector --path-constraint=true --module=onos-e2t-e2setup dir/onos-e2t | tee logs/onos-e2t-e2setup.log 2>&1

rimedo: build
	./constraint_collector --path-constraint=true --module=rimedo dir/rimedo-ts/pkg | tee logs/rimedo.log 2>&1
