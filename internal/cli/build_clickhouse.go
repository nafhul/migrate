// +build clickhouse

package cli

import (
	_ "github.com/ClickHouse/clickhouse-go"
	_ "github.com/nafhul/migrate/v4/database/clickhouse"
)
