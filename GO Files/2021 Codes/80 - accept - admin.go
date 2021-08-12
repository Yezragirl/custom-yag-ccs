{{/* accept Accepting Promotion */}}
{{if eq .User.ID (toInt (dbGet 0 "Promote").Value)}}
{{$name := (dbGet .User.ID "name").Value}}
	{{if $name}}
	{{else}}
	{{$name = (getMember .User.ID).Nick}}
	{{end}}

{{if hasRoleID 588749231150465049}}{{/*Provisional*/}}
	{{$name = (reReplace "✨" $name "🌟")}}
	{{editNickname $name}}
	{{addRoleID 573210843840249889}}{{/*Admin*/}}
	{{removeRoleID 588749231150465049}}{{/*Provisional*/}}
	{{sendMessage nil "Congratulations! You are now a Full Admin!"}}
{{else if hasRoleID 573210843840249889}}{{/*Admin*/}}
	{{$name = (reReplace "🌟" $name "💫")}}
	{{editNickname $name}}
	{{addRoleID 760153772000411739}}{{/*Senior Admin*/}}
	{{sendMessage nil "Congratulations! You are now a Senior Admin!"}}
{{else if hasRoleID 641285945584517131}}
	{{$name = (joinStr "" "✨ " $name)}}   
	{{editNickname $name}}
	{{addRoleID 588749231150465049}}{{/*Provisional*/}}
	{{removeRoleID 641285945584517131}}{{/*Applicant*/}}
	{{sendMessage nil "Congratulations! You are now a Provisional Admin!"}}
{{end}}
{{dbSet .User.ID "name" $name}}
{{execCC 74 nil 0 ""}}
{{end}}
