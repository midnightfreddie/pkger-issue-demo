An example for <https://github.com/markbates/pkger/issues/44>

**After deleting errant /cmd/server/pkged.go and adding /for-pkger.go**

`pkger list`:

    github.com/midnightfreddie/pkger-issue-demo
    > github.com/midnightfreddie/pkger-issue-demo:/cmd/server/html
    > github.com/midnightfreddie/pkger-issue-demo:/cmd/server/html/index.html

`pkger -o /cmd/server` works as expected, putting embedded file(s) into /cmd/server/pkged.go

However, if I delete that pkged.go and then run `pkger`:

> 2020/01/20 10:53:26 exit status 1: can't load package: package github.com/midnightfreddie/pkger-issue-demo: found packages foo (for-pkger.go) and main (pkged.go) in C:\Users\Jim\src\go\src\github.com\midnightfreddie\pkger-issue-demo