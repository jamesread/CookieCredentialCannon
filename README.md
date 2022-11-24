# CookieCredentialCannon

Give people a cookie (with a unique ID), which is then mapped to credentials (optionally by firing them via a canon).

Useful for giving out temporary credentials to students during a workshop.

## Installation in OpenShift

1. Import from Git.
2. Use this repo URL, wait for the build to complete.
3. Add stoage, 1GB is fine, mount it to `/data/`
4. Set the environment variable `CCC_DATA` to `/data`.
6. Use the "terminal" feature of the pod to edit `/data/config.yaml`.

## `config.yaml`

```
server: "http://example.com"
cookieName: "ccc"
credentials:
  - username: user1
    password: user1

  - username: user2
    password: user2

  - username: user3
    password: user3
```
