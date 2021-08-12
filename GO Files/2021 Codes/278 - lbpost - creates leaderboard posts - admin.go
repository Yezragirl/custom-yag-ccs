{{/*lbpost formatted embed for adding Leaderboards to the event leaderboard channel*/}}
{{$embed := cembed 
	"title" "Building Challenge Leaderboard (and Prizes)" 
	"description" "*Building Challenge July 2021 Results*" 
	"color" 3080447
	"fields" (cslice 
		   (sdict "name" "1st Place - 20,000 <:bbs:583143695927214091>" "value" "*<:BuildingChallenge:871949074826395658>Arkward<:BuildingChallenge:871949074826395658>*" "inline" false) 
		   (sdict "name" "2nd Place - 17,500 <:bbs:583143695927214091>" "value" "*The Crows*" "inline" false) 
		   (sdict "name" "3rd Place - 15,000 <:bbs:583143695927214091>" "value" "*Rexual Frustration*" "inline" false)
		   (sdict "name" "4th Place - 12,500 <:bbs:583143695927214091>" "value" "*Tribe of Witches Gold*" "inline" false)
		   (sdict "name" "5th Place - 10,000 <:bbs:583143695927214091>" "value" "*Flak to the Future*" "inline" false)
		   (sdict "name" "Participation - 2,500 <:bbs:583143695927214091>" "value" "*Reptile Dysfunction*\n*Cloud9*" "inline" false)) }}
	{{sendMessage nil $embed}}
	{{deleteTrigger 1}}
	