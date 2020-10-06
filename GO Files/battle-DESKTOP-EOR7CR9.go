{{if (hasRoleID 667754372456775680)}}{{/*Has turn role*/}}
{{$db := dbGet 0 "Battle"}}
{{$battle := sdict}}{{range $k,$v := $db.Value}}{{$battle.Set $k $v}}{{end}}
{{$game := $battle.game}}
{{$player1 := (userArg $battle.p1)}}
{{$player2 := (userArg $battle.p2)}}
{{$P1HP := $battle.hp1}}
{{$P2HP := $battle.hp2}}
{{$CP := $battle.wp}}
{{$WP := $battle.cp}}
{{$damage := (randInt 4 16)}}
{{$shield := $battle.shield}}
{{$action := ""}}
{{$win := ""}}
{{$lose := ""}}
{{$turn := (userArg $battle.turn)}}
{{$countp2 := (dbGet 0 "p2win").Value}}
{{$countp1 := (dbGet 0 "p1win").Value}}
	{{if (eq .Reaction.Emoji.ID 633311599020736531)}}{{/*attack*/}}
		{{if eq (toString $battle.turn) (toString $battle.p1)}}
			{{if eq (toString $shield) "yes"}}
			{{$damage = (toInt (roundCeil (div $damage 2)))}}
			{{$P2HP = (sub $P2HP $damage)}}
			{{$shield = "no"}}
			{{else}}
			{{$P2HP = (sub $P2HP $damage)}}
			{{end}}
			{{takeRoleID $player1 667754372456775680}}
			{{giveRoleID $player2 667754372456775680}}
                                     {{if eq $damage 0}}
			             {{$action = (joinStr "" $player1.Username " missed!")}}
                                     {{else}}
		                     {{$action = (joinStr "" $player1.Username " attacked " $player2.Username " for " (toString $damage) " points of damage!")}}
                                     {{end}}
			{{$turn = $player2}}
		{{else}}
			{{if eq (toString $shield) "yes"}}
			{{$damage = (toInt (roundCeil (div $damage 2)))}}
			{{$P1HP = (sub $P1HP $damage)}}
			{{$shield = "no"}}
			{{else}}
			{{$P1HP = (sub $P1HP $damage)}}
			{{end}}
			{{takeRoleID $player2 667754372456775680}}
			{{giveRoleID $player1 667754372456775680}}
                                     {{if eq $damage 0}}
			             {{$action = (joinStr "" $player2.Username " missed!")}}
                                     {{else}}
		                     {{$action = (joinStr "" $player2.Username " attacked " $player1.Username " for " (toString $damage) " points of damage!")}}
                                     {{end}}
			{{$turn = $player1}}
 		{{end}}
 	
	{{else if (eq .Reaction.Emoji.ID 667754698639540254)}}{{/*shield*/}}
		{{if eq (toString $battle.turn) (toString $battle.p1)}}
		{{$shield = "yes"}}
		{{takeRoleID $player1 667754372456775680}}
		{{giveRoleID $player2 667754372456775680}}
		{{$action = (joinStr "" $player1.Username " took a defensive position.")}}
		{{$turn = $player2}}
		{{else}}
		{{$shield = "yes"}}
		{{takeRoleID $battle.p2 667754372456775680}}
		{{giveRoleID $battle.p1 667754372456775680}}
		{{$action = (joinStr "" $player2.Username " took a defensive position.")}}
		{{$turn = $player1}}
		{{end}}

	{{end}}

	{{if or (le (toInt $P1HP) 0) (le (toInt $P2HP) 0)}}
		 {{if (le (toInt $P1HP) 0)}}
		 {{$win = $player2}}
		 {{$lose = $player1}}
{{$countp2 = (dbIncr 0 "p2win" 1)}}
		 {{else}}
		 {{$win = $player1}}
		 {{$lose = $player2}}
{{$countp1 = (dbIncr 0 "p1win" 1)}}
		 {{end}}
		{{$embed := cembed 
		"Title" "BATTLE OVER"
		"description" "The battle is over! Results below!"
			"fields" (cslice 
		(sdict "name" "Winner" "value" (toString $win.Username) "inline" true)
		(sdict "name" "Loser" "value" (toString $lose.Username) "inline" true)
		(sdict "name" "Final Action" "value" $action "inline" false))
		"thumbnail" (sdict "url" "https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/120/lg/57/crossed-swords_2694.png") }}
		{{editMessage nil $game $embed}}
{{sendMessage nil (joinStr "" "**Current Tally**\n*Player 1* - " (toString (toInt $countp1)) "\n*Player 2* - " (toString (toInt $countp2)))}}
                {{deleteAllMessageReactions nil $game}}
		{{takeRoleID $player1 667754372456775680}}
		{{takeRoleID $player2 667754372456775680}}
		{{dbDel 0 "Battle"}}
	{{else}}
	{{$embed := cembed 
		"Title" "It's BATTLE Time!"
		"description" "When it's your turn, you can choose to either attack <:sword:633311599020736531> or defend <:shield:667754698639540254>. If you defend, their next attack will only do 50% damage. Attacks do between 0 and 15 damage."
			"fields" (cslice 
		(sdict "name" (toString $player1.Username) "value" (joinStr "" "HP: " (toString $P1HP)) "inline" true)
		(sdict "name" (toString $player2.Username) "value" (joinStr "" "HP: " (toString $P2HP)) "inline" true)
		(sdict "name" "Action" "value" $action "inline" false)
		(sdict "name" "Who's Turn Is it?" "value"  $turn.Username "inline" false))
		"thumbnail" (sdict "url" "https://emojipedia-us.s3.dualstack.us-west-1.amazonaws.com/thumbs/120/lg/57/crossed-swords_2694.png") }}
		{{editMessage nil $game $embed}}
               {{deleteAllMessageReactions nil $game}}
               {{addMessageReactions nil $game ":attack:633311599020736531" ":shield:667754698639540254" }}
		{{$battle.Set "turn" (toString $turn.ID)}}{{$battle.Set "hp1" $P1HP}}{{$battle.Set "hp2" $P2HP}}{{$battle.Set "shield" $shield}}
		{{dbSet 0 "Battle" $battle}}
		{{end}}
{{end}}



$Railgun := 0
$Club := "no"
$Riot := "no"


Sword - Does 5 to 15 damage.
Shield - does 0 to 5 damage and Reduces next attack by 20-30%
Medical Brew - Increases HP by 15-30
Railgun - double damage, next round, no damage this round. (this could be tough...)
Club - 40-60% damage, has chance to stun opponent next turn causing them to deal no damage
riot shield- reduces next attack by 75-100%