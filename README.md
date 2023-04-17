A simple message queue server (message broker) tutorial.

In this tutorial, we will create a simple message queue server, producer and consumer. And then modify OpenSource DataBase immuDB to extend its functionality with a message queue server.

1. What is a message queue? It has been used in Event-driven architecture, microservice, etc. Let's look at existing examples:
- Redis:
    * start redis server
    * redis-cli blpop
    * redis-cli rpush
- RabbitMQ:
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
- to connect to the MQ server besides the IP address and the port you need an mq name

4. To make your own MQ server you need the following:
- define your protocol. In this example, the Google protocol buffer will help us a lot. Besides the binary format definition for the protocol, it will generate sources with all code you need, except RPC functions. Explaining gRPC out of this scope. If you don't know I recommend the following article https://grpc.io/docs/languages/go/quickstart/. We will define two RPC functions: MQput for the producer and MQpop for the consumer.
- define your queue. See https://www.educative.io/answers/what-are-channels-in-golang and compare to https://www.rabbitmq.com/tutorials/tutorial-one-go.html. See the difference? Yep, Golang channels is a message queue or a message queue is Golang channels. I think that Golang developers took the principle of how processes communicate with each other and reproduced it in Golang.

5. To patch existing DB (immuDB for example as simple as possible, without cli and authentication):
- pull request the sources
- modify .proto files and rebuild them
- add RPC functions to the server, see the diff https://github.com/codenotary/immudb/compare/master...alexbezu:immudb:master
- don't forget to uncomment the 'replace' statement in go.mod and change the schema in import for send and receive files.
- check the producer and the consumer through the immuDB

Video addendum for this tutorial: TODO:
