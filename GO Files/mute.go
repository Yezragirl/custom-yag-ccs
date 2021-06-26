{{$who := .Message.Author}}
{{$msg := .Message.Content}}
{{$channel := (getChannel nil).Name}}
{{$tz := "America/New_York" }}
{{$location := (newDate 0 0 0 0 0 0 $tz).Location }} 
{{$Time := (currentTime.Add (toDuration "1d"))}}
{{$FormattedTime := $Time.In $location }}
{{$admin := (getMember .User).Nick}}
{{$end := formatTime $FormattedTime "Mon Jan 2,2006 3:04 PM MST" }}
{{if eq .Reaction.Emoji.Name "🔇"}}
{{deleteTrigger}}
{{giveRoleID $who.ID 663076786665685051}}
{{takeRoleID $who.ID 663076786665685051 86400}}
{{sendMessage nil (joinStr "" $who.Mention " has been muted for 24 hours.")}}
{{$capture := exec "ticket open" "Mute"}}
{{$ID := reFind `\d+` (reFind `<#\d+>` $capture)}}
{{execCC 101 $ID 0 (sdict "user" $who)}}
{{$message := sendMessageRetID $ID (joinStr "" $who.Mention ", you have been muted for 24 hours because of the following statement:\n\n" $msg "\n\nIf you feel this mute was unjustified, you can discuss it here. Otherwise, Please hit the <:gcrystal:662677298431918092> below to close this ticket.")}}
{{addMessageReactions $ID $message ":gcrystal:662677298431918092" }}
{{$embed := cembed 
            "color" 000000
            "Title" "Player Mute" 
            "fields" (cslice 
            (sdict "name" "Muted Player" "value" (getMember $who).Nick ) 
            (sdict "name" "Message Channel" "value" $channel ) 
            (sdict "name" "Deleted Message" "value" $msg ) 
            (sdict "name" "Mute Ends" "value" $end ) 
            (sdict "name" "Muted by" "value" $admin ) )
            "thumbnail" (sdict "url" ($who.AvatarURL "1024")) }}

{{sendMessage 587858995012698137 $embed}} 
{{end}}