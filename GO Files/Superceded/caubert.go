count from 0 to N (inclusive), call each number you count i and write down the result using N^2-2*i
{{$N := (toInt (index .CmdArgs 0))}}
{{$countto := (toInt add $N 1)}}

{{range seq 0 $countto}}
i = {{.}}
{{$N}}^2-2*i ={{sub (mult $N $N) (mult 2 .) }}
{{end}}






{{range dbTopEntries "XP" $entries 0}}
{{if (getMember .User.ID)}}
{{(getMember .User.ID).Nick}}
{{else}}
error
{{end}}