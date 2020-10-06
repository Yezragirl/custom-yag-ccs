{{$find := (index .CmdArgs 0)}}
{{$eventID := ""}}

{{range dbTopEntries "event" 100 0}}
{{$event := sdict}}{{range $k,$v := .Value}}{{$event.Set $k $v}}{{end}}
{{if eq (toInt $event.number) (toInt $find) }}
{{$event.number}} - {{$event.name}} - {{.UserID}}
{{$eventID = (toString .UserID)}}
{{end}}
{{end}}