{{/*starter gives starter info to new players*/}}
{{$args := parseArgs 3
	"syntax `-starter Locker Pin @mention`"
(carg "string" "Locker ID")
(carg "string" "pin")
(carg "user" "user")}}
{{$locker := ($args.Get 0)}}
{{$pin := ($args.Get 1)}}
{{$who := ($args.Get 2)}}

{{sendMessage nil "The Cluster password can be found on <#613585635026141185>.\n"}}
{{sendMessage nil (joinStr "" "Please log into our Starter Map, __***Yez's Refuge***__ to collect your starter. **This map is intended as a temporary map for those who want to get a few levels before joining the standard cluster maps,  or new players who are not yet familiar with the games mechanics. If you don't feel like you need to stay, please move on. You can NOT transfer IN to Refuge, only out. Be warned, the dinos on our standard maps are max level 450, with increased damage and resistance. Be prepared before you transfer.**\n\nThe Community Center there can be found along the southern coast at approximately 85, 63. It can be reached by spawning at South 2 or 3, and running along the beach. Check your map. Room " $locker ", Pin " $pin ".")}}
{{sendMessage nil (joinStr "" "Welcome to the cluster!\n\n" $who.Mention ", you'll find your starter info above.\n\nThis channel will be automatically closed after 7 days. If you can not pick up your starter within those 7 days, you can open a general ticket when you are ready and a new starter will be provided.\n")}}
{{$msg := sendMessageRetID nil "Please hit the <:gcrystal:662677298431918092> below when you've collected your items so that we can restock the room for the next new player."}}
{{addMessageReactions nil $msg ":gcrystal:662677298431918092" }}
{{deleteTrigger 1}}
{{sendMessage 653058600230584353 (joinStr "" (getMember .User.ID).Nick " gave " (getMember $who).Nick " their starter.")}}