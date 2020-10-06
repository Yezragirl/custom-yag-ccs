{{if and 
    (eq (toString .Reaction.Emoji.Name) "")
    (eq .Reaction.ChannelID $channel)
    (not .User.Bot)
    (hasRoleID $replyRole)}}
 
{{$embed := (index .ReactionMessage.Embeds 0)}}
 
{{dbSetExpire 0 "convert" $embed 60}}
 
{{$value := (dbGet .User.ID "message").Value}}
 
{{$embed = sdict (dbGet 0 "convert").Value}}
 
{{$embed.Set "fields" (cslice (sdict "name" (joinStr "" "Reply From " .User.String) "value" $value "inline" false))}}
 
{{editMessage .Reaction.ChannelID .Reaction.MessageID (cembed $embed)}}
 
{{/* End of function: Add reply (latest message sent by user who adds reaction "") and has $replyRole to message with "" reaction added to it. /}}
 
{{/ If user without $replyRole, removes reaction */}}
 
{{else if
    (eq .Reaction.ChannelID $channel)}}
 
{{deleteMessageReaction $channel .Reaction.MessageID .User.ID ""}} 
 
{{end}}