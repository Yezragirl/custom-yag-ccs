{{$name := .User.Username}}
{{$nick := (getMember .User.ID).Nick}}
{{$join := (getMember .User.ID).JoinedAt}}
{{$topic := execAdmin "topic"}}

{{$embed := cembed 
           "fields" (cslice 
            (sdict "name" "Name" "value" $name "inline" false) 
			(sdict "name" "Nickname" "value" $nick "inline" false)
			(sdict "name" "Joined At" "value" (toString $join) "inline" false)  
            (sdict "name" "Topic" "value" $topic "inline" false) )
            "thumbnail" (sdict "url" (.User.AvatarURL "1024")) }}
{{sendMessage nil $embed}}




{{/* Reactions needed*/}}
{{$count := 3 }}
{{/* name of your emote*/}}
{{$emoteName := "😂"}}
{{$who := ""}}

{{$match := 0}}
{{/* check all reactions to see if it matches ours */}}
{{range .ReactionMessage.Reactions}}

{{/* check if emote is your emote */}}
{{if eq (toString .Emoji.Name) (toString $emoteName)}}
{{$who = (joinStr " " $who ().Reaction.UserID).Mention)}}
{{/* if emote meets required count and bot has not reacted */}}
{{if eq .Count $count}}
{{if not .Me}}
{{$match = 1}}
{{end}}{{end}}{{end}}{{end}}

{{if eq $match 1}}
{{sleep 30}}
{{sendMessage nil $who}}
{{end}}