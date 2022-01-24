{{/*Player and Clock Code*/}}
{{$Users := (dbGet 0 "newusers").Value}} 
{{editChannelName "631550323332218910" (joinStr "" (( currentTime.UTC.Add ( toDuration ( mult -5 .TimeHour ) ) ).Format "3:04PM") " EDT")}}
{{editChannelName "673190391234691094" (joinStr "" "Members: " .Guild.MemberCount)}}
{{editChannelName "673298791390249064" (joinStr "" "Online: " onlineCount)}}
{{editChannelName "797652063529598986" (joinStr "" $Users " New Players (1/9/21)")}}
