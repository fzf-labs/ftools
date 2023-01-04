SHELL = /bin/bash
GitBranch = $(shell git rev-parse --abbrev-ref HEAD)

.PHONY: mod
# make mod  golang库更新
mod:
	#环境更新-开始
	@go mod download
	@go mod tidy
	#环境更新-结束

.PHONY: fmt
# make fmt  格式化代码
fmt:
	@gofmt -s -w .

.PHONY: vet
# make vet golang官方命令,用于检查代码中的问题.
vet:
	@go vet ./...

.PHONY: ci-lint
# make ci-lint  golang使用最多的第三方静态程序分析工具
ci-lint:
	@golangci-lint run ./...

.PHONY: git-clean
# make git-clean  git clean
git-clean:
	#清除开始
	@git checkout --orphan latest_branch
	@git add -A
	@git commit -am "clean"
	@git branch -D ${GitBranch}
	@git branch -m ${GitBranch}
	@git push -f origin ${GitBranch}
	#清除结束

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

