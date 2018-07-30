
# Build & debug

````
bazel build //:test -c dbg

dlv exec --headless --listen :2345 --api-version 2 bazel-bin/linux_amd64_debug/test

# Start remote debug config in Intellij
````
