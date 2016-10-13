# github-webhook
Generic webhook for GitHub's notifications.

The purpose of this software is to listen on a TCP port and wait for a `push` event via GitHub's [Webhooks](https://developer.github.com/webhooks/).

A shared secret key must be exported as `SECRET_KEY` and it needs to match the key set in the GitHub configuration for the webhook. If the secret does not match the request is ignored and a `401 Unauthorized` response is sent back.

If the secret matches the JSON object is printed to `stdout` and can be piped through [jq](https://stedolan.github.io/jq/) and/or trigger a build process.

```bash
#!/bin/bash

SECRET_KEY=... ./github-webhook | while read json rest; do
  make
done
```
