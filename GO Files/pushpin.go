{{$pattern := `([^-]+-[^-]+)(.+)?` -}}
{{$db := dbGet 0 (toString .Channel.ID)}}
{{$ticket := sdict}}{{range $k,$v := $db.Value}}{{$ticket.Set $k $v}}{{end}}

{{if eq .Channel.ParentID 635305499600093214}}
{{if and (eq .Reaction.Emoji.ID 660605488609755146) (eq (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 2) "")}}
{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has claimed ticket " .Channel.Name )}}
{{editChannelName nil (joinStr "" (getChannel nil).Name "-" .Member.Nick)}}

{{else if (eq .Reaction.Emoji.ID 660605488609755146)}}

	{{if eq $ticket.claimer .User.ID}}
	{{editChannelName nil (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1)}}
	{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has unclaimed ticket " .Channel.Name )}}
	{{dbDel 0 (toString .Channel.ID)}}

	{{else}}

	{{editChannelName nil (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1)}}
	{{dbSet 0 .Channel.ID .User.ID}}
	{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has unclaimed ticket " .Channel.Name )}}
	{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has claimed ticket " .Channel.Name )}}
	{{editChannelName nil (joinStr "" (getChannel nil).Name "-" .Member.Nick)}}
	{{end}}


{{else if and (eq .Reaction.Emoji.ID 660605465721569334) ($db)}}
{{editChannelName nil (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1)}}
{{dbDel 0 .Channel.ID}}
{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has closed ticket " .Channel.Name )}}
{{end}}
{{end}}