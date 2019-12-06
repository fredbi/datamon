**Version: dev**

## datamon version

prints the version of datamon

### Synopsis

Prints the version of datamon. It includes the following components:
	* Semver (output of git describe --tags)
	* Build Date (date at which the binary was built)
	* Git Commit (the git commit hash this binary was built from
	* Git State (when dirty there were uncommitted changes during the build)


```
datamon version [flags]
```

### Options

```
  -h, --help   help for version
```

### Options inherited from parent commands

```
      --upgrade   Upgrades the current version then carries on with the specified command
```

### SEE ALSO

* [datamon](datamon.md)	 - Datamon helps building ML pipelines
