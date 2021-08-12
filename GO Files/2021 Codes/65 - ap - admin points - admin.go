{{/*Admin Points Adding Command*/}}
{{$args := parseArgs 3 "correct syntax is `ap @who amount reason`"
    (carg "user" "admin")
    (carg "int" "amount")
    (carg "string" "reason")}}
{{$points := (dbGet ($args.Get 0).ID "AdminPoints").Value}}
{{$points = add $points ($args.Get 1)}}
{{sendMessage 653058600230584353 (joinStr "" (getMember ($args.Get 0).ID).Nick " earned " (toString (toInt ($args.Get 1))) " points for " ($args.Get 2) ".")}}
{{sendMessage nil (joinStr "" ($args.Get 0).Mention " earned " (toString (toInt ($args.Get 1))) " points for " ($args.Get 2) ".")}}
{{sendMessage nil (joinStr "" ($args.Get 0).Mention " now has " (toString (toInt $points)) " points.")}}
{{dbSet ($args.Get 0).ID "AdminPoints" (toInt $points)}}

