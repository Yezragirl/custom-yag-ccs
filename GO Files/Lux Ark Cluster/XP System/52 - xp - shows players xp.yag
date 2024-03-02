{{/*xp shows player xp info*/}}
{{$user := .User}}

{{if .CmdArgs}}
{{if (userArg (index .CmdArgs 0))}}
{{$user = (userArg (index .CmdArgs 0))}}
{{end}}
{{end}}

{{$roles := (getMember $user.ID).Roles}}
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

{{$xp := (dbGet $user.ID "XP").Value}}
{{$currentlevel := (toInt (roundFloor (sqrt (fdiv $xp 100))))}}
{{$nextlevel := add $currentlevel 1 }}
{{$nextxp := mult (mult $nextlevel $nextlevel) 100}}
{{$neededxp := sub $nextxp $xp}}
{{$embed := cembed 
"color" $color
"fields" (cslice 
    (sdict "name" "Name" "value" $user.Username "inline" false) 
    (sdict "name" "Current Level" "value" (toString $currentlevel) "inline" false) 
    (sdict "name" "Current XP" "value" (toString (toInt $xp)) "inline" false))
"footer" (sdict "text" (joinStr "" "XP needed for next level: " $neededxp))
"thumbnail" (sdict "url" ($user.AvatarURL "1024"))}}
{{sendMessage nil $embed}}