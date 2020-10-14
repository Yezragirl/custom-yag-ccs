{{$color := randInt 111111 999999 }}
 {{$embed := cembed 
 "color" $color
"Title" "Element"
"Description" "Element can be obtained in the following ways:"
 "fields" (cslice 
 (sdict "name" "__**Craft**__" "value" "Unstable Element can be crafted in your inventory using Element Dust. Unstable Element Shards can also be crafted this way." "inline" false) 
 (sdict "name" "__**Found**__" "value" "Element can be found in Red Cave Drops." "inline" false) 
 (sdict "name" "__**Quests**__" "value" "Element can be earned by doing quests with Element as the listed reward." "inline" false) 
 (sdict "name" "__**Bosses**__" "value" "Element can be obtained by fighting the bosses, including the Ice Queen, and the Lava Golem." "inline" false) 
 (sdict "name" "__**Alphas**__" "value" "Element can be found on the bodies of Alpha dinos." "inline" false)
 (sdict "name" "__**Purchased**__" "value" "Element can be purchased from the Arkmart." "inline" false) ) }}
{{sendMessage nil $embed}}
