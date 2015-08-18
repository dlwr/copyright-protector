// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#include <wand/MagickWand.h>
*/
import "C"

type MetricType int

const (
	METRIC_UNDEFINED                          MetricType = C.UndefinedMetric
	METRIC_ABSOLUTE_ERROR                     MetricType = C.AbsoluteErrorMetric
	METRIC_MEAN_ABSOLUTE_ERROR                MetricType = C.MeanAbsoluteErrorMetric
	METRIC_MEAN_ERROR_PER_PIXEL               MetricType = C.MeanErrorPerPixelMetric
	METRIC_MEAN_SQUARED_ERROR                 MetricType = C.MeanSquaredErrorMetric
	METRIC_PEAK_ABSOLUTE_ERROR                MetricType = C.PeakAbsoluteErrorMetric
	METRIC_PEAK_SIGNAL_TO_NOISE_RATIO         MetricType = C.PeakSignalToNoiseRatioMetric
	METRIC_ROOT_MEAN_SQUARED_ERROR            MetricType = C.RootMeanSquaredErrorMetric
	METRIC_NORMALIZED_CROSS_CORRELATION_ERROR MetricType = C.NormalizedCrossCorrelationErrorMetric
	METRIC_FUZZ_ERROR                         MetricType = C.FuzzErrorMetric
)
