# snap-plugin-collector-active-directory
  
What needs to be done:  
5. Try to run with Snap itself with task and get data   
5b. Throw in concurrency if not done already  
6. Make sure README is done and code is commented  
7. Update glide file   
7. Make a pull request to intelsdi org and have Taylor and Joel code review   



# Snap collector plugin - Active Directory
This plugin collects metrics from Windows Active Directory services, including DRA (Directory and Resource Administrator), Kerberos, and LDAP data.

It's used in the [Snap framework](http://github.com:intelsdi-x/snap).

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [Known Issues](#known-issues)
6. [License](#license-and-authors)
7. [Acknowledgements](#acknowledgements)

## Getting Started
### System Requirements 
* [golang 1.7+](https://golang.org/dl/) (needed only for building as code is written in Go)

### Operating systems
All OSs currently supported by this plugin:
* Currently tested on Windows Server 2016

### Installation
#### Download perfmon plugin binary:
You can get the pre-built binaries under the plugin's [release](https://github.com/Snap-for-Windows/snap-plugin-collector-active-directory/releases) page.  For Snap, check [here](https://github.com/intelsdi-x/snap/releases).


#### To build the plugin binary:
Need to create a build script for this plugin still.

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)

## Documentation
There are a number of other resources you can review to learn to use this plugin:

* [Snap Active Direcotry unit tests](https://github.com/Snap-for-Windows/snap-plugin-collector-active-directory/activedirectory/activedirectory_test.go)
* [Snap Active Directory examples](#examples)
* [NTDS object counters](https://technet.microsoft.com/en-us/library/cc961942.aspx)
* To list out all counters available with NTDS object, open up a powershell prompt and use the `(Get-Counter -ListSet NTDS).Paths` command.

### Collected Metrics
Currently, this plugin has the ability to gather the following metrics:

Namespace | Description (optional)
----------|-----------------------
/intel/activedirectory/dra_inbound_bytes | Shows the total number of bytes replicated in. It is the sum of the number of uncompressed bytes (never compressed) and the number of compressed bytes (after compression)
/intel/activedirectory/dra_inbound_objects | Shows the number of objects received from neighbors through inbound replication. A neighbor is a domain controller from which the local domain controller replicates locally
/intel/activedirectory/dra_inbound_values | Shows the total number of object property values received from inbound replication partners. Each inbound object has one or more properties, and each property has zero or more values. Zero values indicates property removal.
/intel/activedirectory/dra_outbound_bytes | Shows the total number of bytes replicated out. It is the sum of the number of uncompressed bytes (never compressed) and the number of compressed bytes (after compression)
/intel/activedirectory/dra_outbound_objects | Shows the number of objects replicated out
/intel/activedirectory/dra_outbound_values | Shows the number of object property values sent to outbound replication partners
/intel/activedirectory/dra_pending_replication_syncs | rate at which pages are read from or written to disk 
/intel/activedirectory/ds_client_binds | percentage of paging file (for virtual memory) being used
/intel/activedirectory/ds_directory_reads | seconds since server last rebooted
/intel/activedirectory/ds_directory_searches | how frequently the processor has to switch from user- to kernel-mode per second
/intel/activedirectory/ds_directory_writes | percentage of elapsed time that all of process threads used the processor (in this case, all the processors) to execute instructions
/intel/activedirectory/kdc_as_requests | percentage of the total usable space on the selected logical disk that is free (in this case, the total of all logical disks on machine)
/intel/activedirectory/kdc_tgs_requests | percentage of the total usable space on the selected logical disk that is free (in this case, the total of all logical disks on machine)
/intel/activedirectory/kerberos_authentications | percentage of the total usable space on the selected logical disk that is free (in this case, the total of all logical disks on machine)
/intel/activedirectory/ldap_bind_time | percentage of the total usable space on the selected logical disk that is free (in this case, the total of all logical disks on machine)
/intel/activedirectory/ldap_client_session | percentage of the total usable space on the selected logical disk that is free (in this case, the total of all logical disks on machine)
/intel/activedirectory/ldap_searches | percentage of the total usable space on the selected logical disk that is free (in this case, the total of all logical disks on machine)
/intel/activedirectory/ldap_successful_binds | percentage of the total usable space on the selected logical disk that is free (in this case, the total of all logical disks on machine)
/intel/activedirectory/ldap_writes | percentage of the total usable space on the selected logical disk that is free (in this case, the total of all logical disks on machine)

### Examples
This is an example running perfmon and writing data to a file. It is assumed that you are using the latest Snap binary and plugins.
It is also assumed that the user has a folder within the C: drive called "SnapLogs".

The example is run from a directory which includes snaptel, snapteld, along with the plugins and task file.

In one terminal window, open the Snap daemon (in this case with logging set to 1 and trust disabled):
```
$ snapteld -l 1 -t 0
```

In another terminal window:
Load perfmon plugin
```
$ snaptel plugin load snap-plugin-collector-perfmon
Plugin loaded
Name: perfmon-collector
Version: 1
Type: collector
Signed: false
Loaded Time: Mon, 20 Feb 2017 11:17:17 MST
```
See available metrics for your system
```
$ snaptel metric list
```

Create a task manifest file (e.g. `task-perfmon.json`):    
```json
{ 
    "version": 1,
    "schedule": {
        "type": "simple",
        "interval": "30s"
    },
    "max-failures": 10,
    "workflow": {
        "collect": {
            "metrics": {
                "/intel/perfmon/memory_committed_bytes": {},
                "/intel/perfmon/memory_available_mbytes": {},
                "/intel/perfmon/processor_time": {}
            },
            "process": [
                {
                    "plugin_name": "passthru-grpc",
                    "process": null,
                    "publish": [
                        {
                            "plugin_name": "mock-file-grpc",
                            "config": {
                                "file": "C:\\SnapLogs\\perfmon_published_revised.log"
                            }
                        }
                    ]
                }
            ]
        }
    }
}
```
Load passthru plugin for processing:
```
$ snaptel plugin load snap-plugin-processor-passthru-grpc
Plugin loaded
Name: passthru-grpc
Version: 1
Type: processor
Signed: false
Loaded Time: Mon, 20 Feb 2017 11:16:37 MST
```

Load file plugin for publishing:
```
$ snaptel plugin load snap-plugin-publisher-mock-file-grpc
Plugin loaded
Name: mock-file-grpc
Version: 1
Type: publisher
Signed: false
Loaded Time: Mon, 20 Feb 2017 11:16:58 MST
```

Create task:
```
$ snaptel task create -t task-perfmon.json
Using task manifest to create task
Task created
ID: 4a156b0f-582f-4a13-8d67-120a2ba72e1d
Name: Task-4a156b0f-582f-4a13-8d67-120a2ba72e1d
State: Running
```

See file output (this is just part of the file):
```
2017-02-20 12:05:57.6877987 -0700 MST|[{intel  } {perfmon  } {processor_time  }]|0.658707496758626|tags[plugin_running_on:DESKTOP-0GETRGO]
2017-02-20 12:05:57.6877987 -0700 MST|[{intel  } {perfmon  } {memory_available_mbytes  }]|10381|tags[plugin_running_on:DESKTOP-0GETRGO]
2017-02-20 12:05:57.6877987 -0700 MST|[{intel  } {perfmon  } {memory_committed_bytes  }]|38.7844268555717|tags[plugin_running_on:DESKTOP-0GETRGO]
2017-02-20 12:06:27.6945658 -0700 MST|[{intel  } {perfmon  } {processor_time  }]|1.25933460733828|tags[plugin_running_on:DESKTOP-0GETRGO]
2017-02-20 12:06:27.6945658 -0700 MST|[{intel  } {perfmon  } {memory_available_mbytes  }]|10382|tags[plugin_running_on:DESKTOP-0GETRGO]
2017-02-20 12:06:27.6945658 -0700 MST|[{intel  } {perfmon  } {memory_committed_bytes  }]|38.8447365115501|tags[plugin_running_on:DESKTOP-0GETRGO]
2017-02-20 12:06:57.6960669 -0700 MST|[{intel  } {perfmon  } {processor_time  }]|1.84033283218797|tags[plugin_running_on:DESKTOP-0GETRGO]
2017-02-20 12:06:57.6960669 -0700 MST|[{intel  } {perfmon  } {memory_available_mbytes  }]|10381|tags[plugin_running_on:DESKTOP-0GETRGO]
2017-02-20 12:06:57.6960669 -0700 MST|[{intel  } {perfmon  } {memory_committed_bytes  }]|38.8223214631021|tags[plugin_running_on:DESKTOP-0GETRGO]
```

Stop task:
```
$ snaptel task stop 4a156b0f-582f-4a13-8d67-120a2ba72e1d
Task stopped:
ID: 4a156b0f-582f-4a13-8d67-120a2ba72e1d
```

### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-perfmon/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-perfmon/pulls).

## Community Support
This repository is one of **many** plugins in **Snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[Snap](http://github.com:intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
* Author: [@mathewlk](https://github.com/mathewlk/)

And **thank you!** Your contribution, through code and participation, is incredibly important to us.
