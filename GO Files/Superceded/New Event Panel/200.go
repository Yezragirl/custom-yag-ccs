{{$eventID := ""}}
{{$channel := "580400941048528900"}}
{{$date := currentTime}}
{{$find := (toInt (index .CmdArgs 0))}}
{{$find}}
{{range dbTopEntries "event" 100 0}}
{{$event := sdict}}{{range $k,$v := .Value}}{{$event.Set $k $v}}{{end}}
{{if eq (toInt $event.number) $find}}
{{$event.number}} - {{$event.name}} - {{.UserID}} this is the right event. 
{{end}}
{{end}}
