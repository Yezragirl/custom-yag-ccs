{{/*hcrules posts the rules for HC as embeds*/}}
{{$embed := cembed 
	"title" "Hardcore Character Level Rewards" 
	"description" "*When you reach these character levels, these are the rewards. Each level can only be redeemed once per player.*" 
	"color" 469732 
	"fields" (cslice 
		   (sdict "name" "**Level 25**" "value" "250 <:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Level 50**" "value" "500 <:bbs:583143695927214091>\nChoice of Thylo or Ravager" "inline" true) 
		   (sdict "name" "**Level 75**" "value" "750 <:bbs:583143695927214091>" "inline" true)
		   (sdict "name" "**Level 100**" "value" "1,000 <:bbs:583143695927214091>\nChoice of Theri or Bear" "inline" true)
		   (sdict "name" "**Level 125**" "value" "1,250 <:bbs:583143695927214091>" "inline" true)
		   (sdict "name" "**Level 150**" "value" "1,500 <:bbs:583143695927214091>\nChoice of Rex or Allo" "inline" true)
		   (sdict "name" "**Level 175**" "value" "1,750 <:bbs:583143695927214091>" "inline" true)
		   (sdict "name" "**Level 200**" "value" "2,000 <:bbs:583143695927214091>\nChoice of Wyvern or Rock Drake" "inline" true) 
		   (sdict "name" "**Level 225**" "value" "2,250 <:bbs:583143695927214091>" "inline" true)
		   (sdict "name" "**Level 250**" "value" "2,500 <:bbs:583143695927214091>\nChoice of Golem or Mosa" "inline" true)) }}
	{{sendMessage nil $embed}}
	
	{{$embed2 := cembed 
	"title" "Hardcore First Tame Rewards" 
	"description" "*When you turn in your first tame of each of these dinos, these are the rewards. Each of these can only be turned in once per player.*" 
	"color" 469732 
	"fields" (cslice        (sdict "name" "**Anglerfish**" "value" "400<:bbs:583143695927214091>" "inline" true)
		   (sdict "name" "**Bulbdog**" "value" "250<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Drake Egg over Level 250**" "value" "500<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Electrophorus**" "value" "400<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Featherlight**" "value" "250<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Giga**" "value" "1,000<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Glowtail**" "value" "250<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Golem**" "value" "1,000<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Hesperonis**" "value" "800<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Karkanos**" "value" "800<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Liopleurodon**" "value" "2,500<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Mantis**" "value" "400<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Megalania**" "value" "400<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Megalosaurus**" "value" "600<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Mosasaur**" "value" "1,000<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Onyx**" "value" "400<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Ovis**" "value" "300<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Plesiosaur**" "value" "600<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Purlovia**" "value" "400<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Quetz**" "value" "600<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Ravager**" "value" "400<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Rex**" "value" "400<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Shinehorn**" "value" "250<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Snow Owl**" "value" "500<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Spino**" "value" "600<:bbs:583143695927214091>" "inline" true)  ) }}
	{{sendMessage nil $embed2}} 
	
	
	{{$embed3 := cembed 
	"title" "Hardcore First Tame Rewards - Continued" 
	"description" "*When you turn in your first tame of each of these dinos, these are the rewards. Each of these can only be turned in once per player.*" 
	"color" 469732 
	"fields" (cslice    
		   (sdict "name" "**Thylo**" "value" "500<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Titanoboa**" "value" "600<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Troodon**" "value" "500<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Trope**" "value" "250<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Tuso**" "value" "1,000<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Unicorn**" "value" "500<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Wooly Rhino**" "value" "400<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Wyvern Egg Over Level 250**" "value" "500<:bbs:583143695927214091>" "inline" true) 
		   (sdict "name" "**Yutyrannus**" "value" "600<:bbs:583143695927214091>" "inline" true) ) }}
	{{sendMessage nil $embed3}} 
	