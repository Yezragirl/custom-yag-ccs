{{/*Timed. opens new Weekly Timer Reset ticket every monday morning at 8 am*/}}
{{$capture := ""}}
{{$ID := ""}}
{{$message := ""}}
{{$type := ""}}

{{$type = "weekly-timer-reset"}}
{{$message = "<@&634598489732546588>, it's timer reset time! Please react with your map so that I can close the ticket whenever every map is done. Deadline is Sunday."}}

{{$capture = exec "ticket open" $type}}
{{$ID = reFind `\d+` (reFind `<#\d+>` $capture)}}
{{$ticket := sendMessageNoEscapeRetID $ID $message}}

{{addMessageReactions $ID $ticket ":kibble:660605488609755146" ":cryopod:631218688099614724" ":Tek_Door:660605465721569334" }}
{{$second := sendMessageNoEscapeRetID $ID (joinStr "" "Mark off your maps here as you do them. When all maps are done, this ticket will close automagically.")}}
{{addMessageReactions $ID $second ":Manticore:640697193052766218" ":Artemis:640697096814592040" ":Gryphon:640697148807184384" ":Zelda:640697492274544643" ":Phoenix:640698407199178771" ":Elysian:640698387506921523" ":Titan:640697323143561276" ":Raiden:655954199476961300" ":Medusa:640697225684713473" ":Beowulf:682592075685953570" ":Osiris:748895819331141674"}}
