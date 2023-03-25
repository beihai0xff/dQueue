// Package server provides the start and dependency registration of the broker server
package server

import (
	"time"

	"github.com/beihai0xff/pudding/app/broker/connector"
	"github.com/beihai0xff/pudding/app/broker/connector/kafka_connector"
	"github.com/beihai0xff/pudding/app/broker/connector/redis_connector"
	"github.com/beihai0xff/pudding/app/broker/pkg/configs"
	"github.com/beihai0xff/pudding/app/broker/storage"
	"github.com/beihai0xff/pudding/app/broker/storage/redis_storage"
	conf "github.com/beihai0xff/pudding/configs"
	"github.com/beihai0xff/pudding/pkg/log"
	"github.com/beihai0xff/pudding/pkg/mq/kafka"
	rdb "github.com/beihai0xff/pudding/pkg/redis"
)

func newQueue(config *conf.BrokerConfig) (storage.DelayStorage, connector.RealTimeConnector) {
	return newDelayStorage(config.Broker), newConnector(config.Connector)
}

// newDelayStorage create a new DelayStorage
func newDelayStorage(broker string) storage.DelayStorage {
	switch broker {
	case "redis":
		// parse Polling delay queue interval
		interval := configs.GetServerConfig().TimeSliceInterval
		t, err := time.ParseDuration(interval)
		if err != nil {
			log.Fatalf("failed to parse '%s' to time.Duration: %v", interval, err)
		}
		log.Infof("timeSlice interval is: %f seconds", t.Seconds())
		return redis_storage.NewDelayStorage(rdb.New(configs.GetRedisConfig()), uint64(t.Seconds()))
	default:
		log.Fatalf("unknown broker type: [%s]", broker)
	}
	return nil
}

// newConnector create a new RealTime Queue Connector
func newConnector(connectorType string) connector.RealTimeConnector {
	switch connectorType {
	case "pulsar":
		log.Fatalf("pulsar connector is not implemented yet")
	case "kafka":
		return kafka_connector.NewConnector(kafka.New(configs.GetKafkaConfig()))
	case "redis":
		return redis_connector.NewConnector(rdb.New(configs.GetRedisConfig()))
	default:
		log.Fatalf("unknown connectorType type: [%s]", connectorType)
	}
	return nil
}
