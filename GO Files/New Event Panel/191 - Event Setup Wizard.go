{{$args := parseArgs 1 "correct syntax is `-create EventName`"
(carg "string" "eventname")}}

{{/*set editor, channel, and start event sdict*/}}
{{$editor := (toString .User.ID)}}
{{$channel := "580400941048528900"}}
{{$event := sdict}}

{{/*generate event number*/}}
{{$number := (dbIncr 0 "currenteventnumber" 1)}}
{{/*set variables*/}}
{{$name := ($args.Get 0)}}
{{$desc  := ""}}
{{$color  := ""}}
{{$style := ""}}
{{$mapname := ""}}
{{$team  := 1}}
{{$types  := ""}}
{{$start := .timehour}}
{{$runner  := cslice}}
{{$participants := cslice}}
{{$waitlist := cslice}}

{{$embed := cembed 
    "title" "Event Panel - Setup in Progress" 
	"footer" (sdict "text" "Event starts") }}

{{/*post embed and retain ID*/}}
{{$msgID := sendMessageRetID $channel $embed}}

{{/*save all variables to sdict*/}}
{{$event.Set "number" $number}}
{{$event.Set "name" $name}}
{{$event.Set "desc"  $desc}}
{{$event.Set "color"  $color}}
{{$event.Set "mapname" $mapname}}
{{$event.Set "style" $style}}
{{$event.Set "team"  $team}}
{{$event.Set "types"  $types}}
{{$event.Set "start" $start}}
{{$event.Set "runner" $runner}}
{{$event.Set "participants" $participants}}
{{$event.Set "waitlist" $waitlist}}

{{/*save sdict to db, save the current event being edited to db*/}}
{{sleep 2}}
{{dbSet $msgID "event" $event}}
{{dbSet 0 "eventeditor" $editor}}
{{dbSet 0 "currenteventediting" (toString $msgID)}}
{{dbSet 0 "EventEditing" "name"}}