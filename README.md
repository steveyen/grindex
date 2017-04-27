grindex - yet another indexing prototype
======================================================

Originally forked from github.com/couchbaselabs/grouter to reuse a
bunch of the bones.

Instead of "hello world", I write proxy/routers when trying out a new
language.  With grindex, I can also play with design ideas for a moxi
"2.0", but learn golang at the same time.

Building
--------

First, set up your GOPATH, like...

    export GOPATH=~/go

Then, get the code...

    go get github.com/steveyen/grindex

Or, old-school...

    mkdir -p ~/go/src/github.com/steveyen
    cd ~/go/src/github.com/steveyen
    git clone git://github.com/steveyen/grindex.git

For developers, to (re-)build it...

    cd grindex
    (cd grindex && go build)

Running
-------

    ./grindex/grindex --help

Workload generation...

    ./grindex/grindex \
        --source=workload:concurrency=200 \
        --target=couchbase://10.3.121.192:8091 \
        --target-concurrency=200

License
-------

Apache 2.0 License
