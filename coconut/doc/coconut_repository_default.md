## coconut repository default

set a git repository as default

### Synopsis

The repository default command sets a git repository
as the default repository for incoming workflow deployment requests.

```
coconut repository default [flags]
```

### Options

```
  -h, --help   help for default
```

### Options inherited from parent commands

```
      --config string            optional configuration file for coconut (default $HOME/.config/coconut/settings.yaml)
      --config_endpoint string   configuration endpoint used by AliECS core as PROTO://HOST:PORT (default "consul://127.0.0.1:8500")
      --endpoint string          AliECS core endpoint as HOST:PORT (default "127.0.0.1:47102")
  -v, --verbose                  show verbose output for debug purposes
```

### SEE ALSO

* [coconut repository](coconut_repository.md)	 - manage git repositories for task and workflow configuration

###### Auto generated by spf13/cobra on 12-Sep-2019