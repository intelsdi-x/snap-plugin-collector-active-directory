package activedirectory

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetCounters(t *testing.T) {
	Convey("Test GetPowershellData", t, func() {

		Convey("When metrics are valid", func() {
			metricNames := []string{
				"dra_inbound_bytes",
				"dra_inbound_objects",
				"dra_inbound_values",
				"dra_outbound_bytes",
				"dra_outbound_objects",
				"dra_outbound_values",
				"dra_pending_replication_syncs",
				"ds_client_binds",
				"ds_directory_reads",
				"ds_directory_searches",
				"ds_directory_writes",
				"kdc_as_requests",
				"kdc_tgs_requests",
				"kerberos_authentications",
				"ldap_bind_time",
				"ldap_client_session",
				"ldap_searches",
				"ldap_successful_binds",
				"ldap_writes"}
			counterDataMap := GetPowershellData(metricNames)

			Convey("Nineteen non-negative counter values should be returned", func() {
				fmt.Println(counterDataMap)
				for _, val := range counterDataMap {
					So(val, ShouldBeGreaterThanOrEqualTo, 0)
				}
			})
		})

		// This test would be the same if the actual counter path given in counters.go was not found on the system
		Convey("When metric names are not valid", func() {
			metricNames := []string{"ra_inbound_bytes", "dra_inbound_objects", "dra_inbound_values"}
			counterDataMap := GetPowershellData(metricNames)

			Convey("One -1 counter value should be returned", func() {
				fmt.Println(counterDataMap)
				for key, val := range counterDataMap {
					if key == "ra_inbound_bytes" {
						So(val, ShouldEqual, -1)
					}
				}
			})
		})

	})

}
