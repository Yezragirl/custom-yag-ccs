{{if eq currentTime.Weekday.String "Monday"}}
{{execCC 42 nil 0 "-giveaway start 23h Monday Madness! -w 1 -c <#589437097602056203>"}}
{{else if eq currentTime.Weekday.String "Tuesday"}}
{{execCC 42 nil 0 "-giveaway start 23h Terrific Tuesday! -w 3 -c <#589437097602056203>"}}
{{else if eq currentTime.Weekday.String "Wednesday"}}
{{execCC 42 nil 0 "-giveaway start 23h Wonderful Wednesday! -w 5 -c <#589437097602056203>"}}
{{else if eq currentTime.Weekday.String "Thursday"}}
{{execCC 42 nil 0 "-giveaway start 23h Fabulous Thursday! -w 2 -c <#589437097602056203>"}}
{{else if eq currentTime.Weekday.String "Friday"}}
{{execCC 42 nil 0 "-giveaway start 23h Freaky Friday! -w 1 -c <#589437097602056203>"}}
{{else if eq currentTime.Weekday.String "Saturday"}}
{{execCC 42 nil 0 "-giveaway start 23h Silly Saturday! -w 4 -c <#589437097602056203>"}}
{{else if eq currentTime.Weekday.String "Sunday"}}
{{execCC 42 nil 0 "-giveaway start 23h Sunday Funday! -w 3 -c <#589437097602056203>"}}
{{end}}