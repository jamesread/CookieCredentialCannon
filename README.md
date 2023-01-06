# CookieCredentialCannon

Give people a cookie (with a unique ID), which is then mapped to credentials (optionally by firing them via a canon).

Useful for giving out temporary credentials to students during a workshop.

## How to use for an OpenShift workshop with basic users from HTPasswd

### Generate config and htpasswd 

./generateUsers.sh

This will generate config.yaml and users.htpasswd. You can edit generateUsers.sh and config.yaml as needed. You probably want to set the "server" in config.yaml to match your OpenShift console address.

With the users.htpasswd, upload it to OpenShift to create the users. Docs: https://docs.openshift.com/container-platform/4.11/authentication/identity_providers/configuring-htpasswd-identity-provider.html

### Install CookieCredentialCannon in OpenShift

1. Import from Git.
2. Use this repo URL, wait for the build to complete.
3. Add stoage, 1GB is fine, mount it to `/data/`
4. Set the environment variable `CCC_DATA` to `/data/`.
6. Use the "terminal" feature of the pod to edit `/data/config.yaml` with the `config.yaml` that you created.

