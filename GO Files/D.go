{{if (eq .Channel.ParentID 635305499600093214)}} 
{{if hasRole 634598489732546588}}
{{else}}
{{$db := dbGet .User.ID "registration"}}
 {{$reg := sdict}}{{range $k,$v := $db.Value}}{{$reg.Set $k $v}}{{end}}
 {{if hasRole 634598489732546588}}
  {{if not $db}}
You can't edit your registration. You don't have one yet. 
{{else}}
      {{if eq $reg.Q "EG"}}
         {{$reg.Set "gt" .Message.Content}}
         {{$reg.Set "Q" "EE"}}
         {{dbSet .User.ID "registration" $reg}}
	 {{else if eq $reg.Q "EN"}}
         {{$reg.Set "name" .Message.Content}}
         {{$reg.Set "Q" "EE"}}
         {{dbSet .User.ID "registration" $reg}}
      {{else if eq $reg.Q "ET"}}
         {{$reg.Set "tribe" .Message.Content}}
         {{$reg.Set "Q" "EE"}}
         {{dbSet .User.ID "registration" $reg}}
      {{else if eq $reg.Q "EP"}}
         {{$reg.Set "pin" .Message.Content}}
         {{$reg.Set "Q" "EE"}}
         {{dbSet .User.ID "registration" $reg}}
      {{else if eq $reg.Q "EPG"}}
         {{$reg.Set "guardian" .Message.Content}}
         {{$reg.Set "Q" "EE"}}
	 {{else if eq $reg.Q "EA"}}
		 {{$msg := sendMessageRetID nil (joinStr "" .User.Mention ", did you have a birthday?")}}
		 {{addMessageReactions nil $msg ":yes:658376626639339533" ":no:658376639322783745"}}
		 {{end}}
	{{$msg := sendMessageRetID nil (joinStr "" .User.Mention ", got it. Do you want to edit something else?")}}
	{{addMessageReactions nil $msg ":yes:658376626639339533" ":no:658376639322783745"}} 
  {{end}}
{{end}}{{end}}




{{ScheduleUniqueCC }}