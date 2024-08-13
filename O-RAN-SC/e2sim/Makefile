lib:
	docker build -t e2sim_lib:latest .
	docker run -d --name e2sim_lib e2sim_lib:latest
	docker cp e2sim_lib:/e2sim/build/e2sim_1.0.0_amd64.deb ./docker_lib/
	docker cp e2sim_lib:/e2sim/build/e2sim-dev_1.0.0_amd64.deb ./docker_lib/
	docker stop e2sim_lib
	docker rm e2sim_lib
	
	cp ./docker_lib/*.deb ./e2sm_examples/kpm_e2sm/

	# rm -rf build && mkdir build
	# cd build && cmake .. && make package && cmake .. -DDEV_PKG=1 && make package 
	# cp ./build/*.deb ./e2sm_examples/kpm_e2sm/