package sqlcache

import (
	"database/sql/driver"
	"testing"
)

func TestNoopHash(t *testing.T) {
	tcs := []struct {
		query    string
		args     []driver.NamedValue
		expected string
	}{
		{
			query: `
			-- @cache-ttl 5
			-- @cache-max-rows 10
			SELECT name, pages FROM books WHERE pages > $1
			`,
			args: []driver.NamedValue{
				{
					Ordinal: 1,
					Value:   10,
				},
			},
			expected: "--@cache-ttl5--@cache-max-rows10SELECTname,pagesFROMbooksWHEREpages>$1:[{ 1 10}]",
		},
	}

	for _, tc := range tcs {
		h, err := NoopHash(tc.query, tc.args)
		if err != nil {
			panic(err)
		}
		if tc.expected != h {
			panic(h)
		}
	}
}
