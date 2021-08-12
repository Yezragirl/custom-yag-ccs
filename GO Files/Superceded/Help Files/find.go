{{$color := randInt 111111 999999 }}
 {{$embed := cembed 
 "color" $color
"Title" "How to Find the Cluster in the Server List"
"Description" "Here's how to find the servers:"
"fields" (cslice
(sdict "name" "__**Step 1**__" "value" "Make sure you’re selecting Survival Evolved and NOT Primitive Plus." "inline" false) 
(sdict "name" "__**Step 2**__" "value" "Change the drop down box to Unofficial PC Servers." "inline" false) 
(sdict "name" "__**Step 3**__" "value" "Check the Password Protected Box." "inline" false) 
(sdict "name" "__**Step 4**__" "value" "In the Search bar, type `Yez`." "inline" false)) }}
{{sendMessage nil $embed}}
