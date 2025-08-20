# Mise Go Backend Tool Version Issue

This repository demonstrates an issue that might arise when using Mise's Go backend.

Basically, some have a dependency on the version of go they were generated with, which means that when
upgrading the go version, the tool must also be recompiled.

For example, if you run a `mise install` (assuming you don't have moq previously installed) and then do this change:

```patch
diff --git a/go.mod b/go.mod
index d6c6c0f..20e5f1c 100644
--- a/go.mod
+++ b/go.mod
@@ -1,3 +1,3 @@
 module github.com/bamorim/mise-go-tool-version-issue

-go 1.24.1
+go 1.25.0
diff --git a/mise.toml b/mise.toml
index bdd8388..3a2c62f 100644
--- a/mise.toml
+++ b/mise.toml
@@ -1,3 +1,3 @@
 [tools]
-go = "1.24.5"
+go = "1.25.0"
 "go:github.com/matryer/moq" = "0.5.3"
```

And then run `mise install` and `go generate ./...`

You'll receive the following error:

```
couldn't load source package: /Users/bamorim/github.com/bamorim/mise-go-tool-version-issue/main.go:1:1: package requires newer Go version go1.25 (application built with go1.24) (and 1 more errors)
moq [flags] source-dir interface [interface2 [interface3 [...]]]
  -fmt string
        go pretty-printer: gofmt, goimports or noop (default gofmt)
  -out string
        output file (default stdout)
  -pkg string
        package name (default will infer)
  -rm
        first remove output file, if it exists
  -skip-ensure
        suppress mock implementation check, avoid import cycle if mocks generated outside of the tested package
  -stub
        return zero values when no mock implementation is provided, do not panic
  -version
        show the version for moq
  -with-resets
        generate functions to facilitate resetting calls made to a mock
Specifying an alias for the mock is also supported with the format 'interface:alias'
Ex: moq -pkg different . MyInterface:MyMock
main.go:3: running "moq": exit status 1
```

The current workaround is to run `mise uninstall go:github.com/matryer/moq` and then `mise install` again.
