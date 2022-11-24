# CookieCredentialCannon

Give people a cookie (with a unique ID), which is then mapped to credentials (optionally by firing them via a canon).

Useful for giving out temporary credentials to students during a workshop.

## config.yaml

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
