{{/*Event*/}}
{{$emojis := cslice "🌟" "✅" "❔"}}
{{if in $emojis .Reaction.Emoji.Name}}

{{$userid := (toString .User.ID)}}
{{$db := dbGet .ReactionMessage.ID "event"}}
{{$event := sdict}}{{range $k,$v := $db.Value}}{{$event.Set $k $v}}{{end}}


{{$DMReminder := toDuration (dbGet 0 "EventDMReminder").Value}}

{{$name := $event.name}}
{{$time := $event.time}}
{{$start := $event.start}}
{{$newtime := $time.Add (toDuration (mult -1 $DMReminder))}}
{{$tz := "America/New_York" }}
	{{$location := (newDate 0 0 0 0 0 0 $tz).Location }} 
	{{$timeloc := $newtime.In $location}}
	{{$timeLeft := (toDuration ($timeloc.Sub currentTime))}}
	{{$til := humanizeDurationSeconds $timeLeft}}


{{$runner  := $event.runner}}
{{$participants := $event.participants}}
{{$waitlist := $event.waitlist}}
{{$runner = (cslice).AppendSlice $runner}}
{{$participants = (cslice).AppendSlice $participants}}
{{$waitlist = (cslice).AppendSlice $waitlist}}
{{if .ReactionAdded}} 
	{{if eq .Reaction.Emoji.Name "🌟" }}
		{{if hasRoleID 573210843840249889}}
			{{$runner = $runner.Append (toString .User.ID)}}
			{{scheduleUniqueCC 197 573896118719610880 $timeLeft.Seconds (joinStr "" .ReactionMessage.ID "-" .User.ID "-runner") (sdict "event" (toString $name) "role" "runner")}}
		{{end}}
	{{else if eq .Reaction.Emoji.Name "✅" }}
	{{$participants = $participants.Append (toString .User.ID)}}
	{{scheduleUniqueCC 197 573896118719610880 $timeLeft.Seconds (joinStr "" .ReactionMessage.ID "-" .User.ID "-participant") (sdict "event" (toString $name) "role" "participant")}}
		{{else if eq .Reaction.Emoji.Name "❔" }}
	{{$waitlist = $waitlist.Append (toString .User.ID)}}
	{{scheduleUniqueCC 197 573896118719610880 $timeLeft.Seconds (joinStr "" .ReactionMessage.ID "-" .User.ID "-waitlist") (sdict "event" (toString $name) "role" "waitlist")}}
	{{end}}
{{else}}
	{{if eq .Reaction.Emoji.Name "🌟" }}
		{{if hasRoleID 573210843840249889}}
			{{$newrunner := cslice}}
				{{range $runner}}
					{{if eq . $userid}}
					{{else}}
					{{$newrunner = $newrunner.Append .}}
					{{end}}
				{{$runner = $newrunner}}
				{{end}}
			{{cancelScheduledUniqueCC 197 (joinStr "" .ReactionMessage.ID "-" .User.ID "-runner")}}
	{{end}}
	{{else if eq .Reaction.Emoji.Name "✅" }}
	{{$newparticipants := cslice}}
		{{range $participants}}
			{{if eq . $userid}}
			{{else}}
			{{$newparticipants = $newparticipants.Append .}}
			{{end}}
		{{$participants = $newparticipants}}
		{{end}}
		{{cancelScheduledUniqueCC 197 (joinStr "" .ReactionMessage.ID "-" .User.ID "-participant")}}
	{{else if eq .Reaction.Emoji.Name "❔" }}
	{{$newwaitlist := cslice}}
		{{range $waitlist}}
			{{if eq . $userid}}
			{{else}}
			{{$newwaitlist = $newwaitlist.Append .}}
			{{end}}
		{{$waitlist = $newwaitlist}}
		{{end}}
		{{cancelScheduledUniqueCC 197 (joinStr "" .ReactionMessage.ID "-" .User.ID "-waitlist")}}
	{{end}}
{{end}}
{{$event.Set "runner" $runner}}
{{$event.Set "participants" $participants}}
{{$event.Set "waitlist" $waitlist}}
{{dbSet .ReactionMessage.ID "event" $event}}
{{execCC 196 nil 0 (sdict "MessageID" (toString .ReactionMessage.ID)) }}
{{end}}