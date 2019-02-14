# importaliaser

An HTTP server to serve `go get`-able urls pointing to repositories hosted elsewhere.

## Usage
Run `importaliaser` with a listen address and the path to a configuration file.

```bash
importaliaser -a :8080 -j ./aliases.json
```

## Configuration

Create a JSON file to store your aliases and some basic config:

```json
{
  "config" : {
    "rootUrl": "https://www.roobre.es",
  },
  "aliases" : {
    "roob.re/goxxy" : {
      "protocol": "git",
      "uri": "https://github.com/roobre/goxxy.git"
    },
    "roob.re/omemo-wget" : {
      "protocol": "git",
      "uri": "https://github.com/roobre/omemo-wget.git"
    },
    "roob.re/importaliaser" : {
      "protocol": "git",
      "uri": "https://github.com/roobre/importaliaser.git"
    }
  }
}
``` 

* **`config`**:
    * **`rootUrl`**: HTTP requests without a path will return a `301 Found` status and a `Location` header pointing to this URL, if set. Otherwise they will just return `404 Not Found`.

* **`aliases`**: Keys define package names:
    * **`protocol`**: The protocol to be used by `go get` to fetch the package.
    * **`uri`**: The URI to be used by `go get` to fetch the package.

### More configuration options

By default, `importaliaser` will return 404 for unknown aliases. However, it can be configured to support speculative aliasing. This is useful if you mainly use one place for hosting your stuff and you don't want to add all of them manually. You can do this by setting the following under the `config` section of the config file:

```json
{
    "config" : {
        "rootUrl": "https://www.roobre.es",
        "speculative": true,
        "speculativeProtocol": "git",
        "speculativeFormat": "https://github.com/roobre/%s.git"
    }
}
```

`importaliaser` will `fmt.Sprintf` your `speculativeFormat` with the *base name* (i.e. without the host part) of the package being fetched as an argument.

* **`config`**:
    * **`speculative`**: Set to true to enable speculative aliasing.
    * **`speculativeProtocol`**: The protocol to be used by `go get` to fetch the package.
    * **`speculativeFormat`**: Template to build `go get`-able URIs.

Config system is extensible but right now only a JSON one is implemented. PRs for TOML/YAML are welcome.