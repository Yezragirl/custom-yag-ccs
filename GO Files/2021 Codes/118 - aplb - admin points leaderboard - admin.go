{{/* Admin Points Leaderboard */}}
{{$out := (printf "%s\t%-20s\t%-10s\n" "Rank" "Name" "Points")}}
{{$rank := 0}}
{{$name := ""}}
{{$entries := 10}}
{{range dbTopEntries "AdminPoints" $entries 0}}
{{if targetHasRoleID .UserID 588749231150465049}}
{{$rank = add $rank 1}}
{{$out = (joinStr "" $out (printf "#%2d:\t%-20s\t%-5d\n" $rank .User.Username (toInt .Value)))}}
{{else}}
{{dbDel .UserID "AdminPoints"}}{{end}}
{{end}}
{{$embed := cembed 
"title" "Admin Point Leaderboard" 
"description" (joinStr "" "*Our Top " $rank " Leaders*\n\n```" $out "```") 
"color" 4645612}}
{{$msgid := sendMessageRetID nil $embed}}