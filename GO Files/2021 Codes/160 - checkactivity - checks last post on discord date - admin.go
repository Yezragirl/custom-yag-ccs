{{/*checkactivity allows anyone to see how long its been since someone spoke in the discord*/}}
{{$args := parseArgs 1 "correct syntax is `checkactivity who`"
    (carg "user" "who")}}
{{$user := $args.Get 0}}
{{$db := dbGet $user.ID "inactive"}}
{{$inactive_time := sub (currentTime.Sub (newDate 2019 11 11 0 0  0)).Seconds $db.Value }}
{{$duration := (toDuration  (mult -1  $inactive_time .TimeSecond))}}
{{humanizeTimeSinceDays ((currentTime).Add $duration)}}
