{{$minimum := toInt (dbGet 0 "minimum").Value}}
{{$maximum := toInt (dbGet 0 "maximum").Value}}
{{$cooldown := toInt (dbGet 0 "cooldown").Value}}
{{if not (dbGet .User.ID "pause")}}
{{dbSetExpire .User.ID "pause" "cooldown" $cooldown}}
{{$xp := toInt ((dbGet .User.ID "XP").Value)}}
{{$currentlevel := roundFloor (sqrt (fdiv $xp 100))}}
{{$amount := randInt 15 25}}
{{$newxp := (toString (add (toInt (dbGet .User.ID "XP").Value) $amount))}}
{{dbSet .User.ID "XP" $newxp}}
{{$newlevel := roundFloor (sqrt (fdiv (toInt $newxp) 100))}}
{{$roles := .Member.Roles}}
{{$pos := 0}}
{{$color := 0}}
{{range .Guild.Roles}}
{{if in $roles .ID}}
{{if and (lt $pos .Position) (.Color)}}
{{$pos = .Position}}
{{$color = .Color}}
{{end}}
{{end}}
{{end}}
{{if gt $newlevel $currentlevel}}
{{giveRoleName .User.ID (joinStr "" "Level " (toInt $newlevel))}}
   {{if gt (toInt $newlevel) 5}}
    {{takeRoleName .User.ID (joinStr "" "Level " (sub (toInt $newlevel) 5))}}
    {{end}}
{{$embed := (cembed 
"title" "You leveled up!" 
"color" $color 
"description" (joinStr "" "Congratulations " .User.Username ",\n you leveled up to **level " (toInt $newlevel) "**!\n\nFor more information on our XP system, type `?xp` in any channel.") 
"thumbnail" (sdict "url" (.User.AvatarURL "1024")))}}

	{{sendMessage nil $embed}}
{{if eq .Channel.ID 584957111872782366}}{{sleep 1}}{{dbSet 18 "counter_count" 1}}{{end}}


{{if eq (toInt $newlevel) 10}}
    {{giveRoleID .User.ID 588031694779449484}}
{{execCC 102 587858995012698137 0 (sdict "user" .User "action" "add role" "content" "<@&588031694779449484>")}}
{{else if eq (toInt $newlevel) 15}}
    {{giveRoleID .User.ID 654027682345910282}}
{{execCC 102 587858995012698137 0 (sdict "user" .User "action" "add role" "content" "<@&654027682345910282>")}}
{{end}}
	{{end}}

{{end}}
	
{{$inactive_roleID := 648706778061602816}}
{{dbSet .User.ID "inactive" (toInt64 (currentTime.Sub (newDate 2019 11 11 0 0  0)).Seconds)}}
          {{if (targetHasRoleID .User.ID $inactive_roleID)}}
                    {{takeRoleID .User.ID $inactive_roleID}}
			{{sendDM "Welcome back! We are so glad you have returned. You have been gone awhile. During that time, you were removed from the registered role, and giving the Misplaced Yezians role. This happens at 60 days of discord inactivity. It's a simple precaution to help keep our cluster secure. To regain access to the server password again, you'll need to reregister. Please do so by going to #welcome and typing `register`."}}
{{execCC 102 587858995012698137 0 (sdict "user" .User "action" "remove role" "content" "<@&648706778061602816>")}}
          {{end}}