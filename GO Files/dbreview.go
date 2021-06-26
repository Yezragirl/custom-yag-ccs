{{$skip := 0 -}}
{{if not .ExecData -}}
	{{$skip = 10}}
{{else -}}
	{{$skip = .ExecData.skip}}
{{end -}}
    {{- range dbTopEntries "bday" 10 $skip -}}
    	{{$who := (userArg .UserID)}}
    	{{$who.username}} - {{.}}
    {{end -}}
  {{if gt (len .ExecData.skip) 0 -}}
    {{execCC .CCID nil 25 (sdict "skip" (add $skip 10)) -}}
  {{end -}}
