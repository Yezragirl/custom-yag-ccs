{{$list = ""}}
{{$args := parseArgs 1 "correct syntax is `-tribe TribeName`"
(carg "string" "Tribe Name")}}
{{$tribe := ($args.Get 0)}}
{{range dbTopEntries "tribe" 100 0}}
{{if reFindAllSubmatches $regex .Value}}
{{$list = joinStr "" $list "\n`" (getMember .UserID).Nick "`"}}
{{end}}
{{$msg := sendMessageRetID nil (cembed "title" (joinStr "" "Players in Tribe " ($args.Get 0)) "description" $list)}}
