# gowiki [WIP]
> A wiki-like web app written in Go.
> Build upon the Go Tutorial found at https://golang.org/doc/articles/wiki/

`gowiki` is a wiki-like Web Application written in Go. It has the ability to edit and create pages which are stored as files and is the result of the Go Tutorial at https://golang.org/doc/articles/wiki/.

The code provided in this repository shouldn't be deployed anywhere. It has (almost) zero security and validation.

---

## Why

After finishing the tutorial I wanted to refactor the code and split it across files and build a bit upon it, eventually adding new features et all - building a wiki is a nice project after all.

## WIP

There are currently a few things half-way implemented, like the `title` input field which doesn't do anything and the general CSS layout (I'm playing around with Grids!).

All-in-all, this is a playground for me. It's not meant to be used anywhere!

## Use

If you want to play around with the wiki code, clone the repository and execute the `wiki.go` file to start the web server.

```
$ cd $GOPATH
$ mkdir -p src/github.com/kevingimbel/ && cd src/github.com/kevingimbel/
$ git clone https://github.com/kevingimbel/gowiki
$ go run wiki.go
```

When everything's up and running, go visit [localhost:8080](http://localhost:8080). You should see a the Front Page. Go to [localhost:8080/view/NotAPage](http://localhost:8080/view/NotAPage) and you're redirected to the Edit Page for a page titles "NotAPage".

## Acknowledgment

The ["Writing Web Applications"](https://golang.org/doc/articles/wiki/) Tutorial from the Go Docs is where most of the code originates from. It's a great starting point into Go Web App Development.
