{{$capture := ""}}
{{$ID := ""}}
{{$message := ""}}
{{$type := ""}}
{{$panel := (toString .ReactionMessage.ID)}}
{{$emoji := ""}}

{{if eq .Reaction.Emoji.Name "1️⃣"}}
{{$type = "mart"}}
{{$message = ", welcome to the ArkMart!\nThank you for creating a ticket!\nSomeone from <@&634598489732546588> will be with you shortly.\nTo help facilitate your assistance, please answer the following questions:\n:one: ***What would you like from the ArkMart?***\n:two: ***What Map would you like to pick up your purchase on?***\n:three: ***How will you be paying? BBS or coupon?***"}}
{{$emoji = "1️⃣"}}

{{else if eq .Reaction.Emoji.Name "2️⃣"}}
{{$type = "auction"}}
{{$message = ", thank you for creating a ticket!\nSomeone from <@&634598489732546588> will be with you shortly.\nTo help facilitate your assistance, please answer the following questions:\n:one: ***Is this regarding the Live Auction, or the Discord Auction?***\n:two: ***If it's regarding the Live Auction, is it about a Donation, or a Proxy Bid?***"}}
{{$emoji = "2️⃣"}}

{{else if eq .Reaction.Emoji.Name "3️⃣"}}
{{$type = "base"}}
{{$message = ", thank you for creating a ticket!\nSomeone from <@&634598489732546588> will be with you shortly.\nTo help facilitate your assistance, please answer the following questions:\n:one: ***What map is your base issue on?***\n:two: ***Please provide coordinates for the issue.***\n:three: ***Please describe the issue.***"}}
{{$emoji = "3️⃣"}}

{{else if eq .Reaction.Emoji.Name "4️⃣"}}
{{$type = "dino"}}
{{$message = ", thank you for creating a ticket!\nSomeone from <@&634598489732546588> will be with you shortly.\nTo help facilitate your assistance, please answer the following questions:\n:one: ***What Map is your Dino Issue on?***\n:two: ***Please provide the coordinates where the issue is.***\n:three: ***Please describe the issue.***"}}
{{$emoji = "4️⃣"}}

{{else if eq .Reaction.Emoji.Name "5️⃣"}}
{{$type = "discord"}}
{{$message = ", thank you for creating a ticket!\nSomeone from <@&634598489732546588> will be with you shortly.\nTo help facilitate your assistance, please answer the following questions:\n:one: ***Is this a problem with a person, a channel, or a bot?***\n:two: ***Please describe the problem.***"}}
{{$emoji = "5️⃣"}}

{{else if eq .Reaction.Emoji.Name "6️⃣"}}
{{$type = "quest"}}
{{$message = ", thank you for creating a ticket!\nSomeone from <@&634598489732546588> will be with you shortly.\n\n***Is this regarding New Player Quests, or Weekly Quests? ***"}}
{{$emoji = "6️⃣"}}

{{else if eq .Reaction.Emoji.Name "7️⃣"}}
{{$type = "trap"}}
{{$message = ", thank you for creating a ticket!\nSomeone from <@&634598489732546588> will be with you shortly.\nTo help facilitate your assistance, please answer the following questions:\n:one: ***What Map is this trap on?***\n:two: ***Please provide coordinates of the trap.***"}}
{{$emoji = "7️⃣"}}

{{else if eq .Reaction.Emoji.Name "8️⃣"}}
{{$type = "tribe"}}
{{$message = ", thank you for creating a ticket!\nSomeone from <@&634598489732546588> will be with you shortly.\nTo help facilitate your assistance, please answer the following questions:\n:one: ***What map is your tribe issue on?***\n:two: ***Please provide coordinates for the issue.***\n:three: ***Please describe the issue.***"}}
{{$emoji = "8️⃣"}}

{{else if eq .Reaction.Emoji.Name "9️⃣"}}
{{$type = "regedit"}}
{{$message = ", this ticket is for you to edit your registration information, privately, where no other server members can see. The only people who have access to this channel are you, and the admin team. If you need assistance, please ping <@&634598489732546588>. They will be happy to help you. Otherwise, you can use '-RegEdit' to begin the process of editing your registration information."}}
{{$emoji = "9️⃣"}}

{{else if eq .Reaction.Emoji.Name "🔟"}}
{{$type = "general"}}
{{$message = ", thank you for creating a ticket!\nSomeone from <@&634598489732546588> will be with you shortly. \n\n***How may we help you today?***"}}
{{$emoji = "🔟"}}

{{end}}
{{$capture = exec "ticket open" $type}}
{{$ID = reFind `\d+` (reFind `<#\d+>` $capture)}}
{{$ticket := sendMessageNoEscapeRetID $ID (joinStr "" .User.Mention $message)}}
{{deleteAllMessageReactions nil $panel}}
{{addMessageReactions nil $panel $emoji}}
{{addMessageReactions $ID $ticket ":Tek_Door:660605465721569334" ":kibble:660605488609755146" ":cryopod:631218688099614724"}}