package main

import (
	"net/http"

	"github.com/BKH7/kafka-simple/realtime/conn"
	"github.com/BKH7/kafka-simple/realtime/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {

	err := conn.NewKafkaConnection(viper.GetString("kafka.server"))
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.Root)
	r.HandleFunc("/health-check", handlers.HealthCheck)

	logrus.Infof("HTTP Serve on 127.0.0.1:%s", viper.GetString("server.port"))
	http.ListenAndServe(viper.GetString("server.port"), r)
}
