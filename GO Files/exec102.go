{{$color := randInt 111111 999999 }}
 {{$embed := cembed 
 "color" $color
"Title" "Chibis"
"Description" "The following List includes all the Chibis that can be purchased from the Arkmart."
 "fields" (cslice 
 (sdict "name" "__**Chibis**__" "value" "Allosaurus/nAnkylosaurus/nArgentavis/nBaryonyx/nBog Spider/nBrontosaurus/nCarbonemys/nCarno/nCastroides/nCherufe/nDaeodon/nDirewolf/nDoedicurus/nEquus/nGasbag/nGiganotosaurus/nGigantopithecus/nIguanodon/nKentrosaurus/nMammoth/nMegalania/nMegaloceros/nMegatherium/nMoschops/nOviraptor/nOvus/nParaceratherium/nParasaur/nPhiomia/nPhoenix/nProcoptodon/nPteranodon/nPulmonoscorpius/nQuetzal/nRaptor/nReaper/nRex/nRhino/nRock Drake/nRock Golem/nRollrat/nSabertooth/nShapeshifter (Large)/nShapeshifter (Small)/nSnow Owl/nSpino/nStego/nTapejara/nTerror Bird/nTherizino/nThylacoleo/nTrike/nWyvern/nYutyrannus" "inline" false)  ) }}
{{sendMessage nil $embed}}



{{$args := parseArgs 2 "proper syntax is -subscribe userid duration"
 (carg "userid" "userid")
 (carg "duration" "sublength") }}

{{$user := ($args.Get 0)}}
{{$duration := ($args.Get 1)}}
{{$expire := currentTime.Add $duration }}
{{$formatted := $expire.Format "1 _2 2006" }}
{{dbSet $user "sublength" $formatted}}

{{sendMessage 669542893438107669 "Sub Added Successfully"}}

{{$duration}}
{{$expire}}
{{$formatted}}



{{$db := dbBottomEntries "XP" 5 0}}
{{if $db}}
	{{range $db}}
	{{dbDel .User.ID "XP"}}
	{{end}}
{{execCC 32 nil 0 nil}}
{{else}}
All XP Cleared. 
{{end}}

{{$entries := dbCount "sublength"}}


{{$duration = (toDuration "2d")}}
{{$compareto := (currentTime.Add $duration).Format "1 _2 2006"}}
{{range dbBottomEntries "sublength" $entries 0}}
{{if eq .Value $compareto}}
Correct
{{else}}
Not Correct
{{end}}{{end}}






"Faffing" "Cack-handed" "Tub-Thumping" "Burping" "Corrugated" "Momentous" "Unrepentant" "Dribbling" "Quacking" "Eye-Watering" "Ironclad" "Thatched" "Horse-bothering" "Indubitable" "Knackered old" "Grotty little" "Ululating" "Pear-shaped" "Pebb le-dashed" "Numpty" "Egregious" "Chundering" "Tinkling" "Preposterous" "Knicker-stealing" "Inconceivable"



"Piss" "Marmite" "Wank" "Bodge" "Cock" "Pants" "Minge" "Ham" "Willy" "Poo" "Tit" "Sock" "Arse" "Pudding" "Knob" "Jam" "Pranny" "Bell-end" "Sollock" "Bum" "Anorak" "Fanny" "Swamp" "Haggis" "Mince" "Muppet"


"Sniffer" "Crumpet" "Kipper" "Trolley" "Womble" "Hammer" "Sandwich" "Bassoon" "Pillock" "Plonker" "Vicar" "Ferret" "Lorry" "Pilot" "Pony" "Stain" "Fluffer" "Spanner" "Biscuit" "Flap" "Scrounger" "Merchant" "Basket" "Wanker" "Monkey" "Bonnet"
