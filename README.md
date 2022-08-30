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

Basic things you need to set:
```yaml
...
rabbit:
  host: "[rabbit address]"
...
test:
  number: "number of tests"
  time_out: "timeout for messages"
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