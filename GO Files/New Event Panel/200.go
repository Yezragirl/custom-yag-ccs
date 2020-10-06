{{$eventID := ""}}
{{$channel := "573897276569813012"}}
{{$date := currentTime}}

{{range dbTopEntries "event" 100 0}}
{{$event := sdict}}{{range $k,$v := .Value}}{{$event.Set $k $v}}{{end}}
{{if $event.time.Before $date}}
{{$event.number}} - {{$event.name}} - {{.UserID}} is old.
{{$eventID = (toString .UserID)}}
{{end}}
{{end}}