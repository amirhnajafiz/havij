# playful-rabbit

Test your RabbitMQT service with **playful rabbit**.

## What is this project for?
This is a golang service that tests your RabbitMQT server and gives
you status about your service.

You can set the configs and then test your MQT server.

## Configs
Copy the example config file:
```shell
cp ./configs/example-config.yaml ./config.yaml
```

Now you can set the RabbitMQT address and some configs
for the testing environment, like the number of tests, message
timeout and ...

## Start
You can run the project on docker with following command:
```shell
docker-compose up -d
```

Now you can get the test results in _log.txt_ file in **playful-rabbit** container, in _app_ directory.
