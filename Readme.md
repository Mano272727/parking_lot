It is necessary that the code should be placed in `$GOPATH/src/` folder.

#### Run

`./parking_lot` runs the tests,
outputs the test results to `output.log` file. (this ensures that the output only has expected values)
built the script.
and run the results.

To enter command line tool.

```
./parking_lot
```

To execute the input file.

```
./parking_lot samples/file_input.txt
```

#### Tests

run tests using

```
go test ./...
```

#### Code Quality

Pre commit hooks:

Pre commit hooks are added to enhance code quality, by automatically checking the code for following standards.

run:

```
pre-commit install
```

for installing the hook to repos pre-commit file.

Install pre-commit from http://pre-commit.com/

gometalinter - https://github.com/alecthomas/gometalinter
