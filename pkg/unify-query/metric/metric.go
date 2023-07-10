// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

package metric

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/otel/trace"

	"github.com/TencentBlueKing/bkmonitor-datalink/pkg/unify-query/log"
)

const (
	SpaceActionRefresh = "refresh"
	SpaceActionGet     = "get"
	SpaceActionParse   = "parse"

	ActionInfo        = "info"
	ActionQuery       = "query"
	ActionLabelValues = "label_values"

	TypeTS     = "ts"
	TypeES     = "es"
	TypePromql = "promql"

	StatusReceived = "received"
	StatusSuccess  = "success"
	StatusFailed   = "failed"
)

var (
	requestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "unify_query",
			Name:      "request_count_total",
			Help:      "request handled count",
		},
		[]string{"action", "type", "status"},
	)

	requestHandleSecondHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "unify_query",
			Name:      "request_handle_seconds",
			Help:      "request handle seconds",
			Buckets:   []float64{0, 0.5, 1, 3, 5, 10, 30},
		},
		[]string{"url"},
	)

	spaceRequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "unify_query",
			Name:      "space_request_count_total",
			Help:      "space request handled count",
		},
		[]string{"action", "space_uid", "status"},
	)

	spaceTableIDFieldRequestCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "unify_query",
			Name:      "space_table_id_field_request_count_total",
			Help:      "space request handled count",
		},
		[]string{"action", "space_uid", "table_id", "metric_name", "status"},
	)
	spaceDetailParseCount = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Namespace: "unify_query",
			Name:      "space_detail_parse_count_total",
			Help:      "space detail parse count",
		},
		[]string{"action", "space_uid", "status"},
	)
	resultTableInfo = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Namespace: "unify_query",
			Name:      "result_table_info",
		},
		[]string{
			"rt_table_id", "rt_bk_biz_id", "rt_data_id",
			"rt_measurement_type", "vm_table_id",
		},
	)
	tsdbRequestSecondHistogram = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "unify_query",
			Name:      "tsdb_request_seconds",
			Help:      "tsdb request seconds",
			Buckets:   []float64{0, 0.5, 1, 3, 5, 10, 30},
		},
		[]string{"space_uid", "tsdb_type"},
	)
)

func TsDBRequestSecond(ctx context.Context, duration time.Duration, params ...string) {
	metric, err := tsdbRequestSecondHistogram.GetMetricWithLabelValues(params...)
	observe(ctx, metric, err, duration, params...)
}

func ResultTableInfoSet(ctx context.Context, value float64, params ...string) {
	metric, err := resultTableInfo.GetMetricWithLabelValues(params...)
	gaugeSet(ctx, metric, err, value, params...)
}

// RequestCountInc http 访问指标
func RequestCountInc(ctx context.Context, params ...string) {
	metric, err := requestCount.GetMetricWithLabelValues(params...)
	counterInc(ctx, metric, err, params...)
}

func RequestSecond(ctx context.Context, duration time.Duration, params ...string) {
	metric, err := requestHandleSecondHistogram.GetMetricWithLabelValues(params...)
	observe(ctx, metric, err, duration, params...)
}

// SpaceRequestCountInc 获取空间
func SpaceRequestCountInc(ctx context.Context, params ...string) {
	metric, err := spaceRequestCount.GetMetricWithLabelValues(params...)
	counterInc(ctx, metric, err, params...)
}

// SpaceTableIDFieldRequestCountInc 空间获取指标
func SpaceTableIDFieldRequestCountInc(ctx context.Context, params ...string) {
	metric, err := spaceTableIDFieldRequestCount.GetMetricWithLabelValues(params...)
	counterInc(ctx, metric, err, params...)
}

// SpaceDetailParseCountInc parse the space detail metric
func SpaceDetailParseCountInc(ctx context.Context, params ...string) {
	metric, err := spaceDetailParseCount.GetMetricWithLabelValues(params...)
	counterInc(ctx, metric, err, params...)
}

func gaugeSet(
	ctx context.Context, metric prometheus.Gauge, err error, value float64, params ...string,
) {
	if err != nil {
		log.Warnf(ctx, "metric gauge: %v failed, error:%s", params, err)
		return
	}

	metric.Set(value)
}

// handleCount
func counterInc(
	ctx context.Context, metric prometheus.Counter, err error, params ...string,
) {
	if err != nil {
		log.Warnf(ctx, "metric counter:%v failed,error:%s", params, err)
		return
	}

	sp := trace.SpanFromContext(ctx).SpanContext()
	if sp.IsSampled() {
		exemplarAdder, ok := metric.(prometheus.ExemplarAdder)
		if ok {
			exemplarAdder.AddWithExemplar(1, prometheus.Labels{
				"traceID": sp.TraceID().String(),
				"spanID":  sp.SpanID().String(),
			})
		} else {
			log.Errorf(ctx, "metric type is wrong: %T, %v", metric, metric)
		}
	} else {
		metric.Inc()
	}
}

func observe(
	ctx context.Context, metric prometheus.Observer, err error, duration time.Duration, params ...string,
) {
	if err != nil {
		log.Warnf(ctx, "metric histogram:%v failed,error:%s", params, err)
		return
	}

	sp := trace.SpanFromContext(ctx).SpanContext()
	if sp.IsSampled() {
		// exemplarObserve 只支持 histograms 类型，使用 summary 会报错
		exemplarObserve, ok := metric.(prometheus.ExemplarObserver)
		if ok {
			exemplarObserve.ObserveWithExemplar(duration.Seconds(), prometheus.Labels{
				"traceID": sp.TraceID().String(),
				"spanID":  sp.SpanID().String(),
			})
		} else {
			log.Errorf(ctx, "metric type is wrong: %T, %v", metric, metric)
		}
	} else {
		metric.Observe(duration.Seconds())
	}

}

// init
func init() {
	prometheus.MustRegister(
		requestCount, requestHandleSecondHistogram, resultTableInfo, tsdbRequestSecondHistogram,
		spaceRequestCount, spaceTableIDFieldRequestCount, spaceDetailParseCount,
	)
}