# ls-proto-deps

ls-proto-deps is a CLI tool to list dependencies of `.proto` files.

## usage

```
./ls-proto-deps protodir/foo.proto protodir/bar.proto
```

with docker:

```
docker run -t --rm -v /path/to/protodir:/protodir ghcr.io/neglect-yp/ls-proto-deps ./ls-proto-deps protodir/foo.proto
```
