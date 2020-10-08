{{$args := parseArgs 1 "map which?"
(carg "string" "user")}}
{{$maps := sdict "manticore" 1 "artemis" 2 "gryphon" 3 "zelda" 4 "phoenix" 5 "elysian" 6 "titan" 7 "raiden" 8 "medusa" 9 "beowulf" 10 "osiris" 11 "sphinx" 12 "silvestria" 13 "setup" 14}}
{{$map := $maps.Get ($args.Get 0)}}
{{if $map}}
{{$mapname := (dbGet $map "name").Value}}
{{$cclocation := (dbGet $map "cclocation").Value}}
{{$coords := (dbGet $map "coords").Value}}
{{$adminlocations := (dbGet $map "adminlocations").Value}}
{{$dinos := (dbGet $map "dinos").Value}}
{{$color := randInt 111111 999999 }}
 {{$embed := cembed 
 "color" $color
"Title" $mapname
"Description" "Each of our PVE maps has a Community Center with basic crafting stations available to the public. That's also where you pick up Arkmart orders, starters, and some include a hatchery!\nAll CC's are equipped with FREE food, FREE water, and FREE honey. Please put in a ticket if you find one of these things lacking from a CC.\nAll maps have public teleporters at all Obis."
"fields" (cslice 
	(sdict "name" "__**Community Center Coordinates**__" "value" $coords)
	(sdict "name" "__**Community Center Location**__" "value" $cclocation)
	(sdict "name" "__**Admin Teleporters and Locations of Interest**__" "value" $adminlocations)
	(sdict "name" "__**Non-Native Dinos**__" "value" $dinos))}}
	{{sendMessage nil $embed}}
{{else}}
{{$color := randInt 111111 999999 }}
 {{$embed := cembed 
 "color" $color
"Title" "Sorry"
"Description" "The map you referenced is either not set up yet, or spelled wrong. Please check `?maps` for the proper spelling of map names and try again."}}
	{{sendMessage nil $embed}}
{{end}}