#!/usr/bin/env bash

main  () {
	args_check "$@"
	echo "Hello, $1"
	exit 0
}

args_check () { 
	if [ $# -ne 1 ]; then
		echo "Usage: error_handling.sh <person>"
		exit 1
	fi
}

main "$@"

