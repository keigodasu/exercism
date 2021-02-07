#!/usr/bin/env bash

while getopts :nlivx option; do
  case "${option}" in
    n) liner=1 ;;
    l) filename=1 ;;
    i) shopt -s nocasematch ;;
    v) invert=1 ;;
    x) entire=1 ;;
    \?) echo Invalid option.; exit 1;;
  esac
done

shift $((OPTIND-1))
pattern=$1
shift
files=("$@")

((entire)) && pattern="^$pattern$"
for file in "${files[@]}"; do
  count=0
  while read -r line
  do
    (( count++ ))
    out=
    if (( invert )); then
      ! [[ ${line} =~ ${pattern} ]] && out="$line"
    else
	[[ ${line} =~ ${pattern} ]] && out="$line"
    fi
    if (( filename )) && [[ -n $out ]]; then echo $file; break; fi
    if (( liner )) && [[ -n $out ]]; then out="$count":$out; fi
    if (( ${#files[@]} > 1 )) && [[ -n $out ]]; then out="$file":$out; fi
    [[ -n $out ]] && echo $out;
  done < "$file"
done

exit 0
