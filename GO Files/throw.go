{{$args := parseArgs 1 "throw who"
(carg "user" "user")}}
{{$target := $args.Get 0}}
{{deleteTrigger 1}}
{{$prob := randInt 1 1001}}
{{$numlist := cslice "Tomatoes" "Narcoberries" "Narcotics" "Tranq Darts" "Energy Brews" "Rock Drake Feathers" "Stones" "Box o’ Chocolates" "Water Balloons" "Gems" "Boogers" "Dollar Bills" "Pennies" "Ninja Stars" "Human Body Parts" "Babies" "Longrass" "Medical Brews" "Updates" "Tide Pods" "Cannon Balls" "Tek Grenades" "Poison Grenades" "C4 Charges"}}
{{$singlelist := cslice "a Knife" "Anime Girls" "Hentai" "an Admin" "Nickelback" "Life Advice" "A Bribe" "An Engagement Ring" "Nothing" "some Monday Blues" "Insults" "Yez" "a Rubiks Cube" "Depression" "Compliments" "A Look" "A Hint" "A Chair" "Crappy Code" "A Stripper" "Sexy Firemen" "A Server Reset" "A Glitch" "A Lifetime of Despair" "A Cardboard Box" "Hate" "Regret" "Shade" "A Dab" "A Dance Move" "An indecent Proposal" "Feces" "A Bola" "Glow Sticks" "Salt" "A Spear" "A Pike" "A Metal Hatchet" "A Torch" "A Metal Pick" "A Chainsaw" "A Boomerang" "A Flame Arrow" "A Tranq Arrow" "A Spear Bolt" "A Tek Gravity Grenade" "A Cluster Grenade" "A Smoke Grenade" "A Bear Trap" "A Cryopod" "Absorbent Substrate" "A Wooden Door" "Ammonite Bile" "Angler Gel" "An Egg" "An Artifact" "Giant Bee Honey" "Kibble" "Beer" "A Birthday Cake" "Hair" "Broth of Enlightenment" "Bug Repellent" "Cementing Paste" "A Chain Bola" "A Chibi" "Citronal" "A Boot" "A Truck" "An Eviction" "A Dead-end Career" "Feelings" "Sack of Flour" "Sand" "Savorout" "A Snowball" "Some Soap" "Soothing Balm" "Spoiled Wyvern Milk" "Sparkpowder" "A Specimen Implant" "A Stimberry" "A Stimulant" "A Survivor’s Trophy" "A Flare" "Fireworks" "A Kiss" "A Wishbone" "A Wooden Cage" "Handcuffs" "Fruit" "A Frying Pan" "A Creampuff" "A Slipper" "A Bra" "Some Dirty Underwear" "Toilet Paper" "A Pillow" "Yourself" "Your Mom" "Trash" "Midgets" "Salty Nuts" "A Purse" "A Cuddly Bear" "An Aggressive Bear" "A Rabid Dog" "Some Balls" "Mayo" "Ketchup" "Spice" "Pizza" "Fertilizer" "A Cone of Shame" "A Congealed Gas Ball" "A Corrupt Heart" "A Charge" "Some Debt" "A Wedding Proposal" "Divorce Papers" "Acid" "Candy" "Element" "Some Hands" "Affection" "Love" "Wood" "A Fishing Rod" "Dye" "A Gacha Crystal" "Gasoline" "A Gift Box" "A Grappling Hook" "An Emote" "An Industrial Grill" "A Mindwipe Tonic" "A Shot" "Nerdy Glasses" "A Note" "An Oil Jar" "An Old Smelly Fish" "Some Old Moldy Cheese" "A Tek Atv" "a Yutyrannus" "a Woolly Rhino" "a Vulture" "a Velonasaur" "a Unicorn" "a Tusoteuthis" "a Troodon" "a Trilobite" "a Triceratops" "a Titanosaur" "a Titanoboa" "a Thylacoleo" "a Thorny Dragon" "a Therizinosaur" "a Terror Bird" "a Tapejara" "a Stegosaurus" "a Spinosaur" "a Snow Owl" "a Shinehorn" "a Seeker" "a shapeshifter" "a Scout" "a Sarco" "a Sabertooth Salmon" "a Sabertooth" "a Roll Rat" "a Rock Elemental" "a Rock Drake" "a Rex" "a Reaper" "a Ravager" "a Raptor" "a Quetzal" "a Purlovia" "a Pulmonoscorpius" "a Pteranodon" "a Procoptodon" "a Poison Wyvern" "a Plesiosaur" "a Piranha" "a Phoenix" "a Phiomia" "a Pelagornis" "a Pegomastax" "a Parasaur" "a Paraceratherium" "a Pachyrhinosaurus" "a Pachy" "an Ovis" "an Oviraptor" "an Otter" "an Onyc" "a Nameless" "a Moschops" "a Mosasaurus" "a Morellatops" "a Microraptor" "a Mesopithecus" "a Megatherium" "a Meganeura" "a Megalosaurus" "a Megalodon" "a Megaloceros" "a Megalania" "a Mantis" "a Manta" "a Managarmr" "a Mammoth" "a Lystrosaurus" "a Lymantria" "a Liopleurodon" "a Lightning Wyvern" "a Leedsichthys" "a Leech" "a Lamprey" "a Kentrosaurus" "a Karkinos" "a Kaprosuchus" "a Kairuku" "a Jug Bug" "a Jerboa" "an Iguanodon" "an Ichthyosaurus" "an Ichthyornis" "an Ice Wyvern" "a Hyaenodon" "a Hesperornis" "a Griffin" "a Glowtail" "a Glowbug" "a Gigantopithecus" "a Giganotosaurus" "a Giant Queen Bee" "a Gasbag" "a Gallimimus" "a Gacha" "a Fire Wyvern" "a Featherlight" "an Eurypterid" "an Equus" "an Enforcer" "an Electrophorus" "a Dunkleosteus" "a Dung Beetle" "a Doedicurus" "a Dodo" "a Direwolf" "a Direbear" "a Diplodocus" "a Diplocaulus" "a Dimorphodon" "a Dimetrodon" "a Dilophosaur" "a Deinonychus" "a Deathworm" "a Daeodon" "a Compy" "a Coelacanth" "a Cnidaria" "a Chalicotherium" "a Castoroides" "a Carnotaurus" "a Carbonemys" "a Bulbdog" "an Brontosaurus" "a Beelzebufo" "a Basilosaurus" "a Basilisk" "a Baryonyx" "an Attack Drone" "an Arthropluera" "an Argentavis" "an Archaeopteryx" "an Araneo" "an Ankylosaurus" "an Anglerfish" "an Ammonite" "an Allosaurus" "an Achatina"}}
{{$item := ""}}
{{$roles := (getMember $target.ID).Roles}}
{{$pos := 0}}
{{$color := 0}}
{{$maxsingle := (len $singlelist)}}
{{$maxnum := (len $numlist)}}
{{$stuff := cslice (index (shuffle $singlelist) (randInt $maxsingle)) (joinStr "" (randInt 100) " " (index (shuffle $numlist) (randInt $maxnum)))}}
{{$stuff2 := cslice (index (shuffle $singlelist) (randInt $maxsingle)) (joinStr "" (randInt 100) " " (index (shuffle $numlist) (randInt $maxnum)))}}
{{$stuff3 := cslice (index (shuffle $singlelist) (randInt $maxsingle)) (joinStr "" (randInt 100) " " (index (shuffle $numlist) (randInt $maxnum)))}}

{{if eq $target.ID .User.ID}}
{{sendMessage nil "You can't throw things at yourself, silly."}}
{{else}}
{{if le $prob 40}}
	{{range .Guild.Roles}}
		{{if in $roles .ID}}
			{{if and (lt $pos .Position) (.Color)}}
			{{$pos = .Position}}
			{{$color = .Color}}
			{{end}}
		{{end}}
	{{end}}
{{$xp := toInt ((dbGet $target.ID "XP").Value)}}
{{$currentlevel := roundFloor (sqrt (fdiv $xp 100))}}
{{$count = randInt 100 501}}
{{$newxp := (toString (add (toInt (dbGet $target.ID "XP").Value) $count))}}
{{dbSet $target.ID "XP" $newxp}}
{{$newlevel := roundFloor (sqrt (fdiv (toInt $newxp) 100))}}
{{$item = (joinStr " " $count "XP")}}

	{{if gt $newlevel $currentlevel}}
	{{giveRoleName $target.ID (joinStr "" "Level " (toInt $newlevel))}}
   		{{if gt (toInt $newlevel) 5}}
    	{{takeRoleName $target.ID (joinStr "" "Level " (sub (toInt $newlevel) 5))}}
    	{{end}}
	{{$embed := (cembed 
	"title" "You leveled up!" 
	"color" $color 
	"description" (joinStr "" "Congratulations " $target.Username ", you leveled up to **level " (toInt $newlevel) "**! For more information on our XP system, type `?xp` in any channel.") 
	"thumbnail" (sdict "url" ($target.AvatarURL "1024")))}}
	

	{{if eq (toInt $newlevel) 10}}
    {{giveRoleID $target.ID 588031694779449484}}
	{{else if eq (toInt $newlevel) 15}}
    {{giveRoleID $target.ID 654027682345910282}}
	{{end}}{{end}}
{{else if le $prob 100}}
{{sendMessage nil "***Triple Throw!***"}}

{{$item = (joinStr "" (index $stuff (randInt 2)) ", " (index $stuff2 (randInt 2)) ", and " (index $stuff3 (randInt 2)))}}


{{else if le $prob 350}}
{{sendMessage nil "***Double Throw!***"}}

{{$item = (joinStr "" (index $stuff (randInt 2)) ", and " (index $stuff2 (randInt 2)))}}

{{else if le $prob 1000}}
{{$item = (index $stuff (randInt 2))}}


{{end}}

{{sendMessage nil (joinStr "" .User.Mention " threw " (lower $item) " at " $target.Mention "!")}}


{{end}}
