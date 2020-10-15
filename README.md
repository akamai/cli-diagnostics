# Akamai CLI for Diagnostic Tools

An [Akamai CLI](https://developer.akamai.com/cli) package for functional testing. With functional testing, you can assert expected conditions to positively confirm that a new business requirement or config version behaves as intended.
## Getting Started
----
### Installing

To install this package, use Akamai CLI:

```sh
$ akamai install diagnostics
```

You may also use this as a stand-alone command by downloading the
[latest release binary](https://github.com/akamai/cli-diagnostic-tools/releases)
for your system, or by cloning this repository and compiling it yourself.

### Compiling from Source (when on github)

If you want to compile the package from source, you will need Go 1.13 or later installed:

1. Fetch the package:  
    `go get github.com/akamai/cli-diagnostic-tools`
  
2. Change to the package directory:
    `cd $GOPATH/src/github.com/akamai/cli-diagnostic-tools`
  
3. Install dependencies:
    `depends on whether glide is used or vendor`
  
4. Compile the binary:
    - Linux/macOS/*nix: `go build -mod=vendor -o akamai-diagnostics`
    - Windows: `go build -mod=vendor -o akamai-diagnostics.exe`
  
5. Move the binary (`akamai-diagnostics` or `akamai-diagnostics.exe`) in to your `PATH`

### Instructions for akamai internal with access to development branch
If you want to compile the package from source, you will need Go 1.13 or later installed:

1. Set up required environment variables such as `GOPATH` by setting them in your `.bashrc` or equivalent.
   ```
   echo 'export GOPATH="$HOME/go"' >> ~/.bashrc'    # Append GOPATH declaration to .bashrc file
   source ~/.bashrc                                 # Reload .bashrc to apply the changes
   ```
2. Build
   ```
   go build
   ```
3. Run the CLI
   ```
   ./akamai-diagnostics
   ```  
## Credentials
----
- Set up your credential files as described in the [authorization](https://developer.akamai.com/introduction/Prov_Creds.html) and [credentials](https://developer.akamai.com/introduction/Conf_Client.html) sections of the getting started guide on developer.akamai.com.
- When working through this process you need to give your API credential the "Diagnostic Tools” Grant. The section in your configuration file should be called ‘diagnostics’.
- You may also use the –section to use the specific section credentials from your .edgerc file
    ```
    [diagnostics]
    client_secret = [CLIENT_SECRET]
    host = [HOST]
    access_token = [ACCESS_TOKEN_HERE]
    client_token = [CLIENT_TOKEN_HERE]
    ```

## Commands
----
- ### verify-ip
    This tool confirms if a certain IP address is that of an edge server
    ```
    $akamai-diagnostics verify-ip <IP address>
    ```

- ### locate-ip
    This tool provides the geographic and network location of any IP address.
    ```
    $akamai-diagnostics locate-ip <IP address>
    ```

- ### dig
    This tool uses the DIG command to provide Domain Name Server (DNS) details for the location of the edge server and hostname or domain name, enabling you to diagnose issues with domain name resolution.
    ```
    $akamai-diagnostics dig <Domain name/ Hostname> <Source server location/ Edge server IP> [-t Query Type]
    ```

- ### mtr
    This tool uses the MTR command to provide information about the route, number of hops, and time that Internet traffic packets take between the edge server and a remote host or destination. The results can show you where network delays are being introduced in the path.
    ```
    $akamai-diagnostics mtr <Domain name/ Destination IP> <Source server location/ Edge server IP> [--resolve-host]
    ```


- ### translate-url
    This tool provides basic information about a specified URL, such as typecode, origin server, CP code, serial number, and TTL for a URL/ARL.
    ```
    $akamai-diagnostics translate-url <URL>
    ```

- ### translate-error-string
    This tool uses the error string from the reference number to fetch a summary and log information for the error that occurred in the original request.
    ```
    $akamai-diagnostics translate-error-string <Error string>
    ```

- ### curl
    This tool uses the CURL command to provide the raw html for a specified URL. Making an HTTP request from an edge server lets you gather information about the HTTP response.
    ```
    $akamai-diagnostics curl <URL> <Source server location/ Edge server IP> [--user-agent Additional user agent]
    ```

- ### debug-url
    This tool provides DNS Information, HTTP Response, Response Header, and Logs for a URL/ARL.
    ```
    $akamai-diagnostics debug-url <URL> [--edge-ip Edge server IP] [--header Request header]
    ```

- ### estats
    This tool provides an understanding of the errors happening in the delivery of websites based on real-time data of traffic of a particular CP code in terms of traffic from clients to edge servers and from edge servers to origin.
    ```
    $akamai-diagnostics estats <URL/ CP code>	
    ```

- ### grep
    This tool uses the GREP command to retrieve and parse logs from an edge server IP address, within the last 48 hours
    ```
    $akamai-diagnostics grep <Edge server IP> <--end-date Date> <--end-time Time> <--duration Duration> [--find-in Header:Value] [--max-lines Maximum log lines to display] <-r|-f|-rf> 
    ```

- ### ghost-locations
    Lists active Akamai edge server locations from which you can run diagnostic tools. Use any id value from the response object for use in other ghost location-based operations.
    ```
    $akamai-diagnostics ghost-locations [--search location]
    ```

- ### user-diagnostics-create-group
    Create group to get user sharable link.
    ```
    $akamai-diagnostics user-diagnostics create-group <Group Name> <Hostname> [flags]
    ```

- ### user-diagnostics-get
    get end user diagnostics result by Id.
    ```
    akamai-diagnostics user-diagnostics get <Diagnostics Link Id> [flags]
    ```

 - ### user-diagnostics-list   
    List all the End User Diagnostic Data.
    ```
    akamai-diagnostics user-diagnostics list [flags]
    ```

### Global Flags
----
- ```--edgerc value```    Location of the credentials file  
- ```--section value```    Section of the credentials file 
- ```--force-color```    Force color to output, when the output is piped or redirected to text file
- ```--json```  Get json output


