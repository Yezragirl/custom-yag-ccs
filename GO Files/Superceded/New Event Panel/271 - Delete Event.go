{{$find := (index .CmdArgs 0)}}
{{$channel := "580400941048528900"}}

{{$EventID := ""}}
{{$runner := cslice}}
{{$participants := cslice}}
{{$waitlist := cslice}}
{{range dbTopEntries "event" 100 0}}
{{$event := sdict}}{{range $k,$v := .Value}}{{$event.Set $k $v}}{{end}}
{{if eq (toInt $event.number) (toInt $find) }}
{{$EventID = (toString .UserID)}}
{{end}}{{end}}
{{$db := dbGet (toInt64 $EventID) "event"}}
{{$event := sdict}}{{range $k,$v := $db.Value}}{{$event.Set $k $v}}{{end}}
{{$number := $event.number}}

{{deleteMessage $channel $EventID}}
{{cancelScheduledUniqueCC 198  (joinStr "" $EventID "-1")}}
{{cancelScheduledUniqueCC 198 (joinStr "" $EventID "-2")}}
{{sendMessage nil (joinStr "" "Event " $number " deleted." )}}
{{$runner = (cslice).AppendSlice $event.runner}}
{{$participants = (cslice).AppendSlice $event.participants}}
{{$waitlist = (cslice).AppendSlice $event.waitlist}}

{{range $runner}}
{{cancelScheduledUniqueCC 197 (joinStr "" $EventID "-" . "-runner")}}
{{end}}
{{range $participants}}
{{cancelScheduledUniqueCC 197 (joinStr "" $EventID "-" . "-participant")}}
{{end}}
{{range $waitlist}}
{{cancelScheduledUniqueCC 197 (joinStr "" $EventID "-" . "-waitlist")}}
{{end}}

{{dbDel (toInt64 $EventID) "event"}}