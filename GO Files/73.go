{{if (eq .Channel.ParentID 635305499600093214)}} 
{{addRoleID 658469834589208616}}
{{$Q := (dbGet .User.ID "Q").Value}} 
{{$admin := (dbGet .User.ID "admin").Value}}
{{$age := (dbGet .User.ID "age").Value}}
{{$gt := (dbGet .User.ID "gt").Value}}
{{$name := (dbGet .User.ID "name").Value}}
{{$pin := (dbGet .User.ID "pin").Value}}
{{$public := (dbGet .User.ID "public").Value}}
{{$ref := (dbGet .User.ID "ref").Value}}
{{$parentguardian := (dbGet .User.ID "parentguardian").Value}}
{{$tribe := (dbGet .User.ID "tribe").Value}} {{if not $tribe}}{{$tribe = "Unregistered"}}{{dbSet .User.ID "tribe" $tribe}}{{end}}
{{$color := randInt 111111 999999 }}
{{$Q = "Edit"}}{{dbSet .User.ID "Q" $Q}}
{{if hasRoleID 645404993863811072}}
{{$embed := cembed 
	"color" $color
	"Title" "Registration Edit Menu"
	"description" "Please select the registration information you would like to edit by reacting to the number associated with that piece of information."
	"fields" (cslice 
	(sdict "name" "Discord Name" "value" .User.String ) 
	(sdict "name" ":one: In Game Name" "value" (toString $name) ) 
	(sdict "name" ":two: Tribe Name" "value" (toString $tribe) ) 
	(sdict "name" ":three: Gamer Tag" "value" (toString $gt) ) 
	(sdict "name" ":four: Pin" "value" (toString $pin) )
	(sdict "name" ":five: Age" "value" (toString (toInt $age)) )
	(sdict "name" ":six: Guardian/Parent" "value" (toString $parentguardian) )
	(sdict "name" ":seven: Done Editing" "value" "Use this to stop editing your registration information." ))
	"thumbnail" (sdict "url" (.User.AvatarURL "1024"))}}
	{{$msg := sendMessageRetID nil $embed}}
	{{addMessageReactions nil $msg "1️⃣" "2️⃣" "3️⃣" "4️⃣" "5️⃣" "6️⃣" "7️⃣"}}
{{else}}
{{$Q = "Edit"}}{{dbSet .User.ID "Q" $Q}}
{{$embed := cembed 
	"color" $color
	"Title" "Registration Edit Menu"
	"description" "Please select the registration information you would like to edit by reacting to the number associated with that piece of information."
	"fields" (cslice 
	(sdict "name" "Discord Name" "value" .User.String "inline" false) 
	(sdict "name" ":one: In Game Name" "value" (toString $name) "inline" false) 
	(sdict "name" ":two: Tribe Name" "value" (toString $tribe) "inline" false) 
	(sdict "name" ":three: Gamer Tag" "value" (toString $gt) "inline" false) 
	(sdict "name" ":four: Pin" "value" (toString $pin) "inline" false)
	(sdict "name" ":five: Done Editing" "value" "Use this to stop editing your registration information." "inline" false))
	"thumbnail" (sdict "url" (.User.AvatarURL "1024")) }}
	{{$msg := sendMessageRetID nil $embed}}
{{addMessageReactions nil $msg "1️⃣" "2️⃣" "3️⃣" "4️⃣" "5️⃣"}}
{{end}}{{end}}