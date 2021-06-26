{{if (eq .Channel.ParentID 635305499600093214)}}
{{$Q := (dbGet .User.ID "Q").Value}}
{{$tribe := (dbGet .User.ID "tribe").Value}}{{if not $tribe}}{{$tribe = "Unregistered"}}{{dbSet .User.ID "tribe" $tribe}}{{end}}
{{$tribe2 := (dbGet .User.ID "tribe2").Value}}{{if not $tribe2}}{{$tribe2 = "Unregistered"}}{{dbSet .User.ID "tribe2" $tribe2}}{{end}}
{{$MsgID := .Reaction.MessageID}}
{{$optionselect := ""}}
{{$tribeedit := (dbGet .Channel.ID "tribeedit").Value}}
{{$fields := cslice}}
{{$addEmotes := cslice}}
{{$rankemotes := sdict "1" "1️⃣" "2" "2️⃣" "3" "3️⃣" "4" "4️⃣" "5" "5️⃣" "6" "6️⃣" "7" "7️⃣" "8" "8️⃣" "9" "9️⃣" "10" "🔟"}}

{{range (index .Message.Embeds 0).Fields -}}
{{$parts := reFindAllSubmatches `([^ ]+) (.+)` .Name -}}
{{$emojipart := index $parts 0 1 -}}
{{if (eq $.Reaction.Emoji.Name $emojipart)}}
	{{$optionselect = index $parts 0 2}}
{{end}}
{{end}}
{{$Q}}

{{if eq $tribeselect "Cancel"}}
	{{removeRoleID 770393574734692363}}
{{dbDel .Channel.ID "tribeedit"}}
{{$msg := sendMessageRetID nil "Tribe Management Cancelled.\n\nPlease hit the <:gcrystal:662677298431918092> below to close this ticket. Or you can type `-regedit` or `-managetribe` to continue making edits."}}
{{addMessageReactions nil $msg ":gcrystal:662677298431918092" }}
{{deleteTrigger 1}}

{{else if eq $Q "TribeMan"}}
	{{$tribeedit = $optionselect}}
	{{$dbkey := print "tribe_" (lower $tribeedit)}}
	{{$tribeinfo := sdict (dbGet 0 $dbkey).Value}}
	{{dbSet .Channel.ID "tribeedit" $tribeedit}}

	{{$color := randInt 111111 999999 }}
	{{$embed := cembed 
	"Title" (joinStr "" "Tribe Management for " $tribeedit)
	"color" $color
	"fields" (cslice 
	(sdict "name" ":one: Transfer Tribe Ownership" "value" "Transfer ownership of this tribe to another tribe member." "inline" false) 
	(sdict "name" ":two: Manage Tribe Members" "value" "Remove Tribe Members." "inline" false) 
	(sdict "name" ":three: Change Tribe Name" "value" "Change the Tribe Name." "inline" false) 
	(sdict "name" ":four: Manage Tribe Bases" "value" "Edit/Remove Tribe Bases." "inline" false)
	(sdict "name" ":five: Cancel" "value" "Cancel Tribe Management." "inline" false))
	"thumbnail" (sdict "url" "https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/46b63d3c-ae67-464c-9a37-670829b2a157/d9eh2nw-f682acd6-3622-461c-8293-e7907bcc11d8.png?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOiIsImlzcyI6InVybjphcHA6Iiwib2JqIjpbW3sicGF0aCI6IlwvZlwvNDZiNjNkM2MtYWU2Ny00NjRjLTlhMzctNjcwODI5YjJhMTU3XC9kOWVoMm53LWY2ODJhY2Q2LTM2MjItNDYxYy04MjkzLWU3OTA3YmNjMTFkOC5wbmcifV1dLCJhdWQiOlsidXJuOnNlcnZpY2U6ZmlsZS5kb3dubG9hZCJdfQ.Tsz77fz2XHArjp78X1vPBVJTrKNZYpUP2cR-lKp_wo8") }}
{{$msg := sendMessageRetID nil $embed}}
{{addMessageReactions nil $msg "1️⃣" "2️⃣" "3️⃣" "4️⃣" "5️⃣"}}

{{else if eq $optionselect "Transfer Tribe Ownership"}} BUILD MEMBER LIST

	{{$dbkey := print "tribe_" (lower $tribeedit)}}
	{{$tribeinfo := sdict (dbGet 0 $dbkey).Value}}
	{{$members := (cslice).AppendSlice $tribeinfo.members}}
	{{- range $members -}}
		{{- $rank = add $rank 1 -}}
		{{- $emote = $rankemotes.Get (toString $rank) -}}
		{{- $addEmotes = $addEmotes.Append $emote -}}
		{{- $fields = $fields.Append
		(sdict "name" (print $emote " " .) "value" "** **")}}
	{{- end}}
	{{$rank = add $rank 1}}
	{{$emote = $rankemotes.Get (toString $rank)}}
	{{$addEmotes = $addEmotes.Append $emote}}
	{{$fields = $fields.Append
	(sdict "name" (print $emote " Cancel") "value" "Cancel tribe management.")}}
	{{$embed := cembed 
	"title" "Which tribe member would you like to transfer ownership to?" 
	"description" "Use the reactions to indicate the new tribe owner." 
	"fields" $fields
	"color" 4645612 }}
	{{$msg := sendMessageRetID nil $embed}}
	{{- range $addEmotes -}}
		{{- addMessageReactions nil $msg . -}}
	{{- end -}}
	{{$Q = "owner"}}
	{{dbSet .User.ID "Q" $Q}}
{{else if eq $tribeselect "Manage Tribe Members"}} BUILD MEMBER LIST

	{{$dbkey := print "tribe_" (lower $tribeselect)}}
	{{$tribeinfo := sdict (dbGet 0 $dbkey).Value}}
	{{$color := randInt 111111 999999 }}
	{{$embed := cembed 
	"Title" (joinStr "" "Tribe Management for " $tribeselect)
	"color" $color
	"fields" (cslice 
	(sdict "name" ":one: Transfer Tribe Ownership" "value" "Transfer ownership of this tribe to another tribe member." "inline" false) 
	(sdict "name" ":two: Manage Tribe Members" "value" "Remove Tribe Members." "inline" false) 
	(sdict "name" ":three: Change Tribe Name" "value" "Change the Tribe Name." "inline" false) 
	(sdict "name" ":four: Manage Tribe Bases" "value" "Edit/Remove Tribe Bases." "inline" false)
	(sdict "name" ":five: Cancel" "value" "Cancel Tribe Management." "inline" false))
	"thumbnail" (sdict "url" "https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/46b63d3c-ae67-464c-9a37-670829b2a157/d9eh2nw-f682acd6-3622-461c-8293-e7907bcc11d8.png?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOiIsImlzcyI6InVybjphcHA6Iiwib2JqIjpbW3sicGF0aCI6IlwvZlwvNDZiNjNkM2MtYWU2Ny00NjRjLTlhMzctNjcwODI5YjJhMTU3XC9kOWVoMm53LWY2ODJhY2Q2LTM2MjItNDYxYy04MjkzLWU3OTA3YmNjMTFkOC5wbmcifV1dLCJhdWQiOlsidXJuOnNlcnZpY2U6ZmlsZS5kb3dubG9hZCJdfQ.Tsz77fz2XHArjp78X1vPBVJTrKNZYpUP2cR-lKp_wo8") }}
{{$msg := sendMessageRetID nil $embed}}
{{addMessageReactions nil $msg "1️⃣" "2️⃣" "3️⃣" "4️⃣" "5️⃣"}}

{{else if eq $tribeselect "Change Tribe Name"}} CHANGE TRIBE NAME

	{{$dbkey := print "tribe_" (lower $tribeselect)}}
	{{$tribeinfo := sdict (dbGet 0 $dbkey).Value}}
	{{$color := randInt 111111 999999 }}
	{{$embed := cembed 
	"Title" (joinStr "" "Tribe Management for " $tribeselect)
	"color" $color
	"fields" (cslice 
	(sdict "name" ":one: Transfer Tribe Ownership" "value" "Transfer ownership of this tribe to another tribe member." "inline" false) 
	(sdict "name" ":two: Manage Tribe Members" "value" "Remove Tribe Members." "inline" false) 
	(sdict "name" ":three: Change Tribe Name" "value" "Change the Tribe Name." "inline" false) 
	(sdict "name" ":four: Manage Tribe Bases" "value" "Edit/Remove Tribe Bases." "inline" false)
	(sdict "name" ":five: Cancel" "value" "Cancel Tribe Management." "inline" false))
	"thumbnail" (sdict "url" "https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/46b63d3c-ae67-464c-9a37-670829b2a157/d9eh2nw-f682acd6-3622-461c-8293-e7907bcc11d8.png?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOiIsImlzcyI6InVybjphcHA6Iiwib2JqIjpbW3sicGF0aCI6IlwvZlwvNDZiNjNkM2MtYWU2Ny00NjRjLTlhMzctNjcwODI5YjJhMTU3XC9kOWVoMm53LWY2ODJhY2Q2LTM2MjItNDYxYy04MjkzLWU3OTA3YmNjMTFkOC5wbmcifV1dLCJhdWQiOlsidXJuOnNlcnZpY2U6ZmlsZS5kb3dubG9hZCJdfQ.Tsz77fz2XHArjp78X1vPBVJTrKNZYpUP2cR-lKp_wo8") }}
{{$msg := sendMessageRetID nil $embed}}
{{addMessageReactions nil $msg "1️⃣" "2️⃣" "3️⃣" "4️⃣" "5️⃣"}}

{{else if eq $tribeselect "Manage Tribe Bases"}} BUILD BASE LIST

	{{$dbkey := print "tribe_" (lower $tribeselect)}}
	{{$tribeinfo := sdict (dbGet 0 $dbkey).Value}}
	{{$color := randInt 111111 999999 }}
	{{$embed := cembed 
	"Title" (joinStr "" "Tribe Management for " $tribeselect)
	"color" $color
	"fields" (cslice 
	(sdict "name" ":one: Transfer Tribe Ownership" "value" "Transfer ownership of this tribe to another tribe member." "inline" false) 
	(sdict "name" ":two: Manage Tribe Members" "value" "Remove Tribe Members." "inline" false) 
	(sdict "name" ":three: Change Tribe Name" "value" "Change the Tribe Name." "inline" false) 
	(sdict "name" ":four: Manage Tribe Bases" "value" "Edit/Remove Tribe Bases." "inline" false)
	(sdict "name" ":five: Cancel" "value" "Cancel Tribe Management." "inline" false))
	"thumbnail" (sdict "url" "https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/46b63d3c-ae67-464c-9a37-670829b2a157/d9eh2nw-f682acd6-3622-461c-8293-e7907bcc11d8.png?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOiIsImlzcyI6InVybjphcHA6Iiwib2JqIjpbW3sicGF0aCI6IlwvZlwvNDZiNjNkM2MtYWU2Ny00NjRjLTlhMzctNjcwODI5YjJhMTU3XC9kOWVoMm53LWY2ODJhY2Q2LTM2MjItNDYxYy04MjkzLWU3OTA3YmNjMTFkOC5wbmcifV1dLCJhdWQiOlsidXJuOnNlcnZpY2U6ZmlsZS5kb3dubG9hZCJdfQ.Tsz77fz2XHArjp78X1vPBVJTrKNZYpUP2cR-lKp_wo8") }}
{{$msg := sendMessageRetID nil $embed}}
{{addMessageReactions nil $msg "1️⃣" "2️⃣" "3️⃣" "4️⃣" "5️⃣"}}








{{end}} 	{{/* End IF Qs */}}
{{end}} 	{{/* End RegEdit Channel Check */}}



