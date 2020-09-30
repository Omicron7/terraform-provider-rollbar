Terraform provider for Rollbar
==============================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)


Requirements
------------

- [Terraform](https://www.terraform.io/downloads.html) 0.12.x
- [Go](https://golang.org/doc/install) 1.13.x+ (to build the provider plugin)


License
-------

This is Free Software, released under the terms of the MIT license.


Status
------

[![Build](https://github.com/jmcvetta/terraform-provider-rollbar/workflows/Build/badge.svg)](https://github.com/jmcvetta/terraform-provider-rollbar/actions)


Debugging
---------

Enable writing debug log to `/tmp/terraform-provider-rollbar.log` by setting an
environment variable:

```
export TERRAFORM_PROVIDER_ROLLBAR_DEBUG=1
terraform apply   # or any command that calls the Rollbar provider
```

### Dev Script

Running `make dev` will:
* Build and install the provider 
* Run `terraform apply` in the `examples` folder with debug logging enabled
* Display the logs on completion.


History
-------

Derived from
[jmcvetta/terraform-provider-rollbar-jmcvetta](https://github.com/jmcvetta/terraform-provider-rollbar-jmcvetta)
and
[babbel/terraform-provider-rollbar](https://github.com/babbel/terraform-provider-rollbar)
