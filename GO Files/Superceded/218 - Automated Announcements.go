{{/*Automated Announcements Code*/}}
{{$ad := 0}}
{{$url := "http"}}
{{$color := randInt 111111 999999 }}
{{$dTime := currentTime}}
{{$time := sdict "Day" $dTime.Day "Month" $dTime.Month "Year" $dTime.Year "Hour" 0 "Min" 0 "Sec" 0 }}
{{if eq $time.Day 25}}
	{{$url = "https://uploads-ssl.webflow.com/5e23ac32992bfa60a92850d9/5f9f7f5445546d617bf0c1d0_COUPON_BUYBACK.png"}}
{{else if or (eq $time.Day 2) (eq $time.Day 15)}}
	{{$ad = (randInt 1 5)}}
	{{if eq $ad 1}}
		{{$url = "https://uploads-ssl.webflow.com/5e23ac32992bfa60a92850d9/5f85dd4cd0fbbe5e5aefa9fa_KEEP_FAMILY_TOGETHER_REAPER.png"}}
	{{else if eq $ad 2}}
		{{$url = "https://uploads-ssl.webflow.com/5e23ac32992bfa60a92850d9/5f85dd4cadfb7963a14e975a_LOVE_THE_SERVER_RAPTOR.png"}}
	{{else if eq $ad 3}}
		{{$url = "https://uploads-ssl.webflow.com/5e23ac32992bfa60a92850d9/5f85dd4c59186dfb35b85628_COME_AND_JOIN_REX.png"}}
	{{else if eq $ad 4}}
		{{$url = "https://uploads-ssl.webflow.com/5e23ac32992bfa60a92850d9/5f85dd4c61966f2e16beeb1e_BECOME_A_SPONSOR_UNDERWATER.png"}}
	{{end}}
{{else if eq $time.Day 8}}
	{{$url = "https://uploads-ssl.webflow.com/5e23ac32992bfa60a92850d9/5f9f8276a96472c1e2fdc2e2_PLAYER_SHOPS.png"}}

{{else if eq $time.Day 10}}
	{{$url = "https://uploads-ssl.webflow.com/5e23ac32992bfa60a92850d9/607494ba3e5b5f7bc729058f_ADMIN_MAP_SWEEPS.png"}}

{{else if eq $time.Day 17}}
	{{$url = "https://uploads-ssl.webflow.com/5e23ac32992bfa60a92850d9/60be3d237cc6a56f37a8a4ee_Admin_Recruitment_-_Admin_application.png"}}


{{end}}
	{{$embed := cembed 
		"color" $color
		"image" (sdict "url" $url) }}

{{sendMessage 573220101977800704 (complexMessage "content" "<@&573210479753691170>" "embed" $embed)}}
