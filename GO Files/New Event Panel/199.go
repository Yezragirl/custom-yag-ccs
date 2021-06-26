{{$find := (index .CmdArgs 0)}}
{{$EventID := ""}}
{{addRoleID 856607038872223764}}
{{range dbTopEntries "event" 100 0}}
{{$event := sdict}}{{range $k,$v := .Value}}{{$event.Set $k $v}}{{end}}
{{if eq (toInt $event.number) (toInt $find) }}
{{$EventID = (toString .UserID)}}
{{end}}{{end}}
{{$db := dbGet (toInt64 $EventID) "event"}}
{{$event := sdict}}{{range $k,$v := $db.Value}}{{$event.Set $k $v}}{{end}}
 
	{{$number := $event.number}}
	{{$name := $event.name}}
	{{$desc  := $event.desc}}
	{{$color  := $event.color}}
	{{$style := $event.style}}
	{{$mapname := $event.mapname}}
	{{$team  := $event.team}}
	{{$types  := $event.types}}
	{{$start := $event.start}}
	{{$time := $event.time}}
	{{$max := $event.max}}


	{{$embed := cembed 
		"title" "Event Edit Panel"
	   "description" "React to the panel with the number of the piece you'd like to edit." 
		"color" $color 
	 		"fields" (cslice 
			(sdict "name" ":one: Name" "value" (toString $name) "inline" true)
			(sdict "name" ":two: Description" "value" (toString $desc) "inline" true)
	 (sdict "name" ":three: Style" "value" (toString $style) "inline" true) 
	 (sdict "name" ":four: Team Size" "value" (toString $team) "inline" true) 
	 (sdict "name" ":five: Type" "value" (toString $types) "inline" true) 
	 (sdict "name" ":six: Map" "value" (toString $mapname) "inline" true) 
	 (sdict "name" ":seven: Max Participants" "value" (toString $max) "inline" true) 
	 (sdict "name" ":eight: Date/Time" "value" (toString $start) "inline" true))	
	 "timestamp" $time }}
{{$msg := sendMessageRetID nil $embed}}

	 {{addMessageReactions nil $msg "1️⃣" "2️⃣" "3️⃣" "4️⃣" "5️⃣" "6️⃣" "7️⃣" "8️⃣" "❌"}}
	 {{dbSet 0 "Event Panel" 1}}
	 {{dbSet 0 "Editing Event" $EventID}}
