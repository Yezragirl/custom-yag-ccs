{{$args := parseArgs 1 "correct syntax is `checkv rule#"
    (carg "int" "Rule Number")}}
{{$rules := cslice 48 56 7 9 14 16 19}}
{{$r := $args.Get 0}}
{{if not (in $rules $r)}}
{{sendMessage nil (joinStr "" .User.Mention ", thats not one of our rules.\nTry again. Use one of the following rule codes:\n48 - Drama/Bad Attitude/Trolling/kiting/land grabbing/stealing\n56 - Building in render/blocking resources\n7 - Mass breeding/dinos on aggressive\n9 - teleporters\n14 - traps\n16 - base size/count\n19 disrespect")}}
{{else}}
{{$db := dbGet (toInt $r) "ruletierinfo"}}
{{if not $db}}
message: oops, no db for that rule
{{else}}
{{$data := sdict}}
{{range $k, $v := $db.Value}}
{{$data.Set $k $v}}
{{end}}
{{$msg := ""}}{{/* but let's just send one message instead of 3 */}}
{{ range seq 1 4 }}
{{$q = $data.Get (joinStr "" "Tier" .)}}
{{$f = $data.Get (joinStr "" "Tier" . "Fine")}}
{{$w = $data.Get (joinStr "" "Tier" . "Warn")}}
{{$msg = joinStr "\n" $msg (joinStr "" "Rule " $r ", Tier " . " set to " $q " violations, a fine of " $f " and the following warning: " $w)}}
{{end}}
{{sendMessage nil $msg}}
{{end}}
{{end}}