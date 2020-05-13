# DO NOT FORGET
# eval $(minikube docker-env)

TAG=latest

images:
	docker build . -f devops/dockerfiles/hats.Dockerfile     -t hats:$(TAG)
	docker build . -f devops/dockerfiles/website.Dockerfile  -t website:$(TAG)
	docker build . -f devops/dockerfiles/frontend.Dockerfile -t frontend:$(TAG)
	docker build . -f devops/dockerfiles/roxie.Dockerfile    -t roxie:$(TAG)
	docker build . -f devops/dockerfiles/soxie.Dockerfile    -t soxie:$(TAG)
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
	--set frontend.socketHost=`minikube ip`:30081 \
	--set hats.tag=$(TAG) \
	--set website.tag=$(TAG) \
	--set roxie.tag=$(TAG) \
	--set soxie.tag=$(TAG) \
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

start-backend: # uses a script to trap killall
	./devops/scripts/start.sh

start-frontend:
	(cd frontend  && \
	REACT_APP_SOCKET_HOST=localhost:8084 \
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
	minikube start --cpus 4 --memory 4096 --vm-driver=virtualbox
	minikube addons enable ingress

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

logs-default:
	kubectl logs -n default -f -l app.kubernetes.io/instance=scorpicode --max-log-requests=20

logs-dev:
	kubectl logs -n dev -f -l app.kubernetes.io/instance=scorpicode
