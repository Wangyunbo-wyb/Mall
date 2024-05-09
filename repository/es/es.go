package es

import (
	conf "Mall/config"
	"fmt"
	elastic "github.com/elastic/go-elasticsearch/v7"
	"github.com/sirupsen/logrus"
	"log"
)

var EsClient *elastic.Client

// InitEs 初始化es
func InitEs() {
	eConfig := conf.Config.Es
	esConn := fmt.Sprintf("http://%s:%s", eConfig.EsHost, eConfig.EsPort)
	cfg := elastic.Config{
		Addresses: []string{esConn},
	}
	client, err := elastic.NewClient(cfg)
	if err != nil {
		log.Panic(err)
	}
	EsClient = client
}

// EsHookLog 初始化log日志
func EsHookLog() *ElasticHook {
	eConfig := conf.Config.Es
	hook, err := NewElasticHook(EsClient, eConfig.EsHost, logrus.DebugLevel, eConfig.EsIndex)
	if err != nil {
		log.Panic(err)
	}
	return hook
}
