# DO NOT FORGET
# eval $(minikube docker-env)

TAG=latest

images:
	docker build . -f devops/dockerfiles/hats.Dockerfile     -t hats:$(TAG)
	docker build . -f devops/dockerfiles/website.Dockerfile  -t website:$(TAG)
	docker build . -f devops/dockerfiles/frontend.Dockerfile -t frontend:$(TAG)
	docker build . -f devops/dockerfiles/roxie.Dockerfile    -t roxie:$(TAG)
	docker build . -f devops/dockerfiles/debugger.Dockerfile -t debugger:latest

#TODO: add wait for services... mongo startup is longer than the default wait timeout
# there is a bug in helm upgrade which intermittently doesn't not accept stdin from sops using: -f -
upgrade: #images
	sops -d ./devops/helmchart/local.sops.yaml > ./devops/helmchart/local.plain.yaml
	helm upgrade --install scorpicode ./devops/helmchart \
	-f devops/helmchart/local.yaml \
	-f devops/helmchart/local.plain.yaml \
	--set common.cacheBuster=`date +%s` \
	--set common.auth0RedirectUri=http://`minikube ip`:30080/callback \
	--set roxie.loginSuccessTarget=http://`minikube ip`:30080/sc \
	--set hats.tag=$(TAG) \
	--set website.tag=$(TAG) \
	--set roxie.tag=$(TAG) \
	--set frontend.tag=$(TAG)

dev:
	kubectl create namespace dev || true
	helm upgrade --install scorpicode ./devops/helmchart \
	-f devops/helmchart/dev.yaml \
	-n dev

go-happy:
	(cd backend && \
	go vet ./... && \
	golint -set_exit_status ./... && \
	go test ./... -short -v && \
	go fmt ./... && \
	go mod tidy)

TEST_ARGS=\
REDIS_ADDRESS=`minikube ip`:`kubectl get svc scorpicode-redis-master -n dev -o json | jq '.spec.ports[0].nodePort'` \
REDIS_PASSWORD=redispassword

test:
	(cd backend && ${TEST_ARGS} go test ./... -v -count=1)

start-backend: #go-happy
	./devops/scripts/start.sh

start-frontend:
	(cd frontend  && \
	npm install   && \
	npm start)

login:
	open http://localhost:8080/login

protobufs:
	(cd backend &&             \
	protoc                     \
	--proto_path=./rpc/hatspb  \
	--twirp_out=./rpc/hatspb/  \
	--go_out=./rpc/hatspb/     \
	rpc/hatspb/hats.proto)

minikube-start:
	minikube delete
	# using virtualbox so that ip will be one of: 192.168.99.100, 192.168.99.101, 192.168.99.102
	# which are registered as callbacks in Auth0
	minikube start --cpus 4 --memory 4096 --vm-driver=virtualbox
	minikube addons enable ingress

# opens the load balancer at http://<minikube ip>:<roxie port>
minikube-service-roxie:
	minikube service roxie

# to get the rabbit management console (opens three tabs, just close two)
minikube-service-rabbit-dev:
	minikube service scorpicode-rabbitmq-ha -n dev

# the .tgz files are committed
helm-dependency-update:
	(cd devops/helmchart && helm dependency update)

logs-local:
	kubectl logs -f -l app.kubernetes.io/instance=scorpicode

logs-dev:
	kubectl logs -n dev -f -l app.kubernetes.io/instance=scorpicode
