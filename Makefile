install:
	go install golang.org/x/tools/cmd/goimports@latest
	go mod vendor

test:
	ENV=test && go test -v ./... -v -count=1

fix-format:
	gofmt -w -s app/ pkg/ cmd/ mocks/ testhelpers
	goimports -w app/ pkg/ cmd/ mocks/ testhelpers

start:
	go run ./cmd/main.go

build:
	GIN_MODE=release go build -o entrypoint app/main.go

dockerhub-build:
	docker build -t blitzshare.bootstrap.node:latest .
	docker tag blitzshare.bootstrap.node:latest iamkimchi/blitzshare.bootstrap.node:latest
	docker push iamkimchi/blitzshare.bootstrap.node:latest

build-docker-run:
	make build-docker
	docker run -t blitzshare.bootstrap.node

k8s-destory:
	kubectl delete namespace bootstrap-ns

k8s-pf:
	kubectl port-forward $(kubectl get pods  | tail -n1 | awk '{print $1}') 8000:80

k8s-apply:
	kubectl apply -f k8s/config/namespace.yaml 
	kubectl apply -f k8s/config/deployment.yaml
	kubectl apply -f k8s/config/service.yaml
	kubectl set image deployment/bootstrap-deployment bootstrap-containers=iamkimchi/blitzshare.bootstrap.node:local-latest -n bootstrap-ns
	kubectl wait -f k8s/config/deployment.yaml --for condition=available
	
# kubectl set image deployment/bootstrap-deployment bootstrap-containers=iamkimchi/blitzshare.bootstrap.node:local-latest
