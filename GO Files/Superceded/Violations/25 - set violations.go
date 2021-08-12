{{$args := parseArgs 4 "correct syntax is `setv rule# tier #-of-violations Fine text`"
    (carg "int" "Rule Number")
    (carg "int" "Tier")
    (carg "int" "Violations")
    (carg "int" "Fine")
    (carg "string" "Text")}}
{{$tiers := cslice 1 2 3 4 5}}
{{$rules := cslice "1" "2" " 3A" "3B" "3C" "4" "5" "6" "7" "8" "9" "10" "11" "12" "13" "14" "15"}}
{{$r := (toString($args.Get 0))}}
{{$n := $args.Get 1}}
{{if not (in $tiers $n)}}
{{sendMessage nil (joinStr "" .User.ID ", there are only 5 tiers.")}}
{{else}}
{{if not (in $rules $r)}}
{{sendMessage nil (joinStr "" .User.Mention ", thats not one of our rules.\nTry again. Use one of the following rule codes:\n1: Earn It\n2: Age\n3A: Player Registration\n3B: Tribe Registration\n3C: Base Registration\n4: Harrassment\n5: Drama/Attitude\n6: Unsportsmanlike Conduct\n7: Building In Render\n8: Base Size/Count\nRule 9: Piping\n10: Mass Breeding\n11: Teleporters\n12: Traps\n13: Dinos Out\n14: Disrespecting an Admin\n15: DMing Admins")}}
{{else}}
{{$Q := (joinStr "" "Tier" $n)}}
{{$F := (joinStr "" "Tier" $n "Fine")}}
{{$W := (joinStr "" "Tier" $n "Warn")}}
{{dbSet (toInt64 $r) $Q ($args.Get 2)}}
{{dbSet (toInt64 $r) $F ($args.Get 3)}}
{{dbSet (toInt64 $r) $W ($args.Get 4)}}
{{sendMessage nil (joinStr "" "Rule " $r ", Tier " $n " set to " ($args.Get 2) " violations, a fine of " ($args.Get 3) " and the following warning: " ($args.Get 4))}}
{{end}}
{{end}}