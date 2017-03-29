package activedirectory

import (
	"bytes"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"

	. "github.com/intelsdi-x/snap-plugin-utilities/logger"
)

/*
 * Takes in a string of metric names requested in task
 * Returns a map of metric names to their values
 */
func GetPowershellData(mts []string) map[string]float64 {
	LogDebug("Begin gathering metric data from system", "metric_count", len(mts))

	runtime.GOMAXPROCS(runtime.NumCPU())
	// Map to store all the metrics with their values to pass to activedirectory.go
	metricValues := make(map[string]float64)
	var mutex = &sync.Mutex{} // will synchronize access to to shared state (metricValues) across multiple goroutines

	// Have powershell command available for each metric name
	argValues := map[string]string{
		"dra_inbound_bytes":             "(get-counter -Counter \"\\NTDS\\DRA Inbound Bytes Total/sec\" -ErrorAction Stop).CounterSamples.CookedValue",
		"dra_inbound_objects":           "(get-counter -Counter \"\\NTDS\\DRA Inbound Objects/sec\" -ErrorAction Stop).CounterSamples.CookedValue",
		"dra_inbound_values":            "(get-counter -Counter \"\\NTDS\\DRA Inbound Values Total/sec\" -ErrorAction Stop).CounterSamples.CookedValue",
		"dra_outbound_bytes":            "(get-counter -Counter \"\\NTDS\\DRA Outbound Bytes Total/sec\" -ErrorAction Stop).CounterSamples.CookedValue",
		"dra_outbound_objects":          "(get-counter -Counter \"\\NTDS\\DRA Outbound Objects/sec\" -ErrorAction Stop).CounterSamples.CookedValue",
		"dra_outbound_values":           "(get-counter -Counter \"\\NTDS\\DRA Outbound Values Total/sec\" -ErrorAction Stop).CounterSamples.CookedValue",
		"dra_pending_replication_syncs": "(get-counter -Counter \"\\NTDS\\DRA Pending Replication Synchronizations\" -ErrorAction Stop).CounterSamples.CookedValue",
		"ds_client_binds":               "(get-counter -Counter \"\\NTDS\\DS Client Binds/sec\" -ErrorAction Stop).CounterSamples.CookedValue",
		"ds_directory_reads":            "(get-counter -Counter \"\\NTDS\\DS Directory Reads/sec\" -ErrorAction Stop).CounterSamples.CookedValue",
		"ds_directory_searches":         "(get-counter -Counter \"\\NTDS\\DS Directory Searches/sec\" -ErrorAction Stop).CounterSamples.CookedValue",
		"ds_directory_writes":           "(get-counter -Counter \"\\NTDS\\DS Directory Writes/sec\" -ErrorAction Stop).CounterSamples.CookedValue",
		"kdc_as_requests":               "(get-counter -Counter \"\\Security System-Wide Statistics\\KDC AS Requests\" -ErrorAction Stop).CounterSamples.CookedValue",
		"kdc_tgs_requests":              "(get-counter -Counter \"\\Security System-Wide Statistics\\KDC TGS Requests\" -ErrorAction Stop).CounterSamples.CookedValue",
		"kerberos_authentications":      "(get-counter -Counter \"\\Security System-Wide Statistics\\Kerberos Authentications\" -ErrorAction Stop).CounterSamples.CookedValue",
		"ldap_bind_time":                "(get-counter -Counter \"\\NTDS\\LDAP Bind Time\" -ErrorAction Stop).CounterSamples.CookedValue",
		"ldap_client_session":           "(get-counter -Counter \"\\NTDS\\LDAP Client Sessions\" -ErrorAction Stop).CounterSamples.CookedValue",
		"ldap_searches":                 "(get-counter -Counter \"\\NTDS\\LDAP Searches/sec\" -ErrorAction Stop).CounterSamples.CookedValue",
		"ldap_successful_binds":         "(get-counter -Counter \"\\NTDS\\LDAP Successful Binds/sec\" -ErrorAction Stop).CounterSamples.CookedValue",
		"ldap_writes":                   "(get-counter -Counter \"\\NTDS\\LDAP Writes/sec\" -ErrorAction Stop).CounterSamples.CookedValue"}
	cmdName := "powershell"
	var wg sync.WaitGroup

	// For each metric the user has requested, wait for responses (goroutines)
	wg.Add(len(mts))

	// Get data for each metric requested concurrently
	for _, metricName := range mts {
		go func(metricName string) {
			defer wg.Done() // defer pushes function call onto a list. Function is executed after surrounding function (goroutine) returns.
			var metricValue float64
			// Command() returns a Cmd struct to execute named program with args, which is then executed by Run() further down
			cmdArg := argValues[metricName]
			cmd := exec.Command(cmdName, cmdArg)
			// Buffer is a variable-sized buffer of bytes with Read and Write methods; needs no initialization
			var counterOut bytes.Buffer
			var stderr bytes.Buffer
			// Stdout and Stderr of exec package specify processes' standard output and error channels
			cmd.Stdout = &counterOut
			cmd.Stderr = &stderr
			// Run() starts the command and waits for it to complete; typically returns error as type *ExitError - this doesn't provide sufficient error detail, so I use Stderr property of Command object as well
			err := cmd.Run()
			// If there is an error with command execution, set metricValue to -1, and return this metric
			if err != nil {
				metricValue = -1
				return
			}
			// counterOut.String() adds a newline for some reason, so it must be removed first
			metricValue, formatErr := strconv.ParseFloat(strings.TrimSpace(counterOut.String()), 64)
			// Check to see if there was an error in parsing the value (this could happen if there are multiple values returned(doing (*) instead of (_total)), if no values are returned, or if the counter cannot be found on the system)
			if formatErr != nil {
				errorMessage := "There was an error with " + metricName
				LogError(errorMessage, "error", formatErr)
				metricValue = -1
			}

			mutex.Lock()
			metricValues[metricName] = metricValue
			mutex.Unlock()

		}(metricName)
	}

	wg.Wait()

	// Return map of requested metrics and their values
	return metricValues

}
