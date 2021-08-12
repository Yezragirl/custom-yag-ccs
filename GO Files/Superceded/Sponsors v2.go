
{{$ap := }}
{{$jm := }}
{{$mc := }}
{{$as := }}
{{$ep := }}
{{$lg := }}
{{$al := }}
{{$capture := }}


{{$color := randInt 111111 999999 }}
{{$embed := cembed 
"color" $color
"Title" "Our Sponsors"
"Description" "The following is a list of our Current Sponsors!"
"fields" (cslice 
(sdict "name" "__**Apprentice Sponsors**__" "value" $ap "inline" false)
(sdict "name" "__**Journeyman Sponsors**__" "value" $jm "inline" false)
(sdict "name" "__**Mastercraft Sponsors**__" "value" $mc "inline" false)
(sdict "name" "__**Ascendant Sponsors**__" "value" $as "inline" false)
(sdict "name" "__**Epic Sponsors**__" "value" $ep "inline" false)
(sdict "name" "__**Legendary Sponsors**__" "value" $lg "inline" false)
(sdict "name" "__**Alpha Sponsors**__" "value" $al "inline" false)) }}

{{if $capture}}
	{{editMessage 573896118719610880 $capture $embed}}
{{else}}
{{$capture = sendMessageNoEscapeRetID 573896118719610880 $embed}}
{{end}}