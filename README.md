# tf-plugins
Terraform Plugins

[![Build Status](https://travis-ci.org/Trility/tf-plugins.svg?branch=master)](https://travis-ci.org/Trility/tf-plugins)

## Install

* Install Go (Minimum 1.8)
* https://golang.org/dl/

```
# Create a workspace for golang projects
mkdir -p ~/golang/src
mkdir -p ~/golang/pkg
mkdir -p ~/golang/bin

# Edit .bashrc to set go paths for user specific
# Golang Exports
export PATH=$PATH:/usr/local/go/bin
export GOROOT=/usr/local/go
export GOPATH=~/golang
export PATH=$PATH:$GOPATH/bin
```

* Build the plugin
```
curl https://glide.sh/get | sh
cd $GOPATH/src/github.com/Trility/
git clone https://github.com/Trility/tf-plugins.git
glide install
go build -o ${GOPATH}/bin/tf-plugins
```

* Add/Edit ~/.terraformrc
```
providers {
    trility = "/path/to/golang/bin/tf-plugins"
}
```
