#!/bin/bash
hugo
sudo rm -R /usr/share/nginx/blog/*
sudo cp -R public/* /usr/share/nginx/blog/

