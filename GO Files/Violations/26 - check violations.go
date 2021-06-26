{{$args := parseArgs 1 "correct syntax is `checkv rule#"
    (carg "int" "Rule Number")}}
    {{$rules := cslice "1" "2" " 3A" "3B" "3C" "4" "5" "6" "7" "8" "9" "10" "11" "12" "13" "14" "15"}}
{{$r := (toString ($args.Get 0))}}
{{if not (in $rules $r)}}
{{sendMessage nil (joinStr "" .User.Mention ", thats not one of our rules.\nTry again. Use one of the following rule codes:\n1: Earn It\n2: Age\n3A: Player Registration\n3B: Tribe Registration\n3C: Base Registration\n4: Harrassment\n5: Drama/Attitude\n6: Unsportsmanlike Conduct\n7: Building In Render\n8: Base Size/Count\nRule 9: Piping\n10: Mass Breeding\n11: Teleporters\n12: Traps\n13: Dinos Out\n14: Disrespecting an Admin\n15: DMing Admins")}}
{{else}}
{{$msg := ""}}{{/* but let's just send one message instead of 3 */}}
{{ range seq 1 6 }}
{{$q = dbGet ((toInt64 $r) (joinStr "" "Tier" .)).Value}}
{{$f = dbGet ((toInt64 $r) (joinStr "" "Tier" . "Fine")).Value}}
{{$w = dbGet ((toInt64 $r) (joinStr "" "Tier" . "Warn")).Value}}
{{$msg = joinStr "\n" $msg (joinStr "" "Rule " $r ", Tier " . " set to " $q " violations, a fine of " $f " and the following warning: " $w)}}
{{end}}
{{sendMessage nil $msg}}
{{end}}
