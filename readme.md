An example for <https://github.com/markbates/pkger/issues/44>

`pkger list`:

    > github.com/midnightfreddie/pkger-issue-demo:/cmd/server/html
    > github.com/midnightfreddie/pkger-issue-demo:/cmd/server/html/index.html

`pkger -o /cmd/server`:

> 2020/01/20 10:42:23 exit status 1: can't load package: package github.com/midnightfreddie/pkger-issue-demo: cannot find module providing package github.com/midnightfreddie/pkger-issue-demo

And the resulting /cmd/server/pkged.go looks like this and has to be deleted before doing other pkger actions:

    package main

    import (
        "github.com/markbates/pkger"
        "github.com/markbates/pkger/pkging/mem"
    )

    var _ = pkger.Apply(mem.UnmarshalEmbed([]byte(`