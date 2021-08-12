{{/*removev removes a violation*/}}
{{$args := parseArgs 2 "correct syntax is `removev rule# who`"
    (carg "int" "Rule Number")
    (carg "user" "Violator")}}
{{$countofviolations := dbIncr ($args.Get 1).ID (toString ($args.Get 0)) -1}}
{{sendMessage nil (joinStr "" "<@" ($args.Get 1).ID ">, 1 violation of rule " (toString ($args.Get 0)) " has been removed.")}}
