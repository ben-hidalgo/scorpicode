
# DO NOT FORGET
# eval $(minikube docker-env)

TAG=$(shell find . -type f -exec md5 {} ';' | md5)

images:
	docker build . -f devops/dockerfiles/hats.Dockerfile     -t hats:$(TAG)
	docker build . -f devops/dockerfiles/site.Dockerfile     -t site:$(TAG)
	docker build . -f devops/dockerfiles/frontend.Dockerfile -t frontend:$(TAG)
	docker build . -f devops/dockerfiles/roxie.Dockerfile    -t roxie:$(TAG)

upgrade:
	helm upgrade --install scorpicode ./devops/helmchart \
	-f devops/helmchart/local.yaml \
	--set hats.tag=$(TAG) \
	--set site.tag=$(TAG) \
	--set roxie.tag=$(TAG) \
	--set frontend.tag=$(TAG)

go-happy:
	(cd backend            && \
	go test ./... -short   && \
	go vet ./...           && \
	go fmt ./...           && \
	go mod tidy)

mksr:
	minikube service roxie

protobufs:
	(cd backend             && \
	protoc                     \
	--proto_path=./rpc/hatspb  \
	--twirp_out=./rpc/hatspb/  \
	--go_out=./rpc/hatspb/     \
	rpc/hatspb/hats.proto)

minikube-start:
	minikube start --cpus 4 --memory 4096
	minikube addons enable ingress
