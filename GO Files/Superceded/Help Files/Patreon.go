{{$color := randInt 111111 999999 }}
 {{$embed := cembed 
 "color" $color
"Title" "How to Link your Patreon to your Discord"
"Description" "If you have a sponsorship set up with Patreon, you should link your Patreon to your Discord to get your role that is associated with your sponsorship tier. Here's how:"
"fields" (cslice
(sdict "name" "__**Step 1**__" "value" "Log into your Patreon account." "inline" false) 
(sdict "name" "__**Step 2**__" "value" "In the top right corner of the screen, click the dropdown menu." "inline" false) 
(sdict "name" "__**Step 3**__" "value" "Click on My Profile Settings drop down menu." "inline" false) 
(sdict "name" "__**Step 4**__" "value" "Click on Apps." "inline" false) 
(sdict "name" "__**Step 5**__" "value" "Click on the Connect button next to the Discord app." "inline" false)) }}
{{sendMessage nil $embed}}