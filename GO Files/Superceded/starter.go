{{$args := parseArgs 3 "syntax `-starter Map Locker Pin`"
(carg "string" "Map")
(carg "string" "Locker ID")
(carg "string" "pin")}}
{{$map := ($args.Get 0)}}
{{$locker := ($args.Get 1)}}
{{$pin := ($args.Get 2)}}
{{$map}}
{{$locker}}
{{$pin}}

{{if eq $map "<:Artemis:640697096814592040 >"}}
{{sendMessage nil (joinStr "" "The Community Center can be reached by spawning at Beach Zone 1, and using the teleporter there (between the two rock arches) to teleporter to Community Center. You can pick up your starter in the starter pick up area, Fridge AND Box " $locker ", Pin " $pin ". Please make sure to check out the <#634509655456088087> and <#575676562872074290> pages! Welcome to the cluster!")}}
{{else if eq $map "<:Manticore:640697193052766218 >"}}
{{sendMessage nil (joinStr "" "The Community Center can be reached by spawning at Jungle 2, and take the teleporter to Green Obelisk/Community Center . The bank is in the field behind the obi. Blue Floor, room " $locker ", Pin " $pin ". Please make sure to check out the <#634509655456088087> and <#575676562872074290> pages! Welcome to the cluster!")}}
{{else if eq $map "<:Gryphon:640697148807184384 >"}}
{{sendMessage nil (joinStr "" "The Community Center can be reached by spawning at Portal, and make your way down the metal walk way until you come to the Teleporter in the tunnel. From there, teleport to Community Center. It should be right above you. You can pick up your starter there, room " $locker ", Pin " $pin ". Please make sure to check out the <#634509655456088087> and <#575676562872074290> pages! Welcome to the cluster!")}}
{{else if eq $map "<:Elysian:640698387506921523 >"}}
{{sendMessage nil (joinStr "" "The Community Center can be reached by spawning at Tropical Island South, and take the Teleporter at Blue Obelisk/ Community Center. You can pick up your starter there, Locker " $locker ", Pin " $pin ". Please make sure to check out the <#634509655456088087> and <#575676562872074290> pages! Welcome to the cluster!")}}
{{else if eq $map "<:Zelda:640697492274544643 >"}}
{{sendMessage nil (joinStr "" "The Community Center can be reached by spawning at South Zone 3, and take the teleporter on the beach to the Community Center. The starter pick up is in the bottom floor. Room " $locker ", Pin " $pin ". Please make sure to check out the <#634509655456088087> and <#575676562872074290> pages! Welcome to the cluster!")}}
{{else if eq $map "<:Phoenix:640698407199178771 >"}}
{{sendMessage nil (joinStr "" "The Community Center can be reached by spawning at South 3, and take the teleporter to the Community Center. Follow the signs at the Community Center for starters. Room " $locker ", Pin " $pin ". Please make sure to check out the <#634509655456088087> and <#575676562872074290> pages! Welcome to the cluster!")}}
{{else if eq $map "<:Medusa:640697225684713473 >"}}
{{sendMessage nil (joinStr "" "Spawn at Sanctuary West and follow the Billboard signs. The Community Center is located near the Green Obelisk in sanctuary. You can pick up your Starter in locker " $locker ", Pin " $pin ". Please make sure to check out the <#634509655456088087> and <#575676562872074290> pages! Welcome to the cluster!")}}
{{else if eq $map "<:Titan:640697323143561276 >"}}
{{sendMessage nil (joinStr "" "Spawn at Small Islands, head towards the shore and take the Teleporter there to Community Center. You can pick up your Starter in the Starter pick up building, Locker " $locker ", Pin " $pin ". Please make sure to check out the <#634509655456088087> and <#575676562872074290> pages! Welcome to the cluster")}}
{{else if eq $map "<:Raiden:655954199476961300 >"}}
{{sendMessage nil (joinStr "" "The Community Center can be reached by spawning at Jungle 2, and take the teleporter to Community Center . The starter pick up is in the top floor. Room " $locker ", Pin " $pin ". Please make sure to check out the <#634509655456088087> and <#575676562872074290> pages! Welcome to the cluster!")}}
{{else if eq $map "<:Silvestria:655954171844886558 >"}}
{{sendMessage nil (joinStr "" "The Community Center can be reached by spawning at Highland South, and run to the Community Center located at the lighthouse. Follow the signs at the Community Center for starters. Room " $locker ", Pin " $pin ". Please make sure to check out the <#634509655456088087> and <#575676562872074290> pages! Welcome to the cluster!")}}
 {{end}}
{{$msg := sendMessageRetID nil "Please hit the <:gcrystal:662677298431918092> below when you've collected your items."}}
{{addMessageReactions nil $msg ":gcrystal:662677298431918092" }}
{{deleteTrigger 1}}