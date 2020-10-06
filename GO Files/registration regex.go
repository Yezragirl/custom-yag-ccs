{{if (eq .Channel.ParentID 635305499600093214)}} 
 {{$db := dbGet .User.ID "registration"}}
 {{$reg := sdict}}{{range $k,$v := $db.Value}}{{$reg.Set $k $v}}{{end}}
 {{if not $db}}
 {{$reg.Set "Q" "A"}}
 {{dbSet .User.ID "registration" $reg}}
 {{sendMessage nil (joinStr "" .User.Mention ", how old are you?")}}
 {{else}} 
 {{if eq $reg.Q "A"}}
 {{if reFind `\D` .Message.Content}}
 {{sendMessage nil (joinStr "" "Sorry " .User.Mention ", thats not a number. I'll ask again, how old are you?")}}
 {{else if gt (toInt .Message.Content) 85}}
 {{sendMessage nil (joinStr "" "Sorry " .User.Mention ", that's an unrealistic number. I'll ask again, how old are you?")}}
 {{else if lt (toInt .Message.Content) 5}}
 {{sendMessage nil (joinStr "" "Sorry " .User.Mention ", that's an unrealistic number. I'll ask again, how old are you?")}}
 {{else}}
 {{$reg.Set "age" (toInt .Message.Content)}}
 {{if ge $reg.age 18}}
 {{$reg.Set "Q" "G"}}
 {{dbSet .User.ID "registration" $reg}}
 {{sendMessage nil (joinStr "" .User.Mention ", what is your Gamer Tag?")}}
 {{else}}
 {{$reg.Set "Q" "PG"}}
 {{dbSet .User.ID "registration" $reg}}
 {{$msg := sendMessageRetID nil (joinStr "" .User.Mention ", is your Parent or Guardian a member of this cluster?")}}
 {{addMessageReactions nil $msg ":yes:658376626639339533" ":no:658376639322783745"}}
 {{end}} 
 {{end}}
 {{else if eq $reg.Q "G"}}
 {{$reg.Set "gt" .Message.Content}}
 {{$reg.Set "Q" "N"}}
 {{dbSet .User.ID "registration" $reg}}
 {{sendMessage nil (joinStr "" .User.Mention ", what is your In Game Name?")}}
 {{else if eq $reg.Q "N"}}
 {{$reg.Set "name" .Message.Content}}
 {{$reg.Set "Q" "R"}}
 {{dbSet .User.ID "registration" $reg}}
 {{$msg := sendMessageRetID nil (joinStr "" .User.Mention ", were you referred by one of our members?")}}
 {{addMessageReactions nil $msg ":yes:658376626639339533" ":no:658376639322783745"}}
 {{else if eq $reg.Q "R2"}}
 {{$reg.Set "ref" .Message.Content}}
 {{$reg.Set "Q" "P"}}
 {{dbSet .User.ID "registration" $reg}}
 {{sendMessage nil (joinStr "" .User.Mention ", what 4 digit pin would you like used whenever a pin may be required? (Quest Vaults, or other situations.)")}}
 {{else if eq $reg.Q "P"}}
 {{$reg.Set "pin" .Message.Content}}
 {{$reg.Set "Q" "S"}}
 {{dbSet .User.ID "registration" $reg}}
 {{$msg := sendMessageRetID nil (joinStr "" .User.Mention ", where would you like to pick up your starter?\n<:Artemis:640697096814592040> Artemis - Valguero (Main Map)\n<:Manticore:640697193052766218> Manticore - Ragnarok\n<:Gryphon:640697148807184384> Gryphon - Aberration\n<:Elysian:640698387506921523> Elysian - Center\n<:Zelda:640697492274544643> Zelda - Island\n<:Phoenix:640698407199178771> Phoenix - Scorched Earth\n<:Medusa:640697225684713473> Medusa - Extinction\n<:Titan:640697323143561276> Titan - Valguero\n<:Raiden:655954199476961300> Raiden - Ragnarok\n<:Beowulf:682592075685953570> Beowulf - Genesis\n⛔ No Starter Needed")}}
 {{addMessageReactions nil $msg ":Artemis:640697096814592040" ":Manticore:640697193052766218" ":Gryphon:640697148807184384" ":Elysian:640698387506921523" ":Zelda:640697492274544643" ":Phoenix:640698407199178771" ":Medusa:640697225684713473" ":Titan:640697323143561276" ":Raiden:655954199476961300" "Beowulf:682592075685953570" "⛔" }}
 {{else if eq $reg.Q "G2"}}
 {{$reg.Set "guardian" .Message.Content}}
 {{$reg.Set "Q" "N2"}}
 {{dbSet .User.ID "registration" $reg}}
 {{sendMessage nil (joinStr "" .User.Mention ", what is your Gamer Tag?")}}
 {{else if eq $reg.Q "N2"}}
 {{$reg.Set "gt" .Message.Content}}
 {{$reg.Set "Q" "T2"}}
 {{dbSet .User.ID "registration" $reg}}
 {{sendMessage nil (joinStr "" .User.Mention ", what is your In Game Name?")}}
 {{else if eq $reg.Q "T2"}}
 {{$reg.Set "name" .Message.Content}}
 {{$reg.Set "Q" "D2"}}
 {{dbSet .User.ID "registration" $reg}}
 {{sendMessage nil (joinStr "" .User.Mention ", what 4 digit pin would you like used whenever a pin may be required? (Quest Vaults, or other situations.)")}}
 {{else if eq $reg.Q "D"}}
 {{$reg.Set "Q" "E"}}
 {{$color := randInt 111111 999999 }}
 {{$embed := cembed 
 "color" $color
 "fields" (cslice 
 (sdict "name" "In Game Name" "value" $reg.name "inline" false) 
 (sdict "name" "Gamer Tag" "value" $reg.gt "inline" false) 
 (sdict "name" "Tribe Name" "value" "Unregistered" "inline" false) )
 "thumbnail" (sdict "url" (.User.AvatarURL "1024")) }}
 {{$reg.Set "public" (toString (sendMessageRetID 658352884701724733 $embed))}}
 {{editNickname $reg.name}}
 {{$embed2 := cembed 
 "color" $color
 "fields" (cslice 
 (sdict "name" "Discord Name" "value" .User.Username "inline" false) 
 (sdict "name" "In Game Name" "value" $reg.name "inline" false) 
 (sdict "name" "Gamer Tag" "value" $reg.gt "inline" false) 
 (sdict "name" "Referral" "value" $reg.ref "inline" false) 
 (sdict "name" "Pin" "value" $reg.pin "inline" false))
 "thumbnail" (sdict "url" (.User.AvatarURL "1024")) }}
 {{$reg.Set "admin" (toString (sendMessageRetID 658354135221141533 $embed2))}}
 {{giveRoleID .User.ID 573210479753691170}}{{giveRoleID .User.ID 607041870966685706}}{{removeRoleID 606566506922377257}}

{{sendMessage 573245421066125324 (joinStr "" "Let's all welcome " .User.Mention " to the Cluster!")}}
 {{sendDM "Welcome to Yez's Ark Cluster! The password can be found on the #current-server-password channel, which you should now have access to. Please be aware that giving out this password is a ***bannable offense*** as is allowing people to piggyback off your login. To find the server make sure you are on unofficial PC sessions and have the password box checked. Do a search for Yez. Please be sure to read over the rules in the #pve-rules channel.\n\n__**Quests**__\nEvery week we offer *quests*, where you *collect*, **craft**, and __tame__ in order to receive rewards like bbs, coupons, element, and max level tames! We also offer ***new player quests*** until you hit level 100, which are designed to teach you how to play this game. If you are interested in participating, head over to #quest-rules, where you can read about how they work, and request a quest vault. \n\n__**Roles**__\nTo keep from pinging everyone, all the time, we've created various *Roles* for each of the things we usually ping people over. You can self-assign these roles by going to the #roles channel.\n\n__**Discord**__\nOur currency system is entirely Discord based, so having Discord is essential. **Leaving the Discord is seen as leaving the server.** If you leave the discord without contacting an admin, we will assume you have left, clearing your bases and confiscating your dinos for auction.\n\nWe hope you enjoy your time here at Yez's Ark Cluster. If you need any assistance, you can head over to #make-a-ticket."}}
 {{if .ExecData}}{{else}}{{sendMessageNoEscape nil (joinStr "" .User.Mention ": Stay turned to this channel! A member of the <@&634598489732546588> will be along as soon as possible to give you the information on how to get your starter.")}}{{end}}
 {{dbSet .User.ID "registration" $reg}}
 {{if eq $reg.ref "None"}}
 {{else}}
 {{sendMessage 573860427721605120 (joinStr "" "Pay referral fee to " $reg.ref " for " $reg.name)}}{{end}}
 {{else if eq $reg.Q "D2"}}
 {{$reg.Set "Q" "E"}}
 {{$reg.Set "pin" .Message.Content}}
 {{dbSet .User.ID "registration" $reg}}
 {{$color := randInt 111111 999999 }}
 {{$embed := cembed 
 "color" $color
 "fields" (cslice 
 (sdict "name" "__***MINOR***__" "value" "_ _" "inline" false) 
 (sdict "name" "In Game Name" "value" $reg.name "inline" false) 
 (sdict "name" "Tribe Name" "value" "Unregistered" "inline" false) 
 (sdict "name" "Gamer Tag" "value" $reg.gt "inline" false) )
 "thumbnail" (sdict "url" (.User.AvatarURL "1024")) }}
 {{$reg.Set "public" (toString (sendMessageRetID 658352884701724733 $embed))}}
 {{editNickname (joinStr "-" $reg.name $reg.guardian)}}
 {{sendMessage 573860427721605120 (joinStr "" "Add Parent Role to  " $reg.guardian " for " $reg.name)}}
 {{$embed2 := cembed 
 "color" $color
 "fields" (cslice 
 (sdict "name" "Discord Name" "value" .User.Username "inline" false) 
 (sdict "name" "__***MINOR***__" "value" (toString $reg.age) "inline" false) 
 (sdict "name" "In Game Name" "value" $reg.name "inline" false) 
 (sdict "name" "Gamer Tag" "value" $reg.gt "inline" false) 
 (sdict "name" "Tribe Name" "value" "Unregistered" "inline" false) 
 (sdict "name" "Parent/Guardian" "value" $reg.guardian "inline" false)
 (sdict "name" "Pin" "value" $reg.pin "inline" false))
 "thumbnail" (sdict "url" (.User.AvatarURL "1024")) }}
 {{$reg.Set "admin" (toString (sendMessageRetID 658354135221141533 $embed2))}}
 {{giveRoleID .User.ID 645404993863811072}}
 {{removeRoleID 606566506922377257}}
 {{dbSet .User.ID "registration" $reg}}
 {{sendMessage nil (joinStr "" .User.Mention ", you are all set. Your parent/guardian should set you up with any gear you need. Please be sure to read through the rules for kids. Be aware that your parent/guardian is responsible for your actions. This channel will close in 60 seconds.") }}
 {{sleep 60}}
{{exec "tickets close" "Registration Complete"}}
 {{end}}{{end}}{{end}}
