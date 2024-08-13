.PHONY: all get-ip build-image run

all: build run

get-ip:
	$(eval export E2T_ADDRESS=$(shell sudo kubectl get svc -n ricplt --field-selector metadata.name=service-ricplt-e2term-sctp-alpha -o jsonpath='{.items[0].spec.clusterIP}'))

build-image: 
	# cd ../.. && make lib
	docker build -t ty3gx/e2sim:latest .
	docker push ty3gx/e2sim:latest

run: get-ip
	docker run -p 19960:19960 -e E2T_ADDRESS=$(E2T_ADDRESS) --name e2sim ty3gx/e2sim:latest


build: build-image

stop:
	docker stop e2sim -t 3
	docker rm e2sim

port:
	sudo netstat -tulnp | grep 19960

