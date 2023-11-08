.PHONY: build-local


build-local:
	@goreleaser release --snapshot --skip=publish --clean
