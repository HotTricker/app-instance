package miniapp

import (
	"context"
	"fmt"
	"os"
	"time"

	"app-instance/internal/pkg/godb"
	"app-instance/internal/pkg/golog"
	"app-instance/internal/pkg/gomq"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/go-redis/redis/v8"
)

type miniApp struct {
	Gin *gin.Engine
	cfg *Config

	Logger   *logrus.Logger
	DB       *godb.DB
	RedisCli *redis.Client
	Mq       *gomq.RabbitMQ
}

var (
	App *miniApp
)

const (
	Version = "V1.0.0.0"
)

func init() {
	App = newApp()
}

func newApp() *miniApp {
	return &miniApp{
		Gin: gin.New(),
	}
}

func (m *miniApp) Init(path string) error {
	if err := loadConfig(path); err != nil {
		return err
	}
	OutputInfo("Log", m.cfg.Log.Path)
	OutputInfo("Serve address", m.cfg.Server.Addr)

	m.registerLog()
	if err := m.registerOrm(); err != nil {
		return err
	}

	if err := m.registerRedis(); err != nil {
		return err
	}

	if err := m.registerMq(); err != nil {
		return err
	}
	return nil
}

func (m *miniApp) Start() error {
	m.Logger.Infoln("App start ...")
	return m.Gin.Run(":" + m.cfg.Server.Addr)
}

func OutputInfo(tag string, value interface{}) {
	fmt.Printf("%-18s    %v\n", tag+":", value)
}

func (m *miniApp) registerLog() error {
	switch m.cfg.Log.Path {
	case "stdout":
		golog.InitLogrusLog(os.Stdout)
	case "":
		golog.InitLogrusLog(os.Stdout)
	default:
		golog.InitLogrusRotateLog(m.cfg.Log.Path)
	}
	m.Logger = golog.Mlogger
	return nil
}

func (m *miniApp) registerOrm() error {
	m.DB = godb.NewDatabase(m.cfg.Db)
	return m.DB.Open()
}

func (m *miniApp) registerRedis() error {
	m.RedisCli = redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    m.cfg.Redis.MasterName,
		SentinelAddrs: m.cfg.Redis.Urls,
	})
	var ctx = context.Background()
	key := "status"
	value := "start"
	ttl := 24 * time.Hour
	err := m.RedisCli.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return err
	} else {
		m.Logger.Infoln("redis start")
	}

	return nil
}

func (m *miniApp) registerMq() error {
	m.Mq = gomq.NewRabbitMQ("queue_publisher", "exchange_publisher", "key1")
	if _, err := m.Mq.Channel.QueueDeclare(
		m.Mq.QueueName, // 队列名
		true,           // 是否持久化
		false,          // 是否自动删除(前提是至少有一个消费者连接到这个队列，之后所有与这个队列连接的消费者都断开时，才会自动删除。注意：生产者客户端创建这个队列，或者没有消费者客户端与这个队列连接时，都不会自动删除这个队列)
		false,          // 是否为排他队列（排他的队列仅对“首次”声明的conn可见[一个conn中的其他channel也能访问该队列]，conn结束后队列删除）
		false,          // 是否阻塞
		nil,            //额外属性
	); err != nil {
		return err
	}

	if err := m.Mq.Channel.ExchangeDeclare(
		m.Mq.Exchange, //交换器名
		"topic",       //exchange type：一般用fanout、direct、topic
		true,          // 是否持久化
		false,         //是否自动删除（自动删除的前提是至少有一个队列或者交换器与这和交换器绑定，之后所有与这个交换器绑定的队列或者交换器都与此解绑）
		false,         //设置是否内置的。true表示是内置的交换器，客户端程序无法直接发送消息到这个交换器中，只能通过交换器路由到交换器这种方式
		false,         // 是否阻塞
		nil,           // 额外属性
	); err != nil {
		return err
	}

	if err := m.Mq.Channel.QueueBind(
		m.Mq.QueueName,
		m.Mq.RoutingKey, // bindkey 用于消息路由分发的key
		m.Mq.Exchange,
		false, // 是否阻塞
		nil,   // 额外属性
	); err != nil {
		return err
	}
	return nil
}
