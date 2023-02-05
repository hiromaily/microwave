
.PHONY: install
install:
	go install

.PHONY: run
run:
	go run main.go 210 500

.PHONY: exec
exec:
	microwave 210 500

