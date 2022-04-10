#!/bin/bash
cd "$(dirname "$0")"
git pull
hugo
rsync -avu --delete public/ /usrshare/nginx/blog
