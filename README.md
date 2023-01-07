# CookieCredentialCannon

Give people a cookie (with a unique ID), which is then mapped to credentials (optionally by firing them via a canon).

Useful for giving out temporary credentials to students during a workshop.

## How to use for an OpenShift workshop with basic users from HTPasswd

### Generate config and htpasswd 

This script will generate `config.yaml` and `users.htpasswd`. You can edit `generateUsers.sh` before generating to customize the username/password format.

```
user@host: ./generateUsers.sh
```

After `config.yaml` has been generated, you will want to edit it, to set the "server" in to match your OpenShift console address.

### Upload `users.htpasswd` as an OAuth provider in OpenShift

With the users.htpasswd, upload it to OpenShift to create the users. This is straightforward and takes just a few minutes. Docs: https://docs.openshift.com/container-platform/4.11/authentication/identity_providers/configuring-htpasswd-identity-provider.html

### Install CookieCredentialCannon in OpenShift

1. Go to the Developer perspective, and Add.
2. Choose "Import from Git".
3. Use this repo URL ("https://github.com/jamesread/CookieCredentialCannon.git") and submit. Wait for the build to complete.
4. Edit the deployment to add PV storage, 1GB is fine, mount it to `/data/`
5. Edit the deployment to set the environment variable `CCC_DATA` to `/data/`.
6. Use the "terminal" feature of the pod to edit `/data/config.yaml` with the `config.yaml` that you created.
7. Visit the Route of the application to get to a webpage, you should be assigned to the first user. Horray :-)

Any issues, or help needed, please raise a GitHub issue.

