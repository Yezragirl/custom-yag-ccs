{{else if eq (toString $Q) "Pin"}}{{$pin = (toString .Message.Content)}}{{dbSet .User.ID "pin" $pin}}{{$Q = "Starter"}}{{dbSet .User.ID "Q" $Q}}

{{$map := cslice (dict "desc" "Artemis - Valguero" "emoji" ":Artemis:640697096814592040") (dict "desc" "Manticore - Ragnarok" "emoji" ":Manticore:640697193052766218") (dict "desc" "Gryphon - Aberration" "emoji" ":Gryphon:640697148807184384") (dict "desc" "Elysian - Center" "emoji" ":Elysian:640698387506921523") (dict "desc" "Zelda - Island" "emoji" ":Zelda:640697492274544643") (dict "desc" "Phoenix - Scorched Earth" "emoji" ":Phoenix:640698407199178771") (dict "desc" "Medusa - Extinction" "emoji" ":Medusa:640697225684713473") (dict "desc" "Titan - Valguero" "emoji" ":Titan:640697323143561276") (dict "desc" "Raiden - Ragnarok" "emoji" ":Raiden:655954199476961300") (dict "desc" "Beowulf - Genesis" "emoji" ":Beowulf:682592075685953570") (dict "desc" "Osiris - Crystal Isles" "emoji" ":Osiris:748895819331141674")}}
{{$out := ", where would you like to pick up your starter?\n"}}
{{$addEmotes := cslice}}

{{range (shuffle $map)}}
	{{$out = print $out "<" .emoji "> " .desc "\n" }}
	{{$addEmotes = $addEmotes.Append .emoji}}
{{end}}
{{$out = print $out "⛔ No Starter Needed"}}
{{$addEmotes = $addEmotes.Append "⛔"}}

{{$msg := sendMessageRetID nil (joinStr "" .User.Mention $out)}}

{{- range $addEmotes -}}
	{{- addMessageReactions nil $msg . -}}
{{- end -}}