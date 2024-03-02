{{/*xpresetall allows for a complete reset of the xp system*/}}
{{$db := dbTopEntries "XP" 5 0}}
{{if $db}}
	{{range $db}}
	{{dbDel .User.ID "XP"}}
	{{end}}
{{execCC 32 nil 0 nil}}
{{else}}
All XP Cleared. 
{{end}}
