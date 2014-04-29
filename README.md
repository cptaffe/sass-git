sass-git
========

shell script for compiling sass, adding files to git, committing with a message, and pushing to a server.

## Install

Open a terminal and type
```sh
git clone http://github.com/cptaffe/sass-git
cd sass-git
cp sass-git /usr/local/bin
```

## Usage

```sh
$ sass-git 404
```
This will compile 404.scss to 404.css, add those files to git, git commit with a default message, and push.

To add a custom commit message use:
```sh
# optional custom message
$ sass-git 404 -m "Updated CSS for 404 page"
```
