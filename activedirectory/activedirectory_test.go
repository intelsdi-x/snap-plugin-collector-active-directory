package activedirectory

import (
	"testing"
	"time"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"

	. "github.com/smartystreets/goconvey/convey"
)

func TestActiveDirectoryCollector(t *testing.T) {
	pm := ActiveDirectoryCollector{}

	Convey("Test ActiveDirectoryCollector", t, func() {
		Convey("Collect dra_inbound_bytes", func() {
			metrics := []plugin.Metric{
				// Create fake dra_inbound_bytes metric to make sure the CollectMetrics function is functioning correctly and returns how it should
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "activedirectory", "dra_inbound_bytes"),
					Config:    map[string]interface{}{"testfloat": float64(100.0211508)},
					Data:      100.0211508,
					Unit:      "float",
					Timestamp: time.Now(),
				},
			}
			mts, err := pm.CollectMetrics(metrics)
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, 100.0211508)
		})
		Convey("Collect dra_inbound_objects", func() {
			metrics := []plugin.Metric{
				// Create fake dra_inbound_objects metric to make sure the CollectMetrics function is functioning correctly and returns how it should
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "activedirectory", "dra_inbound_objects"),
					Config:    map[string]interface{}{"testfloat": float64(0)},
					Data:      0,
					Unit:      "float",
					Timestamp: time.Now(),
				},
			}
			mts, err := pm.CollectMetrics(metrics)
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, 0)
		})
		Convey("Collect dra_inbound_values", func() {
			metrics := []plugin.Metric{
				// Create fake dra_inbound_values metric to make sure the CollectMetrics function is functioning correctly and returns how it should
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "activedirectory", "dra_inbound_values"),
					Config:    map[string]interface{}{"testfloat": float64(0.00077732235)},
					Data:      0.00077732235,
					Unit:      "float",
					Timestamp: time.Now(),
				},
			}
			mts, err := pm.CollectMetrics(metrics)
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, 0.00077732235)
		})
		Convey("Collect dra_outbound_bytes", func() {
			metrics := []plugin.Metric{
				// Create fake dra_outbound_bytes metric to make sure the CollectMetrics function is functioning correctly and returns how it should
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "activedirectory", "dra_outbound_bytes"),
					Config:    map[string]interface{}{"testfloat": float64(1)},
					Data:      1,
					Unit:      "float",
					Timestamp: time.Now(),
				},
			}
			mts, err := pm.CollectMetrics(metrics)
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, 1)
		})
		Convey("Collect dra_outbound_objects", func() {
			metrics := []plugin.Metric{
				// Create fake dra_outbound_objects metric to make sure the CollectMetrics function is functioning correctly and returns how it should
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "activedirectory", "dra_outbound_objects"),
					Config:    map[string]interface{}{"testfloat": float64(10656)},
					Data:      10656,
					Unit:      "float",
					Timestamp: time.Now(),
				},
			}
			mts, err := pm.CollectMetrics(metrics)
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, 10656)
		})
		Convey("Collect dra_outbound_values", func() {
			metrics := []plugin.Metric{
				// Create fake dra_outbound_values metric to make sure the CollectMetrics function is functioning correctly and returns how it should
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "activedirectory", "dra_outbound_values"),
					Config:    map[string]interface{}{"testfloat": float64(36.37)},
					Data:      36.37,
					Unit:      "float",
					Timestamp: time.Now(),
				},
			}
			mts, err := pm.CollectMetrics(metrics)
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, 36.37)
		})
		Convey("Collect dra_pending_replication_syncs", func() {
			metrics := []plugin.Metric{
				// Create fake dra_pending_replication_syncs metric to make sure the CollectMetrics function is functioning correctly and returns how it should
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "activedirectory", "dra_pending_replication_syncs"),
					Config:    map[string]interface{}{"testfloat": float64(0)},
					Data:      0,
					Unit:      "float",
					Timestamp: time.Now(),
				},
			}
			mts, err := pm.CollectMetrics(metrics)
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, 0)
		})
		Convey("Collect ds_client_binds", func() {
			metrics := []plugin.Metric{
				// Create fake ds_client_binds metric to make sure the CollectMetrics function is functioning correctly and returns how it should
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "activedirectory", "ds_client_binds"),
					Config:    map[string]interface{}{"testfloat": float64(1.42)},
					Data:      1.42,
					Unit:      "float",
					Timestamp: time.Now(),
				},
			}
			mts, err := pm.CollectMetrics(metrics)
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, 1.42)
		})
		Convey("Collect ds_directory_reads", func() {
			metrics := []plugin.Metric{
				// Create fake ds_directory_reads metric to make sure the CollectMetrics function is functioning correctly and returns how it should
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "activedirectory", "ds_directory_reads"),
					Config:    map[string]interface{}{"testfloat": float64(668074.8282)},
					Data:      668074.8282,
					Unit:      "float",
					Timestamp: time.Now(),
				},
			}
			mts, err := pm.CollectMetrics(metrics)
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, 668074.8282)
		})
		Convey("Collect ds_directory_searches", func() {
			metrics := []plugin.Metric{
				// Create fake ds_directory_searches metric to make sure the CollectMetrics function is functioning correctly and returns how it should
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "activedirectory", "ds_directory_searches"),
					Config:    map[string]interface{}{"testfloat": float64(24768.5784167836)},
					Data:      24768.5784167836,
					Unit:      "float",
					Timestamp: time.Now(),
				},
			}
			mts, err := pm.CollectMetrics(metrics)
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, 24768.5784167836)
		})
		Convey("Collect ds_directory_writes", func() {
			metrics := []plugin.Metric{
				// Create fake ds_directory_writes metric to make sure the CollectMetrics function is functioning correctly and returns how it should
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "activedirectory", "ds_directory_writes"),
					Config:    map[string]interface{}{"testfloat": float64(0.78)},
					Data:      0.78,
					Unit:      "float",
					Timestamp: time.Now(),
				},
			}
			mts, err := pm.CollectMetrics(metrics)
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, 0.78)
		})
		Convey("Collect kdc_as_requests", func() {
			metrics := []plugin.Metric{
				// Create fake kdc_as_requests metric to make sure the CollectMetrics function is functioning correctly and returns how it should
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "activedirectory", "kdc_as_requests"),
					Config:    map[string]interface{}{"testfloat": float64(50.724595)},
					Data:      50.724595,
					Unit:      "float",
					Timestamp: time.Now(),
				},
			}
			mts, err := pm.CollectMetrics(metrics)
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, 50.724595)
		})
		Convey("Collect kdc_tgs_requests", func() {
			metrics := []plugin.Metric{
				// Create fake kdc_tgs_requests metric to make sure the CollectMetrics function is functioning correctly and returns how it should
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "activedirectory", "kdc_tgs_requests"),
					Config:    map[string]interface{}{"testfloat": float64(50.724595)},
					Data:      50.724595,
					Unit:      "float",
					Timestamp: time.Now(),
				},
			}
			mts, err := pm.CollectMetrics(metrics)
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, 50.724595)
		})
		Convey("Collect kerberos_authentications", func() {
			metrics := []plugin.Metric{
				// Create fake kerberos_authentications metric to make sure the CollectMetrics function is functioning correctly and returns how it should
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "activedirectory", "kerberos_authentications"),
					Config:    map[string]interface{}{"testfloat": float64(50.724595)},
					Data:      50.724595,
					Unit:      "float",
					Timestamp: time.Now(),
				},
			}
			mts, err := pm.CollectMetrics(metrics)
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, 50.724595)
		})
		Convey("Collect ldap_bind_time", func() {
			metrics := []plugin.Metric{
				// Create fake ldap_bind_time metric to make sure the CollectMetrics function is functioning correctly and returns how it should
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "activedirectory", "ldap_bind_time"),
					Config:    map[string]interface{}{"testfloat": float64(50.724595)},
					Data:      50.724595,
					Unit:      "float",
					Timestamp: time.Now(),
				},
			}
			mts, err := pm.CollectMetrics(metrics)
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, 50.724595)
		})
		Convey("Collect ldap_client_session", func() {
			metrics := []plugin.Metric{
				// Create fake ldap_client_session metric to make sure the CollectMetrics function is functioning correctly and returns how it should
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "activedirectory", "ldap_client_session"),
					Config:    map[string]interface{}{"testfloat": float64(50.724595)},
					Data:      50.724595,
					Unit:      "float",
					Timestamp: time.Now(),
				},
			}
			mts, err := pm.CollectMetrics(metrics)
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, 50.724595)
		})
		Convey("Collect ldap_searches", func() {
			metrics := []plugin.Metric{
				// Create fake ldap_searches metric to make sure the CollectMetrics function is functioning correctly and returns how it should
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "activedirectory", "ldap_searches"),
					Config:    map[string]interface{}{"testfloat": float64(50.724595)},
					Data:      50.724595,
					Unit:      "float",
					Timestamp: time.Now(),
				},
			}
			mts, err := pm.CollectMetrics(metrics)
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, 50.724595)
		})
		Convey("Collect ldap_successful_binds", func() {
			metrics := []plugin.Metric{
				// Create fake ldap_successful_binds metric to make sure the CollectMetrics function is functioning correctly and returns how it should
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "activedirectory", "ldap_successful_binds"),
					Config:    map[string]interface{}{"testfloat": float64(50.724595)},
					Data:      50.724595,
					Unit:      "float",
					Timestamp: time.Now(),
				},
			}
			mts, err := pm.CollectMetrics(metrics)
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, 50.724595)
		})
		Convey("Collect ldap_writes", func() {
			metrics := []plugin.Metric{
				// Create fake ldap_writes metric to make sure the CollectMetrics function is functioning correctly and returns how it should
				plugin.Metric{
					Namespace: plugin.NewNamespace("intel", "activedirectory", "ldap_writes"),
					Config:    map[string]interface{}{"testfloat": float64(50.724595)},
					Data:      50.724595,
					Unit:      "float",
					Timestamp: time.Now(),
				},
			}
			mts, err := pm.CollectMetrics(metrics)
			So(mts, ShouldNotBeEmpty)
			So(err, ShouldBeNil)
			So(mts[0].Data, ShouldEqual, 50.724595)
		})
	})

	Convey("Test GetMetricTypes", t, func() {
		pm := ActiveDirectoryCollector{}

		Convey("Collect All Metrics String", func() {
			mt, err := pm.GetMetricTypes(nil)
			So(err, ShouldBeNil)
			So(len(mt), ShouldEqual, 19)
		})
	})

	Convey("Test GetConfigPolicy", t, func() {
		pm := ActiveDirectoryCollector{}
		_, err := pm.GetConfigPolicy()

		Convey("No error returned", func() {
			So(err, ShouldBeNil)
		})
	})

}
