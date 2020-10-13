{{/*command set*/}}
{{$args := parseArgs 2 "correct syntax is `xpset setting new-value`"
    (carg "string" "setting")
    (carg "int" "value")}}
{{dbSet 0 (lower ($args.Get 0)) ($args.Get 1)}}
{{sendMessage nil (joinStr "" ($args.Get 0) " set to " ($args.Get 1))}}
