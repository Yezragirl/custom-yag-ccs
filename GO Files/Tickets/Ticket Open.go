{{$days := 0}}{{$ticketnum := dbIncr 0 "ticketnum" 1}}{{$Q := (dbGet .User.ID "Q").Value}}{{$admin := (dbGet .User.ID "admin").Value}}{{$age := (dbGet .User.ID "age").Value}}{{$gt := (dbGet .User.ID "gt").Value}}{{$name := (dbGet .User.ID "name").Value}}{{$pin := (dbGet .User.ID "pin").Value}}{{$public := (dbGet .User.ID "public").Value}}{{$parentguardian := (dbGet .User.ID "parentguardian").Value}}{{$ref := (dbGet .User.ID "ref").Value}}{{$tribe := (dbGet .User.ID "tribe").Value}}{{$ticketnum := (dbGet 0 "ticketnum").Value}}{{$names_db := (dbGet 0 "names").Value}}{{$names := (cslice).AppendSlice $names_db}}{{$names_new := cslice}}{{$message := ""}}{{$days := 0}}
{{if eq .Reason "vault"}}{{$message = ", thank you for creating a ticket!\nSomeone from <@&634598489732546588> will be set up your quest vault shortly.\nQuests are located in the quest building at the Community Center on Manticore, in the large field behind Green Ob. Someone will close this ticket once your quest vault is set up. If you have any questions about quests, feel free to ask them here."}}
{{else if eq .Reason "mart"}}{{$message = ", welcome to the ArkMart!\nThank you for creating a ticket!\nSomeone from <@&634598489732546588> will be with you shortly.\nTo help facilitate your assistance, please answer the following questions:\n:one: ***What would you like from the ArkMart?***\n:two: ***What Map would you like to pick up your purchase on?***\n:three: ***How will you be paying? BBS or coupon?***"}}
{{else if eq .Reason "auction"}}{{$message = ", thank you for creating a ticket!\nSomeone from <@&634598489732546588> will be with you shortly.\nTo help facilitate your assistance, please answer the following questions:\n:one: ***Is this regarding the Live Auction, or the Discord Auction?***\n:two: ***If it's regarding the Live Auction, is it about a Donation, or a Proxy Bid?***"}}
{{else if eq .Reason "base"}}{{$message = ", thank you for creating a ticket!\nSomeone from <@&634598489732546588> will be with you shortly.\nTo help facilitate your assistance, please answer the following questions:\n:one: ***What map is your base issue on?***\n:two: ***Please provide coordinates for the issue.***\n:three: ***Please describe the issue.***"}}
{{else if eq .Reason "dino"}}{{$message = ", thank you for creating a ticket!\nSomeone from <@&634598489732546588> will be with you shortly.\nTo help facilitate your assistance, please answer the following questions:\n:one: ***What Map is your Dino Issue on?***\n:two: ***Please provide the coordinates where the issue is.***\n:three: ***Please describe the issue.***"}}
{{else if eq .Reason "discord"}}{{$message = ", thank you for creating a ticket!\nSomeone from <@&634598489732546588> will be with you shortly.\nTo help facilitate your assistance, please answer the following questions:\n:one: ***Is this a problem with a person, a channel, or a bot?***\n:two: ***Please describe the problem.***"}}
{{else if eq .Reason "quest"}}{{$message = ", thank you for creating a ticket!\nSomeone from <@&634598489732546588> will be with you shortly.\n\n***Is this regarding New Player Quests, or Weekly Quests? ***"}}
{{else if eq .Reason "trap"}}{{$message = ", thank you for creating a ticket!\nSomeone from <@&634598489732546588> will be with you shortly.\nTo help facilitate your assistance, please answer the following questions:\n:one: ***What Map is this trap on?***\n:two: ***Please provide coordinates of the trap.***"}}
{{else if eq .Reason "tribe"}}{{$message = ", thank you for creating a ticket!\nSomeone from <@&634598489732546588> will be with you shortly.\nTo help facilitate your assistance, please answer the following questions:\n:one: ***What map is your tribe issue on?***\n:two: ***Please provide coordinates for the issue.***\n:three: ***Please describe the issue.***"}}
{{else if eq .Reason "regedit"}}{{$message = ", this ticket is for you to edit your registration information, privately, where no other server members can see. The only people who have access to this channel are you, and the admin team. If you need assistance, please ping <@&634598489732546588>. They will be happy to help you. Otherwise, you can use '-RegEdit' to begin the process of editing your registration information."}}
{{else if eq .Reason "general"}}{{$message = ", thank you for creating a ticket!\nSomeone from <@&634598489732546588> will be with you shortly. \n\n***How may we help you today?***"}}
{{else if eq .Reason "hunt"}}{{$message = ", so you found a flag? Great! Please tell us the ***Map*** and the ***Coordinates*** of the flag. Once we verify that it is indeed a Scavenger Hunt Flag location, your prize will be dispersed."}}
{{else if eq .Reason "hardcore"}}{{$message = ", this is a Hardcore ticket. How can we assist you?"}}
{{else if eq .Reason "smart"}}{{$message = ", welcome to the Sponsor Mart!\nThank you for creating a ticket!\nSomeone from <@&634598489732546588> will be with you shortly.\nTo help facilitate your assistance, please answer the following questions:\n:one: ***What would you like from the Sponsor Mart?***\n:two: ***What Map would you like to pick up your purchase on?***\n:three: ***How will you be paying? BBS or coupon?***\n**Please note that Prices in the Sponsor Mart are NOT discounted.**"}}
{{else if eq .Reason "map"}}{{$message = ", welcome to a Private Map ticket. \nTo begin, please fill out the form at https://docs.google.com/forms/d/e/1FAIpQLSeRkcJdzLWw_MXUYqikDtCMTcdY9hlgeFvMov5ep70Me00RSg/viewform if this is a new map setup. If its just a change, please let the changes you'd like to make below and someone from <@&634598489732546588> will be with you shortly."}}
{{else if eq .Reason "weekly-timer-reset"}}{{$message = ", it's timer reset time! Please react with your map so that I can close the ticket whenever every map is done. Deadline is Sunday."}}
{{else if eq .Reason "violation"}}{{$message = ", this is a violation ticket."}}
{{else if eq .Reason "project"}}{{$message = ", this is a project ticket."}}
{{else if eq .Reason "claimed-dinos"}}{{$message = ", this is a claimed dinos ticket."}}
{{else if eq .Reason "player-issues"}}{{$message = ", this is a player issues ticket."}}
{{else if eq .Reason "training"}}{{$message = ", this is a training ticket."}}
{{else if eq .Reason "Map Sweeps"}}{{$message = ", it's map sweep time! Please react with your map so that I can close the ticket whenever every map is done."}}
{{else if eq .Reason "restock and repin"}}{{$message = ", it's that time again! All community center free items need to be restocked, all turrets and generators checked. All rooms checked and repinned."}}
{{else if eq .Reason "registration"}}
{{if not $name}}{{sendMessage .Channel.ID (joinStr "" "Thanks for joining Yez's Ark Cluster!\n\nJust a few questions, and you'll be registered and ready to play!\nInstructions first. **If emojis appear below the question, please use them to respond. Otherwise, type your answers.** These answers will be editable later, but be sure to check for typos before hitting enter, as your answers affect the direction these questions go.")}}
{{$Q = "Age"}}{{dbSet .User.ID "Q" $Q}}
{{removeRoleID 645404993863811072}}{{removeRoleID 573210479753691170}}{{removeRoleID 607041870966685706}}{{giveRoleID .User.ID 606566506922377257}}
{{sendMessage .Channel.ID (joinStr "" .User.Mention ", how old are you?")}}
{{else}}{{range $names}}{{if ne . $name}}{{$names_new = $names_new.Append .}}{{end}}{{end}}
{{dbSet 0 "names" $names_new}}
{{removeRoleID 645404993863811072}}{{removeRoleID 573210479753691170}}{{removeRoleID 607041870966685706}}{{giveRoleID .User.ID 606566506922377257}}
{{sendMessage .Channel.ID (joinStr "" "Thanks for joining Yez's Ark Cluster!\n\nJust a few questions, and you'll be registered and ready to play!\nInstructions first. **If emojis appear below the question, please use them to respond. Otherwise, type your answers.** These answers will be editable later, but be sure to check for typos before hitting enter, as your answers affect the direction these questions go.\n_ _")}}
{{sendMessage .Channel.ID (joinStr "" "Please be advised that by starting a new registration, any previous registration has been erased. You must complete THIS registration.")}}
{{if $admin}}{{deleteMessage 658354135221141533 $admin}}{{end}}
{{if $public}}{{deleteMessage 658352884701724733 $public}}{{end}}
{{$Q = "Age"}}{{dbSet .User.ID "Q" $Q}}
{{sendMessage .Channel.ID (joinStr "" .User.Mention ", how old are you?")}}
{{end}}

{{end}}
{{if eq .Reason "registration"}}
{{else if or (eq .Reason "Map Sweeps") (eq .Reason "restock and repin")}}
{{$second := sendMessageNoEscapeRetID .Channel.ID (joinStr "" "Mark off your maps here as you do them. When all maps are done, this ticket will close automagically.")}}
{{addMessageReactions .Channel.ID $second ":Haven:796125877633024051" ":Medusa:640697225684713473" ":Sanctuary:797519029715599400" ":Artemis:640697096814592040" ":Hades:797518674616909914" ":Zelda:640697492274544643" ":Manticore:640697193052766218" ":Osiris:748895819331141674" ":Elysian:796160062904729600" ":Everest:811680105218179083" ":Merlin:796116046616592445" ":Beowulf:796127204164632638" ":Raiden:655954199476961300" ":Jupiter:795749721657573446" ":Titan:640697323143561276" ":Gryphon:640697148807184384" ":Phoenix:640698407199178771" ":Fenris:797515621222318091" ":Excelsior:797519872115343360" ":Refuge:797670477081608202"}}
{{else}}
{{$messageID := sendMessageNoEscapeRetID .Channel.ID (joinStr "" .User.Mention $message)}}
{{$daycounter := sendMessageNoEscapeRetID .Channel.ID "\n\nThis ticket was opened today."}}
{{execCC 220 .Channel.ID 86400 (sdict "days" $days "daycounter" $daycounter)}}
{{addMessageReactions .Channel.ID $messageID ":kibble:660605488609755146" ":cryopod:631218688099614724" ":Tek_Door:660605465721569334" }}
{{end}}
{{$ticketnum = add $ticketnum 1}}
{{dbSet 0 "ticketnum" $ticketnum}}