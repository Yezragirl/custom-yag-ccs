{{$capture := ""}}
{{$ID := ""}}
{{$message := ""}}
{{$type := ""}}
{{$panel := (toString .ReactionMessage.ID)}}
{{$emoji := ""}}

{{if eq .Reaction.Emoji.Name "1️⃣"}}
{{$type = "weekly-timer-reset"}}
{{$message = ", it's timer reset time! Please react with your map so that I can close the ticket whenever every map is done."}}
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