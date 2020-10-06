73
{{if (eq .Channel.ParentID 635305499600093214)}} 
{{$db := dbGet .User.ID "registration"}}
{{$reg := sdict}}{{range $k,$v := $db.Value}}{{$reg.Set $k $v}}{{end}}
{{$color := randInt 111111 999999 }}
{{$reg.Set "Q" "Edit"}}
{{if hasRoleID 645404993863811072}}
{{$embed := cembed 
	"color" $color
	"Title" "Registration Edit Menu"
	"description" "Please select the registration information you would like to edit by reacting to the number associated with that piece of information."
	"fields" (cslice 
	(sdict "name" "Discord Name" "value" .User.String "inline" false) 
	(sdict "name" ":one: In Game Name" "value" $reg.name "inline" false) 
	(sdict "name" ":two: Tribe Name" "value" $reg.tribe "inline" false) 
	(sdict "name" ":three: Gamer Tag" "value" $reg.gt "inline" false) 
	(sdict "name" ":four: Pin" "value" $reg.pin "inline" false)
	(sdict "name" ":five: Age" "value" (toString $reg.age) "inline" false)
	(sdict "name" ":six: Guardian/Parent" "value" $reg.guardian "inline" false)
	(sdict "name" ":seven: Done Editing" "value" "Use this to stop editing your registration information." "inline" false))
	"thumbnail" (sdict "url" (.User.AvatarURL "1024"))}}
	{{$msg := sendMessageRetID nil $embed}}
	{{addMessageReactions nil $msg "1️⃣" "2️⃣" "3️⃣" "4️⃣" "5️⃣" "6️⃣" "7️⃣"}}
{{dbSet .User.ID "registration" $reg}}
{{else}}
{{$reg.Set "Q" "Edit"}}
{{$embed := cembed 
	"color" $color
	"Title" "Registration Edit Menu"
	"description" "Please select the registration information you would like to edit by reacting to the number associated with that piece of information."
	"fields" (cslice 
	(sdict "name" "Discord Name" "value" .User.String "inline" false) 
	(sdict "name" ":one: In Game Name" "value" $reg.name "inline" false) 
	(sdict "name" ":two: Tribe Name" "value" $reg.tribe "inline" false) 
	(sdict "name" ":three: Gamer Tag" "value" $reg.gt "inline" false) 
	(sdict "name" ":four: Pin" "value" $reg.pin "inline" false)
	(sdict "name" ":five: Done Editing" "value" "Use this to stop editing your registration information." "inline" false))
	"thumbnail" (sdict "url" (.User.AvatarURL "1024")) }}
	{{$msg := sendMessageRetID nil $embed}}
	{{addMessageReactions nil $msg "1️⃣" "2️⃣" "3️⃣" "4️⃣" "5️⃣"}}
{{dbSet .User.ID "registration" $reg}}
{{end}}{{end}}