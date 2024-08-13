all: build push onboard install

helm:
	docker run --rm -u 0 -it -d -p 8080:8080 -e DEBUG=1 -e STORAGE=local -e STORAGE_LOCAL_ROOTDIR=/charts -v $(pwd)/charts:/charts chartmuseum/chartmuseum:latest

onboard: 
	dms_cli onboard deploy/config.json deploy/schema.json

uninstall:
	dms_cli uninstall kpimon-go ricxapp

install:
	dms_cli install kpimon-go 2.0.1 ricxapp

build: 
	sudo docker build -t localhost:5001/kpimon-go:latest .
	sudo docker run -d --name kpimon-go localhost:5001/kpimon-go:latest
	rm -rf dockerout && mkdir dockerout
	sudo docker cp kpimon-go:/go/src/gerrit.o-ran-sc.org/r/scp/ric-app/kpimon/gofuzzdep-fuzz/metadata ./dockerout/metadata
	sudo docker stop kpimon-go
	sudo docker rm kpimon-go
	docker push localhost:5001/kpimon-go:latest

build-original:
	sudo docker build -f Dockerfile_original  -t localhost:5001/kpimon-go:latest .
	docker push localhost:5001/kpimon-go:latest

push:
	docker push localhost:5001/kpimon-go:latest

test:
	sudo docker run -it --rm --name kpimon-go localhost:5001/kpimon-go:latest