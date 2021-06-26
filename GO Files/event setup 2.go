{{/*set editor, channel, and start event sdict*/}}
{{$editor := (dbGet 0 "eventeditor").Value }}
{{$channel := "573897276569813012"}}
{{$timeString := ""}}
{{$formattedTime := currentTime}}
{{$step := (dbGet 0 "EventEditing").Value}}

{{if eq $editor (toString .User.ID)}}
	{{$evdb := (dbGet 0 "currenteventediting").Value }}
	{{$db := dbGet (toInt64 $evdb) "event"}}
	{{$event := sdict}}{{range $k,$v := $db.Value}}{{$event.Set $k $v}}{{end}}

	{{/*set variables*/}}
	{{$number := $event.number}}
	{{$name := $event.name}}
	{{$desc  := $event.desc}}
	{{$color  := $event.color}}
	{{$style := $event.style}}
	{{$team  := $event.team}}
	{{$types  := $event.types}}
	{{$start := $event.start}}
	{{$runner  := $event.runner}}
	{{$til := (toDuration "1h")}}
	{{$participants := $event.participants}}
	{{$waitlist := $event.waitlist}}
	{{$step := (dbGet 0 "EventEditing").Value}}
	{{$msgID := (dbGet 0 "currenteventediting").Value}}

	{{/*Questions*/}}
	{{if eq (toString $step) "name"}}
		{{$name = (toString .Message.Content)}}
		{{sendMessage nil "Enter short event description." }}
		{{$step = "desc"}}
	{{else if eq (toString $step) "desc"}}
		{{$desc = (toString .Message.Content)}}
		{{sendMessage nil "What is the type of this event? PVE or PVP?" }}
		{{$step = "style"}}
	{{else if eq (toString $step) "style"}}
		{{$style = (toString .Message.Content)}}
		{{if eq $style "PVP"}}
			{{$color = 15158332}}
		{{else}}
			{{$color = 3447003}}
		{{end}}
		{{sendMessage nil "How many people on a team? Enter a number only" }}
		{{$step = "team"}}
	{{else if eq (toString $step) "team"}}
		{{$team = (toInt .Message.Content)}}
		{{sendMessage nil "What type of event is this? List all types, separated by commas. (i.e. Maze, Puzzle, etc.)" }}
		{{$step = "types"}}
	{{else if eq (toString $step) "types"}}
		{{$types = (toString .Message.Content)}}
		{{sendMessage nil "When will this event start? Format Date and Time for your timezone if you have set it using setz." }}
		{{$step = "start"}}
	{{else if eq (toString $step) "start"}}
		{{$timeString = (toString .Message.Content)}}
		{{$step = "time"}}
		{{execCC 185 nil 2 (sdict "timeString" $timeString "CC" .CCID) }}
	{{else if eq (toString $step) "time"}}
		{{/*Time Conversion*/}}
		{{if .ExecData}}
			{{$formattedTime = .ExecData.timeConverted}}
		{{end}}
		{{$tz := "America/New_York" }}{{ $location := (newDate 0 0 0 0 0 0 $tz).Location }} 
		{{$time := $formattedTime.In $location }}
		{{$start = formatTime $time "Mon Jan 2,2006 3:04 PM MST" }}
		{{$til = (toDuration ($start.Sub $currentTime))}}
		{{$step = "complete"}}
		{{$editor = "None"}}
	{{end}}


{{/*Generate New Embed*/}}
{{if eq $step "complete"}}
{{$runnerslist := "\u180E"}}
{{$participantslist := "\u180E"}}
{{$waitlistlist := "\u180E"}}
	{{range $event.runner}}
	{{$runnerslist = print . "\n" $runnerslist}}
	{{end}}
	{{range $event.participants}}
	{{$participantslist = print . "\n" $participantslist}}
	{{end}}
	{{range $event.waitlist}}
	{{$waitlistlist = print . "\n" $waitlistlist}}
	{{end}}
	{{$runnerslist}}
	{{$participantslist}}
	{{$waitlistlist}}

{{$embed := cembed 
    "title" "Upcoming Event"
    "description" (joinStr "" $number " - " $name " \n " $desc) 
    "color" $color 
    "fields" (cslice 
        (sdict "name" "Style" "value" (toString $style) "inline" true) 
        (sdict "name" "Team Size" "value" (toString $team) "inline" true) 
        (sdict "name" "Type" "value" (toString $types) "inline" true) 
        (sdict "name" "Date/Time" "value" (toString $start) "inline" false)
		(sdict "name" "Time Until" "value" (toString $til) "inline" false) 
        (sdict "name" "Event Host" "value" (toString (joinStr "" "```" $runnerslist "```")) "inline" false) 
		(sdict "name" "Particiants" "value" (toString (joinStr "" "```" $participantslist "```")) "inline" false)
		(sdict "name" "Waitlist" "value" (toString (joinStr "" "```" $waitlistlist "```")) "inline" false)) 
    "footer" (sdict "text" "Event starts") 
	"timestamp" $formattedTime }}
{{editMessage $channel (toInt64 $msgID) $embed}}
{{addMessageReactions $channel (toInt64 $msgID) "🌟" "✅" "❔"}} 
{{end}}

{{/*save all variables to sdict*/}}
{{$event.Set "number" $number}}
{{$event.Set "name" $name}}
{{$event.Set "desc"  $desc}}
{{$event.Set "color"  $color}}
{{$event.Set "style" $style}}
{{$event.Set "team"  $team}}
{{$event.Set "types"  $types}}
{{$event.Set "start" $start}}
{{$event.Set "runner" $runner}}
{{$event.Set "participants" $participants}}
{{$event.Set "waitlist" $waitlist}}

{{/*save sdict to db, save the current event being edited to db*/}}
{{dbSet (toInt64 $msgID) "event" $event}}
{{dbSet 0 "eventeditor" $editor}}
{{dbSet 0 "currenteventediting" $msgID}}
{{dbSet 0 "EventEditing" $step}}



{{end}}