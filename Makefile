.PHONY: go-fix-lint
go-fix-lint:
	find . -print | grep --regex '.*\.go$$' | xargs goimports -w -local "github.com/geeeeorge/Go-diary-api"
