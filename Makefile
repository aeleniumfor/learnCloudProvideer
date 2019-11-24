Version=1.26
DockerBuild:
	docker build -t sasenomura/cloud-controller-manager:v$(Version) .
	docker push sasenomura/cloud-controller-manager:v$(Version)