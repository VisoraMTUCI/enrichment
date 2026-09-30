GOLANG_LINT_BIN=$(LOCAL_BIN)/golangci-lint
$(GOLANG_LINT_BIN):
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(LOCAL_BIN) v2.10.1
