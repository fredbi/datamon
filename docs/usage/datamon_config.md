**Version: dev**

## datamon config

Commands to manage the config file

### Synopsis

Commands to manage datamon local CLI config file.

The local datamon configuration file contains the common set of flags that are needed for most commands and do not change across runs,
analogous to "git config ...".

You may force a specific local config file using the $DATAMON_CONFIG environment variable (must be some yaml or json file).


### Options

```
  -h, --help   help for config
```

### Options inherited from parent commands

```
      --config string             Set the config backend store to use (bucket name: do not set the scheme, e.g. 'gs://')
      --context string            Set the context for datamon (default "dev")
      --loglevel string           The logging level. Levels by increasing order of verbosity: none, error, warn, info, debug (default "info")
      --metrics                   Toggle telemetry and metrics collection
      --metrics-password string   Password to connect to the metrics collector backend. Overrides any password set in URL
      --metrics-url string        Fully qualified URL to an influxdb metrics collector, with optional user and password
      --metrics-user string       User to connect to the metrics collector backend. Overrides any user set in URL
      --upgrade                   Upgrades the current version then carries on with the specified command
```

### SEE ALSO

* [datamon](datamon.md)	 - Datamon helps build ML pipelines
* [datamon config set](datamon_config_set.md)	 - Create a local config file

