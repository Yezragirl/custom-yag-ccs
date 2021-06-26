test 1
{{$color := randInt 111111 999999 }}
test 2
 {{$embed := cembed 
 "color" $color
"Title" "How to Link My Discord Account to Patreon"
 "fields" (cslice 
 (sdict "name" "__Link __" "value" "https://discord.com/channels/@me/584230235529019418/828382375742144594" "inline" false) ) }}
test 3
{{sendMessage nil $embed}}