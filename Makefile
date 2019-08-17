
.PHONY: build
build: 
	go build -race .


.PHONY: install
install: 
	go install -race .