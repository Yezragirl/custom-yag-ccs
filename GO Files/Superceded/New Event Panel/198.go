{{$msgID := .ExecData.id}}

{{$db := dbGet (toInt64 $msgID) "event"}}
{{$event := sdict}}{{range $k,$v := $db.Value}}{{$event.Set $k $v}}{{end}}

{{/*set variables*/}}
{{$number := $event.number}}
{{$pcount := 0}}
{{$name := $event.name}}
{{$desc  := $event.desc}}
{{$color  := $event.color}}
{{$style := $event.style}}
{{$mapname := $event.mapname}}
{{$team  := $event.team}}
{{$types  := $event.types}}
{{$max := $event.max}}
{{$start := $event.start}}
{{$time := $event.time}}
{{$runner  := $event.runner}}
{{$participants := $event.participants}}
{{$waitlist := $event.waitlist}}
{{$tz := "America/New_York" }}
{{$location := (newDate 0 0 0 0 0 0 $tz).Location }} 
{{$timeloc := $time.In $location}}
{{$timeLeft := (toDuration ($timeloc.Sub currentTime))}}
{{$til := humanizeDurationSeconds $timeLeft}}


{{$embed := cembed 
    "title" (joinStr "" "EVENT REMINDER! " $number " - **" $name "**")
    "description" $desc 
    "color" $color 
    "fields" (cslice 
        (sdict "name" "__**Style**__" "value" (toString $style) "inline" true) 
        (sdict "name" "__**Team Size**__" "value" (toString $team) "inline" true) 
		(sdict "name" "__**Type**__" "value" (toString $types) "inline" true) 
		(sdict "name" "__**Map**__" "value" (toString $mapname) "inline" true) 
        (sdict "name" "__**Date/Time**__" "value" (toString $start) "inline" false)
		(sdict "name" "__**Time Until**__" "value" (joinStr "" "<t:" $start.Unix ":R>") "inline" false)) 
    "footer" (sdict "text" "Event starts") 
	"timestamp" $time }}

	{{sendMessage 573245421066125324 $embed}}

