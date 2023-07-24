IMG ?= hisunyh/arena-app:latest

# frontend install
install:
	go mod tidy
	cd frontend && npm install --registry=https://registry.npm.taobao.org

# frontend dev
f-dev: 
	cd frontend && npm run dev

# frontend build
f-build: 
	cd frontend && npm run build

# fmt
fmt:
	go fmt ./...

# vet
vet:
	go vet ./...

# test
test: fmt vet
	go test ./... -coverprofile cover.out

# build
build: test
	go build -o ./bin/arena-app main.go

# run
run: build
	./bin/arena-app

# docker build
docker-build: f-build test
	docker build -t $(IMG) .

# docker run
docker-run:
	docker run -it --rm -p 5000:5000 $(IMG)

# docker buildx 
PLATFORMS ?= linux/arm64,linux/amd64,linux/s390x,linux/ppc64le
.PHONY: docker-buildx
docker-buildx: test ## Build and push docker image for the manager for cross-platform support
	# copy existing Dockerfile and insert --platform=${BUILDPLATFORM} into Dockerfile.cross, and preserve the original Dockerfile
	sed -e '1 s/\(^FROM\)/FROM --platform=\$$\{BUILDPLATFORM\}/; t' -e ' 1,// s//FROM --platform=\$$\{BUILDPLATFORM\}/' Dockerfile > Dockerfile.cross
	- docker buildx create --name project-v3-builder
	docker buildx use project-v3-builder
	- docker buildx build --push --platform=$(PLATFORMS) --tag ${IMG} -f Dockerfile.cross .
	- docker buildx rm project-v3-builder
	rm Dockerfile.cross

# docker push
docker-push:
	docker push $(IMG)




	