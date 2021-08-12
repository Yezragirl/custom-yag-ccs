{{$url := "https://uploads-ssl.webflow.com/5e23ac32992bfa60a92850d9/60305dab9d11402d33a27816_Kibble%20Recipes.JPG"}}
{{$color := randInt 111111 999999 }}
 {{$embed := cembed 
 "color" $color
"Title" "Kibble"
"image" (sdict "url" $url) }}
{{$embed}}



