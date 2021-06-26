{{ $start := currentTime }}
{{$color := randInt 111111 999999 }}
 {{$embed := cembed 
 "color" $color
"Title" "Admin Tames"
"Description" "Admin Tames are the admins signature dino, painted in their signature colors, at max level!\nThe following list outlines each admins dino, and colors."
 "fields" (cslice 
 (sdict "name" "__**Yez**__" "value" "*Deep Pink, Albino, and Actual Black* - Rock Drake" "inline" false) 
 (sdict "name" "__**Hatter**__" "value" "*Powder Blue, Jade, and Turquoise* - Maewing" "inline" false) 
 (sdict "name" "__**Gamble**__" "value" "*Albino, Lemon Lime, and Actual Black* - Tek Rex" "inline" false)
 (sdict "name" "__**Rouge**__" "value" "*Albino, Lemon Lime, and Actual Black* - Voidwyrm" "inline" false)
 (sdict "name" "__**Lady Dukes**__" "value" "*Black, Turquoise, and Unused Purple Dye* - Tropical Wyvern" "inline" false)
 (sdict "name" "__**Draygoon**__" "value" "*Glacier Blue, and Black* - Shadowmane" "inline" false)
 (sdict "name" "__**Bug**__" "value" "*Teal, Deep Pink, and Lemon Lime* - R-Thylacoleo" "inline" false)
 (sdict "name" "__**Hermit - Legacy**__" "value" "*Actual Black, and Green* - Velonosaur" "inline" false)
 (sdict "name" "__**Damien - Legacy**__" "value" "*Dark Teal, Glacial Blue, and Mint Green* - X Rex" "inline" false)  
 (sdict "name" "__**Ellie - Legacy**__" "value" "*Glacial Blue, Dark Lavender, and Dark Teal* - Snow Owl" "inline" false) 
 (sdict "name" "__**Josh - Legacy**__" "value" "*Actual Black,  Albino, and Glacial Blue* - Ferox" "inline" false)  ) }}
{{sendMessage nil $embed}}
{{ $dur := currentTime.Sub $start }}
{{ if gt $dur.Seconds 5.0 }}
    {{ sendMessage 587858995012698137 (print "CC #" .CCID " took longer than 5 seconds to run.\n**Duration:** " $dur "\n**Command executed:* `" .Message.Content "`") }}
{{ end }}