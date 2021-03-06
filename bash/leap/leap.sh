#!/usr/bin/env bash

main () {
  arg_validate "$@"
  if [ $(($1 % 4)) -eq 0 ] && [ $(($1 % 100)) -ne 0 ] || [ $(($1 % 400)) -eq 0 ]; then
    echo "true" 
  else
    echo "false"
  fi
}

arg_validate () {
  expr $1 + 1 > /dev/null 2>&1
  RET=$?
  if [ $# -ne 1 ] || [ $RET -gt 1 ]; then
	  echo "Usage: leap.sh <year>"
	  exit 1
  fi
}

main "$@"
