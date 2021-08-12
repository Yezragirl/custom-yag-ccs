{{if (eq .Channel.ParentID 635305499600093214)}} 
{{addRoleID 760904971465654302}}
{{$tribe := (dbGet .User.ID "tribe").Value}} {{if $tribe}}{{$tribe = (cslice).AppendSlice $tribe}}{{else}}{{$tribe = "Unregistered"}}{{dbSet .User.ID "tribe" $tribe}}{{end}}
{{$color := randInt 111111 999999 }}
{{$Q := "Tribe Edit"}}{{dbSet .User.ID "Q" $Q}}
{{if hasRoleID 760904971465654302}}
{{$embed := cembed 
	"color" $color
	"Title" "Tribe Registration Menu"
	"description" "Type the number of the option you'd like.\n:one: Join a Tribe\n:two: Create a New Tribe"
	"thumbnail" (sdict "url" (.User.AvatarURL "1024"))}}
	{{sendMessageRetID nil $embed}}
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