{{$timeString := ""}}
{{if .ExecData}}
	{{$timeString = .ExecData.time}}
	{{else}}
{{/*String to be converted is fed to $timeString variable*/}}
{{$timeString = .StrippedMsg}}
{{/* can be any string variable / string data but should contain only the time in the formats mentioned */}}
{{end}}

{{/* Variable Declarations */}}
{{$dTime := sdict "Day" currentTime.Day "Month" (toInt (printf "%d" currentTime.Month)) "Year" currentTime.Year  "Hour" currentTime.Hour "Min" currentTime.Minute "Sec" currentTime.Second}}
{{$time := sdict "Day" $dTime.Day "Month" $dTime.Month "Year" $dTime.Year "Hour" 0 "Min" 0 "Sec" 0 }}
{{$dateSet := false}} {{$timeSet := false}}
{{$timeConverted := 0}}

{{/* Actual Code */}}
{{/* Fetching Dates */}}
{{/* Fetching dates written in format mm.dd.yyyy or mm/dd/yyyy or mm-dd-yyyy or mm,dd,yyyy */}} 
{{with reFindAllSubmatches `((\s|^)((\d{1,2})(\-|\.|\/|\,)(\d{1,2})(\-|\.|\/|\,)(\d{1,4}))(\s|$))` $timeString}}
{{$time.Set "Month" (toInt (index (index . 0) 4))}}
{{$time.Set "Day" (toInt (index (index . 0) 6))}}
{{$time.Set "Year" (toInt (index (index . 0) 8))}}
{{$dateSet = true}}
{{else}}
{{/* Fetching dates written as a string with both long or short month names supported. Date , Month and Year need not be present together but year must be written in full form(with 4 digits) eg: 20 sept 1am ,2019 is supported */}}
     {{ with (reFindAllSubmatches `(?:[^a-z]|^)(jan((uary)?)|feb((ruary)?)|mar((ch)?)|apr((il)?)|may|jun(e?)|jul(y?)|aug((ust)?)|sep((t(ember)?)?)|oct((ober)?)|nov((ember)?)|dec((ember)?))(?:[^a-z]|$)` (lower $timeString) ) }}
     {{$month := index (index . 0) 1}}
          {{if or (eq $month "jan") (eq $month "january")}}{{$time.Set "Month" 1}}
          {{else if or (eq $month "feb") (eq $month "february")}}{{$time.Set "Month" 2}}
          {{else if or (eq $month "mar") (eq $month "march")}}{{$time.Set "Month" 3}}
          {{else if or (eq $month "apr") (eq $month "april")}}{{$time.Set "Month" 4}}
          {{else if (eq $month "may")}}{{$time.Set "Month" 5}}
          {{else if or (eq $month "jun") (eq $month "june")}}{{$time.Set "Month" 6}}
          {{else if or (eq $month "jul") (eq $month "july")}}{{$time.Set "Month" 7}}
          {{else if or (eq $month "aug") (eq $month "august")}}{{$time.Set "Month" 8}}
          {{else if or (eq $month "sep") (eq $month "sept") (eq $month "september")}}{{$time.Set "Month" 9}}
          {{else if or (eq $month "oct") (eq $month "october")}}{{$time.Set "Month" 10}}
          {{else if or (eq $month "nov") (eq $month "november")}}{{$time.Set "Month" 11}}
          {{else if or (eq $month "dec") (eq $month "december")}}{{$time.Set "Month" 12}}
          {{end}}
     {{$temp:= reReplace `(([^:]|^)((\d+)((:(\d+)){1,2}))((\s?(am|pm))?))|(((\d+))(\s?(am|pm)))` (lower $timeString) ""}}
     {{with (reFindAllSubmatches `(?:\D|^)(\d{1,2})(?:\D|$)` $temp)}}
     {{$time.Set "Day" (toInt (index (index . 0) 1))}}
     {{end}}
     {{with (reFindAllSubmatches `(?:\D|^)(\d{4})(?:\D|$)` $temp)}}
     {{$time.Set "Year" (toInt (index (index . 0) 1))}}
     {{end}}
     {{$dateSet = true}}
     {{end}}
{{end}}

{{/* Fetching dates specified as today or tomorrow and assigning default values to invalid dates */}}
{{if $dateSet}}
{{if not $time.Day}}{{$time.Set "Day" $dTime.Day}}{{end}}
{{if not $time.Month}}{{$time.Set "Month" $dTime.Month}}{{end}}
{{if not $time.Year}}{{$time.Set "Year" $dTime.Year}}{{end}}
{{else}}
{{with reFind `(today)|(tomorrow)` (lower $timeString)}}
{{if eq . "tomorrow"}}
{{$time.Set "Day" (add $dTime.Day 1)}}
{{end}}
{{end}}
{{end}}

{{/* Fetching time specified as hh:mm or hh:mm:ss or hh. Can be followed by am/pm as well. */}}
{{with reFind `(([^:]|^)((\d+)((:(\d+)){1,2}))((\s?(am|pm))?))|(((\d+))(\s?(am|pm)))` (lower $timeString)}}
{{with reFindAllSubmatches `(\d+)` .}}
{{$time.Set "Hour" (toInt (index (index . 0) 0))}}
{{if (gt (len .) 1)}}
{{$time.Set "Min" (toInt (index (index . 1) 0) )}}
{{end}}
{{if (gt (len .) 2)}}
{{$time.Set "Sec" (toInt (index (index . 2) 0))}}
{{end}}
{{end}}
{{with reFind `(am|pm)` .}}
{{if and (eq $time.Hour 12) (eq . "am")}}
{{$time.Set "Hour" 0}}
{{else if and (eq . "pm" ) (lt $time.Hour 12)}}
{{$time.Set "Hour" (add $time.Hour 12)}}
{{end}}
{{end}}
{{$timeSet = true}}
{{end}}

{{/* Setting time to current time when both explicit date and time setting was not done */}}
{{if and (not $timeSet) (not $dateSet)}}
{{$time.Set "Hour" $dTime.Hour}}
{{$time.Set "Min" $dTime.Min}}
{{$time.Set "Sec" $dTime.Sec}}
{{end}}

{{/* Conversion to time.Time datatype */}}
{{$timeConverted = (newDate $time.Year $time.Month $time.Day $time.Hour $time.Min $time.Sec)}}

{{/*timezone adjustment */}}
{{if or $timeSet $dateSet}}
{{$TimeHour := .TimeHour}}
{{with (reFind `(\-)?\d+(:\d+)?` (exec "setz -u") )}}
{{$timeConverted.Add (toDuration (mult -1.0 (toFloat (reReplace ":" . ".")) $TimeHour))}}
{{end}}
{{else}}
{{end}}




{{if .ExecData}}
{{$cc := .ExecData.CC}}
{{execCC $cc nil 0 (sdict "time" $timeConverted)}}
{{else}}
{{$timeConverted}}
{{end}}



