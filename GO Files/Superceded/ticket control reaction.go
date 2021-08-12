{{$pattern := `([^-]+-[^-]+)(.+)?` -}}
{{$db := dbGet (toInt64 .Channel.ID) "ticket_control"}}
{{$ticket := sdict}}{{range $k,$v := $db.Value}}{{$ticket.Set $k $v}}{{end}}


{{if eq .Channel.ParentID 635305499600093214}}
	{{if eq .Reaction.Emoji.ID 660605488609755146}}
		{{if $ticket.hold}}
		{{editChannelName nil (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1)}}
		{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has removed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) " from hold." )}}
		{{$ticket.Del "hold"}}
		{{sendMessage nil (joinStr "" .Member.Nick " will be assisting you.")}}
		{{editChannelName nil (joinStr "" (getChannel nil).Name "-" .Member.Nick)}}
		{{$ticket.Set "claimer" .User.ID}}
		{{dbSet (toInt64 .Channel.ID) "ticket_control" $ticket}}
		{{else}}
			{{if $ticket.claimer}}
				{{if eq $ticket.claimer .User.ID}}
					{{editChannelName nil (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1)}}
					{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has unclaimed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) )}}
					{{sendMessageNoEscape nil (joinStr "" .Member.Nick " will no longer be assisting you. Please be patient, and someone else from <@&634598489732546588> will be with you shortly. ")}}
					{{$ticket.Del "claimer"}}
					{{dbSet (toInt64 .Channel.ID) "ticket_control" $ticket}}
				{{else}}
					{{editChannelName nil (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1)}}
					{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has claimed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) )}}
					{{sendMessage nil (joinStr "" .Member.Nick " will be assisting you.")}}
					{{editChannelName nil (joinStr "" (getChannel nil).Name "-" .Member.Nick)}}
					{{$ticket.Set "claimer" .User.ID}}
					{{dbSet (toInt64 .Channel.ID) "ticket_control" $ticket}}
				{{end}}
			{{else}}
				{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has claimed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) )}}
				{{editChannelName nil (joinStr "" (getChannel nil).Name "-" .Member.Nick)}}
				{{sendMessage nil (joinStr "" .Member.Nick " will be assisting you.")}}
				{{$ticket.Set "claimer" .User.ID}}
				{{dbSet (toInt64 .Channel.ID) "ticket_control" $ticket}}
			{{end}}
		{{end}}
	{{else if eq .Reaction.Emoji.ID 660605465721569334}}
	{{sendMessage nil (joinStr "" .Member.Nick " has closed this ticket. If you have further issues, please open a new ticket. This channel will be deleted in 30 seconds.")}}
	{{sleep 30}}
	{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has closed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) "." )}}
	{{sendMessage 634442883440836609 (joinStr "" .Member.Nick " has closed ticket " .Channel.Name "." )}}
	{{dbDel (toInt64 .Channel.ID) "ticket_control"}}
	{{exec "tickets close" (joinStr "" "Ticket closed by " .User.Username)}}

	{{else if eq .Reaction.Emoji.ID 631218688099614724}}
		{{if $ticket.hold}}
			{{editChannelName nil (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1)}}
			{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has removed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) " from hold." )}}
			{{$ticket.Del "hold"}}
			{{dbSet (toInt64 .Channel.ID) "ticket_control" $ticket}}
		{{else}}
			{{if $ticket.claimer}}	
			{{editChannelName nil (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1)}}
			{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has placed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) " on hold." )}}
			{{sendMessage nil (joinStr "" .Member.Nick " has placed this ticket on hold.")}}
			{{editChannelName nil (joinStr "" (getChannel nil).Name "-hold")}}
			{{$ticket.Set "hold" "yes"}}
			{{$ticket.Del "claimer"}}
			{{dbSet (toInt64 .Channel.ID) "ticket_control" $ticket}}
			{{else}}
			{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has placed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) " on hold." )}}
			{{sendMessage nil (joinStr "" .Member.Nick " has placed this ticket on hold.")}}
			{{editChannelName nil (joinStr "" (getChannel nil).Name "-hold")}}
			{{$ticket.Set "hold" "yes"}}
			{{dbSet (toInt64 .Channel.ID) "ticket_control" $ticket}}
			{{end}}
		{{end}}
	{{end}}
{{end}}