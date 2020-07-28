# DO NOT FORGET
# eval $(minikube docker-env)

docker-build:
	./devops/scripts/docker-build.sh

#TODO: add PIPELINE_ID as first parameter with value `whoami`
docker-push:
	./devops/scripts/docker-push.sh

helm-upgrade-master:
	./devops/scripts/helm-upgrade.sh master

tag-build-push:
	./devops/scripts/write-tags.sh
	./devops/scripts/docker-build.sh
	./devops/scripts/docker-push.sh

#TODO: add wait for services... mongo startup is longer than the default wait timeout
# there is a bug in helm upgrade which intermittently doesn't not accept stdin from sops using: -f -
upgrade-local: #images
	sops -d ./devops/helmchart/local.sops.yaml > ./devops/helmchart/local.plain.yaml
	helm upgrade --install scorpicode ./devops/helmchart \
	-f devops/helmchart/local.yaml \
	-f devops/helmchart/local.plain.yaml \
	--set common.auth0RedirectUri=http://`minikube ip`:30080/callback \
	--set roxie.loginSuccessTarget=http://`minikube ip`:30080/sc \
	--set frontend.socketHost=`minikube ip`:30081 \
	--set hats.tag=$(shell ./devops/scripts/go-checksum.sh hats) \
	--set roxie.tag=$(shell ./devops/scripts/go-checksum.sh roxie) \
	--set soxie.tag=$(shell ./devops/scripts/go-checksum.sh soxie) \
	--set website.tag=$(shell ./devops/scripts/js-checksum.sh website) \
	--set frontend.tag=$(shell ./devops/scripts/js-checksum.sh frontend)

upgrade-dev:
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

test:
	./devops/scripts/integration-tests.sh dev

start-backend: # uses a script to trap killall
	./devops/scripts/start.sh

start-frontend:
	(cd frontend  && \
	REACT_APP_SOCKET_HOST=localhost:8084 \
	REACT_APP_SOCKET_DEBUG=true \
	npm start)

login:
	open http://localhost:8080/login

protobufs:
	(cd backend && protoc --proto_path=./rpc/hatspb --twirp_out=./rpc/hatspb/ --go_out=./rpc/hatspb/ rpc/hatspb/hats.proto)
	(cd backend && protoc --proto_path=./rpc/rmqpb  --twirp_out=./rpc/rmqpb/  --go_out=./rpc/rmqpb/  rpc/rmqpb/rmq.proto)

minikube-start:
	minikube delete
	# using virtualbox so that ip will be one of: 192.168.99.100, 192.168.99.101, 192.168.99.102
	# which are registered as callbacks in Auth0
	minikube start --cpus 4 --memory 4096 --vm-driver=virtualbox --insecure-registry="$(shell ipconfig getifaddr en0):5000"
	minikube addons enable ingress
	minikube addons enable registry

# opens the load balancer at http://<minikube ip>:<roxie port>
minikube-service-roxie-default:
	minikube service roxie -n default

# opens the load balancer at http://<minikube ip>:<soxie port>
minikube-service-soxie-default:
	minikube service soxie -n default

# to get the rabbit management console (opens three tabs, just close two)
minikube-service-rabbit-dev:
	minikube service scorpicode-rabbitmq-ha -n dev

# to get the rabbit management console (opens three tabs, just close two)
minikube-service-rabbit-default:
	minikube service scorpicode-rabbitmq-ha -n default


# the .tgz files are committed
helm-dependency-update:
	(cd devops/helmchart && helm dependency update)

logs-master:
	kubectl logs -n master -f -l app.kubernetes.io/instance=scorpicode --max-log-requests=20

logs-dev:
	kubectl logs -n dev -f -l app.kubernetes.io/instance=scorpicode
