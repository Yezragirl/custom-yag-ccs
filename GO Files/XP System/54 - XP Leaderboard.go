{{/*command lb*/}}
{{$out := (printf "%s\t%-18s\t%-5s\t%-10s\n" "Rank" "Name" "Level" "XP")}}
{{$rank := 0}}
{{$name := ""}}
{{$pagenum := 1}}
{{$currentlevel := 0}}
{{$entries := toInt (dbGet 0 "entries").Value}}
{{range dbTopEntries "XP" $entries 0}}
{{$rank = add $rank 1}}
{{$currentlevel = roundFloor (sqrt (fdiv (toInt .Value) 100))}}

{{$out = (joinStr "" $out (printf "#%2d:\t%-18s\t%-5d\t%-10d\n" $rank .User.Username (toInt $currentlevel) (toInt .Value)))}}

{{end}}
{{$embed := cembed 
"title" "XP Leaderboard" 
"description" (joinStr "" "*Our Top " $rank " Leaders*\n\n```" $out "```") 
"color" 4645612 
"thumbnail" (sdict "url" "https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/120/google/223/sports-medal_1f3c5.png")
"footer" (sdict "text" (joinStr "" "Page " $pagenum))}}
	{{$msgid := sendMessageRetID nil $embed}}
{{addMessageReactions nil $msgid "⬆️" "⬇️"}}
