package miniapp

import (
	"context"
	"fmt"
	"os"
	"time"

	"app-instance/internal/pkg/godb"
	"app-instance/internal/pkg/golog"

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
