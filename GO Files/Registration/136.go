{{$msgID := .ReactionMessage.ID}}
{{$Q := (dbGet .User.ID "Q").Value}}{{$Q}}
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
{{$color := randInt 111111 999999 }}
{{$ticketnum := (dbGet 0 "ticketnum").Value}}
{{if (eq .Channel.ParentID 635305499600093214)}} 
{{if (hasRoleID 607041870966685706)}}{{/*Has Registered Role*/}}
{{if and (eq .Reaction.Emoji.Name "1️⃣") (eq (toString $Q) "Edit")}}
{{$Q = "Edit Name"}}
{{dbSet .User.ID "Q" $Q}}
{{deleteAllMessageReactions nil $msgID}}
{{sendMessage nil (joinStr "" .User.Mention ", what's your *new* **in game name**?")}}
{{else if and (eq .Reaction.Emoji.Name "2️⃣") (eq (toString $Q) "Edit")}}
{{$Q = "Edit Tribe"}}
{{dbSet .User.ID "Q" $Q}}
{{deleteAllMessageReactions nil $msgID}}
{{$embed := cembed 
	"color" $color
	"Title" "Tribe Edit Menu"
	"description" "Please select the tribe information you would like to edit by reacting to the number associated."
	"fields" (cslice 
	(sdict "name" ":one: Tribe 1" "value" (toString $tribe) ) 
	(sdict "name" ":two: Tribe 2" "value" (toString $tribe2) ) 
	(sdict "name" ":three: Done Editing" "value" "Use this to stop editing your registration information." ))
	"thumbnail" (sdict "url" (.User.AvatarURL "1024"))}}
	{{$msg := sendMessageRetID nil $embed}}
	{{addMessageReactions nil $msg "1️⃣" "2️⃣" "3️⃣"}}
	{{deleteAllMessageReactions nil $msgID}}
{{else if and (eq .Reaction.Emoji.Name "1️⃣") (eq (toString $Q) "Edit Tribe")}}
{{$embed := cembed 
	"color" $color
	"Title" "Tribe Edit Menu"
	"description" "Please select the tribe information you would like to edit by reacting to the number associated."
	"fields" (cslice 
	(sdict "name" "Current Tribe 1" "value" (toString $tribe) ) 
	(sdict "name" ":one: Join/Change Tribe 1" "value" "** **" ) 
	(sdict "name" ":two: Leave Tribe 1" "value" "** **" ) 
	(sdict "name" ":three: Create New Tribe" "value" "** **" ) 
	(sdict "name" ":four: Done Editing" "value" "Use this to stop editing your registration information." ))
	"thumbnail" (sdict "url" (.User.AvatarURL "1024"))}}
	{{$msg := sendMessageRetID nil $embed}}
	{{addMessageReactions nil $msg "1️⃣" "2️⃣" "3️⃣" "4️⃣"}}
{{$Q = "First Tribe"}}
{{dbSet .User.ID "Q" $Q}}
{{deleteAllMessageReactions nil $msgID}}
{{else if and (eq .Reaction.Emoji.Name "2️⃣") (eq (toString $Q) "Edit Tribe")}}
{{$embed := cembed 
	"color" $color
	"Title" "Tribe Edit Menu"
	"description" "Please select the tribe information you would like to edit by reacting to the number associated."
	"fields" (cslice 
	(sdict "name" "Current Tribe 2" "value" (toString $tribe2) ) 
	(sdict "name" ":one: Join/Change Tribe 2" "value" "** **" ) 
	(sdict "name" ":two: Leave Tribe 2" "value" "** **" ) 
	(sdict "name" ":three: Create New Tribe" "value" "** **" ) 
	(sdict "name" ":four: Done Editing" "value" "Use this to stop editing your registration information." ))
	"thumbnail" (sdict "url" (.User.AvatarURL "1024"))}}
	{{$msg := sendMessageRetID nil $embed}}
	{{addMessageReactions nil $msg "1️⃣" "2️⃣" "3️⃣" "4️⃣"}}
{{$Q = "Second Tribe"}}
{{dbSet .User.ID "Q" $Q}}
{{deleteAllMessageReactions nil $msgID}}
{{else if  (eq (toString $Q) "Edit First Tribe Join")}}
{{else if  (eq (toString $Q) "Edit Second Tribe Join")}}
{{else if and (eq .Reaction.Emoji.Name "1️⃣") (eq (toString $Q) "First Tribe")}}
{{execCC 222 nil 0 (sdict "tribe" "first" "step" "join")}}
{{deleteAllMessageReactions nil $msgID}}
{{else if and (eq .Reaction.Emoji.Name "2️⃣") (eq (toString $Q) "First Tribe")}}
{{execCC 222 nil 0 (sdict "tribe" "first" "step" "leave")}}
{{deleteAllMessageReactions nil $msgID}}
{{else if and (eq .Reaction.Emoji.Name "3️⃣") (eq (toString $Q) "First Tribe")}}
{{execCC 222 nil 0 (sdict "tribe" "first" "step" "create")}}
{{deleteAllMessageReactions nil $msgID}}
{{else if and (eq .Reaction.Emoji.Name "1️⃣") (eq (toString $Q) "Second Tribe")}}
{{execCC 222 nil 0 (sdict "tribe" "second" "step" "join")}}
{{deleteAllMessageReactions nil $msgID}}
{{else if and (eq .Reaction.Emoji.Name "2️⃣") (eq (toString $Q) "Second Tribe")}}
{{execCC 222 nil 0 (sdict "tribe" "second" "step" "leave")}}
{{deleteAllMessageReactions nil $msgID}}
{{else if and (eq .Reaction.Emoji.Name "3️⃣") (eq (toString $Q) "Second Tribe")}}
{{execCC 222 nil 0 (sdict "tribe" "second" "step" "create")}}
{{deleteAllMessageReactions nil $msgID}}
{{else if and (eq .Reaction.Emoji.Name "3️⃣") (eq (toString $Q) "Edit")}}
{{$Q = "Edit GT"}}
{{dbSet .User.ID "Q" $Q}}
{{deleteAllMessageReactions nil $msgID}}
{{sendMessage nil (joinStr "" .User.Mention ", what's your *new* **Gamer Tag**?")}}
{{else if and (eq .Reaction.Emoji.Name "4️⃣") (eq (toString $Q) "Edit")}}
{{$Q = "Edit Pin"}}
{{dbSet .User.ID "Q" $Q}}
{{deleteAllMessageReactions nil $msgID}}
{{sendMessage nil (joinStr "" .User.Mention ", what's your *new* **Pin**?")}}
{{else if or (and (eq .Reaction.Emoji.Name "5️⃣") (eq (toString $Q) "Edit")) (and (eq .Reaction.Emoji.Name "3️⃣") (eq (toString $Q) "Edit Tribe")) (and (eq .Reaction.Emoji.Name "4️⃣") (eq (toString $Q) "First Tribe")) (and (eq .Reaction.Emoji.Name "4️⃣") (eq (toString $Q) "Second Tribe"))}}
{{deleteAllMessageReactions nil $msgID}}
{{$Q = "Done"}}
{{dbSet .User.ID "Q" $Q}}
{{execCC 74 nil 0 ""}}
{{removeRoleID 658469834589208616}}
{{sleep 3}}
{{exec "tickets close" "Done Editing Registration"}}
{{$ticketnum = sub $ticketnum 1}}
{{dbSet 0 "ticketnum" $ticketnum}}
{{/*Would you like to edit something else?*/}}
{{else if and (eq .Reaction.Emoji.ID 658376626639339533) (eq (toString $Q) "End Edit")}}{{/*yes*/}}
{{deleteAllMessageReactions nil $msgID}}
{{$Q = "Done"}}
{{dbSet .User.ID "Q" $Q}}
{{execCC 73 nil 0 ""}}
{{else if and (eq .Reaction.Emoji.ID 658376639322783745) (eq (toString $Q) "End Edit")}}{{/*no*/}}
{{deleteAllMessageReactions nil $msgID}}
{{$Q = "Done"}}
{{dbSet .User.ID "Q" $Q}}
{{execCC 74 nil 0 ""}}
{{sleep 3}}
{{removeRoleID 658469834589208616}}
{{exec "tickets close" "Done Editing Registration"}}
{{/*Collected*/}}
{{else if (eq .Reaction.Emoji.Name "gcrystal")}}
{{deleteAllMessageReactions nil $msgID}}
{{exec "tickets close" "Picked up Items"}}
{{end}}
{{/*RegEdit-Kid*/}}
{{else if hasRoleID 645404993863811072}}{{/*Has Kids Role*/}}
{{execCC 225 nil 0 (sdict "msgID" $msgID)}}
{{end}}{{end}}


