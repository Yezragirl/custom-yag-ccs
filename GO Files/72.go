{{if (eq .Channel.ParentID 635305499600093214)}}
{{$msgID := .ReactionMessage.ID}}
{{$Q := (dbGet .User.ID "Q").Value}} 
{{$admin := (dbGet .User.ID "admin").Value}}
{{$age := (dbGet .User.ID "age").Value}}
{{$gt := (dbGet .User.ID "gt").Value}}
{{$name := (dbGet .User.ID "name").Value}}
{{$pin := (dbGet .User.ID "pin").Value}}
{{$public := (dbGet .User.ID "public").Value}}
{{$ref := (dbGet .User.ID "ref").Value}}
{{$parentguardian := (dbGet .User.ID "parentguardian").Value}}
{{$tribe := (dbGet .User.ID "tribe").Value}}
{{if and (eq .Reaction.Emoji.ID 658376626639339533) (eq (toString $Q) "Ref")}}
{{deleteAllMessageReactions nil $msgID}}
{{$Q = "Ref2"}}{{dbSet .User.ID "Q" $Q}}
{{sendMessage nil (joinStr "" .User.Mention ", who referred you?")}}
{{else if and (eq .Reaction.Emoji.ID 658376639322783745) (eq (toString $Q) "Ref")}}
{{deleteAllMessageReactions nil $msgID}}
{{$Q = "Pin"}}{{dbSet .User.ID "Q" $Q}}{{$ref = "None"}}{{dbSet .User.ID "ref" "None"}}
{{sendMessage nil (joinStr "" .User.Mention ", what 4 digit pin would you like used whenever a pin may be required? (Quest Vaults, or other situations.)")}}
{{else if and (eq .Reaction.Emoji.ID 640697096814592040) (eq (toString $Q) "Starter")}}
{{deleteAllMessageReactions nil $msgID}}
{{$Q = "Done"}}{{dbSet .User.ID "Q" $Q}}
{{execCC 71 nil 0 ""}}
{{sendMessage nil "Selected Artemis"}}
{{else if and (eq .Reaction.Emoji.ID 640697193052766218) (eq (toString $Q) "Starter")}}
{{deleteAllMessageReactions nil $msgID}}
{{$Q = "Done"}}{{dbSet .User.ID "Q" $Q}}
{{execCC 71 nil 0 ""}}
{{sendMessage nil "Selected Manticore"}}
{{else if and (eq .Reaction.Emoji.ID 640697148807184384) (eq (toString $Q) "Starter")}}
{{deleteAllMessageReactions nil $msgID}}
{{$Q = "Done"}}{{dbSet .User.ID "Q" $Q}}
{{execCC 71 nil 0 ""}}
{{sendMessage nil "Selected Gryphon"}}
{{else if and (eq .Reaction.Emoji.ID 640698387506921523) (eq (toString $Q) "Starter")}}
{{deleteAllMessageReactions nil $msgID}}
{{$Q = "Done"}}{{dbSet .User.ID "Q" $Q}}
{{execCC 71 nil 0 ""}}
{{sendMessage nil "Selected Elysian"}}
{{else if and (eq .Reaction.Emoji.ID 640697492274544643) (eq (toString $Q) "Starter")}}
{{deleteAllMessageReactions nil $msgID}}
{{$Q = "Done"}}{{dbSet .User.ID "Q" $Q}}
{{execCC 71 nil 0 ""}}
{{sendMessage nil "Selected Zelda"}}
{{else if and (eq .Reaction.Emoji.ID 640698407199178771) (eq (toString $Q) "Starter")}}
{{deleteAllMessageReactions nil $msgID}}
{{$Q = "Done"}}{{dbSet .User.ID "Q" $Q}}
{{execCC 71 nil 0 ""}}
{{sendMessage nil "Selected Phoenix"}}
{{else if and (eq .Reaction.Emoji.ID 640697225684713473) (eq (toString $Q) "Starter")}}
{{deleteAllMessageReactions nil $msgID}}
{{$Q = "Done"}}{{dbSet .User.ID "Q" $Q}}
{{execCC 71 nil 0 ""}}
{{sendMessage nil "Selected Medusa"}}
{{else if and (eq .Reaction.Emoji.ID 640697323143561276) (eq (toString $Q) "Starter")}}
{{deleteAllMessageReactions nil $msgID}}
{{$Q = "Done"}}{{dbSet .User.ID "Q" $Q}}
{{execCC 71 nil 0 ""}}
{{sendMessage nil "Selected Titan"}}
{{else if and (eq .Reaction.Emoji.ID 655954199476961300) (eq (toString $Q) "Starter")}}
{{deleteAllMessageReactions nil $msgID}}
{{$Q = "Done"}}{{dbSet .User.ID "Q" $Q}}
{{execCC 71 nil 0 ""}}
{{sendMessage nil "Selected Raiden"}}
{{else if and (eq .Reaction.Emoji.ID 682592075685953570) (eq (toString $Q) "Starter")}}
{{deleteAllMessageReactions nil $msgID}}
{{$Q = "Done"}}{{dbSet .User.ID "Q" $Q}}
{{execCC 71 nil 0 ""}}
{{sendMessage nil "Selected Beowulf"}}
{{else if and (eq .Reaction.Emoji.ID 655954171844886558) (eq (toString $Q) "Starter")}}
{{deleteAllMessageReactions nil $msgID}}
{{$Q = "Done"}}{{dbSet .User.ID "Q" $Q}}
{{execCC 71 nil 0 ""}}
{{sendMessage nil "Selected Silvestria"}}
{{else if and (eq .Reaction.Emoji.Name "⛔") (eq (toString $Q) "Starter")}}
{{deleteAllMessageReactions nil $msgID}}
{{$Q = "Done"}}{{dbSet .User.ID "Q" $Q}}
{{execCC 71 nil 0 (sdict "starter" "none")}}
{{sendMessage nil "You are off and running! The <#613585635026141185> channel should now be visible to you. This channel will close in 60 seconds."}} {{sleep 60}}
{{exec "tickets close" "Registration Complete-No Starter Needed"}}
{{else if and (eq .Reaction.Emoji.ID 658376626639339533) (eq (toString $Q) "ParentGuardian")}}
{{deleteAllMessageReactions nil $msgID}}
{{$Q = "ParentGuardian2"}}{{dbSet .User.ID "Q" $Q}}
{{sendMessage nil (joinStr "" .User.Mention ", who is your parent/guardian?")}}
{{else if and (eq .Reaction.Emoji.ID 658376639322783745) (eq (toString $Q) "ParentGuardian")}}
{{deleteAllMessageReactions nil $msgID}}
{{sendMessage nil (joinStr "" .User.Mention ": I'm sorry, but this cluster is for 18+ only. Please come back when you are over 18.")}}
{{sendMessage 573897276569813012 (joinStr "" .User.Mention " is being kicked for being" $age ". Please take note in case they attempt to register again and lie.\n<@&573210843840249889>")}}
{{sleep 10}}
{{execAdmin "kick" .User.ID "Under age."}}
{{execAdmin "tickets close underage"}}


{{else if (eq .Reaction.Emoji.Name "gcrystal")}}
{{deleteAllMessageReactions nil $msgID}}
{{exec "tickets close" "Picked up Items"}}

{{end}}
{{end}}
