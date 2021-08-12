{{/*eventdb Lists all events in DB */}}
{{$eventID := ""}}
{{$channel := "573897276569813012"}}
{{$date := currentTime}}

{{range dbTopEntries "event" 2 0}}
{{- $event := sdict -}}{{- range $k,$v := .Value -}}{{- $event.Set $k $v -}}{{- end -}}
{{- $eventID = (toString .UserID) -}}
{{.UserID}}
{{.Value}}
{{end}}
