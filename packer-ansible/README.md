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

