# token-up
Automatically update the id-token in your kubeconfig from a new kubeconfig downloaded through Tectonic.

# Test
Be sure to run a test first, as it's parsing the YAML, updating it, and then re-marshalling it to update the config file. The `Config` struct may need to be updated to include optional fields I haven't covered yet.

```sh
go run *.go -test
```

# Build & Use
```sh
go build -o $GOPATH/bin/token-up
```

```sh
$GOPATH/bin/token-up
```

# Assumptions
- Your kubeconfig file is located at $HOME/.kube/config
- New config files are located at $HOME/Downloads