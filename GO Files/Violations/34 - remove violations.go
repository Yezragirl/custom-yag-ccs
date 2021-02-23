{{$args := parseArgs 2 "correct syntax is `remv rule# who`"
    (carg "int" "Rule Number")
    (carg "user" "Violator")}}
{{$countofviolations := dbIncr ($args.Get 1).ID ($args.Get 0) -1}}
{{sendMessage nil (joinStr "" "<@" ($args.Get 1).ID ">, 1 violation of rule " ($args.Get 0) " has been removed.")}}
