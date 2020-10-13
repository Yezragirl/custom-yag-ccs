{{/*reaction add and remove
 Leaderboard Pagination*/}}
{{$entries := toInt (dbGet 0 "entries").Value}}
{{$message := .ReactionMessage}}
{{$page := 0}}
{{if $message.Embeds}}
{{if (eq (index $message.Embeds 0).Title "XP Leaderboard")}}
{{/*If this was triggered by a page turn reaction*/}}
{{$MsgID := .Reaction.MessageID}}
{{$page = toInt (index (split (index .ReactionMessage.Embeds 0).Footer.Text " ") 1)}}
    {{if (eq .Reaction.Emoji.Name "⬇️")}}
    {{$page = add $page 1}}
    {{else if (eq .Reaction.Emoji.Name "⬆️")}}
    {{$page = add $page -1}}
    {{end}}

{{$out := (printf "%s\t%-18s\t%-5s\t%-10s\n" "Rank" "Name" "Level" "XP")}}
{{$pagenum := $page}}
{{$currentlevel := 0}}
{{$nskip := mult $entries (add $page -1)}}
{{$rank := $nskip}}
    {{if lt $pagenum 1}}
    {{else}}
         {{range dbTopEntries "XP" $entries $nskip}}
         {{$rank = add $rank 1}}
         {{$currentlevel = roundFloor (sqrt (fdiv (toInt .Value) 100))}}
         {{$out = (joinStr "" $out (printf "#%2d:\t%-18s\t%-5d\t%-10d\n" $rank .User.Username (toInt $currentlevel) (toInt .Value)))}}
         {{end}}
{{if gt (len $out) 50}}
{{$ord1 := ""}}
{{$ord2 := ""}}
{{$customs := dict 1 "st" 2 "nd" 3 "rd"}}
{{$int1 := (toInt (add $nskip 1))}}
{{$first1 := div (toInt (mod $int1 100)) 10}}
{{$int2 := (toInt (add $nskip $entries))}}
{{$first2 := div (toInt (mod $int2 100)) 10}}
{{if eq $first1 1}}{{$ord1 = "th"}}
{{else if in (seq 1 4) (toInt (mod $int1 10))}}{{$ord1 = index $customs (toInt (mod $int1 10))}}
{{else}}{{$ord1 = "th"}}
{{end}}
{{if eq $first2 1}}{{$ord2 = "th"}}
{{else if in (seq 1 4) (toInt (mod $int2 10))}}{{$ord2 = index $customs (toInt (mod $int2 10))}}
{{else}}{{$ord2 = "th"}}
{{end}}
{{$start := (joinStr "" $int1 $ord1)}}{{$end := (joinStr "" $int2 $ord2)}}


{{$embed := cembed 
"title" "XP Leaderboard" 
"description" (joinStr "" "*Our " $start " through " $end " Leaders*\n\n```" $out "```") 
"color" 4645612 
"thumbnail" (sdict "url" "https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/120/google/223/sports-medal_1f3c5.png")
"footer" (sdict "text" (joinStr "" "Page " $pagenum))}}
	{{editMessage nil $MsgID $embed}}
{{end}}
{{end}}
{{end}}

{{end}}



