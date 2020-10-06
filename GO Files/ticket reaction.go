{{$capture := ""}}
{{$ID := ""}}
{{$message := ""}}
{{$type := ""}}
{{$panel := (toString .ReactionMessage.ID)}}
{{$emoji := ""}}

{{if eq .Reaction.Emoji.Name ":one:"}}
{{$type = "mart"}}
{{$message = ", welcome to the ArkMart!\nThank you for creating a ticket!\nSomeone from @Support Team will be with you shortly.\nTo help facilitate your assistance, please answer the following questions:\n:one: What would you like from the ArkMart?\n:two: What Map would you like to pick up your purchase on?\n:three: How will you be paying? BBS or coupon?"}}
{{$emoji = ":one:"}}

{{else if eq .Reaction.Emoji.Name ":two:"}}
{{$type = "auction"}}
{{$message = ", thank you for creating a ticket!\nSomeone from @Support Team will be with you shortly.\nTo help facilitate your assistance, please answer the following questions:\n:one: Is this regarding the Live Auction, or the Discord Auction?\n:two: If it's regarding the Live Auction, is it about a Donation, or a Proxy Bid?"}}
{{$emoji = ":two:"}}

{{else if eq .Reaction.Emoji.Name ":three:"}}
{{$type = "base"}}
{{$message = ", thank you for creating a ticket!\nSomeone from @Support Team will be with you shortly.\nTo help facilitate your assistance, please answer the following questions:\n:one: What map is your base issue on?\n:two: Please provide coordinates for the issue.\n:three: Please describe the issue."}}
{{$emoji = ":three:"}}

{{else if eq .Reaction.Emoji.Name ":four:"}}
{{$type = "dino"}}
{{$message = ", thank you for creating a ticket!\nSomeone from @Support Team will be with you shortly.\nTo help facilitate your assistance, please answer the following questions:\n:one: What Map is your Dino Issue on?\n:two: Please provide the coordinates where the issue is.\n:three: Please describe the issue."}}
{{$emoji = ":four:"}}

{{else if eq .Reaction.Emoji.Name ":five:"}}
{{$type = "discord"}}
{{$message = ", thank you for creating a ticket!\nSomeone from @Support Team will be with you shortly.\nTo help facilitate your assistance, please answer the following questions:\n:one: Is this a problem with a person, a channel, or a bot?\n:two: Please describe the problem."}}
{{$emoji = ":five:"}}

{{else if eq .Reaction.Emoji.Name ":six:"}}
{{$type = "quest"}}
{{$message = ", thank you for creating a ticket!\nSomeone from @Support Team will be with you shortly.\n\nIs this regarding New Player Quests, or Weekly Quests? "}}
{{$emoji = ":six:"}}

{{else if eq .Reaction.Emoji.Name ":seven:"}}
{{$type = "trap"}}
{{$message = ", thank you for creating a ticket!\nSomeone from @Support Team will be with you shortly.\nTo help facilitate your assistance, please answer the following questions:\n:one: What Map is this trap on?\n:two: Please provide coordinates of the trap."}}
{{$emoji = ":seven:"}}

{{else if eq .Reaction.Emoji.Name ":eight:"}}
{{$type = "tribe"}}
{{$message = ", thank you for creating a ticket!\nSomeone from @Support Team will be with you shortly.\nTo help facilitate your assistance, please answer the following questions:\n:one: What map is your tribe issue on?\n:two: Please provide coordinates for the issue.\n:three: Please describe the issue."}}
{{$emoji = ":eight:"}}

{{else if eq .Reaction.Emoji.Name ":nine:"}}
{{$type = "regedit"}}
{{$message = ", this ticket is for you to edit your registration information, privately, where no other server members can see. The only people who have access to this channel are you, and the admin team. If you need assistance, please ping @Support Team. They will be happy to help you. Otherwise, you can use '-RegEdit' to begin the process of editing your registration information."}}
{{$emoji = ":nine:"}}

{{else if eq .Reaction.Emoji.Name ":keycap_ten:"}}
{{$type = "general"}}
{{$message = ", thank you for creating a ticket!\nSomeone from @Support Team will be with you shortly. \n\nHow may we help you today?"}}
{{$emoji = ":keycap_ten:"}}

{{end}}
{{$capture = exec "ticket open" $type}}
{{$ID = reFind \d+ (reFind <#\d+> $capture)}}
{{$ticket := sendMessageNoEscapeRetID $ID (joinStr "" .User.Mention $message)}}
{{deleteAllMessageReactions nil $panel}}
{{addMessageReactions nil $panel $emoji}}

{{addMessageReactions $ID $ticket ":kibble:660605488609755146" ":cryopod:631218688099614724" ":Tek_Door:660605465721569334" }}