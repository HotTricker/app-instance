package miniapp

import (
	"fmt"
	"os"

	"app-instance/internal/pkg/golog"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type miniApp struct {
	Gin *gin.Engine
	cfg *Config

	Logger *logrus.Logger
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

	return nil
}

func (m *miniApp) Start() error {
	m.Logger.Infoln("App start ...")
	return m.Gin.Run(m.cfg.Server.Addr)
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
