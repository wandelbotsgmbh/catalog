# Catalog

Resource Catalog for [Wandelbots Nova](https://www.wandelbots.com/).
Entries can be easily installed via the [Nova CLI](https://github.com/wandelbotsgmbh/wabocli)

```bash
nova catalog install jupyter
```

## Contributing

If you have a service (no matter if it provides a simple UI or some advanced API) you can add an additional entry in the `catalog` folder.
Please make sure the secrets are available per default in the nova instance.

Any service you provide should be able to handle a `BASE_PATH` environment variable, which gets injected into your container during runtime.
This tells your service the actual URI (e.g. `BASE_PATH=cellname/appname` for `instance.wandelbots.io/cellname/appname`).

## Linting

### YAML linting

[prebuild image](https://hub.docker.com/r/cytopia/yamllint) is used for [yamllint](https://github.com/adrienverge/yamllint)

```bash
docker run --rm -it -v $(pwd):/data cytopia/yamllint -d .yamllint .
```

### Validating Against [schema.json](https://json-schema.org/)

running with docker
```bash
docker run --rm -it -v $(pwd):/app andiikaa/jsonschema schema.json catalog/jupyter/manifest.yaml
```

running when golang is installed
```bash
go install github.com/santhosh-tekuri/jsonschema/cmd/jv@latest
find catalog/*/manifest.yaml -exec sh -c 'jv schema.json {}' \;
```


