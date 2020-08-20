# Run

> Back-up is uploaded by streaming to the Storj network.

The following flags can be used with the `store` command:

* `accesskey` - Connects to the Storj network using a serialized access key instead of an API key, satellite url and encryption passphrase.
* `share` - Generates a restricted shareable serialized access with the restrictions specified in the Storj configuration file.

Once you have built the project you can run the following:

## Get help

```
$ ./connector-pydio --help
```

## Check version

```
$ ./connector-pydio --version
```

## Create backup from Pydio and upload them to Storj

```
$ ./connector-pydio store --pydio <path_to_pydio_config_file> --storj <path_to_storj_config_file>
```

## Create backup files from Pydio and upload them to Storj bucket using Access Key

```
$ ./connector-pydio store --accesskey
```

## Create backup files from Pydio and upload them to Storj and generate a Shareable Access Key based on restrictions in `storj_config.json`

```
$ ./connector-pydio store --share
```
