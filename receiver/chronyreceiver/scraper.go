// Copyright  The OpenTelemetry Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package chronyreceiver // import "github.com/open-telemetry/opentelemetry-collector-contrib/receiver/chronyreceiver"

import (
	"context"

	"github.com/tilinna/clock"
	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"

	"github.com/open-telemetry/opentelemetry-collector-contrib/receiver/chronyreceiver/internal/metadata"
)

type chronyScraper struct {
	mb *metadata.MetricsBuilder
}

func newScraper(ctx context.Context, cfg *Config, set component.ReceiverCreateSettings) *chronyScraper {
	return &chronyScraper{
		mb: metadata.NewMetricsBuilder(cfg.MetricsSettings, set.BuildInfo,
			metadata.WithStartTime(pcommon.NewTimestampFromTime(clock.FromContext(ctx).Now())),
		),
	}
}

func (cs *chronyScraper) scrape(ctx context.Context) (pmetric.Metrics, error) {
	return pmetric.NewMetrics(), nil
}
