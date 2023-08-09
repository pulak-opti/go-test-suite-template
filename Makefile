.PHONY: test prepare
prepare: 
	npm install

test: prepare
	go run main.go & \
	npm test \
	fuser -k tcp/8080
