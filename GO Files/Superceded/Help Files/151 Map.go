{{$color := randInt 111111 999999 }}
 {{$embed := cembed 
 "color" $color
"Title" "Yez's Ark Cluster Maps"
"Description" "For more info on a given map, type `-map name` in any chat channel. Only use the short name though, gryphon, medusa, manticore, etc. We also have Private Access maps for certain sponsorship levels. If you are interested, please type ?sponsor for more information."
"fields" (cslice
(sdict "name" "__**Survival Evolved PVE Maps**__" "value" "Yez's Artemis - PVE Valguero\nYez's Manticore - PVE Ragnarok\nYez's Gryphon - PVE Aberration\nYez's Elysian - PVE Center\nYez's Titan - PVE Valguero\nYez's Zelda - PVE Island\nYez's Phoenix - PVE Scorched Earth\nYez's Medusa - PVE Extinction\nYez's Raiden - PVE Ragnarok\nYez's Beowulf - PVE Genesis\nYez's Osiris - PVE Crystal Isles\nYez's Excelsior - PVE Aberration\nYez's Merlin - PVE Center\nYez's Sanctuary - PVE Crystal Isles\nYez's Hades - PVE Extinction\nYez's Jupiter - PVE Genesis\nYez's Fenris - PVE Island\nYez's Haven - PVE Ragnarok\nYez's Refuge - PVE Island Starter Map\nYez's Central - PVE Valguero" "inline" false) 
(sdict "name" "__**PVP Event Maps**__" "value" "Yez's Sphinx - PVP Ragnarok Event Map (Non-Transferable)\nYez's Venus - PVP Ragnarok Event Map (Transferable)" "inline" false) 
(sdict "name" "__**Specialty Maps**__" "value" "Yez's Grayson - Test Map, Admin Only Access\nYez's Silvestria - Taming Challenge - Ragnarok" "inline" false) 
(sdict "name" "__**Primitive Plus PVE Maps**__" "value" "Yez's Anvil - PVE Primitive Plus Ragnarok\nYez's Zulu - PVE Primitive Plus Crystal Isles" "inline" false) ) }}
{{sendMessage nil $embed}}

