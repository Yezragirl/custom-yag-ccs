{{$args := parseArgs 1 "Promote Who"
	(carg "userid" "User")}}
{{$target := $args.Get 0}}
{{db.Set 0 "Promote" $target}}
{{sendMessage nil (joinStr "" $target.Mention " you have been approved for promotion. Type 'accept' to accept your promotion." )}}




{{if eq .User.ID (toInt (dbGet 0 "Promote").Value)}}
{{$db := dbGet .User.ID "registration"}}
{{$reg := sdict}}{{range $k,$v := $db.Value}}{{$reg.Set $k $v}}{{end}}
 {{if hasRoleID 588749231150465049}}
{{$reg.Set "name" (reReplace "⭐" $reg.name "🌟")}}
   {{editNickname $reg.name}}
   {{dbSet .User.ID "registration" $reg}}
   {{addRoleID 573210843840249889}}
   {{removeRoleID 588749231150465049}}
{{execCC 74 nil 0 ""}}
{{sendMessage nil "Congratulations! You are now a full admin!"}}
{{else}}
   {{$reg.Set "name" (joinStr "" "⭐ " $reg.name)}}   
   {{editNickname $reg.name}}
   {{dbSet .User.ID "registration" $reg}}
   {{addRoleID 588749231150465049}}
{{execCC 74 nil 0 ""}}
{{sendMessage nil "Congratulations! You are now a provisional admin!"}}
{{end}}{{end}}