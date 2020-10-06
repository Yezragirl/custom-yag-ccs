To Give Someone Money
~add-money cash [member] [amount]

To Take Someone's Money
~remove-money bank [member] [amount]

To Give Someone a Coupon
~give-item [member] [quantity] [Item name (use " " if there are spaces in it] 

To Take Someone's Coupon
~take-item [member] [quantity] [Item name (use " " if there are spaces in it] 

Check Someone's Inventory 
~inventory [member]


{{$color := randInt 111111 999999 }}
 {{$embed := cembed 
 "color" $color
"Title" "Common Admin Commands"
"Description" "This commands can only be used by Admins."
 "fields" (cslice 
 (sdict "name" "To Give Someone Money" "value" "~add-money cash [member] [amount]" "inline" false) 
 (sdict "name" "To Take Someone's Money" "value" "~remove-money bank [member] [amount]" "inline" false) 
 (sdict "name" "To Give Someone a Coupon" "value" "~give-item [member] [quantity] [Item name (use " " if there are spaces in it]" "inline" false) 
 (sdict "name" "To Take Someone's Coupon" "value" "~take-item [member] [quantity] [Item name (use " " if there are spaces in it]" "inline" false) 
 (sdict "name" "Check Someone's Inventory " "value" "~inv [member]" "inline" false) 
 (sdict "name" "__**Ellie**__" "value" "*Blue and Royalty* - Dire Bear" "inline" false) 
 (sdict "name" "__**Blythe**__" "value" "*Black and Purple* - Griffin" "inline" false)) }}
{{sendMessage nil $embed}}