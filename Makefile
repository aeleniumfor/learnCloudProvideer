Version=1.32
DockerBuild:
	docker build -t sasenomura/cloud-controller-manager:v$(Version) .
	docker push sasenomura/cloud-controller-manager:v$(Version)