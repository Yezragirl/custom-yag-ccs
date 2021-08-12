{{/*command collected. Allows non-admins to close tickets.*/}}
{{$args := parseArgs 0 "syntax `-collected Map Locker Pin`"
(carg "string" "Map")
(carg "string" "Locker ID")
(carg "string" "pin")}}
{{$map := (toString ($args.Get 0))}}
{{$locker := ($args.Get 1)}}
{{$pin := ($args.Get 2)}}
{{$mapname := ""}}
{{if eq $map "<:Artemis:640697096814592040>"}}
{{$mapname = "Artemis"}}
{{else if eq $map "<:Manticore:640697193052766218>"}}
{{$mapname = "Manticore"}}
{{else if eq $map "<:Gryphon:640697148807184384>"}}
{{$mapname = "Gryphon"}}
{{else if eq $map "<:Elysian:640698387506921523>"}}
{{$mapname = "Elysian"}}
{{else if eq $map "<:Zelda:640697492274544643>"}}
{{$mapname = "Zelda"}}
{{else if eq $map "<:Osiris:748895819331141674>"}}
{{$mapname = "Osiris"}}
{{else if eq $map "<:Beowulf:682592075685953570>"}}
{{$mapname = "Beowulf"}}
{{else if eq $map "<:Phoenix:640698407199178771>"}}
{{$mapname = "Phoenix"}}
{{else if eq $map "<:Titan:640697323143561276>"}}
{{$mapname = "Titan"}}
{{else if eq $map "<:Raiden:655954199476961300>"}}
{{$mapname = "Raiden"}}
{{else if eq $map "<:Medusa:640697225684713473>"}}
{{$mapname = "Medusa"}}
{{end}}

{{if gt (len .Args) 1}}
{{$msg := sendMessageRetID nil (joinStr "" "You can collect your item(s) from Room " $locker " on " $mapname ", using pin " $pin ".\n\nPlease hit the <:gcrystal:662677298431918092> below to close this ticket once you have collected your item(s).")}}
 {{addMessageReactions nil $msg ":gcrystal:662677298431918092" }}
{{deleteTrigger 1}}
{{else}}

 {{$msg := sendMessageRetID nil "Please hit the <:gcrystal:662677298431918092> below to close this ticket."}}
 {{addMessageReactions nil $msg ":gcrystal:662677298431918092" }}
{{deleteTrigger 1}}
{{end}}