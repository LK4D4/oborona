package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var generatorStrings = [][]string{
	{
		"солнечный",
		"траурный",
		"плюшевый",
		"бешеный",
		"памятный",
		"трепетный",
		"базовый",
		"скошенный",
		"преданный",
		"ласковый",
		"пойманный",
		"радужный",
		"огненный",
		"радостный",
		"тензорный",
		"шёлковый",
		"пепельный",
		"ламповый",
		"жареный",
		"загнанный",
	},
	{
		"зайчик",
		"Верник",
		"глобус",
		"ветер",
		"щавель",
		"пёсик",
		"копчик",
		"ландыш",
		"стольник",
		"мальчик",
		"дольшик",
		"Игорь",
		"невод",
		"егерь",
		"пончик",
		"лобстер",
		"жемчуг",
		"кольщик",
		"йогурт",
		"овод",
	},
	{
		"стеклянного",
		"ванильного",
		"резонного",
		"широкого",
		"дешёвого",
		"горбатого",
		"собачьего",
		"исконного",
		"волшебного",
		"картонного",
		"лохматого",
		"арбузного",
		"огромного",
		"запойного",
		"великого",
		"бараньего",
		"вандального",
		"едрёного",
		"парадного",
		"укромного",
	},
	{
		"глаза",
		"плова",
		"Пельша",
		"мира",
		"деда",
		"жира",
		"мема",
		"ада",
		"бура",
		"жала",
		"нёба",
		"гунна",
		"хлама",
		"шума",
		"воза",
		"сала",
		"фена",
		"зала",
		"рака",
		"Глеба",
	},
}

var port = flag.Int("port", 5000, "listen port")

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, GenOborona(generatorStrings))
	})
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))

}