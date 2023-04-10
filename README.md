A simple message queue server (message broker)

In this tutorial we will create a simple message

1. What is a message queue? It has been using in Event driven architecture, microservice etc. Let's look at existing examples:
- Redis:
    * start redis server
    * redis-cli blpop
    * redis-cli rpush
- rabbitmq:
    * start rabbitmq server
    * go run examples.go

2. Why do we need our own MQ broker if it exists already? 
- Extend the functionality of an existing DB (immudb in this example)
- Use features you need

3. Some things you need to pay attention to:
- consumer blocks its own process 
- there are public server, consumer and producer (at least 3 computers)
- producer can push a message even consumer does not exist
- MQ based on a protocol (AMQP for rabbitMQ, Redis serialization protocol (RESP) and Google protobuf in this example)
- to connect to MQ server beside IP and port you need mq name

4. To make your own MQ server you need the following:
- define your protocol. In this example, Google protocol buffer will help us a lot. Besides binary format of a protocol, it will generate sources with all code you need, except RPC functions. Explaining gRPC out of this scope. I recomend the following article https://grpc.io/docs/languages/go/quickstart/. We will define two RPC functions: MQput for the producer and MQpop for the consumer.
- define your queue. See https://www.educative.io/answers/what-are-channels-in-golang and compare to https://www.rabbitmq.com/tutorials/tutorial-one-go.html. See the difference? Yep, Golang channels is a message queue or message queue is Golang channels. I think that Go developers took the principle how processes communicate each other and had reproduced in Golang.

5. To patch existing DB (immuDB for example as simple as possible, without cli and authentication):
- pull request the sources
- modify .proto files and rebuild them
- add RPC functions to the server, see the diff https://github.com/codenotary/immudb/compare/master...alexbezu:immudb:master
- check the producer and the consumer through the immuDB