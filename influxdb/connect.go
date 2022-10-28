package influxdb

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
)

var Client influxdb2.Client

func Connect() {
	Client = influxdb2.NewClient("http://localhost:8086", "9U9sKKDHo5JYW6ZoXCch7dqWk6bZNhVb0O-n7CP13c0NQMrO6sHt_QpUUIYDryjlQNBEcYz6k9dqbtELgmCo_Q==")
}
