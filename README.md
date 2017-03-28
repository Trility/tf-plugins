# tf-plugins
Terraform Plugins

[![Build Status](https://travis-ci.org/Trility/tf-plugins.svg?branch=master)](https://travis-ci.org/Trility/tf-plugins)

## Install

* Get the latest application
```
go get github.com/Trility/tf-plugins
```
* Edit ~/.terraformrc
```
providers {
    trility = "/path/to/golang/bin/tf-plugins"
}
```

## Build

* Update git config to use SSH instead of HTTPS for github
```
git config --global url."git@github.com:".insteadOf "https://github.com/"
```
* Install Glide
```
curl https://glide.sh/get | sh
```
* Install Dependencies
```
glide install
```
* Build
```
go build -o ${GOPATH}/bin/tf-plugins
```

## Usage
Coming soon....
