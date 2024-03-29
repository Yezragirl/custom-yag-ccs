{{if (eq .Channel.ParentID 635305499600093214)}} 
{{$Q := (dbGet .User.ID "Q").Value}} 
{{$admin := (dbGet .User.ID "admin").Value}}
{{$age := (dbGet .User.ID "age").Value}}
{{$gt := (dbGet .User.ID "gt").Value}}
{{$name := (dbGet .User.ID "name").Value}}
{{$pin := (dbGet .User.ID "pin").Value}}
{{$public := (dbGet .User.ID "public").Value}}
{{$ref := (dbGet .User.ID "ref").Value}}
{{$parentguardian := (dbGet .User.ID "parentguardian").Value}}
{{$tribe := (dbGet .User.ID "tribe").Value}}{{if not $tribe}}{{$tribe = "Unregistered"}}{{dbSet .User.ID "tribe" $tribe}}{{end}}
{{$tribe2 := (dbGet .User.ID "tribe2").Value}}{{if not $tribe2}}{{$tribe2 = "Unregistered"}}{{dbSet .User.ID "tribe2" $tribe2}}{{end}}
{{$names_db := (dbGet 0 "names").Value}}
{{$names := (cslice).AppendSlice $names_db}}
{{$tribes_db := (dbGet 0 "tribes").Value}}
{{$tribes := (cslice).AppendSlice $tribes_db}}
{{$lowertribes := cslice}}
{{$finds := cslice}}
{{$out := (printf "%s\t%-18s\n" "Listing" "Tribe")}}
{{$rank := 0}}
{{$tname := "tribe"}}
{{$rankemotes := sdict "1" "1️⃣" "2" "2️⃣" "3" "3️⃣" "4" "4️⃣" "5" "5️⃣" "6" "6️⃣" }}
{{$emote := ""}}
{{$foundMatch := false}}
{{$lowerMC := lower .Message.Content}}
{{- range $tribes -}}
    {{- $lowertribes = $lowertribes.Append (lower .) -}}
{{- end -}}
{{if not $name}}
	You can't edit your registration. You don't have one yet. 
{{else}}
	{{if eq (toString $Q) "Edit GT"}}
		{{$gt = (toString .Message.Content)}}{{dbSet .User.ID "gt" $gt}}
		{{$Q = "End Edit"}}{{dbSet .User.ID "Q" $Q}}
	{{else if eq (toString $Q) "Edit Name"}}
		{{if gt (len (toString .Message.Content)) 25}}
			{{sendMessage nil (joinStr "" "Sorry " .User.Mention ", that's too long for a name. I'll ask again, What would you like to use as your In Game Name?")}}
		{{else if (in $names (toString .Message.Content))}}
			{{sendMessage nil (joinStr "" "Sorry " .User.Mention ", that name is already taken. Since names need to be unique, I'll ask again, what would you like to use as your In Game Name?")}}
		{{else}}
			{{$name = (toString .Message.Content)}}
			{{dbSet .User.ID "name" $name}}
			{{dbSet 0 "names" ($names.AppendSlice (cslice $name))}}
			{{$Q = "End Edit"}}{{dbSet .User.ID "Q" $Q}}
		{{end}}
	{{else if eq (toString $Q) "First Tribe Create"}}
		{{- range $lowertribes -}}
			{{- if not $foundMatch -}}
				{{- if in . $lowerMC -}}
					{{- $foundMatch = true -}}
				{{- end -}}
			{{- end -}}
		{{- end -}}
		{{if $foundMatch}}
			{{sendMessage nil (joinStr "" "Sorry " .User.Mention ", that tribe name is already taken. Since Tribe names need to be unique, I'll ask again, what would you like to use as your Tribe Name?")}}
		{{else}}
			{{$tribe = (toString .Message.Content)}}
			{{dbSet .User.ID "tribe" $tribe}}
			{{dbSet 0 "tribes" ($tribes.Append $tribe)}}
			{{dbSet 0 (joinStr "" "tribe_" (lower $tribe)) (sdict "name" $tribe "owner" .User.ID)}}
			{{$Q = "End Edit"}}{{dbSet .User.ID "Q" $Q}}
			{{sendMessage nil (joinStr "" "Tribe " $tribe ", created. " .User.Mention " set as owner.")}}
		{{end}}
		{{else if eq (toString $Q) "Second Tribe Create"}}
		{{- range $lowertribes -}}
			{{- if not $foundMatch -}}
				{{- if in . $lowerMC -}}
					{{- $foundMatch = true -}}
				{{- end -}}
			{{- end -}}
		{{- end -}}
		{{if $foundMatch}}
			{{sendMessage nil (joinStr "" "Sorry " .User.Mention ", that tribe name is already taken. Since Tribe names need to be unique, I'll ask again, what would you like to use as your Tribe Name?")}}
		{{else}}
			{{$tribe = (toString .Message.Content)}}
			{{dbSet .User.ID "tribe2" $tribe2}}
			{{dbSet 0 "tribes" ($tribes.Append $tribe2)}}
			{{dbSet 0 (joinStr "" "tribe_" (lower $tribe2)) (sdict "name" $tribe2 "owner" .User.ID)}}
			{{$Q = "End Edit"}}{{dbSet .User.ID "Q" $Q}}
			{{sendMessage nil (joinStr "" "Tribe " $tribe2 ", created. " .User.Mention " set as owner.")}}
		{{end}}
	{{else if eq (toString $Q) "First Tribe Join"}}
		{{$finds = dbGetPattern 0 (print "tribe_%" (lower .Message.Content) "%") 5 0}}
		{{$fields := cslice}}
		{{$addEmotes := cslice}}
		{{- range $finds -}}
			{{- $rank = add $rank 1 -}}
			{{- $tribeInfo := sdict .Value -}}
			{{- $tname = $tribeInfo.name -}}
			{{- $emote = $rankemotes.Get (toString $rank) -}}
			{{- $addEmotes = $addEmotes.Append $emote -}}
			{{- $fields = $fields.Append 
            (sdict "name" (print $emote " " $tname) "value" "** **") -}}
		{{- end -}}
		{{$rank = add $rank 1}}
		{{$emote = $rankemotes.Get (toString $rank)}}
		{{$addEmotes = $addEmotes.Append $emote}}
		{{$fields = $fields.Append
		(sdict "name" (print $emote " Try Again") "value" "Don't see what you wanted? Try your search again.")}}
		{{$embed := cembed 
		"title" "Which Tribe would you like to join?" 
		"description" "Use the reactions to indicate the tribe you'd like to join." 
		"fields" $fields
		"color" 4645612 }}
		{{$msg := sendMessageRetID nil $embed}}
		{{- range $addEmotes -}}
			{{- addMessageReactions nil $msg . -}}
		{{- end -}}
		{{$Q = "Edit First Tribe Join"}}{{dbSet .User.ID "Q" $Q}}
	{{else if eq (toString $Q) "Second Tribe Join"}}
		{{$finds = dbGetPattern 0 (print "tribe_%" (lower .Message.Content) "%") 5 0}}
		{{$fields := cslice}}
		{{$addEmotes := cslice}}
		{{- range $finds -}}
			{{- $rank = add $rank 1 -}}
			{{- $tribeInfo := sdict .Value -}}
			{{- $tname = $tribeInfo.name -}}
			{{- $emote = $rankemotes.Get (toString $rank) -}}
			{{- $addEmotes = $addEmotes.Append $emote -}}
			{{- $fields = $fields.Append
			(sdict "name" (print $emote " " $tname) "value" "** **")}}
		{{- end}}
		{{$rank = add $rank 1}}
		{{$emote = $rankemotes.Get (toString $rank)}}
		{{$addEmotes = $addEmotes.Append $emote}}
		{{$fields = $fields.Append
		(sdict "name" (print $emote " Try Again") "value" "Don't see what you wanted? Try your search again.")}}
		{{$embed := cembed 
		"title" "Which Tribe would you like to join?" 
		"description" "Use the reactions to indicate the tribe you'd like to join." 
		"fields" $fields
		"color" 4645612 }}
		{{$msg := sendMessageRetID nil $embed}}
		{{- range $addEmotes -}}
			{{- addMessageReactions nil $msg . -}}
		{{- end -}}
		{{$Q = "Edit Second Tribe Join"}}{{dbSet .User.ID "Q" $Q}}
	{{else if eq (toString $Q) "Edit Pin"}}
		{{$pin = (toString .Message.Content)}}{{dbSet .User.ID "pin" $pin}}
		{{$Q = "End Edit"}}{{dbSet .User.ID "Q" $Q}}
	{{else if eq (toString $Q) "Edit ParentGuardian"}}
		{{$parentguardian = (toString .Message.Content)}}{{dbSet .User.ID "parentguardian" $parentguardian}}
		{{$Q = "End Edit"}}{{dbSet .User.ID "Q" $Q}}
	{{else if eq (toString $Q) "Edit Age"}}
		{{$msg := sendMessageRetID nil (joinStr "" .User.Mention ", did you have a birthday?")}}
		{{addMessageReactions nil $msg ":yes:658376626639339533" ":no:658376639322783745"}}
	{{else if eq (toString $Q) "Edited"}}
	{{$msg := sendMessageRetID nil (joinStr "" .User.Mention ", got it. Do you want to edit something else?")}}
	{{addMessageReactions nil $msg ":yes:658376626639339533" ":no:658376639322783745"}} 
{{end}}{{end}}{{end}}

