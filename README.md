ggoc
====

[![GitHub release](http://img.shields.io/github/release/noissefnoc/ggoc.svg?style=flat-square)][release]
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)][license]

[release]: https://github.com/noissefnoc/ggoc/releases
[license]: https://github.com/noissefnoc/ggoc/blob/master/LICENSE.txt

Generate Google OAuth2 Credential file for batch program.


Usage
-------

### Get client secret from Google Cloud Platform

First you create OAuth client and get client secret from [Google Cloud Console](https://console.cloud.google.com/).

1. Go to Google Cloud Console
1. Click **API Manager**
1. Click **Create Credentials**
    1. Click **Create OAuth Client ID**
    1. Choose **Other**
    1. Fill **Application name**
1. Get the client id and client secret
    1. Download client secret **JSON** from download button

then download client secret file.  (this file path for `-secret` option)

### Run command

After saving client secret you run `ggoc` as follow:

```
$ ggoc -secret CLIENT_SECRET_PATH \
       -credential CREDENTIAL_OUTPUT_PATH \
       -scope GOOGLE_OAUTH_SCOPE
```

The candidate of `GOOGLE_OAUTH_SCOPE` is listed [here](https://developers.google.com/identity/protocols/googlescopes).

### Get OAuth token from browser

OAuth authorization requires

1. 'Go to the following link in your browser then type the authorization code' log and Authorization URL display on console
1. Copy and paste authorization URL on browser address bar
1. OAuth token displays if authorization success.
1. Copy and paste OAuth token to console and hit enter key

After those procedure, you can get credential file detected `-credential` option.


Install
-------

If you are macOS user, you can install via [Homebrew](https://brew.sh/).

```
$ brew tap noissefnoc/homebrew-ggoc
$ brew install ggoc
```

Other OS users can download binary from [release page](https://github.com/noissefnoc/ggoc/releases).

And you can also use `go get`

```
$ go get -u github.com/noissefnoc/ggoc
```

License
-------

MIT

Author
------

[Kota Saito](https://github.com/noissefnoc)
