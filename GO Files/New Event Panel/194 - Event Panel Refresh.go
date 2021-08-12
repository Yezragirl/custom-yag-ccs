{{$channel := "580400941048528900"}}
{{$msgID := .ExecData.MessageID}}
{{$ts := .TimeSecond}}
{{$db := dbGet (toInt64 $msgID) "event"}}
{{$event := sdict}}{{range $k,$v := $db.Value}}{{$event.Set $k $v}}{{end}}
{{$end := "No"}}
	{{$pcount := 0}}
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
	{{$runner  := $event.runner}}
	{{$participants := $event.participants}}
	{{$waitlist := $event.waitlist}}
	{{$tz := "America/New_York" }}
	{{$location := (newDate 0 0 0 0 0 0 $tz).Location }} 
	{{$timeLeft := $time.Sub currentTime}}
	{{$til := humanizeDurationMinutes $timeLeft}}
	{{$runnerslist := "\u180E"}}
			{{$participantslist := "\u180E"}}
			{{$waitlistlist := "\u180E"}}

 	{{if lt $timeLeft (mult .TimeSecond 30)}}
   			{{$timeLeft := $time.Sub currentTime}}
			{{$til = humanizeDurationMinutes $timeLeft}}
			{{range $event.runner}}
				{{$username := or ($mem :=getMember .).Nick $mem.User.Username}}
				{{$runnerslist = print $username "\n" $runnerslist}}
			{{end}}
			{{range $event.participants}}
				{{$username := or ($mem :=getMember .).Nick $mem.User.Username}}
				{{$participantslist = print $username "\n" $participantslist}}
				{{$pcount = add $pcount 1}}
			{{end}}
			{{range $event.waitlist}}
				{{$username := or ($mem :=getMember .).Nick $mem.User.Username}}
				{{$waitlistlist = print $username "\n" $waitlistlist}}
			{{end}}
			{{$embed := cembed 
   			"title" (joinStr "" $number " - **" $name "**")
  			"description" $desc 
   			"color" $color 
	    		"fields" (cslice 
	        	(sdict "name" "Style" "value" (toString $style) "inline" true) 
	        	(sdict "name" "Team Size" "value" (toString $team) "inline" true) 
			(sdict "name" "Type" "value" (toString $types) "inline" true) 
			(sdict "name" "Map" "value" (toString $mapname) "inline" true) 
 			(sdict "name" "Date/Time" "value" (toString $start) "inline" false)
			(sdict "name" "Time Until" "value" "Event Begins in Less than 30 Seconds" "inline" false) 
       			(sdict "name" "Event Runner" "value" (toString (joinStr "" "```" $runnerslist "```")) "inline" false) 
			(sdict "name" (joinStr "" "Participants " $pcount "/" $max) "value" (toString (joinStr "" "```" $participantslist "```")) "inline" false)
			(sdict "name" "Waitlist" "value" (toString (joinStr "" "```" $waitlistlist "```")) "inline" false)) 
    			"footer" (sdict "text" "Event starts") 
			"timestamp" $time }}
			{{editMessage $channel (toInt64 $msgID) $embed}}
			{{addMessageReactions $channel (toInt64 $msgID) "🌟" "✅" "❔"}} 
	 		{{if gt $timeLeft $ts}} {{sleep 1}} {{end}}
			{{$end = "Yes"}}
   	{{else if lt $timeLeft $ts}}

		 {{range $event.runner}}
			{{$username := or ($mem :=getMember .).Nick $mem.User.Username}}
			 {{$runnerslist = print $username "\n" $runnerslist}}
		 {{end}}
		 {{range $event.participants}}
			{{$username := or ($mem :=getMember .).Nick $mem.User.Username}}
			 {{$participantslist = print $username "\n" $participantslist}}
			 {{$pcount = add $pcount 1}}
		 {{end}}
		 {{range $event.waitlist}}
			{{$username := or ($mem :=getMember .).Nick $mem.User.Username}}
			 {{$waitlistlist = print $username "\n" $waitlistlist}}
		 {{end}}
		 {{$embed := cembed 
			 "title" (joinStr "" $number " - **" $name "**")
			 "description" $desc 
			 "color" $color 
			 "fields" (cslice 
				 (sdict "name" "Style" "value" (toString $style) "inline" true) 
				 (sdict "name" "Team Size" "value" (toString $team) "inline" true) 
				 (sdict "name" "Type" "value" (toString $types) "inline" true) 
				 (sdict "name" "Map" "value" (toString $mapname) "inline" true) 
				 (sdict "name" "Date/Time" "value" (toString $start) "inline" false)
				 (sdict "name" "Time Until" "value" "**Event has Begun**" "inline" false) 
				 (sdict "name" "Event Runner" "value" (toString (joinStr "" "```" $runnerslist "```")) "inline" false) 
				 (sdict "name" (joinStr "" "Participants " $pcount "/" $max) "value" (toString (joinStr "" "```" $participantslist "```")) "inline" false)
				 (sdict "name" "Waitlist" "value" (toString (joinStr "" "```" $waitlistlist "```")) "inline" false)) 
				 "footer" (sdict "text" "Event starts") 
				 "timestamp" $time }}
		 {{editMessage $channel (toInt64 $msgID) $embed}}
		 {{addMessageReactions $channel (toInt64 $msgID) "🌟" "✅" "❔"}} 
	{{else}}
		{{$til = (joinStr "" $til)}}

		{{range $event.runner}}
			{{$username := or ($mem :=getMember .).Nick $mem.User.Username}}
			{{$runnerslist = print $username "\n" $runnerslist}}
		{{end}}
		{{range $event.participants}}
			{{$username := or ($mem :=getMember .).Nick $mem.User.Username}}
			{{$participantslist = print $username "\n" $participantslist}}
			{{$pcount = add $pcount 1}}
		{{end}}
		{{range $event.waitlist}}
			{{$username := or ($mem :=getMember .).Nick $mem.User.Username}}
			{{$waitlistlist = print $username "\n" $waitlistlist}}
		{{end}}
		{{$embed := cembed 
    	"title" (joinStr "" $number " - **" $name "**")
    	"description" $desc 
    	"color" $color 
    	"fields" (cslice 
        (sdict "name" "Style" "value" (toString $style) "inline" true) 
        (sdict "name" "Team Size" "value" (toString $team) "inline" true) 
		(sdict "name" "Type" "value" (toString $types) "inline" true) 
		(sdict "name" "Map" "value" (toString $mapname) "inline" true) 
	        (sdict "name" "Date/Time" "value" (toString $start) "inline" false)
		(sdict "name" "Time Until" "value" (joinStr "" "<t:" $start.Unix ":R>") "inline" false) 
	        (sdict "name" "Event Runner" "value" (toString (joinStr "" "```" $runnerslist "```")) "inline" false) 
		(sdict "name" (joinStr "" "Participants " $pcount "/" $max) "value" (toString (joinStr "" "```" $participantslist "```")) "inline" false)
		(sdict "name" "Waitlist" "value" (toString (joinStr "" "```" $waitlistlist "```")) "inline" false)) 
 	   	"footer" (sdict "text" "Event starts") 
		"timestamp" $time }}
		{{editMessage $channel (toInt64 $msgID) $embed}}
		{{addMessageReactions $channel (toInt64 $msgID) "🌟" "✅" "❔"}} 
	{{end}}
		{{if eq $end "No"}}
{{execCC 194 nil 30 (sdict "MessageID" (toInt64 $msgID) "TimeLeft" $time) }}
{{end}}

