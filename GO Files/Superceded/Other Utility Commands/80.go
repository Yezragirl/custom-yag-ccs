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
{{else}}
	{{$name = (joinStr "" "✨ " $name)}}   
	{{editNickname $name}}
	{{addRoleID 588749231150465049}}{{/*Provisional*/}}
	{{removeRoleID 641285945584517131}}{{/*Applicant*/}}
	{{sendMessage nil "Congratulations! You are now a Provisional Admin!"}}
	{{$args := parseArgs 3 "correct syntax is `ap @who amount reason`"
    (carg "user" "admin")
    (carg "int" "amount")
    (carg "string" "reason")}}
{{dbSet .User.ID "AdminPoints" 10}}
{{sendMessage 653058600230584353 (joinStr "" $name " was given 10 starting points.")}}
{{sendMessage 654151821039894547 (joinStr "" .User.Mention " was given 10 starting points.")}}
{{sendMessage 654151821039894547 (joinStr "" .User.Mention " now has 10 points.")}}










	{{end}}
{{dbSet .User.ID "name" $name}}
{{execCC 74 nil 0 ""}}
{{end}}



