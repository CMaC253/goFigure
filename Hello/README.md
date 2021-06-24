# Hello World Notes
1. When your code imports packages contained in other modules, you manage those through your own code's module
2. That module is defined by go.mod
    - that tracks the modules that provide those packages.
    - go.mod file stays with your code, including the source code repo
3. run 'go mod init' giving the name of the module your code will be in
    - $ go mod init example.com/hello
    - in most cases, this will be the repo location where your source code will be in i.e. github.com/mymodule
    