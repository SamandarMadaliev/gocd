#!/bin/sh

folders=(
    cmd
    internal
    pkg
)
rootPath="./test/"
for folder in "${folders[@]}"; do
  mkdir -p "$rootPath$folder"
  chmod 777 "$rootPath$folder"
done
