{{$role := .ExecData.role}}
{{$event := .ExecData.event}}
{{sendDM (joinStr "" "This is your __***one hour***__ reminder, you are scheduled as a *" $role "* for **" $event "**. Please check the #Events-List to see what map this event is on." )}}