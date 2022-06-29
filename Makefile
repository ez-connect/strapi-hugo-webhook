#
# Code generated by `gkgen`
#

.PHONY: data

GOPATH := $(shell go env GOPATH)

# Environments
-include .makerc

# Service
NAME			?= $(shell grep -P -o '(?<=name: )[^\s]+' .config/service.base.yaml)
VERSION			?= $(shell grep -P -o '(?<=version: )[^\s]+' .config/service.base.yaml)
DESCRIPTION		?= $(shell grep -P -o '(?<=description: )[^\s]+' .config/service.base.yaml)
CATEGORY		?= $(shell grep -P -o '(?<=category: )[^\s]+' .config/service.base.yaml)
PACKAGE			?= $(shell grep -P -o '(?<=package: )[^\s]+' .config/service.base.yaml)
PLATFORMS		?= linux windows darwin
ARCH 			?= amd64

# Registry
REGISTRY 		?= registry.gitlab.com
REGISTRY_REPO 	?= free-mind/hub
DOCKERFILE 		?= Dockerfile
DEPLOYMENT_KIND	?= $(shell grep -P -o '(?<=kind: )[\w+]+' .config/service.k8s.yaml | tr '[:upper:]' '[:lower:]')

# OCI
ifndef IMAGE
	ifneq ($(CATEGORY),)
		IMAGE	= $(CATEGORY)-$(NAME)
	else
		IMAGE	= $(NAME)
	endif
endif

TAG				?= $(VERSION)

# Build flags
LD_FLAGS		?= -X $(PACKAGE)/base.BuildDate=$(shell date +%Y-%m-%d) \
	-X $(PACKAGE)/base.Branch=$(shell git rev-parse --abbrev-ref HEAD) \
	-X $(PACKAGE)/base.Hash=$(shell git rev-parse --short HEAD)

D_FLAGS		?= -ldflags="$(LD_FLAGS) -X $(PACKAGE)/base.BuildMode=debug"
P_FLAGS		?= -ldflags="-s -w $(LD_FLAGS) -X $(PACKAGE)/base.BuildMode=production"

# Git
# gitBranch := $(shell git rev-parse --abbrev-ref HEAD)
# gitCommit := $(shell git rev-parse --short HEAD)


# Lists all targets
help:
	@grep -B1 -E "^[a-zA-Z0-9_-]+\:([^\=]|$$)" Makefile \
		| grep -v -- -- \
		| sed 'N;s/\n/###/' \
		| sed -n 's/^#: \(.*\)###\(.*\):.*/\2###\1/p' \
		| column -t -s '###'

#: Removes untracked files from the working tree
clean:
	git clean -fdx

#: Code formatting
fmt:
	go fmt ./... > /dev/null

#: Runs the linters
lint:
	golangci-lint run --fix ./...

#: Copy '/data/*' to '/dist'
data:
ifneq ($(wildcard data/*),)
	mkdir -p dist
	cp data/* dist/
endif

# Builds and exec instead of @go run
# to run on Windows without deal with the Firewall
#: Launchs the service
run: data
	go build -o dist/$(NAME) $(D_FLAGS)
	dist/$(NAME) serve $(args)

#: Launchs the service then watching for changes
watch:
	nodemon -e go --ignore dist/ --exec make run

#: Automates testing the packages
test:
	go test -v ./... -cover

test-clean:
	go clean -testcache

###########################################################
# Generate
###########################################################
#: Parse 'base/pb/*.proto' and generate output
proto:
	protoc \
		-I $(GOPATH)/pkg/mod/github.com/srikrsna/protoc-gen-gotag@v0.6.2 \
		-I base/pb \
		--go_out=base/pb \
		--go_opt=paths=source_relative \
		--go-grpc_out=base/pb \
		--go-grpc_opt=paths=source_relative \
		$(NAME).proto

	protoc \
		-I $(GOPATH)/pkg/mod/github.com/srikrsna/protoc-gen-gotag@v0.6.2 \
		-I base/pb \
		--gotag_out=. \
		$(NAME).proto

#: Parse '.config' and generate output
gen:
	gkgen gen $(args)
	make -s proto
	make -s fmt

#: Cleans output by `gkgen`
gen-clean:
	gkgen -clean .

###########################################################
# Build
###########################################################
#: Build for platfoms defined in the PLATFORMS variable
build: gen data
	-@rm -rf dist/
	@for p in $(PLATFORMS); do \
		echo "Building for $$p"; \
		output=$(NAME); \
		if [ "$$p" = 'windows' ]; then \
			output=$$output.exe; \
		fi; \
		GOOS=$$p GOARCH=$(ARCH) go build -o dist/$$p/$$output $(G_FLAGS); \
		tar -C dist/$$p -zcvf dist/$(NAME)-$$p.tar.gz $$output > /dev/null; \
	done \

#: Launch gRPC web UI
grpcui:
	grpcui -plaintext $(args) localhost:8080

###########################################################
# OCI
###########################################################
#: Builds an OCI image using instructions in 'Dockerfile'
oci:
	podman build -t $(IMAGE):$(VERSION) -f Dockerfile $(args)

#: Pushes an image to a specified location that defined in '.makerc'
oci-push:
	podman login $(REGISTRY)
	podman push $(IMAGE):$(VERSION) $(REGISTRY)/$(REGISTRY_REPO)/$(IMAGE):$(TAG)

###########################################################
# Helm
###########################################################
#: Generates the Helm chart
helm:
	gkgen helm $(args)
	cp .config/service.k8s.yaml .chart/values.yaml
ifneq ($(wildcard .chart/Chart.lock),)
	rm .chart/Chart.lock
endif
	helm dependency build .chart/
	helm lint .chart/

#: Render chart templates locally and write to '.chart/k8s.yaml'
pod: helm
	helm template $(IMAGE) .chart/ > .chart/k8s.yaml

#: Uploads chart to the repo that defined in '.makerc'
package: helm
	helm cm-push .chart/ $(HELM_REPO)

#: Installs the chart to a remote defined in '.makerc'
install:
	helm repo update && helm install $(IMAGE) $(HELM_REPO)/$(IMAGE) -n $(NAMESPACE) --version $(VERSION)

#: Upgrades the release to the current version of the chart
upgrade:
	helm repo update && helm upgrade $(IMAGE) $(HELM_REPO)/$(IMAGE) -n $(NAMESPACE) --version $(VERSION)

#: Restarts the release
restart:
	kubectl rollout restart $(DEPLOYMENT_KIND)/$(IMAGE) -n $(NAMESPACE)

#: Uninstalls the service
uninstall:
	helm uninstall $(IMAGE) -n $(NAMESPACE)
