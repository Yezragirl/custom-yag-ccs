{{$days := .ExecData.days}}{{$daycounter := .ExecData.daycounter}}
{{$days = add $days 1}}
{{editMessage nil $daycounter (joinStr "" "This ticket was opened " $days " ago.")}}
{{execCC 220 nil 86400 (sdict "days" $days "daycounter" $daycounter)}}
{{if ge $days 7}}
{{sendMessage 573897276569813012 (joinStr "" "<@&634598489732546588>, ".Channel.Name " has been open for " $days " days.")}}