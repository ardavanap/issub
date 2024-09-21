# IsSub

Checks if a subdomain exist with fuzzing using a provided possible subdomain wordlist.

	It simply tries to send a http or https (or both,depending on provided flags) request to generated subdomains, if 
ip is resolved and server respond in anyway, it will print it as stdout and is able to write to a specified output file directly.

## flags

- -w
	Path of wordlist of possible subdomain names.
- -d
	Domain, e.g: google.com
- -delay
	Delay in milisecond,between each request. Default is 200 
- -https
	Request on https protocol. https is already default
- -http
	Request on http protocol.
- -http-and-https
	Request on both http and https protocols.
- -o
	Output file destination.
