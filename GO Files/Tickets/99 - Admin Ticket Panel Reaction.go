{{/*Admin Ticket Panel Reaction*/}}
{{$capture := ""}}
{{$ID := ""}}
{{$message := ""}}
{{$type := ""}}
{{$panel := (toString .ReactionMessage.ID)}}
{{$emoji := ""}}

{{if eq .Reaction.Emoji.Name "1️⃣"}}
{{$type = "weekly-timer-reset"}}
{{$message = ", it's timer reset time! Please react with your map so that I can close the ticket whenever every map is done. Deadline is Sunday."}}
{{$emoji = "1️⃣"}}

{{else if eq .Reaction.Emoji.Name "2️⃣"}}
{{$type = "violation"}}
{{$message = ", this is a violation ticket."}}
{{$emoji = "2️⃣"}}

{{else if eq .Reaction.Emoji.Name "3️⃣"}}
{{$type = "project"}}
{{$message = ", this is a project ticket."}}
{{$emoji = "3️⃣"}}

{{else if eq .Reaction.Emoji.Name "4️⃣"}}
{{$type = "claimed-dinos"}}
{{$message = ", this is a claimed dinos ticket."}}
{{$emoji = "4️⃣"}}

{{else if eq .Reaction.Emoji.Name "5️⃣"}}
{{$type = "player-issues"}}
{{$message = ", this is a player issues ticket."}}
{{$emoji = "5️⃣"}}

{{else if eq .Reaction.Emoji.Name "6️⃣"}}
{{$type = "training"}}
{{$message = ", this is a training ticket."}}
{{$emoji = "6️⃣"}}

{{end}}
{{$capture = exec "ticket open" $type}}
{{$ID = reFind `\d+` (reFind `<#\d+>` $capture)}}
{{$ticket := sendMessageNoEscapeRetID $ID (joinStr "" .User.Mention $message)}}
{{deleteAllMessageReactions nil $panel}}
{{addMessageReactions nil $panel $emoji}}

{{addMessageReactions $ID $ticket ":kibble:660605488609755146" ":cryopod:631218688099614724" ":Tek_Door:660605465721569334" }}
{{if eq .Reaction.Emoji.Name "1️⃣"}}
{{$second := sendMessageNoEscapeRetID $ID (joinStr "" "Mark off your maps here as you do them. When all maps are done, this ticket will close automagically.")}}
{{addMessageReactions $ID $second ":Manticore:640697193052766218" ":Artemis:640697096814592040" ":Gryphon:640697148807184384" ":Zelda:640697492274544643" ":Phoenix:640698407199178771" ":Elysian:640698387506921523" ":Titan:640697323143561276" ":Raiden:655954199476961300" ":Medusa:640697225684713473" ":Beowulf:682592075685953570" ":Osiris:748895819331141674"}}
{{end}}


