{{ $tz := "America/New_York" }}
{{ $location := (newDate 0 0 0 0 0 0 $tz).Location }}
{{ $time := currentTime.In $location }}
{{ $formattedtime := formatTime $time "Monday January 2, 2006, 03:04pm MST"}}
{{$eventnumber := 151}}
	{{$eventname := "Dino Races"}}
	{{$eventdesc  := "Who doesn’t love racing round on dinos!!"}}
	{{$color  := 258528}}
	{{$eventstyle := "PVE"}}
	{{$eventteam  := 1}}
	{{$eventtypes  := "Race"}}
	{{$eventstart :=  "22 Aug 2020 21:00"}}
	{{$eventrunner  := (cslice "test")}}
	{{$eventparticipants := (cslice "test")}}
	{{$eventwaitlist := (cslice "test")}}
{{$eventmax := 5}}
{{$countparticipants := 0}}
	{{$embed := cembed 
		"title" (joinStr "-" $eventnumber $eventname) 
		"description" $eventdesc 
		"color" $color 
		"fields" (cslice 
			(sdict "name" "Style" "value" $eventstyle "inline" true) 
			(sdict "name" "Team Size" "value" (toString $eventteam) "inline" true) 
			(sdict "name" "Type" "value" $eventtypes "inline" true) 
			(sdict "name" "Date/Time" "value" (toString $eventstart) "inline" false) 
			(sdict "name" "Event Runner" "value" (joinStr "" "`" $eventrunner "`") "inline" false) 
			(sdict "name" (joinStr "" "Participants " $countparticipants "/" $eventmax)  "value" (joinStr "" "`" $eventparticipants "`") "inline" false)
			(sdict "name" "Wait List" "value" (joinStr "" "`" $eventwaitlist  "`") "inline" false)) 
		"footer" (sdict "text" "Event starts")  }}
		{{$msgID := sendMessageRetID nil $embed}}
		{{addMessageReactions nil $msgID "🌟" "✅" "❔"}} 

		