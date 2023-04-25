package main

import (
	"encoding/json"
	"log"
	"net/http"
	"project/mirrors"
	"time"
)

type response struct {
	FastestURL string        `json:"fastert_url`
	Latency    time.Duration `json:"latency"`
}

func findFastest(urls []string) response {
	urlChan := make(chan string)
	latencyChan := make(chan time.Duration)

	for _, url := range urls {
		mirrorURL := url
		go func() {
			log.Println("Started probing: ", mirrorURL)
			start := time.Now()
			_, err := http.Get(mirrorURL + "/README")
			latency := time.Now().Sub(start) / time.Millisecond
			if err == nil {
				urlChan <- mirrorURL
				latencyChan <- latency
			}
			log.Printf("Got the best mirror: %s with latency: %s  ", mirrorURL, latency)
		}()
	}
	return response{<-urlChan, <-latencyChan}
}

func main() {
	http.HandleFunc("/fastest-mirror", func(w http.ResponseWriter, r *http.Request) {
		response := findFastest(mirrors.MirrorList)
		respJson, _ := json.Marshal(response)
		w.Header().Set("Content-Type", "application/json")
		w.Write(respJson)
	}())
	port := ":8000"
	server := &http.Server{
		Addr:           port,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Printf("starting server on port: &port\n", port)
	log.Fatal(server.ListenAndServe())

	// At this point, i should be able to this buy my self, so next project should be to read and practive as we read
	// not just to finish the book,  but to learn from the book
}

/*

I fucked up, now it's time to unfuck this up
this is defenetely a much better

how ever when it comes to keyboard prices, i wouldn't pay much more than this, and to get a keyboard with more keys
i'd be paying maybe less 400 or 100-200 more, so i am not stressing about the keyboard, if i like and can use
i'll allow myself to keep it.
well, if i am to be honest, my pc keyboard, it's pretty neat, truth be told, it's fine
and i get a double monitor setup (well, double screen)
damn son, that mechanial tho...

well, if i am to be honest, my pc keyboard, it's pretty neat, truth be told, it's fine and i get  a double monitor
setup (Well, double screen
damn son, this mechanical tho...
this is miles better,  brother,
well, this keyboard works fine
well this keyboard works fine
so yeah, 24 dioma monitor and my pc perfect combo, 24 dioma i remember often requires my pc to serve as double scree\
27 not really, 27 can handle mu

The monitor on the other hand, lol, liking it or not, it has to go!!
maybe i should be lookin gat the keyboard, since i want to see witch color  scheme i like the most
, ok
shining one light only, not bad, bring it bnback, ok, yeah,  i kinda like, simple, just one nice cool light
but i won't be using this one, it isn't as one as the others
thios literally just changed the colour, but ok, whoah, this one alternates through all of the available colours, not just the same one blinking
kinda like it, might use this one no cap

uhmm, this one is ok, but am now convinced that i like the one that lights the colours accordingly to they keys am pressing
it just, i just like the effect it feels like i cactually acaused it, while these other modes are just shinign
and this one does kinda of the same thing, but this time vertically, i rather see it horizontally
f

ooh, yeah , i nkow this mode, where it light up the whole row, not just the key yoour press, this one is also fine
but am sure i will  be using it, it's actually nice, mainly when u pay aattetion and reallise that you havent' used the bottom row in a lil while now

mom i missi me mommy, yeah,  i typed mom so that i would use the bottom row more
bad dbitches get good boys and fuck them up
good bitches get good guys fucked bup by bad bitches.
mest, ok this mode is just rrainbow, not really a fan, but am definetely i fan of this modes, where you  damn, i i type fadter i t also sshines fasdter, ok,elat shace a raxce, see yyou can handle my speed, they call me aspeedy gozales
ook yo uwin
jfwhat s the diference between this one and the last one, there's a dif, i just can't spot ti, klol sddwhjoa
this one lights the whole keyboard on pressjfhfjfjfjfk
yeah, tythis was my fav mode, but now i see that i like the one that lits the kwytboad a bit more, but i really like this one
where it only lights up the key i've pressed and then  it goes dark
yeah, i;kll stick with this one i guess, there's

*/
