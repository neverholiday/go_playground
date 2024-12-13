package main

import (
	"context"
	"fmt"
	"time"

	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/api/write"
)

func main() {

	token := "MyInitialAdminToken0=="
	url := "http://localhost:8086"
	client := influxdb2.NewClient(url, token)
	defer client.Close()

	ctx := context.Background()

	s, err := client.Ping(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println("ping status = ", s)

	org := "docs"
	bucket := "home"
	writeAPI := client.WriteAPIBlocking(org, bucket)
	tags := map[string]string{
		"tagname1": "tagvalue1",
	}
	fields := map[string]interface{}{
		"field1": 1,
	}
	point := write.NewPoint("measurement1", tags, fields, time.Now())
	err = writeAPI.WritePoint(ctx, point)

	if err != nil {
		panic(err)
	}

	queryAPI := client.QueryAPI(org)
	query := `from(bucket: "home")
            |> range(start: -10m)
            |> filter(fn: (r) => r._measurement == "measurement1")`
	results, err := queryAPI.Query(ctx, query)
	if err != nil {
		panic(err)
	}
	for results.Next() {
		fmt.Println(results.Record())
	}
	if err := results.Err(); err != nil {
		panic(err)
	}
}
