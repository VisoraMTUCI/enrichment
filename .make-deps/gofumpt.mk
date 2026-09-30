GOFUMPT_BIN=$(LOCAL_BIN)/gofumpt
$(GOFUMPT_BIN):
	GOBIN=$(LOCAL_BIN) go install mvdan.cc/gofumpt@v0.7.0

