{{$args := parseArgs 2 "correct syntax is `clearv rule# who`"
    (carg "int" "Rule Number")
    (carg "userid" "Violator")}}
{{dbSet ($args.Get 1) (toString ($args.Get 0)) 0}}