package msg_nsq

import (
	"fmt"
	"github.com/GYJoker/jgt/common_func"
	"github.com/GYJoker/jgt/config"
	"github.com/GYJoker/jgt/glog"
	"github.com/nsqio/go-nsq"
	"time"
)

type (
	// Manager nsq管理者
	Manager interface {
		// Publish 发布消息
		Publish(topic string, msg []byte) error

		// Subscribe 订阅消息
		Subscribe(topic, channel string, handler HandlerFunc) error
	}

	// Handler 消费者处理接口
	Handler struct {
		callback HandlerFunc
	}

	// HandlerFunc 消费者处理函数
	HandlerFunc func(msg *nsq.Message) error

	nsqManager struct {
		conf     *config.NsqConfig
		producer *nsq.Producer
	}
)

// NewManager 创建一个nsq管理者
func NewManager(conf *config.NsqConfig) Manager {
	m := &nsqManager{
		conf: conf,
	}

	if common_func.StrIsEmpty(conf.Host) || common_func.StrIsEmpty(conf.Port) {
		glog.GetLogger().Error("msg_nsq config is invalid")
		return m
	}

	c := nsq.NewConfig()
	w, err := nsq.NewProducer(fmt.Sprintf("%s:%s", conf.Host, conf.Port), c)

	if err != nil {
		glog.GetLogger().Errorf("create producer failed, err:%v", err)
		return m
	}

	m.producer = w

	return m
}

func (n *nsqManager) Publish(topic string, msg []byte) error {
	if n.producer == nil {
		glog.GetLogger().Error("producer is nil")
		return fmt.Errorf("producer is nil")
	}

	return n.producer.Publish(topic, msg)
}

func (n *nsqManager) Subscribe(topic, channel string, handler HandlerFunc) error {
	if common_func.StrIsEmpty(topic) || common_func.StrIsEmpty(channel) {
		glog.GetLogger().Error("topic or channel is empty")
		return fmt.Errorf("topic or channel is empty")
	}

	c := nsq.NewConfig()
	c.LookupdPollInterval = 15 * time.Second

	consumer, err := nsq.NewConsumer(topic, channel, c)
	if err != nil {
		glog.GetLogger().Errorf("create consumer failed, err:%v", err)
		return err
	}

	consumer.AddHandler(&Handler{callback: handler})

	err = consumer.ConnectToNSQLookupd(fmt.Sprintf("%s:%s", n.conf.Host, n.conf.HttpPort))
	if err != nil {
		glog.GetLogger().Errorf("connect to nsqlookupd failed, err:%v", err)
		return err
	}

	return nil
}

func (h *Handler) HandleMessage(msg *nsq.Message) error {
	return h.callback(msg)
}
