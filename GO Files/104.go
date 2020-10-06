{{$args := parseArgs 4 "syntax `-starter Map Locker Pin @mention`"
(carg "string" "Map")
(carg "string" "Locker ID")
(carg "string" "pin")
(carg "user" "user")}}
{{$map := ($args.Get 0)}}
{{$locker := ($args.Get 1)}}
{{$pin := ($args.Get 2)}}
{{$who := ($args.Get 3)}}

{{sendMessage nil "The Cluster password can be found on <#613585635026141185>.\n"}}
{{if eq $map "<:Artemis:640697096814592040>"}}
{{sendMessage nil (joinStr "" "The Community Center can be reached by spawning at Beach Zone 1, and using the teleporter there (between the two rock arches) to teleporter to Community Center. You can pick up your starter in the starter pick up area on the 3rd floor,  room " $locker ", Pin " $pin ".")}}
{{else if eq $map "<:Manticore:640697193052766218>"}}
{{sendMessage nil (joinStr "" "The Community Center can be reached by spawning at Jungle 2, and take the teleporter to Green Obelisk/Community Center . The bank is in the field behind the obi. Blue Floor, room " $locker ", Pin " $pin ".")}}
{{else if eq $map "<:Gryphon:640697148807184384>"}}
{{sendMessage nil (joinStr "" "The Community Center can be reached by spawning at Portal, and make your way down the metal walk way until you come to the Teleporter in the tunnel. From there, teleport to Community Center. It should be right above you. You can pick up your starter there, room " $locker ", Pin " $pin ".")}}
{{else if eq $map "<:Elysian:640698387506921523>"}}
{{sendMessage nil (joinStr "" "The Community Center can be reached by spawning at Tropical Island South, and take the Teleporter at Blue Obelisk to Community Center. You can pick up your starter there, Locker " $locker ", Pin " $pin ".")}}
{{else if eq $map "<:Zelda:640697492274544643>"}}
{{sendMessage nil (joinStr "" "The Community Center can be reached by spawning at South Zone 3, and take the teleporter on the beach to the Community Center. The starter pick up is in the bottom floor. Room " $locker ", Pin " $pin ".")}}
{{else if eq $map "<:Osiris:748895819331141674>"}}
{{sendMessage nil (joinStr "" "The Community Center can be reached by spawning Central, and making your way to 48, 53. The Community Center sits on a plateau along the river where the red forest meets the plains. Room " $locker ", Pin " $pin ".")}}
{{else if eq $map "<:Beowulf:682592075685953570>"}}
{{sendMessage nil (joinStr "" "The Community Center can be reached by spawning at Bog East, and taking the teleporter there to the Community Center. The starter pick up is in the second floor of the Starter & Arkmart Pick up building. Room " $locker ", Pin " $pin ".")}}
{{else if eq $map "<:Phoenix:640698407199178771>"}}
{{sendMessage nil (joinStr "" "The Community Center can be reached by spawning at South 3, and take the teleporter to the Community Center. Follow the signs at the Community Center for starters. Room " $locker ", Pin " $pin ".")}}
{{else if eq $map "<:Medusa:640697225684713473>"}}
{{sendMessage nil (joinStr "" "Sanctuary East, turn to the North and head towards the transmitter beam. The Community Center is located atop a waterfall against the back of a large city structure, there are walkways that resemble incomplete mobius strips in front, over the water. You can pick up your Starter in locker " $locker ", Pin " $pin ".")}}
{{else if eq $map "<:Titan:640697323143561276>"}}
{{sendMessage nil (joinStr "" "Spawn at Small Islands. The Community Center is located on the island. Look for transmitter beams. You can pick up your Starter in the Starter pick up building, Locker " $locker ", Pin " $pin ".")}}
{{else if eq $map "<:Raiden:655954199476961300>"}}
{{sendMessage nil (joinStr "" "The Community Center can be reached by spawning at Jungle 2, and take the teleporter to Community Center . The starter pick up is in the first floor. Room " $locker ", Pin " $pin ".")}}
{{end}}
{{sendMessage nil (joinStr "" "Welcome to the cluster!\n\n" $who.Mention ", you'll find your starter info above.\n\nThis channel will be automatically closed after 7 days. If you can not pick up your starter within those 7 days, you can open a general ticket when you are ready and a new starter will be provided.\n")}}
{{$msg := sendMessageRetID nil "Please hit the <:gcrystal:662677298431918092> below when you've collected your items so that we can restock the room for the next new player."}}
{{addMessageReactions nil $msg ":gcrystal:662677298431918092" }}
{{deleteTrigger 1}}