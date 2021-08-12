{{/*xpcleanup clears people who are no longer here from the xp system*/}}
{{$skip := 0 -}}
{{$count := 0}}
{{if not .ExecData -}}
	{{$skip = 0}}
{{else -}}
	{{- $skip = .ExecData.skip -}}
{{end -}}
    {{- range dbTopEntries "XP" 20 $skip -}}
		{{- if (getMember .UserID) -}}
			{{ .User.Username}} is still here.
		{{else}}
			{{.UserID}} no longer here...clearing XP.
			{{dbDel .UserID "XP" -}}
		{{end}}
	{{- $count = add $count 1 -}}
    {{end -}}
  {{if ge  (toInt $count) 1 -}}
    {{execCC .CCID nil 10 (sdict "skip" (add $skip 20)) -}}
  {{end -}}
