{{$id := dbGet 0 "ready-detection"}}{{/*Retrieve Ready Detection Message ID*/}}
{{if $id}} {{/*Does Ready Detection Exist?*/}}
	{{if eq .ReactionMessage.ID (toInt64 $id.Value)}}{{/*Was this reaction to the ready detection embed?*/}}
		{{if eq .Reaction.Emoji.ID 734019999802458164}}{{/*Was this reaction the check mark?*/}}
			{{addRoleID newRoleID}}{{/*Give Ready Role*/}}
			{{removeRoleID newRoleID 21600}}{{/*Ready Role removed after 6 hours*/}}
			{{sendMessage }}
		{{else if eq .Reaction.Emoji.ID 734027152676356157}}{{/*Was this reaction the red Cross?*/}}
			{{removeRoleID newRoleID}}{{/*Remove Ready Role*/}}
		{{end}}
	{{end}}
{{end}}