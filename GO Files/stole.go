{{$args := parseArgs 1 "steal from who?"
(carg "user" "user")}}
{{$target := $args.Get 0}}
{{$prob := randInt 1 1001}}
{{$singlelist := cslice "your money" "all your money" "your wallet" "your phone" "your cash" "your dignity" "your girlfriend" "your boyfriend" "your guy" "your girl" "your husband" "your wife" "your partner" "your friend" "your date" "your gold" "your dinos" "your BBS" "your car" "your truck" "your tek atv" "your glider suit" "your element dust gacha" "your element" "your tv" "your pc" "your xbox" "your ark game" "your controller" "your headset" "your dog" "your cat" "your pizza" "your food" "your donut" "your coffee" "your energy brew" "your gacha crystal" "your ID" "your keys" "your cryopod " "your rex" "your raptor" "your giga" "your wyvern" "your wyvern milk" "your griffin" "your jerboa" "your equus" "your argy" "your ascendant longneck rifle" "your ascendant crossbow" "your armor" "your hazard suit" "your rock drake" "your scuba gear" "your tek gear" "your saddle" "your box o’ chocolates " "all of your fertilized eggs" "all of your kibble " "your beer" "your wine" "your cereal" "your kiss" "your bacon" "your cooked meat" "your candy" "your clothes" "your tobacco" "your smokes" "your drink" "your turret" "your feces" "your love" "your jewelry " "your bicycle " "your tide pods" "your rims" "your alcohol" "your fajitas" "your jackpot winnings" "your balls" "your gold" "your loot" "your watch" "your underwear" "your childhood" "your shoes" "your pants" "your shirt" "your mojo" "your last name" "your job" "your tampax" "your bra" "your joy" "your pride" "your winnings" "your perfect tame" "your pteranodon " "your thyla" "your chibi" "your lawnmower " "your lunch money" "your fishing rod" "your tranq darts" "your gems" "your hentai" "your insults" "your cardboard box" "your cardboard sign" "your hatchet" "your pick" "your dermis" "your artifact" "your trophy" "your Yutyrannus" "your Woolly Rhino" "your Vulture" "your Velonasaur" "your Unicorn" "your Tusoteuthis" "your Troodon" "your Trilobite" "your Triceratops" "your Titanosaur" "your Titanoboa" "your Thylacoleo" "your Thorny Dragon" "your Therizinosaur" "your Terror Bird" "your Tapejara" "your Stegosaurus" "your Spinosaur" "your Snow Owl" "your Shinehorn" "your Seeker" "your Scout" "your Sarco" "your Sabertooth Salmon" "your Sabertooth" "your Roll Rat" "your Rock Elemental" "your Rock Drake" "your Rex" "your Reaper" "your Ravager" "your Raptor" "your Quetzal" "your Purlovia" "your Pulmonoscorpius" "your Pteranodon" "your Procoptodon" "your Poison Wyvern" "your Plesiosaur" "your Piranha" "your Phoenix" "your Phiomia" "your Pelagornis" "your Pegomastax" "your Parasaur" "your Paraceratherium" "your Pachyrhinosaurus" "your Pachy" "your Ovis" "your Oviraptor" "your Otter" "your Onyc" "your Nameless" "your Moschops" "your Mosasaurus" "your Morellatops" "your Microraptor" "your Mesopithecus" "your Megatherium" "your Meganeura" "your Megalosaurus" "your Megalodon" "your Megaloceros" "your Megalania" "your Mantis" "your Manta" "your Managarmr" "your Mammoth" "your Lystrosaurus" "your Lymantria" "your Liopleurodon" "your Lightning Wyvern" "your Leedsichthys" "your Leech" "your Lamprey" "your Kentrosaurus" "your Karkinos" "your Kaprosuchus" "your Kairuku" "your Jug Bug" "your Jerboa" "your Iguanodon" "your Ichthyosaurus" "your Ichthyornis" "your Ice Wyvern" "your Hyaenodon" "your Hesperornis" "your Griffin" "your Glowtail" "your Glowbug" "your Gigantopithecus" "your Giganotosaurus" "your Giant Queen Bee" "your Gasbag" "your Gallimimus" "your Gacha" "your Fire Wyvern" "your Featherlight" "your Eurypterid" "your Equus" "your Enforcer" "your Electrophorus" "your Dunkleosteus" "your Dung Beetle" "your Doedicurus" "your Dodo" "your Direwolf" "your Direbear" "your Diplodocus" "your Diplocaulus" "your Dimorphodon" "your Dimetrodon" "your Dilophosaur" "your Deinonychus" "your Deathworm" "your Daeodon" "your Compy" "your Coelacanth" "your Cnidaria" "your Chalicotherium" "your Castoroides" "your Carnotaurus" "your Carbonemys" "your Bulbdog" "your Brontosaurus" "your Beelzebufo" "your Basilosaurus" "your Basilisk" "your Baryonyx" "your Attack Drone" "your Arthropluera" "your Argentavis" "your Archaeopteryx" "your Araneo" "your Ankylosaurus" "your Anglerfish" "your Ammonite" "your Allosaurus" "your Achatina"}}
{{$item := ""}}
{{$roles := (getMember $target.ID).Roles}}
{{$pos := 0}}
{{$color := 0}}
{{$maxsingle := (len $singlelist)}}


{{if eq $target.ID .User.ID}}
{{sendMessage nil "You can't steal things from yourself, silly."}}
{{else}}



{{if le $prob 100}}
{{sendMessage nil "***Triple Steal!***"}}

{{$item = (joinStr "" (index ($singlelist) (randInt $maxsingle)) ", " (index ($singlelist) (randInt $maxsingle)) ", and " (index ($singlelist) (randInt $maxsingle)))}}


{{else if le $prob 300}}
{{sendMessage nil "***Double Steal!***"}}

{{$item = (joinStr "" (index ($singlelist) (randInt $maxsingle)) ", and " (index ($singlelist) (randInt $maxsingle)))}}

{{else if le $prob 1000}}
{{$item = (index ($singlelist) (randInt $maxsingle))}}


{{end}}

{{sendMessage nil (joinStr "" $target.Mention ", " .User.Mention " stole " (lower $item) "!")}}


{{end}}





