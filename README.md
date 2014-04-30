sass-git
========

Shell script for compiling sass, adding files to git, committing with a message, and optionally pushes. Currently errors for whatever stage means that there is no need to do those things (Your repo is up to date). I will specify errors eventually.



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
