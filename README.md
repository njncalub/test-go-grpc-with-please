# test-go-grpc-with-please

Testing [gRPC](https://grpc.io/) using [Go](https://golang.org/) and [Please](https://please.build).

## Running

```bash
# Run the program
$ plz run //src:main
Hello, World!

# Run the tests
$ plz test //src:...
//src/greetings:greetings_test 1 test run in 0s; 1 passed [cached]
1 test target and 1 test run; 1 passed.
Total time: 1.05s real, 0s compute.
```
