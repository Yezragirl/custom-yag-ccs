{{if (eq .Channel.ParentID 635305499600093214)}} 
	{{$Q := (dbGet .User.ID "Q").Value}}{{$admin := (dbGet .User.ID "admin").Value}}{{$age := (dbGet .User.ID "age").Value}}{{$gt := (dbGet .User.ID "gt").Value}}{{$name := (dbGet .User.ID "name").Value}}{{$pin := (dbGet .User.ID "pin").Value}}{{$public := (dbGet .User.ID "public").Value}}{{$ref := (dbGet .User.ID "ref").Value}}{{$parentguardian := (dbGet .User.ID "parentguardian").Value}}
	{{if not $parentguardian}}{{$parentguardian = "Not a Minor"}}{{dbSet .User.ID "parentguardian" $parentguardian}}{{end}}{{$tribe := (dbGet .User.ID "tribe").Value}}{{$color := randInt 111111 999999 }}
	{{if not $tribe}}{{$tribe = "Unregistered"}}{{dbSet .User.ID "tribe" $tribe}}{{end}}
	{{if eq (toString $Q) "Done"}}
	{{embed := cembed}}
	{{embed2 := cembed}}
		{{if eq $parentguardian "Not a Minor"}}
			{{$embed = cembed 
			"color" $color
			"fields" (cslice 
			(sdict "name" "In Game Name" "value" (toString $name) ) 
			(sdict "name" "Gamer Tag" "value" (toString $gt) ) 
			(sdict "name" "Tribe Name" "value" (toString $tribe) ) )
			"thumbnail" (sdict "url" (.User.AvatarURL "1024")) }}
			{{$embed2 = cembed 
			"color" $color
			"fields" (cslice 
			(sdict "name" "Discord Name" "value" .User.Username ) 
			(sdict "name" "In Game Name" "value" (toString $name) ) 
			(sdict "name" "Gamer Tag" "value" (toString $gt) ) 
			(sdict "name" "Referral" "value" (toString $ref) ) 
			(sdict "name" "Pin" "value" (toString $pin) ))
			"thumbnail" (sdict "url" (.User.AvatarURL "1024")) }}
		{{else}}
			{{$embed = cembed 
			"color" $color
			"fields" (cslice 
			(sdict "name" "__***MINOR***__" "value" "_ _" ) 
			(sdict "name" "In Game Name" "value" (toString $name) ) 
			(sdict "name" "Tribe Name" "value" (toString $tribe) ) 
			(sdict "name" "Gamer Tag" "value" (toString $gt) ) )
			"thumbnail" (sdict "url" (.User.AvatarURL "1024")) }}
			{{$embed2 = cembed 
			"color" $color
			"fields" (cslice 
			(sdict "name" "Discord Name" "value" .User.Username ) 
			(sdict "name" "__***MINOR***__" "value" (toString (toInt $age)) ) 
			(sdict "name" "In Game Name" "value" (toString $name) ) 
			(sdict "name" "Gamer Tag" "value" (toString $gt) ) 
			(sdict "name" "Tribe Name" "value" (toString $tribe) ) 
			(sdict "name" "Parent/Guardian" "value" (toString $parentguardian) )
			(sdict "name" "Pin" "value" (toString $pin) ))
			"thumbnail" (sdict "url" (.User.AvatarURL "1024")) }}
		{{end}}
		{{$public = (toString (sendMessageRetID 658352884701724733 $embed))}}
		{{dbSet .User.ID "public" $public}}
		{{$admin = (toString (sendMessageRetID 658354135221141533 $embed2))}}
		{{dbSet .User.ID "admin" $admin}}
		{{execCC 213 nil 604800 ""}}
		{{sendDM "Welcome to Yez's Ark Cluster! The password can be found on the #current-server-password channel, which you should now have access to. Please be aware that giving out this password is a ***bannable offense*** as is allowing people to piggyback off your login. To find the server make sure you are on unofficial PC sessions and have the password box checked. Do a search for Yez. Please be sure to read over the rules in the #pve-rules channel.\n\n__**Quests**__\nEvery week we offer *quests*, where you *collect*, **craft**, and/or __tame__ in order to receive rewards like bbs, coupons, element, and max level tames! We also offer ***new player quests*** until you hit level 100, which are designed to teach you how to play this game. If you are interested in participating, head over to #quest-rules, where you can read about how they work, and request a quest vault. \n\n__**Roles**__\nTo keep from pinging everyone, all the time, we've created various *Roles* for each of the things we usually ping people over. You can self-assign these roles by going to the #roles channel.\n\n__**Discord**__\nOur currency system is entirely Discord based, so having Discord is essential. **Leaving the Discord is seen as leaving the server.** If you leave the discord without contacting an admin, we will assume you have left, clearing your bases and confiscating your dinos for auction.\n\nWe hope you enjoy your time here at Yez's Ark Cluster. If you need any assistance, you can head over to #make-a-ticket."}}
		{{if eq $parentguardian "Not a Minor"}}
			{{giveRoleID .User.ID 573210479753691170}}{{giveRoleID .User.ID 607041870966685706}}{{removeRoleID 606566506922377257}}
			{{editNickname (toString $name)}}
			{{sendMessage 573245421066125324 (joinStr "" "Let's all welcome " .User.Mention " to the Cluster!")}}
			{{if not (eq $ref "None")}}
				{{$message := sendMessageRetID 573860427721605120 (joinStr "" "Pay referral fee to " $ref " for " $name)}}{{addMessageReactions 573860427721605120 $message "✔️" "❌"}}
			{{end}}
		{{else}}
			{{giveRoleID .User.ID 645404993863811072}}{{removeRoleID 606566506922377257}}
			{{editNickname (joinStr "" $name "(M)")}}
			{{sendMessage 646033978721173523 (joinStr "" "Let's all welcome " .User.Mention " to the Cluster!")}}
			{{$message := sendMessageRetID 573860427721605120 (joinStr "" "Add Parent Role to  " $parentguardian " for " $name)}}{{addMessageReactions 573860427721605120 $message "✔️" "❌"}}
		{{end}}
	{{end}}
{{end}}

