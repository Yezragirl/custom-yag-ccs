{{deleteResponse 5}}
{{deleteTrigger 5}}
{{if .CmdArgs}}
{{dbSet .User.ID "afk" (joinStr " " .CmdArgs)}}
{{.User.Mention}}, set your AFK to `{{joinStr " " .CmdArgs}}`!
{{$name := .User.Username}}
{{if .Member.Nick}}
  {{$name = .Member.Nick}}
{{end}}
{{editNickname (joinStr " " "[AFK]" $name)}}
{{end}}


{{$adjustedtime := currentTime.sub (toDuration (mult 5 .TimeMinute))
	{{$db := dbGet .User.ID "afk"}}
	{{if and $db (gt $db.CreatedAt (currentTime.Sub (toDuration (mult -5 .TimeMinute))))}}
	{{dbDel .User.ID "afk"}}
	{{.User.Mention}}, removed AFK!
	{{ $matches := reFindAllSubmatches `(AFK)` .Member.Nick }}
	{{ if $matches }}
	{{editNickname (reReplace "(\\[AFK\\] )" .Member.Nick "")}}
	{{end}}
	{{end}}


{{$db := dbGet .User.ID "afk"}}
{{if $db}}
{{dbDel .User.ID "afk"}}
{{.User.Mention}}, removed AFK!
{{ $matches := reFindAllSubmatches `(AFK)` .Member.Nick }}
{{ if $matches }}
{{editNickname (reReplace "(\\[AFK\\] )" .Member.Nick "")}}
{{end}}
{{end}}