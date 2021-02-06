#!/usr/bin/env bash

main () {
  declare -A colors=([black]=0 [brown]=1 [red]=2 [orange]=3 [yellow]=4 [green]=5 [blue]=6 [violet]=7 [grey]=8 [white]=9) 
  declare -A units=( [black]=" " [brown]="0 " [red]="00 " [orange]=" kilo" [yellow]="0 kilo" [green]="00 kilo" [blue]=" mega" [violet]="0 mega" [grey]="00 mega" [white]=" giga" )
  declare -A alts=( [black]="0 " [brown]="00 " [red]=" kilo" [orange]="0 kilo" [yellow]="00 kilo" [green]=" mega" [blue]="0 mega" [violet]="00 mega" [grey]=" giga" [white]="0 giga" )

 [[ -z "${colors[${1,,}]}" || -z "${colors[${2,,}]}"  || -z "${units[${3,,}]}" ]] && echo "invalid color" && exit 1

 [[ ${1,,} != "black" ]] && init=${colors[${1,,}]} || init=""
 [[ ${2,,} != "black" ]] && rest=${colors[${2,,}]}${units[${3,,}]} || rest=${alts[${3,,}]}
 echo "$init${rest}ohms"

}

main "$@"
