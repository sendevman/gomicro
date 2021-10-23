module github.com/asim/go-micro/plugins/broker/segmentio/v4

go 1.16

require (
	github.com/asim/go-micro/plugins/broker/kafka/v4 v4.0.0-20211019191242-9edc569e68bb
	github.com/asim/go-micro/plugins/codec/segmentio/v4 v4.0.0-20211019191242-9edc569e68bb
	github.com/google/uuid v1.2.0
	github.com/segmentio/kafka-go v0.4.16
	go-micro.dev/v4 v4.2.1
)

replace (
	github.com/asim/go-micro/plugins/broker/kafka/v4 => ../kafka
	github.com/asim/go-micro/plugins/codec/segmentio/v4 => ../../codec/segmentio
	go-micro.dev/v4 => ../../../../go-micro
)
