{{$pattern := `([^-]+-[^-]+)(.+)?` -}}
{{$db := dbGet 0 (toString .Channel.ID)}}
{{$ticket := sdict}}{{range $k,$v := $db.Value}}{{$ticket.Set $k $v}}{{end}}


IF CHANNEL IS IN TICKET CATEGORY
{{if eq .Channel.ParentID 635305499600093214}}
	IF REACTION IS CLAIM
	{{if eq .Reaction.Emoji.ID 660605488609755146}}
		IF TICKET IS ON HOLD
		{{if $ticket.hold}}
		{{editChannelName nil (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1)}}
		{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has removed ticket " .Channel.Name " from hold." )}}
		{{$ticket.Del "hold"}}
		{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has claimed ticket " .Channel.Name )}}
		{{editChannelName nil (joinStr "" (getChannel nil).Name "-" .Member.Nick)}}
		{{$ticket.Set "claimer" .User.ID}}
		{{dbSet 0 (toString .Channel.ID) $ticket}}
		ELSE
		{{else}}



			IF TICKET HAS ALREADY BEEN CLAIMED
			{{if $ticket.claimer}}
				IF REACTOR IS CLAIMER
				{{if eq $ticket.claimer .User.ID}}
					THAN UNCLAIM TICKET, REMOVE CLAIMANTS NAME FROM CHANNEL NAME
					{{editChannelName nil (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1)}}
					{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has unclaimed ticket " .Channel.Name )}}
					{{$ticket.Del "claimer"}}
					{{dbSet 0 (toString .Channel.ID) $ticket}}
				ELSE IF REACTOR IS NOT CLAIMER
				{{else}}
					THAN CHANGE CLAIMANT TO NEW CLAIMER, CHANGE CHANNEL NAME TO REFLECT
					{{editChannelName nil (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1)}}
					{{sendMessage 653058600230584353 (joinStr "" (getMember ($ticket.claimer)).Nick " has unclaimed ticket " .Channel.Name )}}
					{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has claimed ticket " .Channel.Name )}}
					{{editChannelName nil (joinStr "" (getChannel nil).Name "-" .Member.Nick)}}
					{{$ticket.Set "claimer" .User.ID}}
					{{dbSet 0 (toString .Channel.ID) $ticket}}
				END
				{{end}}
			ELSE IF TICKET HAS NOT ALREADY BEEN CLAIMED
			{{else}}
				THAN CLAIM TICKET, CHANGE CHANNEL NAME TO REFLECT
				{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has claimed ticket " .Channel.Name )}}
				{{editChannelName nil (joinStr "" (getChannel nil).Name "-" .Member.Nick)}}
				{{$ticket.Set "claimer" .User.ID}}
				{{dbSet 0 (toString .Channel.ID) $ticket}}
			END
			{{end}}
		END
		{{end}}
	ELSE IF REACTION IS CLOSE
	{{else if eq .Reaction.Emoji.ID 660605465721569334}}
		THAN CLOSE TICKET
		{{exec "tickets close" (joinStr "" "Ticket closed by " .User.Username)}}
	ELSE IF REACTION IS HOLD
	{{else if eq .Reaction.Emoji.ID 631218688099614724}}
		IF TICKET IS ALREADY ON HOLD
		{{if $ticket.hold}}
			THAN REMOVE HOLD FROM CHANNEL NAME, RETURN CLAIMERS NAME TO CHANNEL NAME
			{{editChannelName nil (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1)}}
			{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has removed ticket " .Channel.Name " from hold." )}}
			{{$ticket.Del "hold"}}
			{{dbSet 0 (toString .Channel.ID) $ticket}}
		ELSE IF TICKET IS NOT ON HOLD
		{{else}}
			IF ITS CLAIMED
			{{if $ticket.claimer}}	
			{{editChannelName nil (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1)}}
			{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has unclaimed ticket " .Channel.Name )}}
			{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has placed ticket " .Channel.Name " on hold." )}}
			{{editChannelName nil (joinStr "" (getChannel nil).Name "-hold")}}
			{{$ticket.Set "hold" "yes"}}
			{{$ticket.Del "claimer"}}
			{{dbSet 0 (toString .Channel.ID) $ticket}}

		
		
			THAN PLACE TICKET ON HOLD, CHANGE CHANNEL NAME
			{{else}}
			{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has placed ticket " .Channel.Name " on hold." )}}
			{{editChannelName nil (joinStr "" (getChannel nil).Name "-hold")}}
			{{$ticket.Set "hold" "yes"}}
			{{dbSet 0 (toString .Channel.ID) $ticket}}
		END
		{{end}}
	END
	{{end}}
END	
{{end}}



{{$pattern := `([^-]+-[^-]+)(.+)?` -}}
{{$db := dbGet 0 (toString .Channel.ID)}}
{{$ticket := sdict}}{{range $k,$v := $db.Value}}{{$ticket.Set $k $v}}{{end}}


{{if eq .Channel.ParentID 635305499600093214}}
	{{if eq .Reaction.Emoji.ID 660605488609755146}}
		{{if $ticket.hold}}
		{{editChannelName nil (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1)}}
		{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has removed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) " from hold." )}}
		{{$ticket.Del "hold"}}
		{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has claimed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) )}}
		{{sendMessage nil (joinStr "" .Member.Nick " will be assisting you.")}}
		{{editChannelName nil (joinStr "" (getChannel nil).Name "-" .Member.Nick)}}
		{{$ticket.Set "claimer" .User.ID}}
		{{dbSet 0 (toString .Channel.ID) $ticket}}
		{{else}}
			{{if $ticket.claimer}}
				{{if eq $ticket.claimer .User.ID}}
					{{editChannelName nil (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1)}}
					{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has unclaimed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) )}}
					{{sendMessageNoEscape nil (joinStr "" .Member.Nick " will no longer be assisting you. Please be patient, and someone else from <@&634598489732546588> will be with you shortly. ")}}
					{{$ticket.Del "claimer"}}
					{{dbSet 0 (toString .Channel.ID) $ticket}}
				{{else}}
					{{editChannelName nil (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1)}}
					{{sendMessage 653058600230584353 (joinStr "" (getMember ($ticket.claimer)).Nick " has unclaimed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) )}}
					{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has claimed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) )}}
					{{sendMessage nil (joinStr "" .Member.Nick " will be assisting you.")}}
					{{editChannelName nil (joinStr "" (getChannel nil).Name "-" .Member.Nick)}}
					{{$ticket.Set "claimer" .User.ID}}
					{{dbSet 0 (toString .Channel.ID) $ticket}}
				{{end}}
			{{else}}
				{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has claimed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) )}}
				{{editChannelName nil (joinStr "" (getChannel nil).Name "-" .Member.Nick)}}
				{{sendMessage nil (joinStr "" .Member.Nick " will be assisting you.")}}
				{{$ticket.Set "claimer" .User.ID}}
				{{dbSet 0 (toString .Channel.ID) $ticket}}
			{{end}}
		{{end}}
	{{else if eq .Reaction.Emoji.ID 660605465721569334}}
	{{sendMessage nil (joinStr "" .Member.Nick " has closed this ticket. If you have further issues, please open a new ticket. This channel will be deleted in 30 seconds.")}}
	{{sleep 30}}
	{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has closed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) "." )}}
	{{sendMessage 634442883440836609 (joinStr "" .Member.Nick " has closed ticket " .Channel.Name "." )}}
	{{exec "tickets close" (joinStr "" "Ticket closed by " .User.Username)}}

	{{else if eq .Reaction.Emoji.ID 631218688099614724}}
		{{if $ticket.hold}}
			{{editChannelName nil (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1)}}
			{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has removed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) " from hold." )}}
			{{$ticket.Del "hold"}}
			{{dbSet 0 (toString .Channel.ID) $ticket}}
		{{else}}
			{{if $ticket.claimer}}	
			{{editChannelName nil (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1)}}
			{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has unclaimed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) )}}
			{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has placed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) " on hold." )}}
			{{sendMessage nil (joinStr "" .Member.Nick " has placed this ticket on hold.")}}
			{{editChannelName nil (joinStr "" (getChannel nil).Name "-hold")}}
			{{$ticket.Set "hold" "yes"}}
			{{$ticket.Del "claimer"}}
			{{dbSet 0 (toString .Channel.ID) $ticket}}
			{{else}}
			{{sendMessage 653058600230584353 (joinStr "" .Member.Nick " has placed ticket " (index (index (reFindAllSubmatches $pattern .Channel.Name) 0) 1) " on hold." )}}
			{{sendMessage nil (joinStr "" .Member.Nick " has placed this ticket on hold.")}}
			{{editChannelName nil (joinStr "" (getChannel nil).Name "-hold")}}
			{{$ticket.Set "hold" "yes"}}
			{{dbSet 0 (toString .Channel.ID) $ticket}}
			{{end}}
		{{end}}
	{{end}}
{{end}}