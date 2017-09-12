BRANCH ?= "master"
GO_BUILDER_IMAGE ?= "vidsyhq/go-builder"
PACKAGES ?= "."
PATH_BASE ?= "/go/src/github.com/vidsy"
REPONAME ?= "fbintegration"
SSH_KEY_NAME ?= "id_circleci_github"
VERSION ?= $(shell cat ./VERSION)

DEFAULT: test

build:
	@go build "${PACKAGES}"

build-ci:
	@docker run \
	-it \
	--rm \
	-v "${CURDIR}/..":${PATH_BASE} \
	-w ${PATH_BASE}/${REPONAME} \
	--entrypoint=go \
	${GO_BUILDER_IMAGE} build "${PACKAGES}"

check-version:
	@echo "=> Checking if VERSION exists as Git tag..."
	(! git rev-list ${VERSION})

install-dependencies:
	@docker run \
	-v /home/ubuntu/.ssh/${SSH_KEY_NAME}:/root/.ssh/${SSH_KEY_NAME} \
	-v "${CURDIR}":${PATH_BASE}/${REPONAME} \
	-w ${PATH_BASE}/${REPONAME} \
	${GO_BUILDER_IMAGE}

push-tag:
	@echo "=> New tag version: ${VERSION}"
	git checkout ${BRANCH}
	git pull origin ${BRANCH}
	git tag ${VERSION}
	git push origin ${BRANCH} --tags

test:
	@go test "${PACKAGES}"

test-ci:
	@docker run \
	-v "${CURDIR}/..":${PATH_BASE} \
	-w ${PATH_BASE}/${REPONAME} \
	--entrypoint=go \
	${GO_BUILDER_IMAGE} test "${PACKAGES}" -cover

vet:
	@go vet "${PACKAGES}"

vet-ci:
	@docker run \
	-it \
	--rm \
	-v "${CURDIR}/..":${PATH_BASE} \
	-w ${PATH_BASE}/${REPONAME} \
	--entrypoint=go \
	${GO_BUILDER_IMAGE} vet "${PACKAGES}"