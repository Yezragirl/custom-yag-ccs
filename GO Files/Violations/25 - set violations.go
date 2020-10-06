{{$args := parseArgs 4 "correct syntax is `setv rule# tier #-of-violations Fine text`"
    (carg "int" "Rule Number")
    (carg "int" "Tier")
    (carg "int" "Violations")
    (carg "int" "Fine")
    (carg "string" "Text")}}
{{$tiers := cslice 1 2 3}}
{{$rules := cslice 48 56 7 9 14 16 17 19 20}}
{{$r := $args.Get 0}}
{{$n := $args.Get 1}}
{{if not (in $tiers $n)}}
{{sendMessage nil (joinStr "" .User.ID ", there are only 3 tiers.")}}
{{else}}
{{if not (in $rules $r)}}
{{sendMessage nil (joinStr "" .User.Mention ", thats not one of our rules.\nTry again. Use one of the following rule codes:\n48 - Drama/Bad Attitude/Trolling/kiting/land grabbing/stealing\n56 - Building in render/blocking resources\n7 - Mass breeding/dinos on aggressive\n9 - teleporters\n14 - traps\n16 - base size/count\n17 Base Registration\n19 disrespect\n20 Messaging Admin Staff")}}
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