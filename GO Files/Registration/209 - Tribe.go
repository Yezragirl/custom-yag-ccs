{{$list := ""}}
{{$tribe := (joinStr " " .CmdArgs)}}
{{$regex := print `(?i)` $tribe}}
{{range dbTopEntries "tribe" 100 0}}
{{if reFindAll $regex .Value}}
{{$list = (joinStr "" $list "\n`" (getMember .UserID).Nick "`")}}
{{end}}{{end}}
{{$msg := sendMessageRetID nil (cembed "title" (joinStr "" "Players in Tribe " $tribe) "description" $list)}}
