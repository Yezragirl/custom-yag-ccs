			{{$args := parseArgs 1 "regrefresh for who"
			(carg "userid" "user")}}
			 {{$user := (userArg ($args.Get 0))}}
			 
			 {{$Q := (dbGet .User.ID "Q").Value}} 
			 {{$admin := (dbGet .User.ID "admin").Value}}
			 {{$age := (dbGet .User.ID "age").Value}}
			 {{$gt := (dbGet .User.ID "gt").Value}}
			 {{$name := (dbGet .User.ID "name").Value}}
			 {{$pin := (dbGet .User.ID "pin").Value}}
			 {{$public := (dbGet .User.ID "public").Value}}
			 {{$ref := (dbGet .User.ID "ref").Value}}
			 {{$parentguardian := (dbGet .User.ID "parentguardian").Value}}
			 {{$tribe := (dbGet .User.ID "tribe").Value}}

			{{$embed := cembed 
				"Title" "Registration Edit Menu"
				"description" "Please select the registration information you would like to edit by reacting to the number associated with that piece of information."
				"fields" (cslice 
				(sdict "name" "Discord Name" "value" $user.String ) 
				(sdict "name" ":one: In Game Name" "value" $name ) 
				(sdict "name" ":two: Tribe Name" "value" $tribe ) 
				(sdict "name" ":three: Gamer Tag" "value" $gt ) 
				(sdict "name" ":four: Pin" "value" $pin )
				(sdict "name" ":five: Done Editing" "value" "Use this to stop editing your registration information." ))
				"thumbnail" (sdict "url" ($user.AvatarURL "1024")) }}
{{sendMessage nil $embed}}