Version:=1.24
DockerBuild:
	docker build -t sasenomura/cloud-controller-manager:v$(Version) .
	docker push sasenomura/cloud-controller-manager:v$(Version)