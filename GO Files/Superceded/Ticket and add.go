```go
{{$capture := exec "ticket open" "Reason"}}
{{$ID := reFind `\d+` (reFind `<#\d+>` $capture)}}
```
Gets you the ID of the channel of the Ticket

```go
{{$args := parseArgs 2 "correct syntax is `adduser @who #channel`"
     (carg "user" "user")
     (carg "channel" "channel")}}
{{execCC 101 ($args.Get 1).ID 0 (sdict "user" ($args.Get 0))}}```

{{execAdmin "tickets adduser" (.ExecData.user).ID}}