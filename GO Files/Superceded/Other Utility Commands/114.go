{{$name := (dbGet .User.ID "name").Value}}

{{if hasRoleID 588749231150465049}}{{/*Provisional*/}}
	{{$name = (reReplace "✨ " $name "")}}
	{{removeRoleID 588749231150465049}}{{/*Provisional*/}}
{{else if hasRoleID 573210843840249889}}{{/*Admin*/}}
	{{$name = (reReplace "🌟 " $name "")}}
	{{removeRoleID 573210843840249889}}{{/*Admin*/}}
{{else if hasRoleID 760153772000411739}}{{/*Senior Admin*/}}
	{{$name = (reReplace "💫 " $name "")}}   
	{{removeRoleID 760153772000411739}}{{/*Senior Admin*/}}
	{{removeRoleID 573210843840249889}}{{/*Admin*/}}
	{{sendMessage nil "Congratulations! You are now a Provisional Admin!"}}
{{end}}
{{editNickname $name}}
{{dbSet .User.ID "name" $name}}
{{removeRoleID 634598489732546588}}{{/*Support Team*/}}
{{execCC 74 nil 0 ""}}
{{sendDM "You have stepped down as Admin of Yez's Ark Cluster."}}
{{sendMessage 653058600230584353 (joinStr "" .User.Mention ", has stepped down as admin.")}}
{{sendMessage 573897276569813012 (joinStr "" .User.Mention "***, has stepped down as admin.***")}}
{{deleteTrigger 1}}

