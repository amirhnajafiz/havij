<h1 align="center">
Carrot
</h1>

Test your **RabbitMQ** service with a **Carrot**. Carrot checks your
rabbitMQ service by subscribing and publishing over topics with Golang workers. After
that it will give you metrics that shows your rabbitMQ service status.

You can set the configs and then test your MQT server.

### Configs
Copy the example config file:
```shell
cp ./configs/example-config.yaml ./config.yaml
```

Basic things you need to set:
```yaml
...
rabbit:
  host: "[string] RabbitMQ service address"
...
consumers: "[int] number of subscribers"
providers: "[int] number of publishers"
time_out: "[int] timeout for messages"
```

Now you can set the **RabbitMQT** address and some configs
for the testing environment, like the number of tests, message
timeout and ...

## Start
### Go
Execute the main go file:
```shell
go run main.go
```

### Docker
Build and run with docker:
```shell
docker build . -t orange-carrot
docker run -d -p 2112:2112 orange-carrot
```

Metrics will be exposed as **Prometheus** metrics over ```localhost:2112/metrics```.

## Test
You can run the project on docker (with a RabbitMQ service) with following command:
```shell
docker-compose up -d
```

Now you can get the test results in _log.txt_ file in **carrot** container, in _app_ directory.

Testing result sample:
```shell
2022/04/14 10:53:13 start testing
2022/04/14 10:53:13 [storage 1970-01-01 03:30:00 +0330 +0330][duration 42.324375ms][timeout false]: Consequatur aut perferendis voluptatem sit accusantium. 
2022/04/14 10:53:13 [storage 1970-01-01 03:30:00.001 +0330 +0330][duration 46.428959ms][timeout false]: Voluptatem aut accusantium consequatur sit perferendis. 
2022/04/14 10:53:13 [storage 1970-01-01 03:30:00.002 +0330 +0330][duration 50.0605ms][timeout false]: Voluptatem consequatur aut sit perferendis accusantium. 
2022/04/14 10:53:13 [storage 1970-01-01 03:30:00.003 +0330 +0330][duration 53.646917ms][timeout false]: Perferendis accusantium consequatur sit voluptatem aut. 
2022/04/14 10:53:13 [storage 1970-01-01 03:30:00.004 +0330 +0330][duration 56.380917ms][timeout false]: Consequatur accusantium perferendis sit aut voluptatem. 
2022/04/14 10:53:13 [storage 1970-01-01 03:30:00.005 +0330 +0330][duration 59.531125ms][timeout false]: Aut voluptatem perferendis sit accusantium consequatur. 
2022/04/14 10:53:13 [storage 1970-01-01 03:30:00.006 +0330 +0330][duration 62.425875ms][timeout false]: Sit accusantium aut consequatur voluptatem perferendis. 
2022/04/14 10:53:13 [storage 1970-01-01 03:30:00.007 +0330 +0330][duration 65.155208ms][timeout false]: Perferendis aut sit accusantium voluptatem consequatur. 
2022/04/14 10:53:13 [storage 1970-01-01 03:30:00.008 +0330 +0330][duration 67.785292ms][timeout false]: Voluptatem accusantium aut consequatur sit perferendis. 
2022/04/14 10:53:13 [storage 1970-01-01 03:30:00.009 +0330 +0330][duration 69.869208ms][timeout false]: Accusantium voluptatem sit perferendis aut consequatur. 
```
