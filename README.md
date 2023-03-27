# MINUS NA PLUS API
[HOME PAGE](http://minus-na-plus.nabank.tech)

![example workflow](https://github.com/minusnaplus/api/actions/workflows/sec-scan.yml/badge.svg)
![example workflow](https://github.com/minusnaplus/api/actions/workflows/golangci-lint.yml/badge.svg)
![example workflow](https://github.com/minusnaplus/api/actions/workflows/mathlogic-test.yml/badge.svg)
![example workflow](https://github.com/minusnaplus/api/actions/workflows/publish.yml/badge.svg)


# Create a Hetzner snapshot using Packer - provisioning using Ansible

The snapshot can then be used to quickly stand a cluster using Terraform

You can customize the Ansible playbook to do more customisations.

## Usage
* Install [Packer](https://www.packer.io/docs/install) and [Ansible](https://docs.ansible.com/ansible/latest/installation_guide/intro_installation.html)
* Get an API token from Hetzner Cloud
* Set Hetzner token environment variable `export HCLOUD_TOKEN=xxx`
* Customize packer.json
* Customize the Ansible playbook
* Run packer to create the snapshot

```bash
deploy_user_name
packer build packer.json

```
## API
This is a simple API with mathematical operations designed to test Fiber(go) & Node pipelines. Additionally, it serves as a proof of concept for testing the performance of the Node Fastify vs Golang Fiber framework.

![Alt text](diagram.drawio.png "network diagram")
### Localhost API Requests
```bash
curl http://erdos.localhost/v1/healthy
```
### Public API Requests
```bash
curl http://minus-na-plus.nabank.tech/v1/api/healthy
```
### Secure API Requests
* Adds two numbers together.
* /add ``` GET: curl --cookie "access_token=leaked-key-123" "http://minus-na-plus.nabank.tech/v1/api/add?x=0&y=1" ```
*  Subtracts the second number from the first.
* /sub ``` GET: curl --cookie "access_token=leaked-key-123" "http://minus-na-plus.nabank.tech/v1/api/sub?x=4&y=2" ```
*  Multiplies two numbers together.
* /mul ``` GET: curl --cookie "access_token=leaked-key-123" "http://minus-na-plus.nabank.tech/v1/api/mul?x=5&y=3" ```
* Multiplies two numbers together.
* /div ``` GET: curl --cookie "access_token=leaked-key-123" "http://minus-na-plus.nabank.tech/v1/api/div?x=10&y=2" ```
* Returns the mathematical constant e.
* /e ``` GET: curl --cookie "access_token=leaked-key-123" "http://minus-na-plus.nabank.tech/v1/api/e" ```

####  Interested? You can find more information about the API on SWAGGER. 

[![Alt text](front_test_app_screen.png "Some fun front edt test")](http://minus-na-plus.nabank.tech)



