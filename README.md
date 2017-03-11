# Snap collector plugin - Active Directory
This plugin collects metrics from Windows Active Directory services, including DRA (Directory and Resource Administrator), Kerberos, and LDAP data.

It's used in the [Snap framework](http://github.com/intelsdi-x/snap).

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
* [powershell 3.0+](https://www.microsoft.com/en-us/download/details.aspx?id=50395): In order to use the Get-Counter powershell cmdlet which this plugin requires, powershell v.3.0+ must be installed. Powershell 2.0 does not contain this cmdlet. WMF (Windows Management Framework 5.0) is recommended as it contains the most recent version of powershell, but WMF 3.0+ is acceptable as well. If the plugin does not run, this is most likely the cause.
* [golang 1.7+](https://golang.org/dl/): Needed only for building as code is written in Go
* [glide 0.12.3+](http://glide.sh/): Required for developers in order to install correct package dependency versions.
 
### Operating systems
All OSs currently supported by this plugin:
* Currently tested on Windows Server 2016

### Installation
#### Download active-directory plugin binary:
You can get the pre-built binaries under the plugin's [release](https://github.com/intelsdi-x/snap-plugin-collector-active-directory/releases) page.  For Snap, check [here](https://github.com/intelsdi-x/snap/releases).


#### To build the plugin binary:
Build script for this plugin pending.  
For now, build manually:  
1. Download the plugin with `go get github.com/intelsdi-x/snap-plugin-collector-active-directory`
2. Navigate to the snap-plugin-collector-active-directory folder in your Go-Workspace
3. Use Glide to install correct dependency versions with `glide install`
4. Build the snap-plugin-collector-active-directory executable with `go install`
5. The plugin executable should now be located at $GOPATH\bin

### Configuration and Usage
* Set up the [Snap framework](https://github.com/intelsdi-x/snap/blob/master/README.md#getting-started)
* Due to current overhead issues with powershell, it is recommended that you use a task interval of 30 seconds or higher to ensure minimum failures when integrating with Snap. It is also recommended that you do not try to gather more than 16 metrics at a time, as this will lead to a substantial increase in failures. These are only beta limitations and will be addressed in future releases.

## Documentation
There are a number of other resources you can review to learn to use this plugin:

* [Snap Active Directory unit tests](https://github.com/intelsdi-x/snap-plugin-collector-active-directory/activedirectory/activedirectory_test.go)
* [Snap Active Directory examples](#examples)
* [NTDS object counters](https://technet.microsoft.com/en-us/library/cc961942.aspx)
* To list out all counters available with NTDS object, open up a powershell prompt and use the `(Get-Counter -ListSet NTDS).Paths` command.

### Collected Metrics
Currently, this plugin has the ability to gather the following metrics:

Namespace | Description (optional)
----------|-----------------------
/intel/activedirectory/dra/inbound_bytes | total number of bytes (per second) received through replication; the sum of the number of bytes of uncompressed data (never compressed) and compressed data (after compression)
/intel/activedirectory/dra/inbound_objects | number of objects received (per second) through inbound replication from replication partners
/intel/activedirectory/dra/inbound_values | total number of object properties received(per second) from replication partners; each inbound object has one or more properties, and each propertiy has zero or more values; a value of zero indicates that the property is to be removed
/intel/activedirectory/dra/outbound_bytes | total number of bytes sent per second; sum of number of bytes of uncompressed data (never compressed) and compressed data (after compression)
/intel/activedirectory/dra/outbound_objects | number of objects sent (per second) through outbound replication to replication partners
/intel/activedirectory/dra/outbound_values | total number of values of object properties sent (per second) to replication partners
/intel/activedirectory/dra/pending_replication_syncs | number of directory synchronizations that are queued for this server that are not yet processed; helps in determining replication backlog - larger the number, larger the backlog
/intel/activedirectory/ds/client_binds | number of ntdsapi.dll binds per second serviced by this DC
/intel/activedirectory/ds/directory_reads | number of directory reads per second
/intel/activedirectory/ds/directory_searches | number of directory searches per second
/intel/activedirectory/ds/directory_writes | number of directory writes per second
/intel/activedirectory/kdc/as_requests | number of Authentication Server (AS) requests serviced by the Kerberos Key Distribution Center (KDC) per second; AS requests are used by clients to obtain a ticket-granting ticket
/intel/activedirectory/kdc/tgs_requests | number of Ticket Granting Server (TGS) requests serviced by the KDC per second; TGS requests are used by the client to obtain a ticket to a resource
/intel/activedirectory/kerberos/authentications | number of times per second that clients use a client ticket to this domain controller to authenticate to this domain controller
/intel/activedirectory/ldap/bind_time | time (in milliseconds) required for the completion of the last successful LDAP binding
/intel/activedirectory/ldap/client_session | number of sessions of connected LDAP clients
/intel/activedirectory/ldap/searches | number of search operations per second performed by LDAP clients
/intel/activedirectory/ldap/successful_binds | number of LDAP bindings (per second) that occurred successfully
/intel/activedirectory/ldap/writes | rate at which LDAP clients perform write operations

### Examples
This is an example running active-directory and writing data to a file. It is assumed that you are using the latest Snap binary and plugins.
It is also assumed that the user has a folder within the C: drive called "SnapLogs".

The example is run from a directory which includes `snaptel`, `snapteld`, along with the plugins and task file.

In one terminal window, open the Snap daemon (in this case with logging set to 1 and trust disabled):
```
$ snapteld -l 1 -t 0
```

In another terminal window:
Load active-directory plugin:
```
$ snaptel plugin load snap-plugin-collector-active-directory
Plugin loaded
Name: activedirectory-collector
Version: 1
Type: collector
Signed: false
Loaded Time: Mon, 20 Feb 2017 11:17:17 MST
```
See available metrics for your system
```
$ snaptel metric list
```

Create a task manifest file (e.g. `task-active-directory.json`):    
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
                "/intel/activedirectory/ds/directory_reads": {},
                "/intel/activedirectory/ldap/client_session": {},
                "/intel/activedirectory/kerberos/authentications": {}
            },
            "process": [
                {
                    "plugin_name": "passthru-grpc",
                    "process": null,
                    "publish": [
                        {
                            "plugin_name": "mock-file-grpc",
                            "config": {
                                "file": "C:\\SnapLogs\\activedirectory_log.log"
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
$ snaptel task create -t task-active-directory.json
Using task manifest to create task
Task created
ID: 4a156b0f-582f-4a13-8d67-120a2ba72e1d
Name: Task-4a156b0f-582f-4a13-8d67-120a2ba72e1d
State: Running
```

See file output (this is just part of the file):
```
2017-03-08 09:39:48.4358386 -0800 PST|[{intel  } {activedirectory  } {kerberos  } {authentications  }]|0|tags[plugin_running_on:WIN-7RME9THVMTT]
2017-03-08 09:39:48.4358386 -0800 PST|[{intel  } {activedirectory  } {ldap  } {client_session  }]|6|tags[plugin_running_on:WIN-7RME9THVMTT]
2017-03-08 09:39:48.4358386 -0800 PST|[{intel  } {activedirectory  } {ds  } {directory_reads  }]|0|tags[plugin_running_on:WIN-7RME9THVMTT]
2017-03-08 09:40:18.4377212 -0800 PST|[{intel  } {activedirectory  } {kerberos  } {authentications  }]|0|tags[plugin_running_on:WIN-7RME9THVMTT]
2017-03-08 09:40:18.4377212 -0800 PST|[{intel  } {activedirectory  } {ldap  } {client_session  }]|6|tags[plugin_running_on:WIN-7RME9THVMTT]
2017-03-08 09:40:18.4377212 -0800 PST|[{intel  } {activedirectory  } {ds  } {directory_reads  }]|0|tags[plugin_running_on:WIN-7RME9THVMTT]
2017-03-08 09:40:48.5216239 -0800 PST|[{intel  } {activedirectory  } {kerberos  } {authentications  }]|0|tags[plugin_running_on:WIN-7RME9THVMTT]
2017-03-08 09:40:48.5216239 -0800 PST|[{intel  } {activedirectory  } {ldap  } {client_session  }]|6|tags[plugin_running_on:WIN-7RME9THVMTT]
```

Stop task:
```
$ snaptel task stop 4a156b0f-582f-4a13-8d67-120a2ba72e1d
Task stopped:
ID: 4a156b0f-582f-4a13-8d67-120a2ba72e1d
```

### Roadmap
There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-active-directory/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-active-directory/pulls).

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
