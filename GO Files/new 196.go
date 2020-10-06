{{$msgID := .ExecData.MessageID}}
{{$channel := "573897276569813012"}}
{{$db := dbGet (toInt64 $msgID) "event"}}
{{$event := sdict}}{{range $k,$v := $db.Value}}{{$event.Set $k $v}}{{end}}
{{$number := $event.number}}
	{{$pcount := 0}}
	{{$name := $event.name}}
	{{$desc  := $event.desc}}
	{{$color  := $event.color}}
	{{$style := $event.style}}
	{{$team  := $event.team}}
	{{$types  := $event.types}}
	{{$max := $event.max}}
	{{$start := $event.start}}
	{{$runner  := $event.runner}}
	{{$participants := $event.participants}}
	{{$waitlist := $event.waitlist}}
{{$til := (toDuration "1h")}}
{{$runner = (cslice).AppendSlice $runner}}
{{$participants = (cslice).AppendSlice $participants}}
{{$waitlist = (cslice).AppendSlice $waitlist}}
{{$runnerslist := "\u180E"}}
{{$participantslist := "\u180E"}}
{{$waitlistlist := "\u180E"}}
	{{range $event.runner}}
	{{$username := (getMember .).Nick}}
	{{$runnerslist = print $username "\n" $runnerslist}}
	{{end}}
	{{range $event.participants}}
	{{$username := (getMember .).Nick}}
	{{$participantslist = print $username "\n" $participantslist}}
	{{$pcount = add $pcount 1}}
	{{end}}
	{{range $event.waitlist}}
	{{$username := (getMember .).Nick}}
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
        (sdict "name" "Date/Time" "value" (toString $start) "inline" false)
		(sdict "name" "Time Until" "value" (toString $til) "inline" false) 
        (sdict "name" "Event Runner" "value" (toString (joinStr "" "```" $runnerslist "```")) "inline" false) 
		(sdict "name" (joinStr "" "Particiants " $pcount "/" $max) "value" (toString (joinStr "" "```" $participantslist "```")) "inline" false)
		(sdict "name" "Waitlist" "value" (toString (joinStr "" "```" $waitlistlist "```")) "inline" false)) 
    "footer" (sdict "text" "Event starts") 
	"timestamp" $start }}
{{editMessage $channel (toInt64 $msgID) $embed}}
{{addMessageReactions $channel (toInt64 $msgID) "🌟" "✅" "❔"}} 

{{/*save all variables to sdict*/}}
{{$event.Set "number" $number}}
{{$event.Set "name" $name}}
{{$event.Set "desc"  $desc}}
{{$event.Set "color"  $color}}
{{$event.Set "style" $style}}
{{$event.Set "team"  $team}}
{{$event.Set "max" $max}}
{{$event.Set "types"  $types}}
{{$event.Set "start" $start}}
{{$event.Set "runner" $runner}}
{{$event.Set "participants" $participants}}
{{$event.Set "waitlist" $waitlist}}

{{/*save sdict to db, save the current event being edited to db*/}}
{{dbSet (toInt64 $msgID) "event" $event}}