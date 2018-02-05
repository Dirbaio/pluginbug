#!/bin/bash

echo If we run the plugin directly, everthing works fine:

go run ./plugin.v1/main.go

echo
echo But if we compile the plugin as a plugin, and try to dynamically
echo load it from another package, we see that the init funcs are not run.
go build -buildmode=plugin -o plugin.so ./plugin.v1/
go run host/main.go

echo
echo If you rename plugin.v1 directoy to pluginv1 without the dot, it works.
