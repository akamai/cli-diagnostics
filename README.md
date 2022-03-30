#Edge Diagnostics CLI (Beta)

- [Get started with the Edge Diagnostics CLI](#get-started-with-the-edge-diagnostics-cli)
  - [Install Edge Diagnostics CLI](#install-edge-diagnostics-cli)
    - [Install using the Akamai CLI (recommended)](#install-using-the-akamai-cli-recommended)
    - [Install from binaries](#install-from-binaries)
    - [Compile from Source](#compile-from-source)
  - [Work with multiple accounts](#work-with-multiple-accounts)
  - [API credentials](#api-credentials)
    - [Using the .edgerc file](#using-the-edgerc-file)
    - [Setting environmental variables](#setting-environmental-variables)
  - [JSON support](#json-support)
- [Operation categories](#operation-categories)
- [Available operations and commands](#available-operations-and-commands)
  - [Check available commands and options](#check-available-commands-and-options)
  - [List edge server locations](#list-edge-server-locations)
  - [List GTM hostnames](#list-gtm-hostnames)
  - [List IP Acceleration hostnames](#list-ip-acceleration-hostnames)
  - [Generate a diagnostic link](#generate-a-diagnostic-link)
  - [List diagnostic links](#list-diagnostic-links)
  - [Get collected diagnostic data](#get-collected-diagnostic-data)
  - [Translate an error string](#translate-an-error-string)
  - [Verify an IP](#verify-an-ip)
  - [Locate an IP network](#locate-an-ip-network)
  - [Verify an IP and locate its network](#verify-an-ip-and-locate-its-network)
  - [Translate an Akamaized URL](#translate-an-akamaized-url)
  - [Get error statistics](#get-error-statistics)
  - [Get logs](#get-logs)
  - [Request content with cURL](#request-content-with-curl)
  - [Get domain details with dig](#get-domain-details-with-dig)
  - [Test network connectivity with MTR](#test-network-connectivity-with-mtr)
  - [Run the URL health check](#run-the-url-health-check)
  - [Run the Connectivity problems scenario](#run-the-connectivity-problems-scenario)
  - [Run the Content problems scenario](#run-the-content-problems-scenario)
- [Available flags](#available-flags)
  - [edgerc](#edgerc)
  - [section](#section)
  - [help](#help)
  - [version](#version)
  - [json](#json)
- [Exit codes](#exit-codes)
- [Windows 10 2018 version](#windows-10-2018-version)
- [Notice](#notice)

# Get started with the Edge Diagnostics CLI

Use the Edge Diagnostics CLI to identify, analyze, and troubleshoot common content delivery network issues that your users may encounter. 

> **_NOTE:_** This CLI is based on the Edge Diagnostics API. The previous version, the Diagnostic Tools CLI, is now deprecated and will be discontinued on September 30, 2022. We encourage you to migrate to this new Edge Diagnostics CLI to avoid any inconvenience after the discontinuation of the Diagnostic Tools CLI.

## Install Edge Diagnostics CLI
There are three ways in which you can install the CLI. 

### Install using the Akamai CLI (recommended)

To install this CLI, you need the [Akamai CLI](https://github.com/akamai/cli) package manager. Once you install the Akamai CLI, run this command:

`akamai install diagnostics` 

**Stay up to date**

To make sure you always use the latest version of the CLI, run this command:  

`akamai update diagnostics`  

### Install from binaries

Follow the instructions for your operating system.

**Linux and macOS**

Once you download the appropriate binary for your system, make it executable, and make it available in your `$PATH`. Run the following commands:

```sh
$ chmod +x ~/Downloads/akamai-diagnostics-<VERSION>-<PLATFORM>
$ mv ~/Downloads/akamai-diagnostics-<VERSION>-<PLATFORM> /usr/local/bin/akamai-diagnostics
$ akamai-diagnostics help
```

**Windows**

Once you download the appropriate binary for your system and add the binary location to the Path variable, simply execute the binary from the command line. For example:

```sh
$ akamai-diagnostics help
```
### Compile from Source

**Prerequisite:** Make sure you install Go 1.17 or later.

To compile Edge Diagnostics CLI from source:

1. Change the working directory:

    ```sh
    $ cd $GOPATH
    ```

2. Fetch the package:

    ```sh
    $ git clone https://github.com/akamai/cli-diagnostics.git
    ```

3.  Go to the package directory:

    ```sh
    $ cd cli-diagnostics
    ```
4. Compile the binary:  

  - For Linux, macOS, and other Unix-based systems, run: `go build -o akamai-diagnostics`
  - For Windows, run: `go build -o akamai-diagnostics.exe`

5. Move the `akamai-diagnostics` or `akamai-diagnostics.exe` binary so that it's available in your `$PATH`.

## Work with multiple accounts
To diagnose an issue, you may need to switch between different accounts. To do this, run the required operation with the `--account-key` flag followed by the account ID of your choice. 

For example: `akamai diagnostics --account-key 1-1TJZFB ipa-hostnames`

## API credentials
The Edge Diagnostics CLI requires your authentication credentials to execute any command. To create your credentials, see [Create a quick API client](https://techdocs.akamai.com/developer/docs/set-up-authentication-credentials#create-a-quick-api-client).
You can provide these API credentials to the CLI in two ways, either through the `.edgerc` file or by setting the environmental variables.

### Using the .edgerc file
Akamai-branded packages use an `.edgerc` file for standard EdgeGrid authentication. To set up your `.edgerc` file, see [Add credential to .edgerc file](https://techdocs.akamai.com/developer/docs/set-up-authentication-credentials#add-credential-to-edgerc-file). We recommend to change the `default` header name in the `.edgerc` file to `diagnostics`.

### Setting environmental variables
You can set these environmental variables to provide your API credentials:
- `AKAMAI_DIAGNOSTICS_HOST`. 
- `AKAMAI_DIAGNOSTICS_CLIENT_TOKEN`.
- `AKAMAI_DIAGNOSTICS_CLIENT_SECRET`.
- `AKAMAI_DIAGNOSTICS_ACCESS_TOKEN`.

Make sure that the environment variable `AKAMAI_EDGERC_SECTION` is set to `diagnostics`.

## JSON support 
Instead of creating requests directly in CLI, you can request to run a command on a specific JSON file. 

For example, instead of such MTR command:
 `akamai diagnostics mtr --source bangalore-india --destination www.akamai.com --ip-version IPv4 --port 443 --packet-type icmp`

You can run 
`akamai diagnostics mtr < mtr_data_file.json` command. Where the JSON file has all the required data for the MTR request. 

To create the JSON files, you can use the documentation for the appropriate [API operation](https://techdocs.akamai.com/edge-diagnostics/reference/edge-diagnostics-api-1). Add your values to `BODY PARAMS` fields, copy the body of your request from the CURL code sample, and save it as a JSON file. 

# Operation categories
Check these descriptions to learn more about the available categories of tools and tools they include.

- **General**. These operations fetch data required by other operations.
  - [List edge server locations](#list-edge-server-locations)
  - [List GTM hostnames](#list-gtm-hostnames)
  - [List IP Acceleration hostnames](#list-ip-acceleration-hostnames)
- **Client to Edge**. These operations let you primarily diagnose issues in the traffic between a client and edge servers. You can also use these operations for other issues.
  - [Generate a diagnostic link](#generate-a-diagnostic-link)
  - [List diagnostic links](#list-diagnostic-links)
  - [Get collected diagnostic data](#get-collected-diagnostic-data)
  - [Translate an error string](#translate-an-error-string)
- **Edge**. These operations get basic data from an edge server.
  - [Verify an IP](#verify-an-ip)
  - [Locate an IP network](#locate-an-ip-network)
  - [Verify an IP and locate its network](#verify-an-ip-and-locate-its-network)
  - [Translate an Akamaized URL](#translate-an-akamaized-url)
  - [Get error statistics](#get-error-statistics)
  - [Get logs](#get-logs)
- **Edge to Origin**. These operations let you primarily diagnose issues in the traffic between an edge and origin server. You can also use these operations for other issues.
  - [Request content with cURL](#request-content-with-curl)
  - [Get domain details with dig](#get-domain-details-with-dig)
  - [Test network connectivity with MTR](#test-network-connectivity-with-mtr)
- **Problem scenarios**. If you're not sure where the problem is and which operations could help you identify the issue, use *Problem scenarios*. These are several tools stitched up. You provide the input once, and the scenario runs several tools to provide you with the data and suggestions necessary to identify the cause of the issue.
  - [Run the URL health check](#run-the-url-health-check)
  - [Run the Connectivity problems scenario](#run-the-connectivity-problems-scenario)
  - [Run the Content problems scenario](#run-the-content-problems-scenario)

# Available operations and commands

## Check available commands and options

The `help` command returns a list of available commands and flags with their descriptions.

**Command**: `help [COMMAND]` where `COMMAND` is the command you need help with. 

**Examples**: 
- `akamai diagnostics help`
- `akamai diagnostics help verify-ip`

## List edge server locations
The `edge-locations` command lists active edge server locations you can use to run [Request content with cURL](#request-content-with-curl), [Get domain details with dig](#get-domain-details-with-dig), [Test network connectivity with MTR](#test-network-connectivity-with-mtr),[Run the URL health check](#run-the-url-health-check), [Run the Connectivity problems scenario](#run-the-connectivity-problems-scenario) and [Run the Content problems scenario](#run-the-content-problems-scenario) operations.

**Command**: `edge-locations [--search REGION]` where the `--search` flag filters edge server locations for a particular `REGION`. This flag is optional.

**Examples**: 
- `akamai diagnostics edge-locations`
- `akamai diagnostics edge-locations --search india`

**Expected output**: The response includes locations of active edge servers you can run [curl](#request-content-with-curl), [dig](#get-domain-details-with-dig), [mtr](#test-network-connectivity-with-mtr), [url-health-check](#run-the-url-health-check), 
[connectivity-problem](#run-the-connectivity-problems-scenario) and [content-problem](#run-the-content-problems-scenario) commands. For more details, you can check the [API response](https://techdocs.akamai.com/edge-diagnostics/reference/get-edge-locations) description.


## List GTM hostnames
The `gtm-hostnames` command lists GTM properties you have access to. You can also request their test and target IPs. The returned values can be used to run [Get domain details with dig](#get-domain-details-with-dig) and [Test network connectivity with MTR](#test-network-connectivity-with-mtr) operations for a GTM hostname.

**Command**: `gtm-hostnames [--test-target-ip GTM_HOSTNAME]` where the `--test-target-ip` flag requests the test and target IPs for a `GTM_HOSTNAME`. This flag is optional.

**Examples**: 
- `akamai diagnostics gtm-hostnames`
- `akamai diagnostics gtm-hostnames --test-target-ip www-origin.20000puzzles.akadns.net`

**Expected output**: The response includes all GTM properties you have access to and can run `dig` and `mtr` commands for. For more 
details, you can check the [API response](https://techdocs.akamai.com/edge-diagnostics/reference/get-gtm-properties) description.

## List IP Acceleration hostnames
The `ipa-hostnames` command lists IP Acceleration (IPA) hostnames you have access to and can [generate a diagnostic link](#generate-a-diagnostic-link) to collect diagnostic data for.

**Command**: `ipa-hostnames`

**Examples**: `akamai diagnostics ipa-hostnames`

**Expected output**: The response includes all IPA hostnames you can gather diagnostic data for.

## Generate a diagnostic link
The `user-diagnostics create` command generates a diagnostic link for you to share with users of a particular URL or IP Acceleration (IPA) hostname experiencing 
similar issues. Each link is valid for 7 days or 50 submissions. 
After a user clicks the link, Edge Diagnostics gathers necessary diagnostic data and asks the user to submit them. Once submitted, you 
can check the collected data with the [Get collected diagnostic data](#get-collected-diagnostic-data) operation.
To get the list of IPA hostnames you can create the diagnostic link for, run the [List IP Acceleration hostnames](#list-ip-acceleration-hostnames) operation.

**Command**:
- To generate a diagnostic link for a specific URL: `user-diagnostics create --url URL [--notes "NOTE"]` 
- To generate a diagnostic link for a specific IPA hostname: `user-diagnostics create --ipa-hostname IPA_HOSTNAME [--notes "NOTE"]`

Where:
- `URL` is the fully qualified URL experiencing issues. It needs to contain protocol, hostname, path, and string parameters (if applicable). You need to provide one of these flags: either `--url` or `--ipa-hostname`.
- `IPA_HOSTNAME` is the IPA hostname experiencing issues. To get the list of IPA hostnames you have access to, run the [List IP Acceleration hostnames](#list-ip-acceleration-hostnames) operation. You need to provide one of these flags: either `--ipa-hostname` or `--url`.
- the `--notes` flag adds the `NOTE` about the link to be generated or issues users of the URL or IPA are experiencing. Notes can have up to 400 characters. This flag is optional. 

**Examples**: 
- `akamai diagnostics user-diagnostics create --url https://www.akamai.com`
- `akamai diagnostics user-diagnostics create --url https://www.akamai.com --notes "Diagnostic data of users of www.akamai.com".`
- `akamai diagnostics user-diagnostics create --ipa-hostname www.akamai.com --notes "Tokyo olympics"`

**Expected output**: The response includes the link you can send to users experiencing issues 
with the hostname. When a user clicks the link, Edge Diagnostics gathers necessary data 
and asks the user to submit them. The response includes also `link_id` you can use to [get collected diagnostic data](#get-collected-diagnostic-data).

## List diagnostic links
The `user-diagnostics list` command lists generated diagnostic links used to collect user diagnostic data for hostnames experiencing issues. It also lists the number of data submitted with each link.

**Command**: `user-diagnostics list [--url URL] [--user USER] [--active]`, where:

- the `url` flag filters the list for the `URL` of either a hostname or an IP Acceleration hostname with issues. 
- the `user` flag filters the list for the `USER` who created the diagnostic link. 
- the `active` flag filters the list for only active links. 
Active links can be further shared with users to collect more diagnostic data. Each link is valid for 7 days or 50 submissions.

All three filters are optional and can be used together.

**Examples**: 
- `akamai diagnostics user-diagnostics list`
- `akamai diagnostics user-diagnostics list --active`
- `akamai diagnostics user-diagnostics list --url https://www.akamai.com/`
- `akamai diagnostics user-diagnostics list --user jsmith`
- `akamai diagnostics user-diagnostics list --url https://akamai.com/ --user jsmith --active`

**Expected output**: The response is the list with an overview of created diagnostic links. 
To get the details of a specific link and collected data, run the [Get collected diagnostic data](#get-collected-diagnostic-data) operation using the returned `link_id`.


## Get collected diagnostic data
The `user-diagnostics get` command returns diagnostic data submitted by users using a diagnostic link with a specific ID.

**Command**: `user-diagnostics get LINK_ID [--mtr] [--dig] [--curl]`, where:

- `LINK_ID` is the identifier for the generated link. To get the identifier, run [List diagnostic links](#list-diagnostic-links) or create it with the [Generate a diagnostic link](#generate-a-diagnostic-link) operation.
- the `--mtr` flag includes MTR data in the response.
- the `--dig` flag includes `dig` data in the response.
- the `--curl` flag includes cURL data in the response.

The `--mtr`, `--dig`, and `curl` flags are optional and can be used together.

**Examples**: 
- `akamai diagnostics user-diagnostics get ab123c`
- `akamai diagnostics user-diagnostics get ab123c --mtr --curl`
- `akamai diagnostics user-diagnostics get ab123c --dig --mtr`

**Expected output**: The response includes a table with the collected data. For more details, you can check the [API response](https://techdocs.akamai.com/edge-diagnostics/reference/get-user-diagnostic-data-group-records) description. 

## Translate an error string
The `translate-error-string` command fetches summary and logs for an error with a specific reference code. You can also translate any Global Request Number (GRN) generated by Property Manager's Global Request Number behavior. Data is available from either the last 6 or 24 hours depending on the server and traffic conditions.

All error reference codes are valid for 48 hours since their occurrence.

**Command**: `translate-error-string ERROR_STRING` where `ERROR_STRING` is the alphanumeric part of the error reference code you want to 
get the data for.

**Examples**: `akamai diagnostics translate-error-string 9.6f64d440.1318965461.2f2b078`

**Expected outcome**: The response includes the error translation in the JSON format. For more details, you can check the [API response](https://techdocs.akamai.com/edge-diagnostics/reference/error-translator#post-error-translator) description.

## Verify an IP

The `verify-ip` command checks whether a specific IP address belongs to an edge server.

**Command**: `verify-ip IP_ADDRESS ...` where `IP_ADDRESS` is the IP address you want to verify. You can enter multiple data.

**Examples**: `akamai diagnostics verify-ip 123.123.123.123 2001:db8::2:1`

**Expected output**: The response notes whether the IP address is for an edge server. 


## Locate an IP network
The `locate-ip` command provides network geolocation and details for an edge server IP address.

**Command**: `locate-ip IP_ADDRESS ...` where `IP_ADDRESS` is the IP address you want to get the data for. You can enter multiple data.

**Examples**: `akamai diagnostics locate-ip 123.123.123.123 2001:db8::2:1`

**Expected output**: The response includes network geolocation data for a network of the IP address. For more details, you can check the [API response](https://techdocs.akamai.com/edge-diagnostics/reference/post-locate-ip) description.


## Verify an IP and locate its network
The `verify-locate-ip` command verifies if an IP address belongs to an edge server and gets geolocation data for its network. 

**Command**: `verify-locate-ip IP_ADDRESS` where `IP_ADDRESS` is the IP address you want to get the data for.

**Examples**: `akamai diagnostics verify-locate-ip 123.123.123.123`

**Expected output**: The response includes the IP's 
geographic and network location data. For more details, you can check the [API response](https://techdocs.akamai.com/edge-diagnostics/reference/post-verify-locate-ip) description.


## Translate an Akamaized URL
The `translate-url` command provides basic information about an Akamaized URL, such as typecode, cache key hostname, CP code, serial number, and TTL. A URL 
becomes an Akamaized URL (ARL) once it's on an edge server.

**Command**: `translate-url URL`, where `URL` is the fully qualified, Akamaized URL you want to get the details of.

**Examples**: `akamai diagnostics translate-url http://www.akamai.com`

**Expected outcome**: The response includes details about the requested URL. For more details, you can check the [API response](https://techdocs.akamai.com/edge-diagnostics/reference/translated-url#post-translated-url) description.


## Get error statistics
The `estats` command provides statistics of errors happening for a URL or a content provider (CP) code's traffic. The tool also returns logs for 
the edge server response to a client and the edge server forward request to an origin server, and error details. Error data returned is based on a nine-second traffic sample from the last two minutes.

**Command**: 
- To run `estats` for a specific URL: `estats --url URL [--logs] [--enhanced-tls | --standard-tls] [--edge-errors] [--origin-errors]`
- To run `estats` for a specific CP code: `estats --cp-code CP_CODE [--logs] [--enhanced-tls | --standard-tls] [--edge-errors] [--origin-errors]`

Where:
- `URL` is the fully qualified URL you want to get the data for. You need to provide one of these flags: either `--url` or `--cp-code`.
- `CP_CODE` is the CP code you want to get the data for. You need to provide one of these flags: either `--cp-code` or `--url`.
- the `--logs` flag includes the `grep` logs in the response. 
- the `--enhanced-tls` or `--standard-tls` flags specify the delivery type of the resource you want to get the data for, Enhanced TLS hostname or Standard TLS 
hostname respectively. Without this filter, Edge Diagnostics checks the type of delivery used by your resource and returns data for it. 
If your resource uses both delivery types, then Edge Diagnostics returns data for the type which got all data collected faster. If you choose the 
delivery type not used by your resource, then the results are empty. If you want to add this filter, you can run the [List edge hostnames](https://developer.akamai.com/api/core_features/edge_hostnames/v1.html#getedgehostnames) operation in
 [Edge Hostnames API](https://developer.akamai.com/api/core_features/edge_hostnames/v1.html) to confirm the delivery type used by your 
 resource. It is returned as the `securityType` value. The `--enhanced-tls` and `--standard-tls` flags are optional but mutually exclusive. You can use only one of them at a time.
- the `--edge-errors` flag filters the data for the traffic between a customer and an edge server.
- the `--origin-errors` flag filters the data for the traffic between the edge server and the origin.

The `--logs`, `--edge-errors`, and `--origin-errors` flags are optional and can be used together.

**Examples**: 
- `akamai diagnostics estats --url https://www.akamai.com`
- `akamai diagnostics estats --url https://www.akamai.com --logs --standard-tls`
- `akamai diagnostics estats --cp-code 12345 --logs --standard-tls`

**Expected outcome**: The response includes error statistics in the JSON format. For more details, you can check the [API response](https://techdocs.akamai.com/edge-diagnostics/reference/post-estats) description. 

## Get logs 
The `grep` command uses GREP to fetch log lines from an IP address within the Akamai network. You can use parameters to filter the data. Logs provide 
low-level details on how each request was handled. You can use this information to troubleshoot caching and performance issues, and to 
ensure the correct Akamai features were applied to the traffic. Data is available from either the last 6 or 24 hours depending on the server and traffic conditions.

**Command**: 
- To run `grep` for a specific CP code: `grep EDGE_IP "START_TIME" "END_TIME" --cp-code CP_CODE ... [--client-ip "CLIENT_IP" ...] [--user-agent "USER_AGENT" ...] [--http-status-code HTTP_STATUS_CODE ... | --error-status-codes] [--arl ARL ...] [-r] [-f]`
- To run `grep` for a specific hostname: `grep EDGE_IP START_TIME END_TIME --hostname HOSTNAME ... [--client-ip CLIENT_IP ...] [--user-agent USER_AGENT ...] [--http-status-code HTTP_STATUS_CODE ... | --error-status-codes] [--arl ARL ...] [-r] [-f]`
  
Where:

- `EDGE_IP` is the edge server IP address you want to get logs for. You can use the edge server IP address value from the `answerSection` array of the `dig` command response or the `ip` value from the `edgeIps` array in the collected diagnostic data.
- `START_TIME` is the ISO 8601 timestamp for a point of time in the past when the log search window starts. Data is available from either the last 6 or 24 hours depending on the server and traffic conditions.
- `END_TIME` is the ISO 8601 timestamp for a point of time in the past when the log search window ends. We recommend 10 minute periods to ensure that data are fetched quickly and you get the most relevant logs.
- `CP_CODE` is the CP code you want to get the logs for. You need to provide one of these flags: either `--cp-code` or `--hostname`. This flag accepts multiple values.
- `HOSTNAME` is the hostname you want to get the logs for. You need to provide one of these flags: either `--hostname` or `--cp-code`. This flag accepts multiple values.
- `CLIENT_IP` is the client IPs you want to filter the logs by. The `--client-ip` flag is optional and accepts multiple values.
- `USER_AGENT` is the user agent you want to filter the logs by. The `--user-agent` flag is optional and accepts multiple values.
- `HTTP_STATUS_CODE` is the HTTP status code you want to filter the logs by.  The `--http-status-code` flag is optional and accepts multiple values. 
- the `--error-status-codes` flag returns all log lines with HTTP status codes for errors (all except: 100, 101, 102, 122, 200, 201, 202, 203, 204, 205, 206, 207, 226, 300, 301, 302, 303, 304, 305, 306, 307, and 404).
- `ARL` is the Akamaized URL you want to filter the logs by. The `--arl` flag is optional and accepts multiple values.
- `-r | -f | -rf` is the type of log lines you want to filter. The possible values are: `-r` for client requests to an edge server, `-f` for forward requests from an edge server to the origin, or `-rf` for both.

**Examples**: 
- `akamai diagnostics grep 123.123.123.123  "2021-04-13T14:27:06.000Z" "2021-04-13T14:59:06.000Z" --host-header "www.akamai.com" --client-ip "123.123.123.123" --http-status-code "400, 401" -rf`
- `akamai diagnostics grep 123.123.123.123  "2021-04-13T14:27:06.000Z" "2021-04-13T14:59:06.000Z" --cp-code 12345 --client-ip "123.123.123.123" --http-status-code "400, 401" -rf`

**Expected outcome**: The response includes a standard `grep` response in the JSON format. For more details, you can check the [API response](https://techdocs.akamai.com/edge-diagnostics/reference/get-grep-request) description.

## Request content with cURL
The `curl` command requests content to provide the raw HTML for a URL, including response headers. You can run this 
operation either for a specific location or an edge IP. 

>**_NOTE:_** Currently, Request content with cURL doesn't support requests for origin hostnames. We're working on enabling these requests soon.

**Command**: 
- To run `curl` from a specific location: `curl URL [--client-location CLIENT_LOCATION] [--ip-version IPv4|IPv6] [--request-header REQUEST_HEADER...] [--run-from-site-shield-map]`
- To run `curl` from a specific edge server: `curl URL [--edge-server-ip EDGE_SERVER_IP] [--ip-version IPv4|IPv6] [--request-header REQUEST_HEADER...] [--run-from-site-shield-map]`

Where:
- `URL` is the fully qualified URL you want to get the data for. 
- `CLIENT_LOCATION` is a unique identifier for an edge server location closest to your users. To get this value, run the [List edge server locations](#list-edge-server-locations) operation first. Provide one of these flags: either `--client-location` or `--edge-server-ip`. This flag is optional.
- `EDGE_SERVER_IP` is an IP address of an edge server you want to run `curl` from. Provide one of these flags: either `--edge-server-ip` or `--client-location`. This flag is optional.
- the `--ip_version` flag specifies the IP version you want to use to run the operation. The available values are `IPV4` or `IPV6`. It's set to `IPV4` by default. This flag is optional.
- `REQUEST_HEADER` is a customized header for the `curl` request in the format `"header: value"`. This `--request-header` flag is optional and accepts multiple values.
- the `--run-from-site-shield-map` flag uses the entered location or edge server IP to find its Site Shield map and runs the tool using the map. This flag is optional. 

If you don't provide neither `--client-location` nor `--edge-server-ip`, Edge Diagnostics will run the `curl` command using a random location.

**Examples**: 
- `akamai diagnostics curl http://www.example.com --client-location bangalore-india --ip-version IPv4 --request-header "accept:text/html"`
- `akamai diagnostics curl http://www.example.com --edge-server-ip 123.123.123.123 --ip-version IPv4 --request-header "accept:text/html"`
- `akamai diagnostics curl http://www.example.com --edge-server-ip 123.123.123.123`

**Expected outcome**: The response includes the `curl` response. For more details, you can check the [API response](https://techdocs.akamai.com/edge-diagnostics/reference/post-curl) description. 


## Get domain details with dig
The  `dig` command provides DNS details for a hostname or a domain name, or a GTM hostname. The results may help you diagnose issues with domain name resolutions. 
You can run this operation either for a specific location or an edge IP.

>**_NOTE:_** Currently, Get domain details with dig doesn't support requests for origin hostnames. We're working on enabling these requests soon.

**Command**: 
- To run `dig` from a particular location: `dig --hostname HOSTNAME [--client-location CLIENT_LOCATION]  [--query-type QUERY_TYPE] [--gtm]`
- To run `dig` from a particular edge server: `dig --hostname HOSTNAME [-e EDGE_SERVER_IP]  [--query-type QUERY_TYPE] [--gtm]`

Where:

- `HOSTNAME` is either a hostname or a domain name you want to get the DNS details for. 
- `CLIENT_LOCATION` is a unique identifier for an edge server location closest to your users. To get this value, run the [List edge server locations](#list-edge-server-locations) operation first. Provide one of these flags: either `--client-location` or `--edge-server-ip`. This flag is optional.
- `EDGE_SERVER_IP` is an IP address of an edge server you want to run `dig` from. Provide one of these flags: either `--edge-server-ip` or `--client-location`. This flag is optional.
- `QUERY_TYPE` is the DNS record type you want to get. The available values are: `A` for IPv4 address record, `AAAA` for IPv6 address record, `SOA` for Start of Authority record, `CNAME` for Canonical Name record, `PTR` for Pointer record, `MX` for Mail Exchanger record, `NS` for Nameserver record, `TXT` for Text record, `SRV` for Service Location record, `CAA` for Certificate Authority Authorization record, and `ANY` for all associated records available. This flag is optional, set to `A` by default.
- the `--gtm` flag marks the `HOSTNAME` as a GTM hostname.

If you don't provide neither `--client-location` nor `--edge-server-ip`, Edge Diagnostics will run the `dig` command using a random location.

**Examples**: 
- `akamai diagnostics dig --hostname www.akamai.com --client-location bangalore-india`
- `akamai diagnostics dig --hostname www.akamai.com --client-location bangalore-india --query-type NS`
- `akamai diagnostics dig --hostname www.akamai.com --edge-server-ip 123.123.123.123 --query-type NS`
- `diagnostics dig --hostname www.akamai.com --client-location bangalore-india --query-type NS --gtm`

**Expected outcome**: The response includes a standard `dig` response. For more details, you can check the [API response](https://techdocs.akamai.com/edge-diagnostics/reference/post-dig) description. 


## Test network connectivity with MTR
The `mtr` command uses MTR to provide information about the packets loss and latency between an edge server IP address or location and a remote destination. You can run this operation also for a GTM hostname. 

>**_NOTE:_** Currently, Test network connectivity with MTR doesn't support requests for origin hostnames. We're working on enabling these requests soon.


**Command**: 
- To run MTR between a location and a hostname: `mtr --source SOURCE_LOCATION --destination HOSTNAME [--ip-version IPv4|IPv6] [--port 80|443] [--packet-type TCP|ICMP]`
- To run MTR between a location and an IP address: `mtr --source SOURCE_LOCATION --destination DESTINATION_IP [--ip-version IPv4|IPv6] [--port 80|443] [--packet-type TCP|ICMP]`
- To run MTR between an edge server IP address and a hostname: `mtr --source SOURCE_IP --destination HOSTNAME [--ip-version IPv4|IPv6] [--port 80|443] [--packet-type TCP|ICMP]`
- To run MTR between an edge server IP address and an IP: `mtr --source SOURCE_IP --destination DESTINATION_IP [--ip-version IPv4|IPv6] [--port 80|443] [--packet-type TCP|ICMP]`
- To run MTR for a GTM hostname: `mtr --source SOURCE_IP --destination DESTINATION_IP --gtm-hostname GTM_HOSTNAME [--packet-type TCP|ICMP]`
- To run MTR on a Site Shield hostname and a destination IP: `mtr --destination DESTINATION_IP --site-shield-hostname HOSTNAME [--source SOURCE_LOCATION|SOURCE_IP][--ip-version IPv4|IPv6] [--port 80|443] [--packet-type TCP|ICMP]`

Where:

- `SOURCE_LOCATION` is a unique identifier for an edge server location closest to your users. To get this value, run the [List edge server locations](#list-edge-server-locations) operation first. 
- `HOSTNAME` is the MTR target hostname.
- `SOURCE_IP` is the edge server IP address you want to run MTR from. You can use the edge server IP address value from the `answerSection` array in the [Get domain details with dig](#get-domain-details-with-dig) operation response. If you run MTR for a GTM hostname, use the Test IP value returned by the [List GTM hostnames](#list-gtm-hostnames) operation.
- `DESTINATION_IP` is the MTR destination. If you run MTR for a GTM hostname, use the Target IP value returned by the [List GTM hostnames](#list-gtm-hostnames) operation.
- the `--ip_version` flag specifies the IP version for MTR to use, either `IPV4` or `IPV6`. It's set to `IPV4` by default. Omit when running MTR for a GTM hostname.
- the `--port` flag specifies the port number for MTR to use, either `80` or `443`. By default it's set to `80`. Omit when running MTR for a GTM hostname.
- the `--packet_type` flag specifies the packet type for MTR to use, either `tcp` or `icmp`. By default it's set to `tcp`.
- `GTM_HOSTNAME` is the GTM hostname you want to run MTR for. To get the list of GTM hostnames you have access to, run the [List GTM hostnames](#list-gtm-hostnames) operation.
- the `--site-shield-hostname` flag uses the provided `HOSTNAME` to run the tool using the Site Shield map. 

**Examples**: 
- `akamai diagnostics mtr --source bangalore-india --destination www.akamai.com --ip-version IPv4 --port 443 --packet-type icmp`
- `akamai diagnostics mtr --source bangalore-india --destination 121.121.121.121 --ip-version IPv4`
- `akamai diagnostics mtr --source 123.123.123.123 --hostname www.akamai.com --ip-version IPv4 --port 443 --packet-type icmp`
- `akamai diagnostics mtr --source 123.123.123.123  --destination 121.121.121.121`
- `akamai diagnostics mtr --source 2.2.2.2 -–destination 1.1.1.1 --gtm-hostname www.akamai.com --packet-type icmp`


**Expected outcome**: The response includes a standard `mtr` response. For more details, you can check the [API response](https://techdocs.akamai.com/edge-diagnostics/reference/post-mtr) description. 

## Run the URL health check
The `url-health-check` operation simultaneously runs [Translate an Akamaized URL](#translate-an-akamaized-url), [Request content with cURL](#request-content-with-curl), [Get domain details with dig](#get-domain-details-with-dig), [Test network connectivity with MTR](#test-network-connectivity-with-mtr), and [Get logs](#get-logs) operations for a URL.

**Command**: `url-health-check URL [--client-location LOCATION] [--edge-server-ip EDGE_SERVER_IP] [--ip-version IPv4|IPv6] [--port 80|443] [--packet-type TCP|ICMP] [--request-header REQUEST_HEADER...] [-q QUERY_TYPE] [--run-from-site-shield-map] [--logs] [--network-connectivity]`, where:

- `URL` is the fully qualified URL you want to run the health check for.
- `LOCATION` is a unique identifier for an edge server location closest to your users. To get this value, run the [List edge server locations](#list-edge-server-locations) operation first. This flag is optional. 
- `EDGE_SERVER_IP` is the IP address of the edge server you want to serve traffic from. You can use the edge server IP address value from the `answerSection` array in the [Get domain details with dig](#get-domain-details-with-dig) operation response. 
- the `--ip_version` flag specifies the IP version for the URL health check to use, either `IPV4` or `IPV6`. By default it's set to `IPv4`. 
- the `--port` flag specifies the port number for the URL health check to use, either `80` and `443`. By default it's set to `80`.
- the `--packet_type` flag specifies the packet type for the `mtr` to use, either `tcp` or `icmp`. By default it's set to `tcp`.
- `REQUEST_HEADER` is a customized header for the `curl` request in the format `"header: value"`. The `--request-header` flag accepts multiple values.
- `QUERY_TYPE` is the DNS record type you want to get. The available values are: `A` for IPv4 address record, `AAAA` for IPv6 address record, `SOA` for Start of Authority record, `CNAME` for Canonical Name record, `PTR` for Pointer record, `MX` for Mail Exchanger record, `NS` for Nameserver record, `TXT` for Text record, `SRV` for Service Location record, `CAA` for Certificate Authority Authorization record, and `ANY` for all associated records available. This flag is optional, set to `A` by default.
- the `--logs` flag runs also the [Get logs](#get-logs) operation.
- the `--network-connectivity` flag runs also the [Test network connectivity with MTR](#test-network-connectivity-with-mtr) operation.
- the `--run-from-site-shield-map` flag uses the entered location or edge server IP to find its Site Shield map and runs the tool using the map.

The `--edge-server-ip`, `--ip-version`, `--port`, `--packet-type`, `--request-header`, `--query-type`, `--logs`, and `--network-connectivity` flags are optional.

**Examples**: 
- `akamai diagnostics url-health-check http://www.example.com --client-location bangalore-india`
- `akamai diagnostics url-health-check http://www.example.com --client-location bangalore-india --edge-server-ip 123.123.123.123 --port 80 --packet-type TCP --ip-version IPV4 --logs --network-connectivity  --request-header "X-Location: NGDT"`


**Expected outcome**: The response includes `grep`, `dig`, `curl`, `mtr`, and `translate-url` details for the URL in the JSON format. For more details, you can check the [API response](https://techdocs.akamai.com/edge-diagnostics/reference/post-url-health-check) description.


## Run the Connectivity problems scenario
The `connectivity-problem` operation simultaneously runs [Request content with cURL](#request-content-with-curl), [Test network connectivity with MTR](#test-network-connectivity-with-mtr), and [Get logs](#get-logs) operations for a URL.

**Command**: `connectivity-problem URL [--client-location LOCATION] [--edge-server-ip EDGE_SERVER_IP] [--client-ip CLIENT_IP] [--request-header REQUEST_HEADER...] [--ip-version IPv4|IPv6] [--port 80|443] [--packet-type TCP|ICMP] [--run-from-site-shield-map]`, where:

- `URL` is the fully qualified URL you want to run the Connectivity problems scenario for.
- `LOCATION` is a unique identifier for an edge server location closest to your users. To get this value, run the [List edge server locations](#list-edge-server-locations) operation first. This flag is optional.  
- `EDGE_SERVER_IP` is the IP address of the edge server you want to serve traffic from. You can use the edge server IP address value from the `answerSection` array in the [Get domain details with dig](#get-domain-details-with-dig) operation response.
- `CLIENT_IP` is the client IP for the Connectivity problems scenario to use as the MTR destination. You can use the `clientIpv4` or `clientIpv6` values from the diagnostic data. To check this values, run the [Get collected diagnostic data](#get-collected-diagnostic-data) operation.
- `REQUEST_HEADER` is a customized header for the `curl` request in the format `"header: value"`. The `--request-header` flag accepts multiple values.
- the `--ip_version` flag specifies the IP version for the Connectivity problems scenario to use, either `IPV4` or `IPV6`. By default it's set to `IPv4`. 
- the `--port` flag specifies the port number for the Connectivity problems scenario to use, either `80` and `443`. By default it's set to `80`.
- the `--packet_type` flag specifies the packet type for the `mtr` to use, either `tcp` or `icmp`. By default it's set to `tcp`.
- the `--run-from-site-shield-map` flag uses the entered location or edge server IP to find its Site Shield map and runs the tool using the map.

The `--edge-server-ip`, `--client-ip`, `--ip-version`, `--port`, `--packet-type`, `--request-header`, and `--run-from-site-shield-map` flags are optional.


**Examples**: 
- `akamai diagnostics connectivity-problem http://www.example.com --client-location bangalore-india`
- `akamai diagnostics connectivity-problem http://www.example.com --client-location bangalore-india --edge-server-ip 123.123.123.123 --client-ip 123.123.123.123 --request-header accept:text/html --port 80 --packet-type TCP --ip-version IPV4`

**Expected outcome**: The response includes `grep`, `curl`, and `mtr` details for the URL in the JSON format. For more details, you can check the [API response](https://techdocs.akamai.com/edge-diagnostics/reference/post-connectivity-problems) description.

## Run the Content problems scenario 
The `content-problem` operation simultaneously runs [Request content with cURL](#request-content-with-curl) and [Get logs](#get-logs) operations for a URL.

**Command**: `content-problem URL [--client-location LOCATION] [--edge-server-ip EDGE_SERVER_IP] [--request-header REQUEST_HEADER...] [--ip-version IP_VERSION] [--run-from-site-shield-map]`, where:

- `URL` is the fully qualified URL you want to run the Connectivity problems scenario for.
- `LOCATION` is a unique identifier for an edge server location closest to your users. To get this value, run the [List edge server locations](#list-edge-server-locations) operation first. This flag is optional.  
- `EDGE_SERVER_IP` is the IP address of the edge server you want to serve traffic from. You can use the edge server IP address value from the `answerSection` array in the [Get domain details with dig](#get-domain-details-with-dig) operation response.
- `REQUEST_HEADER` is a customized header for the `curl` request in the format `"header: value"`. The `--request-header` flag accepts multiple values.
- the `--ip_version` flag specifies the IP version for the problem scenario to use, either `IPV4` or `IPV6`. By default it's set to `IPv4`. 
- the `--run-from-site-shield-map` flag uses the entered location or edge server IP to find its Site Shield map and runs the tool using the map.

The `--edge-server-ip`, `--request-header`, and `--run-from-site-shield-map` flags are optional.


**Examples**: 
- `akamai diagnostics content-problem http://www.example.com --client-location bangalore-india --run-from-site-shield-map`
- `akamai diagnostics content-problem http://www.example.com --client-location bangalore-india --edge-ip 123.123.123.123 --request-header accept:text/html --ip-version IPV4`

**Expected outcome**: The response includes `grep` and `curl` details for the URL in the JSON format. For more details, you can check the [API response](https://techdocs.akamai.com/edge-diagnostics/reference/post-content-problems) description.

# Available flags

You can use the following flags with all the listed commands.

## edgerc 
Use the `--edgerc` flag to provide a new location of `.edgerc` file. This file contains the API credentials required to run all commands.

The CLI takes the default value of the `--edgerc` flag from the `AKAMAI_EDGERC` environment variable. If the variable is not set, then the CLI uses your home directory:
- **Linux**: `/home/{username}/.edgerc`
- **macOS**: `/Users/{username}/.edgerc`
- **Windows**: `C:\Users\{username}\.edgerc`  

**Command**: 
To add the `--edgerc` flag to a specific command: `$akamai diagnostics --edgerc EDGERC_PATH [command]`, where `EDGERC_PATH` is the new path to the file.


**Example**: 
`$akamai diagnostics --edgerc "/Users/new_user/Downloads/.edgerc" edge-locations`

## section
Use the `--section` flag to provide a new section name. The section name specifies which section of API credentials to read from the `.edgerc` file. The CLI takes the default value of the `--section` flag from the `AKAMAI_EDGERC_SECTION` environment variable. If the variable is not set, then the CLI uses `diagnostics` as the section name.

**Command**: 
To add the `--section` flag to a specific command:`$akamai diagnostics --section SECTION_NAME [command]`, where `SECTION_NAME` is the new section name.

**Example**: 
`$akamai diagnostics --section default gtm-hostnames`

## help
The `--help` flag returns help for a command. 

## version
The `--version` flag returns the version.

## json
The `--json` flag returns the information in JSON format. You can also set this flag to be `true` always by setting the environment variable `AKAMAI_OUTPUT_JSON` to `true`.

# Exit codes
When you complete an operation, the CLI generates one of these exit codes:

- `0` (Success) - Indicates that the latest command or script executed successfully.
- `1` (Configuration error) - Indicates an error while loading the CLI.
- `2` - Indicates an error related to command arguments, missing flags, or mismatch exception.
- `3` - Indicates a parsing error in API request and response.
- `100-199` - Indicates a 4xx HTTP error. The exit code equals 4xx – 300.
- `200-255` : Indicates a 5xx HTTP error. The exit code equals 5xx – 300.

# Windows 10 2018 version
If you're using Windows 10, 2018 version and you're having problems running the Edge Diagnostics 
CLI, we recommend you try the following work-around. In the downloaded repository, add the `.exe` 
suffix to the `akamai-diagnostics` executable file.

# Notice

Copyright © 2022 Akamai Technologies, Inc.

Your use of Akamai's products and services is subject to the terms and provisions outlined in [Akamai's legal policies](https://www.akamai.com/us/en/privacy-policies/).