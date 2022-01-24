{{/*clearcooldown*/}}
{{$db := dbTopEntries "pause" 5 0}}
{{if $db}}
	{{range $db}}
	{{dbDel .User.ID "pause"}}
	{{end}}
{{execCC 61 nil 0 nil}}
{{else}}
All Cooldowns Cleared. 
{{end}}
