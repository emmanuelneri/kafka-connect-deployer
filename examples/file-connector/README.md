# file-connector deploy example

`docker-compose up`

```
emmanuelneri@MacBook-Pro-de-Emmanuel file-connector % docker logs -f file-connector_kafka-connect-deployer_1
2021/05/31 01:00:43 Deployer starting....
2021/05/31 01:00:43 waiting 15s before start
2021/05/31 01:00:58 start deploy file-sink.json
[http-client] 2021/05/31 01:00:58 [DEBUG] POST http://kafka-connect:8083/connectors
[http-client] 2021/05/31 01:00:58 [ERR] POST http://kafka-connect:8083/connectors request failed: Post "http://kafka-connect:8083/connectors": dial tcp 172.31.0.4:8083: connect: connection refused
[http-client] 2021/05/31 01:00:58 [DEBUG] POST http://kafka-connect:8083/connectors: retrying in 1s (20 left)
[http-client] 2021/05/31 01:00:59 [ERR] POST http://kafka-connect:8083/connectors request failed: Post "http://kafka-connect:8083/connectors": dial tcp 172.31.0.4:8083: connect: connection refused
[http-client] 2021/05/31 01:00:59 [DEBUG] POST http://kafka-connect:8083/connectors: retrying in 2s (19 left)
[http-client] 2021/05/31 01:01:01 [ERR] POST http://kafka-connect:8083/connectors request failed: Post "http://kafka-connect:8083/connectors": dial tcp 172.31.0.4:8083: connect: connection refused
[http-client] 2021/05/31 01:01:01 [DEBUG] POST http://kafka-connect:8083/connectors: retrying in 4s (18 left)
[http-client] 2021/05/31 01:01:05 [ERR] POST http://kafka-connect:8083/connectors request failed: Post "http://kafka-connect:8083/connectors": dial tcp 172.31.0.4:8083: connect: connection refused
[http-client] 2021/05/31 01:01:05 [DEBUG] POST http://kafka-connect:8083/connectors: retrying in 8s (17 left)
2021/05/31 01:01:13 file-sink.json - status: 201 Created 
2021/05/31 01:01:13 start deploy file-source.json
[http-client] 2021/05/31 01:01:13 [DEBUG] POST http://kafka-connect:8083/connectors
2021/05/31 01:01:13 file-source.json - status: 201 Created 
emmanuelneri@MacBook-Pro-de-Emmanuel file-connector % 
emmanuelneri@MacBook-Pro-de-Emmanuel file-connector % 
emmanuelneri@MacBook-Pro-de-Emmanuel file-connector % 
emmanuelneri@MacBook-Pro-de-Emmanuel file-connector % 
emmanuelneri@MacBook-Pro-de-Emmanuel file-connector % docker logs -f file-connector_kafka-connect-deployer_1
2021/05/31 01:03:11 Deployer starting....
2021/05/31 01:03:11 waiting 15s before start
2021/05/31 01:03:26 start deploy file-sink.json
2021/05/31 01:03:26 [DEBUG] POST http://kafka-connect:8083/connectors
2021/05/31 01:03:26 [ERR] POST http://kafka-connect:8083/connectors request failed: Post "http://kafka-connect:8083/connectors": dial tcp 172.31.0.4:8083: connect: connection refused
2021/05/31 01:03:26 [DEBUG] POST http://kafka-connect:8083/connectors: retrying in 1s (20 left)
2021/05/31 01:03:27 [ERR] POST http://kafka-connect:8083/connectors request failed: Post "http://kafka-connect:8083/connectors": dial tcp 172.31.0.4:8083: connect: connection refused
2021/05/31 01:03:27 [DEBUG] POST http://kafka-connect:8083/connectors: retrying in 2s (19 left)
2021/05/31 01:03:29 [ERR] POST http://kafka-connect:8083/connectors request failed: Post "http://kafka-connect:8083/connectors": dial tcp 172.31.0.4:8083: connect: connection refused
2021/05/31 01:03:29 [DEBUG] POST http://kafka-connect:8083/connectors: retrying in 4s (18 left)
2021/05/31 01:03:33 [ERR] POST http://kafka-connect:8083/connectors request failed: Post "http://kafka-connect:8083/connectors": dial tcp 172.31.0.4:8083: connect: connection refused
2021/05/31 01:03:33 [DEBUG] POST http://kafka-connect:8083/connectors: retrying in 8s (17 left)
2021/05/31 01:03:41 file-sink.json - status: 201 Created 
2021/05/31 01:03:41 start deploy file-source.json
2021/05/31 01:03:41 [DEBUG] POST http://kafka-connect:8083/connectors
2021/05/31 01:03:41 file-source.json - status: 201 Created 
```