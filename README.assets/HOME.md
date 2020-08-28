## Flow Diagram

![](https://github.com/storj-thirdparty/connector-pydio/blob/master/README.assets/arch.drawio.png)


## Config Files

There are two config files that contain Storj network and Pydio connection information. The tool is designed so you can specify a config file as part of your tooling/workflow.

##### ```pydio_property.json```

Inside the ```./config``` directory there is a ```pydio_property.json``` file, with following information about your Pydio instance:

* url - URL of Pydio Instance
* user - Username of Pydio Instance account
* password - Password of Pydio Instance account
* bucket - ```io``` (Fixed S3 Bucket Name of Pydio)
* region - ```us-east-1``` (Preferable Region Name of Pydio Instance)
* apiKey - ```gateway``` (Fixed S3 API Key of Pydio)
* apiSecret - ```gatewaysecret``` (Fixed S3 Secret Access Key of Pydio)

##### ```storj_config.json```

Inside the ```./config``` directory a ```storj_config.json``` file, with Storj network configuration information in JSON format:

* apiKey - API Key created in Storj Satellite GUI
* satelliteURL - Storj Satellite URL
* encryptionPassphrase - Storj Encryption Passphrase.
* bucket - Name of the bucket to upload data into.
* uploadPath - Path on Storj Bucket to store data (optional) or "" or "/" (mandatory)
* serializedAccess - Serialized access shared while uploading data used to access bucket without API Key
* allowDownload - Set true to create serialized access with restricted download
* allowUpload - Set true to create serialized access with restricted upload
* allowList - Set true to create serialized access with restricted list access
* allowDelete - Set true to create serialized access with restricted delete
* notBefore - Set time that is always before notAfter
* notAfter - Set time that is always after notBefore

## Run

The following flags can be used with the ```store``` command:

* `accesskey` - Connects to the Storj network using a serialized access key instead of an API key, satellite url and encryption passphrase .
* `shared` - Generates a restricted shareable serialized access with the restrictions specified in the Storj configuration file.

Once you have built the project you can run the following:

##### Get help

```
$ ./connector-pydio --help
```

##### Check version

```
$ ./connector-pydio version
```

##### Create backup from Pydio and upload them to Storj

```
$ ./connector-pydio store --pydio <path_to_pydio_config_file> --storj <path_to_storj_config_file>
```

##### Create backup files from Pydio and upload them to Storj bucket using Access Key

```
$ ./connector-pydio store --accesskey
```

##### Create backup files from Pydio and upload them to Storj and generate a Shareable Access Key based on restrictions in ```storj_config.json```.

```
$ ./connector-pydio store --share
```
		
## Testing

* The project has been tested on the following operating systems:
```
	* Windows
		* Version: 10 Pro
		* Processor: Intel(R) Core(TM) i3-5005U CPU @ 2.00GHz 2.00GHz
		* Pydio version: Pydio cells

	* macOS Catalina
		* Version: 10.15.4
		* Processor: 2.5 GHz Dual-Core Intel Core i5
		* Pydio version: Pydio cells

	* ubuntu
		* Version: 16.04 LTS
		* Processor: AMD A6-7310 APU with AMD Radeon R4 Graphics × 4
		* Pydio version: Pydio cells
```		
