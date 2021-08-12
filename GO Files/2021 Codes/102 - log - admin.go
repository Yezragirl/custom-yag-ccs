{{/*log for logging admin channel*/}}
{{$user := .ExecData.user}}
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
{{$text1 := ""}}
{{$text2 := ""}}
    {{if eq .ExecData.action "remove role"}}
          {{$text1 = "was removed from"}}
          {{$text2 = "role"}}
    {{else if eq .ExecData.action "add role"}}
            {{$text1 = "was given"}}
          {{$text2 = "role"}}
{{end}}
{{$embed := cembed 
            "Title" .ExecData.action  
"description" (joinStr " "  $user.Mention $text1 .ExecData.content $text2)           
"color" $color
             "thumbnail" (sdict "url" ($user.AvatarURL "1024")) }}
{{sendMessage 587858995012698137 $embed}}