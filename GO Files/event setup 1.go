{{/*set editor, channel, and start event sdict*/}}
{{$editor := (toString .User.ID)}}
{{$channel := "573897276569813012"}}
{{$event := sdict}}

{{/*generate event number*/}}
{{$number := (dbIncr 0 "currenteventnumber" 1)}}
{{/*set variables*/}}
{{$name := ""}}
{{$desc  := ""}}
{{$color  := ""}}
{{$style := ""}}
{{$team  := 1}}
{{$types  := ""}}
{{$start := .timehour}}
{{$runner  := cslice}}
{{$participants := cslice}}
{{$waitlist := cslice}}

{{$embed := cembed 
    "title" "Event Panel - Setup in Progress" 
	"footer" (sdict "text" "Event starts") }}
	
{{sendMessage nil "Event Creator Started\nWhat would you like to call this event?" }}


{{/*post embed and retain ID*/}}
{{$msgID := sendMessageRetID $channel $embed}}

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
{{dbSet $msgID "event" $event}}
{{dbSet 0 "eventeditor" $editor}}
{{dbSet 0 "currenteventediting" (toString $msgID)}}
{{dbSet 0 "EventEditing" "name"}}