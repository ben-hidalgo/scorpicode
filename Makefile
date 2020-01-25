# DO NOT FORGET
# eval $(minikube docker-env)

TAG=latest

images:
	docker build . -f devops/dockerfiles/hats.Dockerfile     -t hats:$(TAG)
	docker build . -f devops/dockerfiles/website.Dockerfile  -t website:$(TAG)
	docker build . -f devops/dockerfiles/frontend.Dockerfile -t frontend:$(TAG)
	docker build . -f devops/dockerfiles/roxie.Dockerfile    -t roxie:$(TAG)

upgrade: #images
	helm upgrade --install scorpicode ./devops/helmchart \
	-f devops/helmchart/local.yaml \
	--set common.cacheBuster=`date +%s` \
	--set hats.tag=$(TAG) \
	--set website.tag=$(TAG) \
	--set roxie.tag=$(TAG) \
	--set frontend.tag=$(TAG)

dry-run:
	helm upgrade --install --debug --dry-run scorpicode ./devops/helmchart \
	-f devops/helmchart/local.yaml \
	--set common.cacheBuster=`date +%s` \
	--set hats.tag=$(TAG) \
	--set website.tag=$(TAG) \
	--set roxie.tag=$(TAG) \
	--set frontend.tag=$(TAG)

go-happy:
	(cd backend            && \
	go test ./... -short   && \
	go vet ./...           && \
	go fmt ./...           && \
	go mod tidy)

start:
	./devops/scripts/start.sh

mksr:
	minikube service roxie

protobufs:
	(cd backend &&             \
	protoc                     \
	--proto_path=./rpc/hatspb  \
	--twirp_out=./rpc/hatspb/  \
	--go_out=./rpc/hatspb/     \
	rpc/hatspb/hats.proto)

minikube-start:
	minikube delete
	minikube start --cpus 4 --memory 4096
	minikube addons enable ingress
