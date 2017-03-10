package activedirectory

import (
	"fmt"
	"time"

	"github.com/intelsdi-x/snap-plugin-lib-go/v1/plugin"
)

// Collector implementation. Needed even if empty. Create an empty struct to use as receiver type of methods below
type ActiveDirectoryCollector struct {
}

func stringInNamespace(givenString string) bool {
	availableMetrics := []string{"dra_inbound_bytes",
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
	for _, metricName := range availableMetrics {
		if metricName == givenString {
			return true
		}
	}
	return false
}

/*
* CollectMetrics collects metrics for testing.
* CollectMetrics() is called by Snap when a task (which is collecting one+ of the metrics returned from the GetMetricTypes()) is started.
* Input: A slice of all the metric types being collected.
* Output: A slice (list) of the collected metrics as plugin.Metric with their values and an error if failure.
 */
func (ActiveDirectoryCollector) CollectMetrics(mts []plugin.Metric) ([]plugin.Metric, error) {
	metrics := []plugin.Metric{} // Create a slice of MetricType objects. This is where the metrics requested by the task will be stored
	// Iterate through metrics first time to create slice of metric names
	metricNames := make([]string, 0)
	for _, mt := range mts {
		metricNames = append(metricNames, mt.Namespace[len(mt.Namespace)-1].Value)
	}
	// Get metric data from powershell script if data has not been set already (for testing). -1 means there was an error getting that metric from system
	counterData := GetPowershellData(metricNames)
	// Iterate through each of the metrics specified by task to collect
	for idx, mt := range mts {
		// Make sure the metric given in the Task is actually a metric provided by this plugin
		if stringInNamespace(mt.Namespace[len(mt.Namespace)-1].Value) {
			mts[idx].Timestamp = time.Now() // Set metric timestamp
			// Make sure config hasn't been set for testing (SEE activedirectory_test.go)
			if val, err := mt.Config.GetFloat("testfloat"); err == nil {
				mts[idx].Data = val
			} else {
				mts[idx].Data = counterData[mt.Namespace[len(mt.Namespace)-1].Value] // Set metric value
			}
			metrics = append(metrics, mts[idx])
		} else {
			return nil, fmt.Errorf("Invalid metric: %v", mt.Namespace.Strings())
		}
	}
	return metrics, nil
}

/*
 * GetMetricTypes returns a list of available metric types
 * GetMetricTypes() is called when this plugin is loaded in order to populate the "metric catalog" (where Snap
 * stores all of the available metrics for each plugin)
 * Input: Config info. This information comes from global Snap config settings
 * Output: A slice (list) of all plugin metrics, which are available to be collected by tasks
 */
func (ActiveDirectoryCollector) GetMetricTypes(cfg plugin.Config) ([]plugin.Metric, error) {
	// slice to store list of all available active directory metrics
	mts := []plugin.Metric{}

	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "activedirectory", "dra_inbound_bytes"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "activedirectory", "dra_inbound_objects"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "activedirectory", "dra_inbound_values"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "activedirectory", "dra_outbound_bytes"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "activedirectory", "dra_outbound_objects"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "activedirectory", "dra_outbound_values"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "activedirectory", "dra_pending_replication_syncs"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "activedirectory", "ds_client_binds"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "activedirectory", "ds_directory_reads"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "activedirectory", "ds_directory_searches"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "activedirectory", "ds_directory_writes"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "activedirectory", "kerberos_authentications"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "activedirectory", "kdc_as_requests"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "activedirectory", "kdc_tgs_requests"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "activedirectory", "ldap_bind_time"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "activedirectory", "ldap_client_session"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "activedirectory", "ldap_searches"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "activedirectory", "ldap_successful_binds"),
		Version:   1,
	})
	mts = append(mts, plugin.Metric{
		Namespace: plugin.NewNamespace("intel", "activedirectory", "ldap_writes"),
		Version:   1,
	})

	return mts, nil
}

/*
 * GetConfigPolicy() returns the config policy for this plugin
 *   A config policy allows users to provide configuration info to the plugin and is provided in the task. Here we define what kind of config info this plugin can take and/or needs.
 */
func (ActiveDirectoryCollector) GetConfigPolicy() (plugin.ConfigPolicy, error) {
	policy := plugin.NewConfigPolicy()

	// This rule is simply for unit testing
	policy.AddNewFloatRule([]string{"random", "float"},
		"testfloat",
		false,
		plugin.SetMaxFloat(1000.0),
		plugin.SetMinFloat(0.0))

	// For now, assuming that active directory has no configs. May need to add some if permissions becomes an issue.
	return *policy, nil
}
