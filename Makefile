install:
	go install golang.org/x/tools/cmd/goimports@latest
	go mod vendor

test:
	ENV=test && go test -v ./... -v -count=1

fix-format:
	gofmt -w -s app/ cmd
	goimports -w app/ cmd

start:
	go run ./cmd/main.go

build:
	go build -o entrypoint app/main.go

dockerhub-build:
	docker buildx build --platform linux/amd64 -t blitzshare.bootstrap.node:latest .
	docker tag blitzshare.bootstrap.node:latest iamkimchi/blitzshare.bootstrap.node:latest
	docker push iamkimchi/blitzshare.bootstrap.node:latest
	

k8s-destory:
	kubectl delete namespace blitzshare-bootstrap-ns

k8s-pf:
	kubectl port-forward $(kubectl get pods  | tail -n1 | awk '{print $1}') 8000:80

k8s-apply:
	kubectl apply -f k8s/namespace.yaml 
	kubectl apply -f k8s/deployment.yaml
	kubectl apply -f k8s/service.yaml
	kubectl rollout restart deployment/blitzshare-bootstrap-deployment --namespace blitzshare-bootstrap-ns
	kubectl wait -f k8s/deployment.yaml --for condition=available
