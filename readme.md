An example for <https://github.com/markbates/pkger/issues/44>

`pkger list`:

    > github.com/midnightfreddie/pkger-issue-demo:/cmd/server/html
    > github.com/midnightfreddie/pkger-issue-demo:/cmd/server/html/index.html

`pkger -o /cmd/server`:

> 2020/01/20 10:42:23 exit status 1: can't load package: package github.com/midnightfreddie/pkger-issue-demo: cannot find module providing package github.com/midnightfreddie/pkger-issue-demo
