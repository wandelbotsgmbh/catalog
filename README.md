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

### Structure

* `catalog` is the folder, where all available entries are added
    * `catalog/your-app-name/` is the folder for your app
    * `catalog/your-app-name/manifest.yaml` is the manifest for the entry, which follows [schema.json](schema.json)
    * `catalog/your-app-name/icon.png` icon which will be shown in graphical app managers

The same structure is mirrored in `catalog/_preview` for the preview store.

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

## Updating Version And Image For Existing Entries

If you want to have your CI to publish new versions to the existing catalog entries, 
one could use [yq](https://github.com/mikefarah/yq)

```bash
# update the image
$ yq -i -I4 '.app.containerImage.image = "someimage:123"' catalog/doom/manifest.yaml

# update the version 
$ yq -i -I4 '.version = "1.2.3"' catalog/doom/manifest.yaml
```

## Updating Catalog Entries in CI/CD

The following shows an example on how to trigger update jobs via curl:
```bash
IMAGE="registry/imagename:1.2.3"
VERSION="1.2.3"
PREVIEW="0" # Set to 1 to publish in preview store
ENTRY="entry-name"
# needs token with "Contents" repository permissions (write)
TOKEN="github_pat_..."

curl -L \
  -X POST \
  -H "Accept: application/vnd.github.v3+json" \
  -H "Authorization: Bearer ${TOKEN}" \
  https://api.github.com/repos/wandelbotsgmbh/catalog/dispatches \
  -d "{\"event_type\": \"update-entry\", 
       \"client_payload\": {
         \"image\":\"${IMAGE}\",
         \"version\":\"${VERSION}\",
         \"preview\":\"${PREVIEW}\",
         \"entry\":\"${ENTRY}\"
      }
    }"
```
