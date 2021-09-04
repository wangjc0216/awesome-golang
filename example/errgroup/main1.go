package main

import (
	"context"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)
func main() {
	logger, _ := zap.NewProduction()
	log := logger.Sugar()

	url := "http://m3-pro.meetwhale.com:2345/notify"
	method := "POST"



	group ,_ := errgroup.WithContext(context.Background())
	for i:=0;i< 500;i++ {
		group.Go(func() error {
			time.Sleep(20*time.Millisecond)
			payload := strings.NewReader(`{
    "notify_type":"feishu",
    "data":[
         {
        "labels": {
            "alertname": "group_offline",
            "department": "wop",
            "project": "armani",
            "level": "high",
            "instance": "192.168.66.139:9100",
            "job": "node-exporter",
            "mode": "user"
        },
        "annotations": {
            "description": "",
            "rule_id": "101",
            "source_id": "1",
            "summary": "some summary 1 "
        },
        "endsAt": "2021-09-11T20:11:26.172+08:00",
        "generatorURL": "http://localhost:9190/graph?g0.expr=rate%28node_cpu_seconds_total%5B1m%5D%29+%3E+0\u0026g0.tab=1"
    }
    ],
    "receivers":[
        "https://open.feishu.cn/open-apis/bot/v2/hook/65c3cdeb-7fc3-4dac-8004-bae61a64714a",
        "https://open.feishu.cn/open-apis/bot/v2/hook/218a9b11-8ce3-41d7-a23c-a652e9a87eae"
    ]
}`)
			client := &http.Client {
			}
			req, err := http.NewRequest(method, url, payload)

			if err != nil {
				log.Error(err)
				return err
			}
			req.Header.Add("Content-Type", "application/json")

			res, err := client.Do(req)
			if err != nil {
				log.Error(err)
				return err
			}
			defer res.Body.Close()
			_, err = ioutil.ReadAll(res.Body)
			if err != nil {
				log.Error(err)
				return err
			}
			log.Info(res.StatusCode)
			return nil
		})
	}
	if err := group.Wait(); err != nil {
		log.Error(err)
	}
	time.Sleep(time.Minute)
}