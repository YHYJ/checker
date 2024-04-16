#!/usr/bin/env bash

: << !
Name: create-hook-link.sh
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2023-09-18 13:37:04

Description: 创建 git 钩子

Attentions:
-

Depends:
-
!

hook_work_dir=".git/hooks"
mkdir -p "$hook_work_dir"

hook_base_dir="hooks"
path=$(ls "$hook_base_dir")
for file in $path; do
  ln -sf ../../$hook_base_dir/"$file" "$hook_work_dir/$file"
done
