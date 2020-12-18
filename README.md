# Diagnostic Tools CLI

* [Get Started with the Diagnostic Tools CLI](#get-started-with-the-diagnostic-tools-cli)

* [Install Diagnostic Tools CLI](#install-diagnostic-tools-cli)

* [Stay up to date](#stay-up-to-date)

* [Concepts](#concepts)

* [Available commands](#available-commands)

    * [help](#help)

    * [list](#list)

    * [user-diagnostics create-group](#user-diagnostics-create-group)

    * [user-diagnostics list](#user-diagnostics-list)

    * [user-diagnostics get](#user-diagnostics-get)

    * [verify-ip](#verify-ip)

    * [locate-ip](#locate-ip)

    * [dig](#dig)

    * [mtr](#mtr)

    * [translate-url](#translate-url)
    
    * [translate-error-string](#translate-error-string)

    * [curl](#curl)

    * [grep](#grep)

    * [debug-url](#debug-url)

    * [estats](#estats)

    * [ghost-locations](#ghost-locations)

* [Available flags](#available-flags)

    * [--edgerc value](#--edgerc-value)

    * [--section value](#--section-value)

    * [--help](#--help)

    * [--version](#--version)

    * [--json](#--json)

* [Windows 10 2018 version](#windows-10-2018-version)  

* [Notice](#notice)

# Get Started with the Diagnostic Tools CLI

The Diagnostic Tools CLI lets you identify, analyze, and troubleshoot common content delivery network issues that your users may encounter. 

## Install Diagnostic Tools CLI

To install this CLI, you need the [Akamai CLI](https://github.com/akamai/cli) package manager. Once you install the Akamai CLI, run this command:

`akamai install diagnostics` 

## Stay up to date

To make sure you always use the latest version of the CLI, run this command:  

`akamai update diagnostics`  


## Concepts

To learn more about the concepts behind this CLI and the Diagnostic Tools API, see the Diagnostic Tools API [Overview](https://developer.akamai.com/api/core_features/diagnostic_tools/v2.html#overview).


## Available commands 

### `help`
This command returns a list of available commands and flags with their descriptions.


### `list`
This command lists available commands with their short descriptions. 


### `user-diagnostics create-group`
This command creates a group in Diagnostic Tools for a hostname you want to gather diagnostic data for. It also generates a diagnostic link 
that you can send to end users of a group’s hostname or URL. When end users click the link, the tool gathers the necessary diagnostic data to submit.

**Command**: `user-diagnostics create-group group_name hostname` where:

- `group_name` is the name you assign to the group.
- `hostname` is the hostname you need to gather diagnostic data for.

**Example**: `akamai diagnostics user-diagnostics create-group test_group www.akamai.com`

**Expected output**: The response includes the link you can send to end users experiencing issues with the hostname. When an end user 
clicks the link, Diagnostic Tools gathers necessary data and asks the user to submit them. Each link is active for 7 days and has a limit of 50 submissions. 


### `user-diagnostics list`
This command lists the diagnostic groups that gather user data for hostnames experiencing issues. It also lists the generated links and the number of submitted data.

**Command**: `user-diagnostics list`

**Expected output**: The response includes basic information about groups including names, hostnames, number of submitted data, and generated links. If the status of the link is `expired`, it means that you can't gather more data for this group. For more details, you can check the [API response](https://developer.akamai.com/api/core_features/diagnostic_tools/v2.html#dd80149d) description.

### `user-diagnostics get`
This command returns diagnostic data submitted by end users.

**Command**: `get link_id` where `link_id` is the identifier for the generated link. To get the identifier, run the `user-diagnostics list` command.

**Example**: `akamai diagnostics get 2661`

**Expected output**: The response includes a table with the collected data, such as timestamp (UTC), client IP preferred, client DNS IPv4, client DNS IPv6, user agent, cookie, protocol, connected cipher, client IPv4, and client IPv6. For more details, you can check the [API response](https://developer.akamai.com/api/core_features/diagnostic_tools/v2.html#d3c2cd5e) description. 

### `verify-ip`
This command checks whether the specified IP address is part of the Akamai edge network.

**Command**: `verify-ip IP_address` where `IP_address` is the IP address you want to verify.

**Example**: `akamai diagnostics verify-ip 123.123.123.123`

**Expected output**: The response includes information whether the IP is part of the Akamai edge network or not. 

### `locate-ip`
This command provides the geographic and network location of an IP address within the Akamai network. This operation’s requests are limited to 500 per day.

**Command**: ` locate-ip IP_address` where `IP_address` is the IP address you want to get the location data for.

**Example**: `akamai diagnostics locate-ip 123.123.123.123`

**Expected output**: The response includes the IP's geographic and network location data. For more details, you can check the [API response](https://developer.akamai.com/api/core_features/diagnostic_tools/v2.html#f4e74e4b) description.

### `dig`
This command returns DNS details for the location of an Akamai edge server and the hostname or the domain name. You can use it to diagnose issues with the DNS resolution.

**Command**: `dig hostname source_server_location/edge_server_IP --type query_type` where:

- `hostname` is either a hostname or a domain name you want to get the DNS details for. 
- `source_server_location/edge_server_IP` is either one of the locations listed by the `ghost-locations` command or an edge server IP.
- `query_type` is one of the following: `A`, `AAAA`, `SOA`, `CNAME`, `MX`, `PTR`, `NS`. The default is `A`. 

The `--type` flag is optional.

**Example**: `akamai diagnostics dig www.akamai.com bangalore-india --type NS`

**Expected outcome**: The response includes a standard `dig` response. For more details, you can check the [API response](https://developer.akamai.com/api/core_features/diagnostic_tools/v2.html#0b02e065) description. 

### `mtr`
This command provides information about the route, number of hops, and time that Internet traffic packets take between the Akamai edge server and a remote host or destination. You can use it to diagnose network delays issues.

**Command**: `mtr domain_name/destination_IP source_server_location/edge_server_IP --resolve-hostname` where:

- `domain_name/destination_IP` is either a domain name or a destination IP.
- `source_server_location/edge_server_IP` is either one of the locations listed by the `ghost-locations` command or an edge server IP. 

The `--resolve-hostname` flag is optional. Without it, the outcome includes only IP addresses.

**Example**: `akamai diagnostics mtr www.akamai.com bangalore-india --resolve-hostname`

**Expected outcome**: The response includes a standard `mtr` response. For more details, you can check the [API response](https://developer.akamai.com/api/core_features/diagnostic_tools/v2.html#576a9e79) description. 

### `translate-url`
This command gets high-level information about an Akamai-optimized URL (ARL), such as its time to live, origin server, and associated CP code. 

**Command**: ` translate-url URL` where `URL` is the URL with a protocol you want to get information about.

**Example**: `akamai diagnostics translate-url http://www.akamai.com`

**Expected outcome**: The response includes basic metadata about the URL. For more details, you can check the [API response](https://developer.akamai.com/api/core_features/diagnostic_tools/v2.html#5f95c48d) description.  

### `translate-error-string`
This command provides a summary and logs for the error that occurred in the original request. It uses the error string from the reference number to provide this data.

**Command**: `translate-error-string error_string` where `error_string` is the error reference number you want to translate.

**Example**: `akamai translate-error-string 9.6f64d440.1318965461.2f2b078`

**Expected outcome**: The response includes the error translation. For more details, you can check the [API response](https://developer.akamai.com/api/core_features/diagnostic_tools/v2.html#0fbbc293) description.

### `curl`
This command runs `curl` which provides raw HTML for a URL within the Akamai network. You can use it to gather information about the HTTP response. 

**Command**: `curl URL source_server_location/edge_server_IP --user-agent additional-user-agent` where:

- `URL` is the URL within the Akamai network you want to get data for.
- `source_server_location/edge_server_IP` is either one of the locations listed by the `ghost-locations` command or an edge server IP. 
- `additional-user-agent` can be one of the following: `android`, `firefox`, `iphone`, `mobile`, `chrome`, `msie`, `msie9`, `msie10`, `safari`, `safari/5`, `safari/6`, `webkit`, `webkit/5`, `webkit/6`.

**Example**: `akamai diagnostics curl http://www.akamai.com bangalore-india --user-agent chrome`

**Expected outcome**: The response includes the `curl` response. For more details, you can check the [API response](https://developer.akamai.com/api/core_features/diagnostic_tools/v2.html#859e3dca) description. 

### `grep`
This command lists log lines from an IP address within the Akamai network. You can use parameters to filter the data. Logs provide 
low-level details on how each request was handled. You can use this information to troubleshoot caching and performance issues, and to 
ensure the correct Akamai features were applied to the traffic. Data is available for 48 hours after the 
traffic occurs.

**Command**: `grep edge_server_IP --end-date date --end-time time --duration duration --find-in Header:Value --max-lines maximum_log_lines_to_display -r | -f | -rf` where:

- `edge_server_IP` is the edge server IP you want to get logs for.
- `date` is the date in the past when the log search window ends, in <YYYY-MM-DD> format.
- `time` is the time when the log search window ends. Enter in <HH-MM-SS> format using the UTC time zone.
- `duration` is the number of minutes before the `end-date` and `end-time` you want to collect logs for. The default is 30, and 360 is the maximum.
- `Header:Value` is the HTTP status code you want to filter the logs for. Possible values are: `Host Header`, `User Agent`, `HTTP Status Code`, `ARL`, `CP Code` and `Client IPonly`. 
- `maximum_log_lines_to_display` is the maximum number of log lines to include in the results, `200` by default and a maximum of `1000`.
- `-r | -f | -rf` is the type of log lines you want to filter. The possible values are: `-r` for incoming client requests, `-f` for requests to other edge servers or to the origin
 server, or `-rf` for both. Any `-f` log type specifying a `Forward-IP` in the `10.x.x.x` range means the request was forwarded to another edge server.

Only the `--end-date`, `--end-time`, `--duration` flags are mandatory. The `--find-in` flag can be used multiple times.

**Example**: `akamai diagnostics grep 123.123.123.123 --end-date 2020-07-20 --end-time 11:45:03 --duration 90 --max-lines 200 --find-in ”HTTP Status Code:400” --find-in ”CP code:12345” -rf`

**Expected outcome**: The response includes a `grep` response. For more details, you can check the [API response](https://developer.akamai.com/api/core_features/diagnostic_tools/v2.html#802809bc) description.

### `debug-url`
This command provides DNS information, HTTP response, response headers, and logs for a URL on Akamai edge servers. 

**Command**: `debug-url URL --edge-ip edge_server_IP --header request_header` where:

- `URL` is the URL you want to gather data for.
- `edge_server_IP` is the edge server IP address to test the URL against, otherwise a random server by default.
- `request_header` is any additional header to add to the request. 

The `--edge-ip`, `--header` flags are optional. The `--header` flag can be used multiple times.

**Example**: `akamai diagnostics debug-url http://www.akamai.com --edge-ip 123.123.123.123 --header accept:text/html`

**Expected outcome**: The response includes HTTP and DNS information for the URL. For more details, you can check the [API response](https://developer.akamai.com/api/core_features/diagnostic_tools/v2.html#b51d1d78) description.

### `estats`
This command returns error statistics on a CP code’s traffic from Akamai edge servers to both clients and the origin. 

**Command**: `estats URL/CP_code` where `URL/CP_code` is either a URL or a CP code you want to get statistics for.

**Example**: `akamai diagnostics estats https://www.akamai.com`

**Expected outcome**: The response includes error statistics. For more details, you can check the [API response](https://developer.akamai.com/api/core_features/diagnostic_tools/v2.html#010c57a6) description.  

### `ghost-locations`
Lists active Akamai edge servers in a particular location from which you can run diagnostic tools.

**Command**: `ghost-locations --search location`, where `location` is the location you want to check Akamai edge servers for.

**Example**: `akamai diagnostics ghost-locations --search india`

**Expected outcome**: The response includes a list of Akamai edge servers locations.

## Available flags

You can use the following flags with all the listed commands.

### `--edgerc value`
This flag returns the location of the credentials file.  

### `--section value`
This flag returns the section name of the credentials file. 

### `--help`
This flag returns help for a command. 

### `--version`
This flag returns the version.

### `--json`
This flag returns the information in JSON format.

# Windows 10 2018 version
If you're using Windows 10, 2018 version and you're having problems running the Diagnostic Tools 
CLI, we recommend you try the following work-around. In the downloaded repository, add the `.exe` 
suffix to the `akamai-diagnostics` executable file.

# Notice

Copyright © 2018-2020 Akamai Technologies, Inc.

Your use of Akamai's products and services is subject to the terms and provisions outlined in [Akamai's legal policies](https://www.akamai.com/us/en/privacy-policies/).