#!/usr/bin/env bash

: << !
Name: sync-hook.sh
Author: YJ
Email: yj1516268@outlook.com
Created Time: 2024-04-16 16:51:07

Description: 从 System 同步 git 钩子

Attentions:
-

Depends:
-
!

script_dir=$(dirname "$0")                              # 本脚本所在路径
hooks_dir="${script_dir}/hooks"                         # git 钩子路径
system_hooks_dir="${HOME}/Documents/Repos/System/hooks" # 本地 System 存储库 git 钩子路径
repo_name="https://github.com/YHYJ/System.git"          # 云端 System 存储库路径

if [ -d "${system_hooks_dir}" ]; then
  echo "Copying hooks from 'System' repo"
  mkdir -p "${hooks_dir}"
  cp -r "${system_hooks_dir}/"* "${hooks_dir}"
  echo "Hooks copied successfully."
else
  echo "Downloading hooks from 'System' repo"
  git clone "${repo_name}" temp
  mkdir -p "${hooks_dir}"
  cp -r temp/hooks/* "${hooks_dir}"
  rm -rf temp
  echo "Hooks downloaded successfully."
fi
