{{if (hasRoleID 667754372456775680)}}{{/*Has turn role*/}}
{{$db := dbGet 0 "Battle"}} {{$battle := sdict}}{{range $k,$v := $db.Value}}{{$battle.Set $k $v}}{{end}}{{$game := $battle.game}}{{$player1 := (userArg $battle.p1)}}{{$player2 := (userArg $battle.p2)}}{{$P1MAX := toInt ((dbGet $player1.ID "BattleHP").Value)}}{{if lt $P1MAX 25}}{{else}}{{$P1MAX = 25}}{{end}}{{$P2MAX := toInt ((dbGet $player2.ID "BattleHP").Value)}}{{if lt $P2MAX 25}}{{else}}{{$P2MAX = 25}}{{end}}{{$P1HP := $battle.hp1}}{{$P2HP := $battle.hp2}}{{$damage := 0}}{{$shield := (toString $battle.shield)}}{{$riot := (toString $battle.riot)}}{{$rail1 := (toString $battle.rail1)}}{{$rail2 := (toString $battle.rail2)}}{{$club := (toString $battle.club)}}{{$cp := (userArg $battle.cp)}}{{$wp := (userArg $battle.wp)}}{{$heal := 0}}{{$hpchange := 0}}{{$NEWMAX := 0}}{{$action := ""}}{{$win := ""}}{{$lose := ""}}{{$countp2 := (dbGet 0 "p2win").Value}}{{$countp1 := (dbGet 0 "p1win").Value}}
	{{if (eq .Reaction.Emoji.ID 633311599020736531)}}{{/*sword*/}}
		{{$damage = (randInt 5 11)}}
		{{$action = (joinStr "" $cp.Username " swung their sword at " $wp.Username " doing " (toString $damage) " points of damage!")}}
			{{if eq $shield "yes"}}
				{{$damage = roundCeil (mult (fdiv (randInt 50 76) 100) $damage)}}
				{{$action = (joinStr "" $action "\n Reduced to " (toInt $damage) " by " $wp.Username "'s shield!")}}
			{{else if eq $riot "yes"}}
				{{$damage = roundCeil (mult (fdiv (randInt 0 26) 100) $damage)}}
				{{$action = (joinStr "" $action "\n Reduced to " (toInt $damage) " by " $wp.Username "'s riot shield!")}}
			{{else if eq $club "yes"}}
				{{$damage = roundCeil (mult $damage (randInt 0 2))}}
					{{if eq (toInt ($damage)) 0}}
						{{$action = (joinStr "" $cp.Username " was too disoriented from their clubbing to hit " $wp.Username "!")}}
					{{end}}
			{{end}}{{$riot = "no"}}
			{{$club = "no"}}{{$shield = "no"}}
	{{else if (eq .Reaction.Emoji.ID 667754698639540254)}}{{/*shield*/}}
		{{$damage = (randInt 0 6)}}
		{{$action = (joinStr "" $cp.Username " held their shield high, swinging their sword wildly at " $wp.Username " doing " (toInt $damage) " points of damage!")}}
			{{if eq $shield "yes"}}
				{{$damage = roundCeil (mult (fdiv (randInt 50 76) 100) $damage)}}
				{{$action = (joinStr "" $action "\n Reduced to " (toInt $damage) " by " $wp.Username "'s shield!")}}
			{{else if eq $riot "yes"}}
				{{$damage = roundCeil (mult (fdiv (randInt 0 26) 100) $damage)}}
				{{$action = (joinStr "" $action "\n Reduced to " (toInt $damage) " by " $wp.Username "'s riot shield!")}}
				{{$riot = "no"}}
			{{else if eq $club "yes"}}
				{{$damage = roundCeil (mult $damage (randInt 0 2))}}
					{{if eq (toInt ($damage)) 0}}
						{{$action = (joinStr "" $cp.Username " was too disoriented from their clubbing to hit " $wp.Username "!")}}
						{{$club = "no"}}
					{{end}}
			{{end}}
			{{$shield = "yes"}} 
	{{else if (eq .Reaction.Emoji.ID 670664127701581829)}}{{/*Med*/}}
		{{$heal = (randInt 5 11)}}	
		{{$action = (joinStr "" $cp.Username " takes a step back, and chugs a Medical Brew, healing for " (toInt $heal) " points of health!")}}
				{{$shield = "no"}}	{{$riot = "no"}}	{{$club = "no"}}
	{{else if (eq .Reaction.Emoji.ID 670664127588466708)}}{{/*Rail*/}}
			{{if and (eq $cp.ID $player1.ID) (eq $rail1 "charged")}}
				{{$damage = roundCeil (mult (randInt 5 11) 2)}}
				{{$action = (joinStr "" $cp.Username " blasted " $wp.Username " with their railgun, doing " (toInt $damage) " points of damage!")}}
				{{if eq $shield "yes"}}
					{{$action = (joinStr "" $action "\n " $wp.Username "'s shield was completely shattered and offered no protection!")}}
					{{$shield = "no"}}
				{{else if eq $riot "yes"}}
					{{$damage = roundCeil (mult (fdiv (randInt 50 76) 100) $damage)}}
					{{$action = (joinStr "" $action "\n Reduced to " (toInt $damage) " by " $wp.Username "'s riot shield!")}}
					{{$riot = "no"}}
				{{end}}
			{{$rail1 = "discharged"}}
			{{else if and (eq $cp.ID $player2.ID) (eq $rail2 "charged")}}
				{{$damage = toInt (roundCeil (mult (randInt 5 11) 2))}}
				
				{{$action = (joinStr "" $cp.Username " blasted " $wp.Username " with their railgun, doing " (toInt $damage) " points of damage!")}}		
				{{if eq $shield "yes"}}
					{{$action = (joinStr "" $action "\n " $wp.Username "'s shield was completely shattered and offered no protection!")}}
					{{$shield = "no"}}
				{{else if eq $riot "yes"}}
					{{$damage = roundCeil (mult (fdiv (randInt 50 76) 100) $damage)}}
					{{$action = (joinStr "" $action "\n Reduced to " (toInt $damage) " by " $wp.Username "'s riot shield!")}}
					{{$riot = "no"}}
				{{end}}
			{{$rail2 = "discharged"}}
			{{else if (eq $cp.ID $player2.ID)}}
				{{$action = (joinStr "" $cp.Username " is charging their railgun.")}}
				{{$rail2 = "charging"}}
				{{$riot = "no"}}{{$shield = "no"}}
			{{else if (eq $cp.ID $player1.ID)}}
				{{$action = (joinStr "" $cp.Username " is charging their railgun.")}}
				{{$rail1 = "charging"}}
				{{$riot = "no"}}{{$shield = "no"}}
			{{end}}
			{{$club = "no"}}
	{{else if (eq .Reaction.Emoji.ID 670664127324356608)}}{{/*Club*/}}
		{{$damage = toInt (roundCeil (div (randInt 5 11) 2))}}
		{{$action = (joinStr "" $cp.Username " swung at " $wp.Username " hard with their club, doing " (toInt $damage) " points of damage!")}}
				{{if eq $shield "yes"}}
				{{$damage = roundCeil (mult (fdiv (randInt 50 76) 100) $damage)}}
				{{$action = (joinStr "" $action "\n Reduced to " (toInt $damage) " by " $wp.Username "'s shield!")}}
			{{else if eq $riot "yes"}}
				{{$damage = roundCeil (mult (fdiv (randInt 0 26) 100) $damage)}}
				{{$action = (joinStr "" $action "\n Reduced to " (toInt $damage) " by " $wp.Username "'s riot shield!")}}
			{{else if eq $club "yes"}}
				{{$damage = roundCeil (mult $damage (randInt 0 2))}}
					{{if eq (toInt ($damage)) 0}}
						{{$action = (joinStr "" $cp.Username " was too disoriented from their clubbing to hit " $wp.Username "!")}}
					{{end}}
			{{end}}
		{{$club = "yes"}}{{$shield = "no"}}{{$riot = "no"}}
	{{else if (eq .Reaction.Emoji.ID 670664127806570496)}}{{/*Riot*/}}
		{{$action = (joinStr "" $cp.Username " took a defensive position behind a Riot Shield.")}}
		{{$riot = "yes"}}
		{{$shield = "no"}}
		{{$club = "no"}}
	{{end}}
	{{$hpchange = (mult $damage -1)}}
	{{if (eq $cp.ID $player1.ID)}}
		{{$P2HP = (add $P2HP $hpchange)}}{{if gt (toInt $P2HP) $P2MAX}}{{$P2HP = $P2MAX}}{{end}}
		{{$P1HP = (add $P1HP $heal)}}{{if gt (toInt $P1HP) $P1MAX}}{{$P1HP = $P1MAX}}{{end}}
	{{else if (eq $cp.ID $player2.ID)}}
		{{$P1HP = (add $P1HP $hpchange)}}{{if gt (toInt $P1HP) $P1MAX}}{{$P1HP = $P1MAX}}{{end}}
		{{$P2HP = (add $P2HP $heal)}}{{if gt (toInt $P2HP) $P2MAX}}{{$P2HP = $P2MAX}}{{end}}
	{{end}}
		{{if eq $rail1 "charging"}}
			{{$rail1 = "charged"}}
		{{end}}
		{{if eq $rail2 "charging"}}
			{{$rail2 = "charged"}}
		{{end}}
	{{if or (le (toInt $P1HP) 0) (le (toInt $P2HP) 0)}}
		 {{if (le (toInt $P1HP) 0)}}
		 {{$win = $player2}}
		 {{$lose = $player1}}
		{{$NEWMAX = add (toInt $P2MAX) 1}}
		{{dbSet $player2.ID "BattleHP" $NEWMAX}}
		{{dbSet $player1.ID "BattleHP" $P1MAX}}
		 {{else}}
		 {{$win = $player1}}
		 {{$lose = $player2}}
		{{$NEWMAX = add (toInt $P1MAX) 1}}
		{{dbSet $player1.ID "BattleHP" $NEWMAX}}
		{{dbSet $player2.ID "BattleHP" $P2MAX}}
		 {{end}}
		{{$embed := cembed 
		"Title" "BATTLE OVER"
		"description" "The battle is over! Results below!"
			"fields" (cslice 
		(sdict "name" "Winner" "value" (toString $win.Username) "inline" true)
		(sdict "name" "Loser" "value" (toString $lose.Username) "inline" true)
		(sdict "name" "Final Action" "value" $action "inline" false))
		"thumbnail" (sdict "url" "https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/120/lg/57/crossed-swords_2694.png") }}
		{{deleteMessage nil $game}}
		{{takeRoleID $player1 667754372456775680}}
		{{takeRoleID $player2 667754372456775680}}
		{{sendMessage nil (joinStr "" "1 Point added to " $win.Mention "'s Max HP. New Max HP - " (toInt $NEWMAX) ".")}}
				{{sendMessage nil $embed}}
		{{dbDel 0 "Battle"}}
	{{else}}
	{{$embed := cembed 
		"Title" "It's BATTLE Time!"
		"description" "When it's your turn, you can choose to either attack with a <:sword:633311599020736531>,  a <:Club:670664127324356608>,  or a <:Railgun:670664127588466708>. Or you can defend with a <:shield:667754698639540254>, or <:Riot_Shield:670664127806570496>. You could also drink a <:Medical_Brew:670664127701581829>. Each option has its own benefits and drawbacks. See `?battle` for more info."
			"fields" (cslice 
		(sdict "name" (toString $player1.Username) "value" (joinStr "" "HP: " (toString $P1HP)) "inline" true)
		(sdict "name" (toString $player2.Username) "value" (joinStr "" "HP: " (toString $P2HP)) "inline" true)
		(sdict "name" "Action" "value" $action "inline" false)
		(sdict "name" "Who's Turn Is it?" "value"  $wp.Username "inline" false))
		"thumbnail" (sdict "url" "https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/120/lg/57/crossed-swords_2694.png") }}
		{{deleteMessage nil $game}}
			{{$newgame := sendMessageRetID nil $embed}}  
			{{takeRoleID $cp 667754372456775680}}
			{{giveRoleID $wp 667754372456775680}}
               {{addMessageReactions nil $newgame ":attack:633311599020736531" ":shield:667754698639540254" ":Riot_Shield:670664127806570496" ":Club:670664127324356608" ":Railgun:670664127588466708" ":Medical_Brew:670664127701581829" }}
		{{$battle.Set "cp" (toString $wp.ID)}}{{$battle.Set "wp" (toString $cp.ID)}}{{$battle.Set "hp1" $P1HP}}{{$battle.Set "hp2" $P2HP}}{{$battle.Set "shield" $shield}}{{$battle.Set "game" (toString $newgame)}}{{$battle.Set "riot" $riot}}{{$battle.Set "rail1" $rail1}}{{$battle.Set "rail2" $rail2}}{{$battle.Set "club" $club}}
				{{dbSet 0 "Battle" $battle}}
		{{end}}
{{end}}