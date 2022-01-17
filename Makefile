install:
	go install golang.org/x/tools/cmd/goimports@latest
	go mod vendor

test:
	ENV=test && go test -v ./... -v -count=1

fix-format:
	gofmt -w -s app/ cmd/
	goimports -w app/ cmd/

start:
	go run ./cmd/main.go

build:
	go build -o entrypoint app/main.go

k8s-apply:
	kubectl apply -f k8s/namespace.yaml 
	kubectl apply -f k8s/deployment.yaml
	kubectl apply -f k8s/service.yaml
	kubectl rollout restart deployment/blitzshare-bootstrap-deployment --namespace blitzshare-ns
	kubectl wait -f k8s/deployment.yaml --for condition=available
