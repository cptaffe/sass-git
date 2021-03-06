sass-git
========

Golang utility for compiling sass, adding files to git, committing with a message, and optionally pushes. View its [project page](http://cptaffe.github.io/sass-git).

## Install

Open a terminal and type
```sh
git clone http://github.com/cptaffe/sass-git
cd sass-git/
go build sass-git.go
cp sass-git /usr/local/bin
```

## Usage

```sh
$ sass-git 404
```
This will compile 404.scss to 404.css, add those files to git, git commit with a default message, and push.

To add a custom commit message or push use:
```sh
# optional custom message
$ sass-git 404 -m "Updated CSS for 404 page"
# pushing to server
$ sass-git 404 -p -m "Updated CSS for 404 page"
```

## License

If you've never seen an MIT license, check out the LICENSE.

## Authors

+ Connor Taffe (@cptaffe)
