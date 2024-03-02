{{/*shop generates the shops rating embed for new player shops*/}}
{{$color := randInt 111111 999999 }}
	{{$embed := cembed 
	"Title" "Player Shop Rating"
	"color" $color
	"Description" "Please use the reactions below to describe your experience with this shop.\n🥇 for Great Service\n🥈 for Decent Service\n🥉 for Ok Service\n👎 for Bad Service"
	"footer" (sdict "text" "Yez's Ark Cluster")}}
{{$msg := sendMessageRetID nil $embed}}
{{addMessageReactions nil $msg "🥇" "🥈" "🥉" "👎"}}

{{deleteTrigger 1}}