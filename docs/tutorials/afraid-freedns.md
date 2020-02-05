# Setting up afraid.org freedns

NOTE: it is probably possible to use the free subdomains in freedns but this is probably a bad idea because you could end up clashing with other users.  This implementation also does not do any "idle updates" to stop them from expiring.  These are not problems if you are using the paid service and your own domain at afraid.org, which is the intent behind this provider.

Afraid.org freedns provider support was added via [this PR](https://github.com/kubernetes-sigs/external-dns/pull/1384), thus you need to use a release where this pr is included. This should be at least v0.?.??

Requirements :
- Username and password to afraid.org login (there's not an API key system available)

Limitations :
- Only A records are currently supported (any other kind of record is hacky)

