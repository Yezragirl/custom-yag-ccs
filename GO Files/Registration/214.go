{{if (eq .Channel.ParentID 635305499600093214)}} 
	{{$Q := (dbGet .User.ID "Q").Value}}{{$admin := (dbGet .User.ID "admin").Value}}{{$age := (dbGet .User.ID "age").Value}}{{$gt := (dbGet .User.ID "gt").Value}}{{$name := (dbGet .User.ID "name").Value}}{{$pin := (dbGet .User.ID "pin").Value}}{{$public := (dbGet .User.ID "public").Value}}{{$ref := (dbGet .User.ID "ref").Value}}{{$parentguardian := (dbGet .User.ID "parentguardian").Value}}
	{{if not $parentguardian}}{{$parentguardian = "Not a Minor"}}{{dbSet .User.ID "parentguardian" $parentguardian}}{{end}}{{$tribe := (dbGet .User.ID "tribe").Value}}
	{{if not $tribe}}{{$tribe = "Unregistered"}}{{dbSet .User.ID "tribe" $tribe}}{{end}}
	{{if eq (toString $Q) "Done"}}
		{{$color := randInt 111111 999999 }}
		{{$embed := cembed 
		"color" $color
		"fields" (cslice 
		(sdict "name" "In Game Name" "value" (toString $name) ) 
		(sdict "name" "Gamer Tag" "value" (toString $gt) ) 
		(sdict "name" "Tribe Name" "value" (toString $tribe) ) )
		"thumbnail" (sdict "url" (.User.AvatarURL "1024")) }}
		{{$public = (toString (sendMessageRetID 658352884701724733 $embed))}}
		{{dbSet .User.ID "public" $public}}
		{{editNickname (toString $name)}}
		{{$embed2 := cembed 
		"color" $color
		"fields" (cslice 
		(sdict "name" "Discord Name" "value" .User.Username ) 
		(sdict "name" "In Game Name" "value" (toString $name) ) 
		(sdict "name" "Gamer Tag" "value" (toString $gt) ) 
		(sdict "name" "Referral" "value" (toString $ref) ) 
		(sdict "name" "Pin" "value" (toString $pin) ))
		"thumbnail" (sdict "url" (.User.AvatarURL "1024")) }}
		{{$admin = (toString (sendMessageRetID 658354135221141533 $embed2))}}
		{{dbSet .User.ID "admin" $admin}}
		{{giveRoleID .User.ID 573210479753691170}}{{giveRoleID .User.ID 607041870966685706}}{{removeRoleID 606566506922377257}}
		{{execCC 213 nil 604800 ""}}
		{{sendMessage 573245421066125324 (joinStr "" "Let's all welcome " .User.Mention " to the Cluster!")}}
		{{sendDM "Welcome to Yez's Ark Cluster! The password can be found on the #current-server-password channel, which you should now have access to. Please be aware that giving out this password is a ***bannable offense*** as is allowing people to piggyback off your login. To find the server make sure you are on unofficial PC sessions and have the password box checked. Do a search for Yez. Please be sure to read over the rules in the #pve-rules channel.\n\n__**Quests**__\nEvery week we offer *quests*, where you *collect*, **craft**, and __tame__ in order to receive rewards like bbs, coupons, element, and max level tames! We also offer ***new player quests*** until you hit level 100, which are designed to teach you how to play this game. If you are interested in participating, head over to #quest-rules, where you can read about how they work, and request a quest vault. \n\n__**Roles**__\nTo keep from pinging everyone, all the time, we've created various *Roles* for each of the things we usually ping people over. You can self-assign these roles by going to the #roles channel.\n\n__**Discord**__\nOur currency system is entirely Discord based, so having Discord is essential. **Leaving the Discord is seen as leaving the server.** If you leave the discord without contacting an admin, we will assume you have left, clearing your bases and confiscating your dinos for auction.\n\nWe hope you enjoy your time here at Yez's Ark Cluster. If you need any assistance, you can head over to #make-a-ticket."}}
		{{if .ExecData}}
			{{if eq .ExecData.starter "none"}}

			{{else}}
				{{sendMessageNoEscape nil (joinStr "" .User.Mention ": Stay turned to this channel! A member of the <@&634598489732546588> will be along as soon as possible to give you the information on how to get your starter.")}}
			{{end}}
		{{end}}
		{{if not (eq $ref "None")}}
			{{$message := sendMessageRetID 573860427721605120 (joinStr "" "Pay referral fee to " $ref " for " $name)}}{{addMessageReactions 573860427721605120 $message "✔️" "❌"}}
		{{end}}
	{{else if eq (toString $Q) "Pin2"}}
		{{$pin = (toString .Message.Content)}}
		{{dbSet .User.ID "pin" $pin}}
		{{$Q = "Done"}}{{dbSet .User.ID "Q" $Q}}
		{{$color := randInt 111111 999999 }}
		{{$embed := cembed 
		"color" $color
		"fields" (cslice 
		(sdict "name" "__***MINOR***__" "value" "_ _" ) 
		(sdict "name" "In Game Name" "value" (toString $name) ) 
		(sdict "name" "Tribe Name" "value" (toString $tribe) ) 
		(sdict "name" "Gamer Tag" "value" (toString $gt) ) )
		"thumbnail" (sdict "url" (.User.AvatarURL "1024")) }}
		{{$public = (toString (sendMessageRetID 658352884701724733 $embed))}}
		{{dbSet .User.ID "public" $public}}
		{{editNickname (joinStr "-" $name $parentguardian)}}
		{{$message := sendMessageRetID 573860427721605120 (joinStr "" "Add Parent Role to  " $parentguardian " for " $name)}}{{addMessageReactions 573860427721605120 $message "✔️" "❌"}}
		{{$embed2 := cembed 
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
		{{$admin = (toString (sendMessageRetID 658354135221141533 $embed2))}}
		{{dbSet .User.ID "admin" $admin}}
		{{giveRoleID .User.ID 645404993863811072}}
		{{removeRoleID 606566506922377257}}
		{{sendMessage nil (joinStr "" .User.Mention ", you are all set. Your parent/guardian should set you up with any gear you need. Please be sure to read through the rules for kids. Be aware that your parent/guardian is responsible for your actions. This channel will close in 60 seconds.") }}
		{{sleep 60}}{{exec "tickets close" "Registration Complete"}}
	{{end}}
{{end}}

