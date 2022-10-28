package influxdb

import (
	"github.com/ario-team/glassnode-api/config"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var Client influxdb2.Client

func Connect() {
	Client = influxdb2.NewClient(config.Viper.GetString("INFLUX_URL"), config.Viper.GetString("INFLUX_TOKEN"))
}
