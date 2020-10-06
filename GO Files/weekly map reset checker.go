
{{$Manticore := 0}}
{{$Artemis := 0}}
{{$Gryphon := 0}}
{{$Zelda := 0}}
{{$Phoenix := 0}}
{{$Elysian := 0}}
{{$Titan := 0}}
{{$Raiden:= 0}}
{{$Medusa := 0}}
{{$Beowulf := 0}}
{{$count := 2}}

{{range .ReactionMessage.Reactions}}

{{if eq (toString .Emoji.ID) (toString 640697193052766218)}}
{{if eq .Count $count}}

{{$Manticore = 1}}
{{end}}
{{else if eq (toString .Emoji.ID) (toString 640697096814592040)}}
{{if eq .Count $count}}

{{$Artemis = 1}}
{{end}}
{{else if eq (toString .Emoji.ID) (toString 640697148807184384)}}
{{if eq .Count $count}}

{{$Gryphon = 1}}
{{end}}
{{else if eq (toString .Emoji.ID) (toString 640697492274544643)}}
{{if eq .Count $count}}

{{$Zelda = 1}}
{{end}}
{{else if eq (toString .Emoji.ID) (toString 640698407199178771)}}
{{if eq .Count $count}}

{{$Phoenix = 1}}
{{end}}
{{else if eq (toString .Emoji.ID) (toString 640698387506921523)}}
{{if eq .Count $count}}

{{$Elysian = 1}}
{{end}}
{{else if eq (toString .Emoji.ID) (toString 640697323143561276)}}
{{if eq .Count $count}}

{{$Titan = 1}}
{{end}}
{{else if eq (toString .Emoji.ID) (toString 655954199476961300)}}
{{if eq .Count $count}}

{{$Raiden = 1}}
{{end}}
{{else if eq (toString .Emoji.ID) (toString 640697225684713473)}}
{{if eq .Count $count}}

{{$Medusa = 1}}
{{end}}
{{else if eq (toString .Emoji.ID) (toString 682592075685953570)}}
{{if eq .Count $count}}

{{$Beowulf = 1}}
{{end}}{{end}}{{end}}


{{if and (eq 1 $Manticore) (eq 1 $Artemis) (eq 1 $Gryphon) (eq 1 $Zelda) (eq 1 $Phoenix) (eq 1 $Elysian) (eq 1 $Titan) (eq 1 $Raiden) (eq 1 $Medusa) (eq 1 $Beowulf)}}

{{exec "tickets close" "All Weekly Map Resets Complete"}}
{{end}}



{{$channel := (getChannel nil).Name}}
{{if reFind `weekly-timer-reset$` $channel}}
{{$reactions := cslice 640697193052766218 640697096814592040 640697148807184384 640697492274544643 640698407199178771 640698387506921523 640697323143561276 655954199476961300 640697225684713473 682592075685953570}}
{{$count := 0}}{{$allMatch := true}}

{{range .ReactionMessage.Reactions}}
  {{if and (in $reactions .Emoji.ID) .Me (eq .Count 2)}}
      {{$count = add $count 1}}
  {{else}}
      {{$allMatch = false}}
  {{end}}
{{end}}


{{if and $allMatch (eq $count 10)}}

{{exec "tickets close" "All Weekly Map Resets Complete"}}
{{end}}
{{end}}
