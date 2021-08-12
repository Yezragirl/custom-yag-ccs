{{/*listener for AFK*/}}
{{$id := index (index (reFindAllSubmatches `<@!?(\d+)>` .Message.Content) 0) 1}}
{{$color := randInt 111111 16777216}}
{{with (userArg (toInt64 $id))}}
{{$user := .}}
{{with (dbGet (toInt64 $id) "afk")}}
{{sendMessage nil (cembed "author" (sdict "name" (joinStr "" $user.Username " is away!") "icon_url" ($user.AvatarURL "256")) "thumbnail" (sdict "url" "https://media0.giphy.com/media/5ZZTgiEbqLIR7bbqvz/source.gif") "description" .Value "footer" (sdict "text" "Away since") "timestamp" .UpdatedAt "color" $color)}}
{{end}}
{{end}}