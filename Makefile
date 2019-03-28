# 2019 Rackspace, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License"): you may
# not use this file except in compliance with the License. You may obtain
# a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
# WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
# License for the specific language governing permissions and limitations
# under the License.

# Variables to be used throughout the makefile. Most of these will be passed in as
# ldflags when building/compiling the application.
app_name := servicenow
version := $(git describe --tags)
hash := $(shell git rev-parse HEAD)
timestamp := $(shell date -u '+%Y-%m-%d_%I:%M:%S%p')

.PHONY: all lint release test test_clean_dir test_run update_pages

# Default call to GNU make. Basically displays that the variables needed to run the
# subsequent make functions are being pulled correctly
all:
	@echo $(app_name): v$(version)
	@echo "\tgit hash: $(hash)"
	@echo "\ttimestamp: $(timestamp)"
	@echo "\nTop-level commands:\n\tmake lint\n\tmake build\n\tmake test"

# Small function to run golint for the packages directories
lint:
	goimports -w -l *.go
	gometalinter --config=.gometalinter.json ./...

release: gorelease

gorelease:
	git tag $(version)
	goreleaser --rm-dist

# Test related actions
# This action should run all tests related to the nethealthAPI, generate readable test coverage results,
# then update gh-pages with the latest information.
test: test_clean_dir test_run update_pages

test_clean_dir:
	rm -rf coverage/*

# Run test for the servicenow package.  If there a out coverage profile is built, convert it to a readable
# format to be used in gh-pages
test_run:
	go test -v github.com/josh5276/$(app_name) -coverprofile=coverage/result.out
	if test -f coverage/result.out; \
	then go tool cover -html=coverage/result.out -o coverage/index.html; \
	fi

update_pages:
	git add coverage
	git commit -m "Make tests: $(timestamp) - $(hash)" -- coverage
	git subtree push --prefix coverage origin gh-pages
