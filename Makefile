COMPONENT := github.com/mandelsoft/ocmdemoinstaller

IMAGE                                          = mandelsoft/ocmdemoinstaller

REPO_ROOT                                      := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
VERSION                                        = $(shell cat $(REPO_ROOT)/VERSION)
COMMIT                                         = $(shell git rev-parse HEAD)
EFFECTIVE_VERSION                              = $(VERSION)-$(COMMIT)

.PHONY: ctf
ctf: ca
	ocm transfer ca gen/ca gen/ctf
	
.PHONY: ca
ca: build gen
	ocm create ca -f $(COMPONENT) "$(VERSION)" mandelsoft gen/ca
	ocm add resources gen/ca VERSION="$(VERSION)" COMMIT="$(COMMIT)" IMAGE="$(IMAGE):$(VERSION)" resources.yaml

.PHONY: build
build:
	docker build -t $(IMAGE):$(VERSION) .

.PHONY: dummy
dummy: commit ca

.PHONY: patch
patch: clean incpatch ctf

.PHONY: minor
minor: clean incminor ctf

.PHONY: major
major: clean incmajor ctf


.PHONY: release-patch
release-patch: clean incpatch release

.PHONY: release-minor
release-minor: clean incminor release

.PHONY: release-major
release-major: clean incmajor release


.PHONY: incpatch
incpatch:
	semver -i patch $(VERSION) | tee VERSION

.PHONY: incminor
incminor:
	semver -i minor $(VERSION) | tee VERSION

.PHONY: incmajor
incmajor:
	semver -i major $(VERSION) | tee VERSION

.PHONY: push 
push: ctf
	ocm transfer ctf -f gen/ctf ghcr.io/mandelsoft/cnudie

.PHONY: gen
gen:
	mkdir -p gen

.PHONY: image
image:
	docker build -t $(IMAGE):$(VERSION)

.PHONY: commit
commit: image
	git add .
	git commit -m "release $(VERSION)"

.PHONY: release
release: commit push
	
.PHONY: info
info:
	@echo "VERSION:  $(VERSION)"
	@echo "COMMIT;   $(COMMIT)"

.PHONY: clean
clean:
	rm -rf gen
