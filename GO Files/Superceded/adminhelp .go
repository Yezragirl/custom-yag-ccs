{{$color := randInt 111111 999999 }}
 {{$embed := cembed 
 "color" $color
"Title" "So You Want to Be an Admin..."
"Description" "This help file details how to become an admin, and the admin teams general philosophy."
 "fields" (cslice 
 (sdict "name" "__**Measurable Requirements**__" "value" "\n**Rep Points**\nA minimum of 5 rep points are required to be considered for admin. Rep points prove that you are helpful since they are given when people thank you.\n\n**Violations**\nYou must have no outstanding violations to become an admin. If you have any violations, discuss them with an admin in a ticket to see what can be done to clear them.\n\n**Discord Level**\nTo make sure that our admins aren't just random people, and to ensure they are familiar with the server, and the discord, we require at least Discord level 15. Once you reach level 15, the <#626884431163949056> channel will be available to you." "inline" false) 
 (sdict "name" "__**Non-Measurable Requirements**__" "value" "Clearly some knowledge/personality traits are required to make a good admin. While we can teach anyone how to use admin commands, and work with discord bots, the things we can't teach are:\nFriendliness\nHelpfulness\nPatience\nTolerance" "inline" false) 
 (sdict "name" "__**Team and Philosophy**__" "value" "The Admins on this server aren't just a collection of individuals, we are a family. Entering into that family isn't something you should do lightly. It comes with responsibilities, as well as advantages, but none of it is without expectations." "inline" false) 
 (sdict "name" "__**Responsibilities**__" "value" "Some people may not be aware of the requirements of admining. So here is a partial list of the things that admins are responsible for:\nAnswering player questions with politeness and knowledge\nHandling Support Tickets (in a timely manner)\nRunning/Assisting with Events\nMap Decay Timer Resets\nMap Sweeps for Decayed Bases and Traps\nSpecial Projects" "inline" false) 
 (sdict "name" "__**Commitment**__" "value" "Admins work hard, but they also play. We don't work our admins to the bone. Below are the approximate average time commitments needed:\nProvisional Admin: 10-12 Hours a week\nThis includes training, and some time consuming up front tasks we ask all provisionals to do.\n\nFull Admin: 8-10 Hours a week" "inline" false) 
) }}
{{sendMessage nil $embed}}