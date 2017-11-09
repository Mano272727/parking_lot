It is necessary that the code should be placed in `$GOPATH/src/` folder.

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
