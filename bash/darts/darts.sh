#!/usr/bin/env bash

main() {
  (( $# == 2 )) || fail
  str::isFloat "$1" && str::isFloat "$2"  || fail

  score "$@"
}

fail() {
  echo "failed"
  exit 1
}

score() {
  local bools fmt

  printf -v fmt 'sqrt(%s^2 + %s^2) <= %%d\n' "$@"
 mapfile -t bools < <(printf "$fmt" 1 5 10 | bc -l)

  case "${bools[*]}" in
      "1 1 1") echo 10 ;;
      "0 1 1") echo  5 ;;
      "0 0 1") echo  1 ;;
      "0 0 0") echo  0 ;;
  esac
}

str::isFloat() {
    [[ $1 == ?([-+])+([[:digit:]])?(.*([[:digit:]])) ]] ||
    [[ $1 == ?([-+])*([[:digit:]]).+([[:digit:]]) ]]
}

main "$@"