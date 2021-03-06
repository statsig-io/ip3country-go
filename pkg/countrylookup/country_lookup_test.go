package countrylookup

import (
	"math/rand"
	"testing"
	"time"
)

func TestHardcodedIPs(t *testing.T) {
	var tests = []struct {
		ip      string
		country string
	}{
		{"1.1.1.1", "US"},
		{"2.2.2.2", "FR"},
		{"3.3.3.3", "US"},
		{"4.4.4.4", "US"},
		{"5.5.5.5", "DE"},
		{"6.6.6.6", "US"},
		{"7.7.7.7", "US"},
		{"8.8.8.8", "US"},
		{"9.9.9.9", "US"},
		{"11.11.11.11", "US"},
		{"12.12.12.12", "US"},
		{"13.13.13.13", "US"},
		{"14.14.14.14", "JP"},
		{"15.15.15.15", "US"},
		{"16.16.16.16", "US"},
		{"17.17.17.17", "US"},
		{"18.18.18.18", "US"},
		{"19.19.19.19", "US"},
		{"20.20.20.20", "US"},
		{"21.21.21.21", "US"},
		{"22.22.22.22", "US"},
		{"23.23.23.23", "US"},
		{"24.24.24.24", "US"},
		{"25.25.25.25", "GB"},
		{"26.26.26.26", "US"},
		{"27.27.27.27", "CN"},
		{"28.28.28.28", "US"},
		{"29.29.29.29", "US"},
		{"30.30.30.30", "US"},
		{"31.31.31.31", "MD"},
		{"41.41.41.41", "EG"},
		{"42.42.42.42", "KR"},
		{"45.45.45.45", "CA"},
		{"46.46.46.46", "RU"},
		{"49.49.49.49", "TH"},
		{"101.101.101.101", "TW"},
		{"110.110.110.110", "CN"},
		{"111.111.111.111", "JP"},
		{"112.112.112.112", "CN"},
		{"150.150.150.150", "KR"},
		{"200.200.200.200", "BR"},
		{"202.202.202.202", "CN"},
		{"45.85.95.65", "CH"},
		{"58.96.74.25", "AU"},
		{"88.99.77.66", "DE"},
		{"25.67.94.211", "GB"},
		{"27.67.94.211", "VN"},
		{"27.62.93.211", "IN"},
	}
	lookup := New()
	for _, tt := range tests {
		t.Run(tt.ip, func(t *testing.T) {
			ans, ok := lookup.LookupIp(tt.ip)
			if !ok {
				t.Errorf("Failed to lookup ip %s", tt.ip)
			}
			if ans != tt.country {
				t.Errorf("Expected country %s but got %s", tt.country, ans)
			}
		})
	}
}

func TestRandom(t *testing.T) {
	lookup := New()
	rand.Seed(time.Now().UnixNano())
	for i := 1; i < len(lookup.ip_ranges); i++ {
		max := lookup.ip_ranges[i] - 1
		min := lookup.ip_ranges[i-1]

		expected := lookup.country_table[i]
		if lookup_ip(lookup, min) != expected {
			t.Errorf("Expected %s but got %s", expected, lookup_ip(lookup, min))
		}
		if lookup_ip(lookup, max) != expected {
			t.Errorf("Expected %s but got %s", expected, lookup_ip(lookup, max))
		}

		for j := 1; j < 100; j++ {
			random_ip := uint64(rand.Float64()*float64(max-min) + float64(min))
			if lookup_ip(lookup, random_ip) != expected {
				t.Errorf("Expected %s but got %s", expected, lookup_ip(lookup, random_ip))
			}
		}
	}
}

func lookup_ip(lookup *CountryLookup, ip_number uint64) string {
	result, ok := lookup.LookupNumericIp(ip_number)
	if !ok {
		return "--"
	}
	return result
}
