{{$color := randInt 111111 999999 }}
 {{$embed := cembed 
 "color" $color
"Title" "How Do I Get to an Event?"
"Description" "Most of our events are held on our event map, Sphinx. It's best to ask if the event is being held elsewhere."
 "fields" (cslice 
 (sdict "name" "__**How Do I Get to Sphinx?**__" "value" "If you've never been to an event before: 1. Log out of your current map. 2. Log in to Sphinx. 3. Spawn at Jungle 2. 4. Find the sign there, and follow the instructions. 5. Teleport to the event." "inline" false)) }}
{{sendMessage nil $embed}}


{{$args := parseArgs 1 "map which?"
(carg "string" "user")}}
{{$maps := sdict "manticore" 1 "artemis" 2 "gryphon" 3 "zelda" 4 "phoenix" 5 "elysian" 6 "titan" 7 "raiden" 8 "medusa" 9 "beowulf" 10 "osiris" 11 "sphinx" 12 "silvestria" 13}}
{{$map := $maps.Get ($args.Get 0)}}
{{$mapname := (dbGet $map "name").Value}}
{{$cclocation := (dbGet $map "cclocation").Value}}
{{$coords := (dbGet $map "coords").Value}}
{{adminlocations := (dbGet $map "adminlocations").Value}}
{{$dinos := (dbGet $map "dinos").Value}}
{{$color := randInt 111111 999999 }}
 {{$embed := cembed 
 "color" $color
"Title" $mapname
"Description" "Each of our PVE maps has a Community Center with basic crafting stations available to the public. That's also where you pick up Arkmart orders, starters, and some include a hatchery!\nAll CC's are equiped with FREE cryopods, FREE food, FREE water, and FREE honey. Please put in a ticket if you find one of these things lacking from a CC."
"fields" (cslice 
	(sdict "name" "__**Community Center Coordinates**__" "value" $coords)
	(sdict "name" "__**Community Center Location**__" "value" $cclocation)
	(sdict "name" "__**Admin Teleporters and Locations of Interest**__" "value" $adminlocations)
	(sdict "name" "__**Non-Native Dinos**__" "value" $dinos))}}
	{{sendMessage nil $embed}}



	{{$twitchnick := slice .URL 22}}
	{{$embed := sdict
		"title" "Stream Announcement"
		"color" 0xE91E63
		"description" (print .User.Mention " is currently streaming! Check it out: " .URL)
		"fields" (cslice 
			(sdict "name" "Now Playing" "value" .Game )
			(sdict "name" "Stream Status" "value" .StreamTitle ))
		"thumbnail" (sdict "url" (.User.AvatarURL "1024"))
		"image" (sdict "url" (joinStr "" "https://static-cdn.jtvnw.net/previews-ttv/live_user_" $twitchnick "-1920x1080.jpg"))
		"footer" (sdict "text" "Streamed at" "icon_url" (.User.AvatarURL "256"))
		"timestamp" currentTime
	}}
	{{sendMessageNoEscape nil (complexMessage "content" (mentionRoleID 745517062281232465) "embed" (cembed $embed))}}