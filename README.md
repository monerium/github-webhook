# github-webhook
Generic webhook for GitHub's notifications.

The purpose of this software is to listen on a TCP port and wait for a `push` event via GitHub's [Webhooks](https://developer.github.com/webhooks/).

### Configuration

A shared secret key must be exported as `SECRET_KEY` and it needs to match the key set in the GitHub configuration for the webhook. If the secret does not match the request is ignored and a `401 Unauthorized` response is sent back.

The listening port must be exported as `PORT`.

If the secret matches the JSON object is printed to `stdout` and can be piped through [jq](https://stedolan.github.io/jq/) and/or trigger a build process.

```bash
#!/bin/bash

SECRET_KEY=... ./github-webhook | while read json rest; do
  make
done
```

### Caveats

When using [daemontools](https://cr.yp.to/daemontools.html) to manage the process a simple unix pipe like the one above will leave the `github-webhook` hanging when trying to bring the process down.

A simple solution is to use [socat](http://www.dest-unreach.org/socat/) to pipe `github-webhook`'s _stdout_ to a shell script.

```bash
#!/bin/sh

exec 2>&1
exec socat -u EXEC:'envdir env setuidgid nobody ./github-webhook' EXEC:'envdir env setuidgid nobody ./build.sh'
```

With `build.sh`:
```bash
#!/bin/bash

while read json; do
  reponame=$(echo $json | jq --raw-output '.repository.name')
  echo $reponame
  # ...
done
```
