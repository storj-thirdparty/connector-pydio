# Config Files

> There are two config files that contain Storj network and Pydio connection information. The tool is designed so you can specify a config file as part of your tooling/workflow.

## `pydio_property.json`

Inside the `./config` directory there is a `pydio_property.json` file, with following information about your Pydio instance:

* `url` - URL of Pydio Instance
* `user` - Username of Pydio Instance account
* `password` - Password of Pydio Instance account
* `bucket` - `io` (Fixed S3 Bucket Name of Pydio)
* `region` - `us-east-1` (Preferable Region Name of Pydio Instance)
* `apiKey` - `gateway` (Fixed S3 API Key of Pydio)
* `apiSecret` - `gatewaysecret` (Fixed S3 Secret Access Key of Pydio)

## `storj_config.json`

Inside the `./config` directory a `storj_config.json` file, with Storj network configuration information in JSON format:

* `apiKey` - API Key created in Storj Satellite GUI(mandatory)
* `satelliteURL` - Storj Satellite URL(mandatory)
* `encryptionPassphrase` - Storj Encryption Passphrase(mandatory)
* `bucketName` - Name of the bucket to upload data into(mandatory)
* `uploadPath` - Path on Storj Bucket to store data (optional) or "" or "/" (mandatory)
* `serializedAccess` - Serialized access shared while uploading data used to access bucket without API Key (mandatory while using *accesskey* flag)
* `allowDownload` - Set *true* to create serialized access with restricted download (mandatory while using *share* flag)
* `allowUpload` - Set *true* to create serialized access with restricted upload (mandatory while using *share* flag)
* `allowList` - Set *true* to create serialized access with restricted list access
* `allowDelete` - Set *true* to create serialized access with restricted delete
* `notBefore` - Set time that is always before *notAfter*
* `notAfter` - Set time that is always after *notBefore*