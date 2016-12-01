PATH_BASE ?= "/go/src/github.com/vidsy"
GO_BUILDER_IMAGE ?= "vidsyhq/go-builder"
REPONAME ?= "fbintegration"
SSH_KEY_NAME ?= "id_circleci_github"
BRANCH = "master"
VERSION = $(shell cat ./VERSION)

DEFAULT: test

install-dependencies:
	@docker run \
	-v /home/ubuntu/.ssh/${SSH_KEY_NAME}:/root/.ssh/${SSH_KEY_NAME} \
	-v "${CURDIR}":${PATH_BASE}/${REPONAME} \
	-w ${PATH_BASE}/${REPONAME} \
	${GO_BUILDER_IMAGE}

check-version:
	@echo "=> Checking if VERSION exists as Git tag..."
	(! git rev-list ${VERSION})

push-tag:
	@echo "=> New tag version: ${VERSION}"
	git checkout ${BRANCH}
	git pull origin ${BRANCH}
	git tag ${VERSION}
	git push origin ${BRANCH} --tags

test:
	@echo "No tests yet :("

test-ci: test

test-coverage:
	@echo "No tests yet :("
