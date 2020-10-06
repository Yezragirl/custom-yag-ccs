
{{$db := dbGet .User.ID "registration"}}
{{$reg := sdict}}{{range $k,$v := $db.Value}}{{$reg.Set $k $v}}{{end}}
 {{if hasRoleID 588749231150465049}}
{{$reg.Set "name" (reReplace "⭐" $reg.name "")}}
   {{editNickname $reg.name}}
   {{dbSet .User.ID "registration" $reg}}
   {{removeRoleID 588749231150465049}}
{{execCC 74 nil 0 ""}}
{{else}}
   {{$reg.Set "name" (reReplace "🌟" $reg.name "")}}
   {{editNickname $reg.name}}
   {{dbSet .User.ID "registration" $reg}}
   {{removeRoleID 573210843840249889}}
{{execCC 74 nil 0 ""}}
{{end}}{{end}}{{sendDM "You have stepped down as Admin of Yez's Ark Cluster."}}