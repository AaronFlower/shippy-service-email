## Event Driven Architeture
业务场景，创建用户之后发邮件。传统的系统设计是，在用户创建后去调用发邮件的接口。(有缺点)

而基于 Event-Driven 的设计是，在用户创建后， Publish 一个事件，邮件服务订阅这个事件。通过 Pub/Sub 可以实现
服务之间的解耦(Decouple)。

- Event Driven Architecture
- NATS
- Pub/Sub
- go-micro broker

### How to run
```
make build && make run
```

then waiting the enent

```
2018/12/20 08:02:31 Listening on [::]:50054
2018/12/20 08:02:31 Broker Listening on [::]:42601
2018/12/20 08:02:31 Registering node: go.micro.srv.email-7299e66e-83c2-4b5a-826b-03890a111712
2018/12/20 08:02:41 Picked up a new message
2018/12/20 08:02:41 Sending email to: Aaron Flower
2018/12/20 08:02:46 Picked up a new message
2018/12/20 08:02:46 Sending email to: Aaron Flower
```
