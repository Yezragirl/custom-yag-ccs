{{if (eq .Channel.ParentID 635305499600093214)}} 
{{$names_db := (dbGet 0 "names").Value}}
{{$names := (cslice).AppendSlice $names_db}}
	{{$Q := (dbGet .User.ID "Q").Value}}{{$admin := (dbGet .User.ID "admin").Value}}{{$age := (dbGet .User.ID "age").Value}}{{$gt := (dbGet .User.ID "gt").Value}}{{$name := (dbGet .User.ID "name").Value}}{{$pin := (dbGet .User.ID "pin").Value}}{{$public := (dbGet .User.ID "public").Value}}{{$ref := (dbGet .User.ID "ref").Value}}{{$parentguardian := (dbGet .User.ID "parentguardian").Value}}
	{{if not $parentguardian}}{{$parentguardian = "Not a Minor"}}{{dbSet .User.ID "parentguardian" $parentguardian}}{{end}}{{$tribe := (dbGet .User.ID "tribe").Value}}
	{{if not $tribe}}{{$tribe = "Unregistered"}}{{dbSet .User.ID "tribe" $tribe}}{{end}}
	{{if not $Q}}
		{{$Q = "Edit Name"}}{{dbSet .User.ID "Q" $Q}}
		{{sendMessage nil (joinStr "" .User.Mention ", how old are you?")}}
	{{else}} 
		{{if eq (toString $Q) "Age"}}
			{{if reFind `\D` (toString .Message.Content)}}{{sendMessage nil (joinStr "" "Sorry " .User.Mention ", thats not a number. I'll ask again, how old are you?")}}
			{{else if gt (toInt .Message.Content) 85}}{{sendMessage nil (joinStr "" "Sorry " .User.Mention ", that's an unrealistic number. I'll ask again, how old are you?")}}
			{{else if lt (toInt .Message.Content) 5}}{{sendMessage nil (joinStr "" "Sorry " .User.Mention ", that's an unrealistic number. I'll ask again, how old are you?")}}
			{{else}}
				{{$age = (toInt .Message.Content)}}
				{{dbSet .User.ID "age" $age}}
				{{if ge $age 18}}
					{{$Q = "GT"}}{{dbSet .User.ID "Q" $Q}}
					{{sendMessage nil (joinStr "" .User.Mention ", what is your Gamer Tag?")}}
				{{else}}
					{{$Q = "ParentGuardian"}}{{dbSet .User.ID "Q" $Q}}
					{{$msg := sendMessageRetID nil (joinStr "" .User.Mention ", is your Parent or Guardian a member of this cluster?")}}
					{{addMessageReactions nil $msg ":yes:658376626639339533" ":no:658376639322783745"}}
				{{end}}
			{{end}}
		{{else if eq (toString $Q) "GT"}}{{$gt = (toString .Message.Content)}}{{dbSet .User.ID "gt" $gt}}{{$Q = "Name"}}{{dbSet .User.ID "Q" $Q}}{{sendMessage nil (joinStr "" .User.Mention ", what would you like to use as your In Game Name?")}}
		{{else if eq (toString $Q) "Name"}}
			{{if gt (len (toString .Message.Content)) 25}}{{sendMessage nil (joinStr "" "Sorry " .User.Mention ", that's too long for a name. I'll ask again, What would you like to use as your In Game Name?")}}
			{{else if (in $names (toString .Message.Content))}}{{sendMessage nil (joinStr "" "Sorry " .User.Mention ", that name is already taken. Since names need to be unique, I'll ask again, what would you like to use as your In Game Name?")}}
			{{else}}
			{{$name = (toString .Message.Content)}}
			{{dbSet .User.ID "name" $name}}
			{{dbSet 0 "names" ($names.AppendSlice (cslice $name))}}
			{{$Q = "Ref"}}
			{{dbSet .User.ID "Q" $Q}}
			{{$msg := sendMessageRetID nil (joinStr "" .User.Mention ", were you referred by one of our members?")}}
			{{addMessageReactions nil $msg ":yes:658376626639339533" ":no:658376639322783745"}}
			{{end}}
		{{else if eq (toString $Q) "Ref2"}}{{$ref = (toString (toString .Message.Content))}}{{dbSet .User.ID "ref" $ref}}{{$Q = "Pin"}}{{dbSet .User.ID "Q" $Q}}{{sendMessage nil (joinStr "" .User.Mention ", what 4 digit pin would you like used whenever a pin may be required? (Quest Vaults, or other situations.)")}}
		{{else if eq (toString $Q) "Pin"}}{{$pin = (toString .Message.Content)}}{{dbSet .User.ID "pin" $pin}}{{$Q = "Starter"}}{{dbSet .User.ID "Q" $Q}}
			{{$msg := sendMessageRetID nil (joinStr "" .User.Mention ", where would you like to pick up your starter?\n<:Artemis:640697096814592040> Artemis - Valguero (Main Map)\n<:Manticore:640697193052766218> Manticore - Ragnarok\n<:Gryphon:640697148807184384> Gryphon - Aberration\n<:Elysian:640698387506921523> Elysian - Center\n<:Zelda:640697492274544643> Zelda - Island\n<:Phoenix:640698407199178771> Phoenix - Scorched Earth\n<:Medusa:640697225684713473> Medusa - Extinction\n<:Titan:640697323143561276> Titan - Valguero\n<:Raiden:655954199476961300> Raiden - Ragnarok\n<:Beowulf:682592075685953570> Beowulf - Genesis\n<:Osiris:748895819331141674> Osiris - Crystal Isles\n⛔ No Starter Needed")}}{{addMessageReactions nil $msg ":Artemis:640697096814592040" ":Manticore:640697193052766218" ":Gryphon:640697148807184384" ":Elysian:640698387506921523" ":Zelda:640697492274544643" ":Phoenix:640698407199178771" ":Medusa:640697225684713473" ":Titan:640697323143561276" ":Raiden:655954199476961300" "Beowulf:682592075685953570" ":Osiris:748895819331141674" "⛔" }}
		{{else if eq (toString $Q) "ParentGuardian2"}}
			{{$parentguardian = (toString .Message.Content)}}{{dbSet .User.ID "parentguardian" $parentguardian}}{{$Q = "GT2"}}{{dbSet .User.ID "Q" $Q}}{{sendMessage nil (joinStr "" .User.Mention ", what is your Gamer Tag?")}}
		{{else if eq (toString $Q) "GT2"}}{{$gt = (toString .Message.Content)}}{{dbSet .User.ID "gt" $gt}}{{$Q = "Name2"}}{{dbSet .User.ID "Q" $Q}}{{sendMessage nil (joinStr "" .User.Mention ", what is your In Game Name?")}}
		{{else if eq (toString $Q) "Name2"}}{{$name = (toString .Message.Content)}}{{dbSet .User.ID "name" $name}}{{$Q = "Pin2"}}{{dbSet .User.ID "Q" $Q}}{{sendMessage nil (joinStr "" .User.Mention ", what 4 digit pin would you like used whenever a pin may be required? (Quest Vaults, or other situations.)")}}
		{{else if eq (toString $Q) "Done"}}
			{{execCC 214 nil 0 ""}}
		{{else if eq (toString $Q) "Pin2"}}
			{{execCC 214 nil 0 ""}}
		{{end}}
	{{end}}
{{end}}
