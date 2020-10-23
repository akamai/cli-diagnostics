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
[latest release binary](https://github.com/akamai/cli-diagnostic-tools/releases) for your system, or by cloning this repository and compiling it yourself.

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
    Verifies whether the specified IP address is part of the Akamai edge network.
    ```
    $akamai diagnostics verify-ip <IP address>
    ```

- ### locate-ip
    Provides the geographic and network location of an IP address within the Akamai network.
    ```
    $akamai diagnostics locate-ip <IP address>
    ```

- ### dig
    Runs DIG on a hostname or a domain name to return DNS details for the location of an Akamai edge server and the hostname or the domain name. You can use it to diagnose issues with the DNS resolution.
    ```
    $akamai diagnostics dig <Domain name/Hostname> <Ghost location/edge server IP address> [--type Query type]
    ```

- ### mtr
    Runs MTR to check connectivity between an Akamai edge server and a remote host or destination. You can use it to diagnose network delays issues.
    ```
    $akamai diagnostics mtr <Domain Name/Destination IP> <Ghost location/edge server IP address> [--resolve-hostname]
    ```


- ### translate-url
    Provides high-level information about an Akamai-optimized URL (ARL), such as its time to live, origin server, and associated CP code.
    ```
    $akamai diagnostics translate-url <URL>
    ```

- ### translate-error-string
    Provides information about an error string from the reference number produced by Akamai edge servers when a request to retrieve content fails.
    ```
    $akamai diagnostics translate-error-string <error string>
    ```

- ### curl
    Runs CURL to provide a raw HTML for a URL within the Akamai network. You can use it to gather information about the HTTP response.
    ```
    $akamai diagnostics curl <URL> <Ghost Location/edge server IP address> [--user-agent Additional user agent]
    ```

- ### debug-url
    Provides DNS information, HTTP response, response headers, and logs for a URL on Akamai edge servers.
    ```
    $akamai diagnostics debug-url <URL> [--edge-ip Edge IP address] [--header headername:value]
    ```

- ### estats
    Provides error statistics on a CP code’s traffic from clients to Akamai edge servers and from Akamai edge servers to origin.
    ```
    $akamai diagnostics estats <URL/CP Code>	
    ```

- ### grep
    Runs GREP to retrieve and parse logs for an IP address within the Akamai network using flags to filter the data. Data is available for 48 hours after the traffic occurs.
    ```
    $akamai diagnostics grep <Edge server IP> <--end-date Date> <--end-time Time> <--duration Duration> [--find-in Header:Value ...] [--max-lines Maximum log lines to display] <-r | -f | -rf> 
    ```

- ### ghost-locations
    Lists active Akamai edge server locations from which you can run diagnostic tools.
    ```
    $akamai diagnostics ghost-locations [--search location]
    ```

- ### user-diagnostics-create-group
    Creates a group for a hostname you want to gather diagnostic data for. It also generates a diagnostic link that you can send to end users of the group’s hostname or URL. When end users click the link, the tool gathers necessary diagnostic data to submit.
    ```
    $akamai diagnostics user-diagnostics create-group <Group Name> <Hostname>
    ```

- ### user-diagnostics-get
    Lists end users' diagnostic data submitted using a diagnostic link.
    ```
    akamai diagnostics user-diagnostics get <Diagnostics Link Id>
    ```

 - ### user-diagnostics-list   
    Lists all groups created to gather diagnostic data of end users of hostnames experiencing issues together with the generated links and number of collected data.
    ```
    akamai diagnostics user-diagnostics list
    ```

### Global Flags
----
- ```--edgerc value```    Location of the edgegrid credentials file.
- ```--section value```    Section name in the credentials file.
- ```--force-color```    Force color to non-tty output.
- ```--json```  Get JSON output.


