{{$Step := toInt (dbGet 0 "Event Panel").Value}}
{{$EventID := (dbGet 0 "Editing Event").Value}}
{{$db := dbGet (toInt64 $EventID) "event"}}
{{$event := sdict}}{{range $k,$v := $db.Value}}{{$event.Set $k $v}}{{end}}
 {{$hannel := "573897276569813012"}}
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
	{{$timeLeft = $time.Sub currentTime}}
	{{$max := $event.max}}
	{{$runner  := $event.runner}}
	{{$participants := $event.participants}}
	{{$waitlist := $event.waitlist}}
	{{$pcount := 0}}
	{{$tz := "America/New_York" }}
	{{$location := (newDate 0 0 0 0 0 0 $tz).Location }} 
	{{$timestring := ""}}
	
	{{if eq $Step 1}}
	{{$Step = 0}}
	{{dbSet 0 "Event Panel" $Step}}
	{{$name = (toString .Message.Content)}}{{$Step := toInt (dbGet 0 "Event Panel").Value}}
	{{$formattedTime := ""}}
{{if .ExecData}}
	{{$formattedTime = .ExecData.timeConverted}}
	{{$Step = 9}}
	{{end}}
{{$EventID := (dbGet 0 "Editing Event").Value}}
{{$db := dbGet (toInt64 $EventID) "event"}}
{{$event := sdict}}{{range $k,$v := $db.Value}}{{$event.Set $k $v}}{{end}}
 {{$channel := "580400941048528900"}}

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
	{{$timeLeft := $time.Sub currentTime}}
	{{$til := humanizeDurationMinutes $timeLeft}}
	{{$max := $event.max}}
	{{$runner  := $event.runner}}
	{{$participants := $event.participants}}
	{{$waitlist := $event.waitlist}}
	{{$pcount := 0}}
	{{$tz := "America/New_York" }}
	{{$location := (newDate 0 0 0 0 0 0 $tz).Location }} 
	{{$timeString := ""}}
	
	{{if eq $Step 1}}
	{{$Step = 0}}
	{{dbSet 0 "Event Panel" $Step}}
	{{$name = (toString .Message.Content)}}
	Name Updating...
	{{else if eq $Step 2}}
	{{$Step = 0}}
	{{dbSet 0 "Event Panel" $Step}}
	{{$desc = (toString .Message.Content)}}
	Description Updating...
	{{else if eq $Step 3}}
	{{$Step = 3}}
	{{dbSet 0 "Event Panel" $Step}}
	{{$style = (toString .Message.Content)}}
	Style Updating...
	{{else if eq $Step 4}}
	{{$Step = 4}}
	{{dbSet 0 "Event Panel" $Step}}
	{{$team = (toInt .Message.Content)}}
	Team Size Updating...
	{{else if eq $Step 5}}
	{{$Step = 5}}
	{{dbSet 0 "Event Panel" $Step}}
	{{$types = (toString .Message.Content)}}
	Event Type Updating...
	{{else if eq $Step 6}}
	{{$Step = 6}}
	{{dbSet 0 "Event Panel" $Step}}
	{{$mapname = (toString .Message.Content)}}
	Map Updating...
	{{else if eq $Step 7}}
	{{$Step = 7}}
	{{dbSet 0 "Event Panel" $Step}}
	{{$max = (toInt .Message.Content)}}
	Max Participants Updating...
	{{else if eq $Step 8}}
	{{$Step = 8}}
	{{dbSet 0 "Event Panel" $Step}}
	{{$timeString = (toString .Message.Content)}}
	{{$Step = 9}}
	{{execCC 185 nil 1 (sdict "timeString" $timeString "CC" .CCID) }}
	{{else if eq $Step 9}}
	{{/*Time Conversion*/}}
		{{$time = $formattedTime.In $location }}
	Time: {{$time}}
	{{$start = formatTime $time "Mon Jan 2, 2006 3:04 PM MST" }}
	Start: {{$start}}
	{{$timeLeft = $time.Sub currentTime}}
	{{$til = humanizeDurationMinutes $timeLeft}}
	{{end}}

	{{/*save all variables to sdict*/}}
{{$event.Set "number" $number}}
{{$event.Set "name" $name}}
{{$event.Set "desc"  $desc}}
{{$event.Set "color"  $color}}
{{$event.Set "style" $style}}
{{$event.Set "mapname" $mapname}}
{{$event.Set "team"  $team}}
{{$event.Set "max" $max}}
{{$event.Set "types"  $types}}
{{$event.Set "start" $start}}
{{$event.Set "time" $time}}
{{$event.Set "runner" $runner}}
{{$event.Set "participants" $participants}}
{{$event.Set "waitlist" $waitlist}}

{{/*save sdict to db, save the current event being edited to db*/}}
{{dbSet (toInt64 $EventID) "event" $event}}
{{removeRoleID 856607038872223764}}




	Name Updating...
	{{else if eq $Step 2}}
	{{$Step = 0}}
	{{dbSet 0 "Event Panel" $Step}}
	{{$desc = (toString .Message.Content)}}
	Description Updating...
	{{else if eq $Step 3}}
	{{$Step = 3}}
	{{dbSet 0 "Event Panel" $Step}}
	{{$style = (toString .Message.Content)}}
	{{else if eq $Step 4}}
	{{$Step = 4}}
	{{dbSet 0 "Event Panel" $Step}}
	{{$team = (toInt .Message.Content)}}
	{{else if eq $Step 5}}
	{{$Step = 5}}
	{{dbSet 0 "Event Panel" $Step}}
	{{$types = (toString .Message.Content)}}
	{{else if eq $Step 6}}
	{{$Step = 6}}
	{{dbSet 0 "Event Panel" $Step}}
	{{$mapname = (toString .Message.Content)}}
	{{else if eq $Step 7}}
	{{$Step = 7}}
	{{dbSet 0 "Event Panel" $Step}}
	{{$max = (toInt .Message.Content)}}
	{{else if eq $Step 8}}
	{{$Step = 8}}
	{{dbSet 0 "Event Panel" $Step}}
	{{$timeString = (toString .Message.Content)}}
	{{$Step = "time"}}
	{{execCC 185 nil 1 (sdict "timeString" $timeString "CC" .CCID) }}
	{{else if eq (toString $Step) "time"}}
	{{/*Time Conversion*/}}
	{{if .ExecData}}
		{{$formattedTime = .ExecData.timeConverted}}
	{{end}}
	{{$time = $formattedTime.In $location }}
	{{$start = formatTime $time "Mon Jan 2, 2006 3:04 PM MST" }}
	{{$timeLeft = $time.Sub currentTime}}
	{{$til = humanizeDurationMinutes $timeLeft}}
	{{end}}

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

	{{/*save all variables to sdict*/}}
{{$event.Set "number" $number}}
{{$event.Set "name" $name}}
{{$event.Set "desc"  $desc}}
{{$event.Set "color"  $color}}
{{$event.Set "style" $style}}
{{$event.Set "mapname" $mapname}}
{{$event.Set "team"  $team}}
{{$event.Set "max" $max}}
{{$event.Set "types"  $types}}
{{$event.Set "start" $start}}
{{$event.Set "time" $time}}
{{$event.Set "runner" $runner}}
{{$event.Set "participants" $participants}}
{{$event.Set "waitlist" $waitlist}}

{{/*save sdict to db, save the current event being edited to db*/}}
{{dbSet (toInt64 $EventID) "event" $event}}
{{execCC 194 nil 30 (sdict "MessageID" (toInt64 $EventID) "TimeLeft" $time) }}
{{end}}



