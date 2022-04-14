# creating config file
config:
	cp ./configs/example-config.yaml ./config.yaml

# build service
build:
	cd cmd
	go build -o ./runner
# starting application
start:
	./runner