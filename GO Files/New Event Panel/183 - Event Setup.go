{{/*Event System Setup - Command eventsetup*/}}
{{$ReminderOne := toDuration (dbGet 0 "EventReminderOne").Value}}
{{$ReminderTwo := toDuration (dbGet 0 "EventReminderTwo").Value}}
{{$DMReminder := toDuration (dbGet 0 "EventDMReminder").Value}}
{{$Runners := toInt (dbGet 0 "EventRunners").Value}}

{{if not $ReminderOne}}
{{$ReminderOne = toDuration ("6h")}}
{{dbSet 0 "EventReminderOne" $ReminderOne}}
{{end}}

{{if not $ReminderTwo}}
{{$ReminderTwo = toDuration ("3h")}}
{{dbSet 0 "EventReminderTwo" $ReminderTwo}}
{{end}}

{{if not $DMReminder}}
{{$DMReminder =  toDuration ("1h")}}
{{dbSet 0 "EventDMReminder" $DMReminder}}
{{end}}

{{if not $Runners}}
{{$Runners = 3}}
{{dbSet 0 "EventRunners" $Runners}}
{{end}}

{{$embed := cembed 
"title" "Events Setup Menu" 
"description" "**This Event System is using the following settings.**" 
"color" 4645612 
"fields" (cslice 
       (sdict "name" "**Time of First Reminder before Event**" "value" (toString $ReminderOne) "inline" false) 
       (sdict "name" "**Time of Second Reminder before Event**" "value" (toString $ReminderTwo) "inline" false) 
       (sdict "name" "**Time of DM Reminder before Event**" "value" (toString $DMReminder) "inline" false)
       (sdict "name" "**Max Number of Event Runners Allowed**" "value" (toString $Runners) "inline" false) ) 
"footer" (sdict "text" "To change any of these settings, use -set [setting] [new value].")}}
{{sendMessage nil $embed}}