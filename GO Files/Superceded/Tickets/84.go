{{/*Ticket Claimer Reaction*/}}
{{ $start := currentTime }}
{{$pattern := `([^-]+-[^-]+)(.+)?` -}}
{{$db := dbGet (toInt64 .Channel.ID) "ticket_control"}}
{{$ticketnum := (dbGet 0 "ticketnum").Value}}
{{$ticket := sdict}}{{range $k,$v := $db.Value}}{{$ticket.Set $k $v}}{{end}}

{{$ticketcount := (dbGet .User.ID "tickets").Value}}
{{if lt (toInt $ticketcount) 0}}
	{{$ticketcount = 0}}
{{end}}
{{if eq .Channel.ParentID 635305499600093214}}
	{{if eq .Reaction.Emoji.ID 660605488609755146}}
		{{if $ticket.hold}}
		{{editChannelName nil (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1)}}
		{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has removed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) " from hold." )}}
		{{$ticket.Del "hold"}}
		{{sendMessage nil (joinStr "" .Member.Nick " will be assisting you.")}}
		{{editChannelName nil (joinStr "" (getChannel nil).Name "-" .Member.Nick)}}
		{{$ticket.Set "claimer" .User.ID}}
		{{$ticketcount = add $ticketcount 1}}
		{{dbSet (toInt64 .Channel.ID) "ticket_control" $ticket}}
			{{if hasRoleID 588749231150465049}} 
				{{if eq (toInt $ticketcount) 3}}
					{{$points := dbIncr .User.ID "AdminPoints" 1}}
					{{sendMessage 653058600230584353 (joinStr "" (getMember .User.ID).Nick " earned 1 point for tickets.")}}
					{{sendMessage 654151821039894547 (joinStr "" .User.Mention " earned 1 point for tickets.")}}
					{{sendMessage 654151821039894547 (joinStr "" .User.Mention " now has " (toString (toInt $points)) " points.")}}
					{{dbSet .User.ID "tickets" 0}}
				{{else}}
					{{dbSet .User.ID "tickets" $ticketcount}}
				{{end}}
			{{end}}
		{{else}}
			{{if $ticket.claimer}}
				{{if eq $ticket.claimer .User.ID}}
					{{editChannelName nil (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1)}}
					{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has unclaimed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) )}}
					{{sendMessageNoEscape nil (joinStr "" .Member.Nick " will no longer be assisting you. Please be patient, and someone else from <@&634598489732546588> will be with you shortly. ")}}
					{{$ticket.Del "claimer"}}
					{{$ticketcount = sub $ticketcount 1}}
					{{dbSet (toInt64 .Channel.ID) "ticket_control" $ticket}}
					{{dbSet .User.ID "tickets" $ticketcount}}
				{{else}}
					{{editChannelName nil (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1)}}
					{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has claimed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) )}}
					{{sendMessage nil (joinStr "" .Member.Nick " will be assisting you.")}}
					{{editChannelName nil (joinStr "" (getChannel nil).Name "-" .Member.Nick)}}
					{{$unclaimer := $ticket.Get "claimer"}}
					{{$x := dbIncr $unclaimer "tickets" -1}}
					{{$ticket.Set "claimer" .User.ID}}
					{{$ticketcount = add $ticketcount 1}}
					{{dbSet (toInt64 .Channel.ID) "ticket_control" $ticket}}
					{{if hasRoleID 588749231150465049}} 
						{{if eq (toInt $ticketcount) 3}}
						{{$points := dbIncr .User.ID "AdminPoints" 1}}
						{{sendMessage 653058600230584353 (joinStr "" (getMember .User.ID).Nick " earned 1 point for tickets.")}}
						{{sendMessage 654151821039894547 (joinStr "" .User.Mention " earned 1 point for tickets.")}}
						{{sendMessage 654151821039894547 (joinStr "" .User.Mention " now has " (toString (toInt $points)) " points.")}}
						{{dbSet .User.ID "tickets" 0}}
						{{else}}
						{{dbSet .User.ID "tickets" $ticketcount}}
						{{end}}
					{{end}}
				{{end}}
			{{else}}
				{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has claimed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) )}}
				{{editChannelName nil (joinStr "" (getChannel nil).Name "-" .Member.Nick)}}
				{{sendMessage nil (joinStr "" .Member.Nick " will be assisting you.")}}
				{{$ticket.Set "claimer" .User.ID}}
				{{$ticketcount = add $ticketcount 1}}
				{{dbSet (toInt64 .Channel.ID) "ticket_control" $ticket}}
				{{if hasRoleID 588749231150465049}} 
					{{if eq (toInt $ticketcount) 3}}
					{{$points := dbIncr .User.ID "AdminPoints" 1}}
					{{sendMessage 653058600230584353 (joinStr "" (getMember .User.ID).Nick " earned 1 point for tickets.")}}
					{{sendMessage 654151821039894547 (joinStr "" .User.Mention " earned 1 point for tickets.")}}
					{{sendMessage 654151821039894547 (joinStr "" .User.Mention " now has " (toString (toInt $points)) " points.")}}
					{{dbSet .User.ID "tickets" 0}}
					{{else}}
					{{dbSet .User.ID "tickets" $ticketcount}}
					{{end}}
				{{end}}
			{{end}}
		{{end}}
	{{else if eq .Reaction.Emoji.ID 660605465721569334}}
	{{sendMessage nil (joinStr "" .Member.Nick " has closed this ticket. If you have further issues, please open a new ticket. This channel will be deleted in 30 seconds.")}}
	{{sleep 30}}
	{{$ticketnum = sub $ticketnum 1}}
	{{dbSet 0 "ticketnum" $ticketnum}}
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
{{dbSet .User.ID "tickets" $ticketcount}}



{{$channel := (getChannel nil).Name}}
{{if reFind `weekly-timer-reset$` $channel}}
{{$reactions := cslice 796125877633024051 640697225684713473 797519029715599400 640697096814592040 797518674616909914 640697492274544643 640697193052766218 748895819331141674 796160062904729600 796123308059000882 796116046616592445 796127204164632638 655954199476961300 795749721657573446 640697323143561276 640697148807184384 640698407199178771 797515621222318091 797519872115343360 797670477081608202}}
{{$count := 0}}{{$allMatch := true}}

{{range .ReactionMessage.Reactions}}
{{if and (in $reactions .Emoji.ID) .Me (ge .Count 2)}}
	{{$count = add $count 1}}
{{else}}
	{{$allMatch = false}}
{{end}}
{{end}}

{{if and $allMatch (eq $count 20)}}

{{exec "tickets close" "All Weekly Map Resets Complete"}}
{{end}}

{{end}}

{{ $dur := currentTime.Sub $start }}
{{ if gt $dur.Seconds 5.0 }}
	{{ sendMessage 587858995012698137 (print "CC #" .CCID " took longer than 5 seconds to run.\n**Duration:** " $dur "\n**Command executed:* `" .Message.Content "`") }}
{{ end }}	

