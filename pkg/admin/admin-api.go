/**
 * Copyright 2017 ~ 2025 the original author or authors<Wanglsir@gmail.com, 983708408@qq.com>.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package admin

import (
	"net/http"

	config "xcloud-webconsole/pkg/config"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
)

type adminMetricCollector struct {
	mysqlMetric        *prometheus.Desc
	ssh2DispatchMetric *prometheus.Desc
}

func newAdminMetricCollector() *adminMetricCollector {
	m1 := make(map[string]string)
	m1["env"] = "prod"
	v := []string{"hostname"}
	return &adminMetricCollector{
		mysqlMetric:        prometheus.NewDesc("fff_metrics", "Show metrics a for mysql", nil, nil),
		ssh2DispatchMetric: prometheus.NewDesc("bbb_metrics", "Show metrics a bar occu", v, m1),
	}
}

func (collect *adminMetricCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collect.ssh2DispatchMetric
	ch <- collect.mysqlMetric
}

func (collect *adminMetricCollector) Collect(ch chan<- prometheus.Metric) {
	var metricValue float64

	if 1 == 1 {
		metricValue = 1
	}

	ch <- prometheus.MustNewConstMetric(collect.mysqlMetric, prometheus.GaugeValue, metricValue)
	ch <- prometheus.MustNewConstMetric(collect.ssh2DispatchMetric, prometheus.CounterValue, metricValue, "kk")
}

// ServeStart ...
func ServeStart() {
	adminMetricCollector := newAdminMetricCollector()
	prometheus.MustRegister(adminMetricCollector)

	log.Info("Starting prometheus exporter on port: " + config.GlobalConfig.Admin.Listen)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(config.GlobalConfig.Admin.Listen, nil))
}
