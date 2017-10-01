package main

import (
	"fmt"
	"unicode"
	"os"
	"github.com/PuerkitoBio/goquery"
	"strings"
)

func isChinese(str string) bool {
	for _, r := range str {
		if unicode.Is(unicode.Scripts["Han"], r) {
			return true
		}
	}
	return false
}

func query(word string){
	var url string
	var isChn bool
	isChn=isChinese(word)
	if isChn {
		url = "http://dict.youdao.com/w/eng/%s"
	} else {
		url = "http://dict.youdao.com/w/%s"
	}

	doc, err := goquery.NewDocument(fmt.Sprintf(url, word))

	if err != nil {
		red("Query failed with err: %s", err.Error())
		os.Exit(1)
	}

	if !isChn{
		printPronounce(doc)
		printMeaning(doc)
		printSentence(doc)
	}else {
		printPronounceCN(doc)
		printMeaningCN(doc)
		printSentence(doc)
	}
}

func printPronounce(doc *goquery.Document){
	var pronounce string
	doc.Find("div.baav > span.pronounce").Each(func(i int, s *goquery.Selection) {
			p := s.Text()
			p =strings.Replace(p," ","",-1)
			ss:=strings.Split(p,"\n")
			for _,s:=range ss{
				pronounce+=fmt.Sprintf("%s ",s)
			}
	})

	red(pronounce)
}

func printPronounceCN(doc *goquery.Document){
	red(doc.Find("h2.wordbook-js span.phonetic").Text())
}

func printMeaning(doc *goquery.Document){
	doc.Find("div#phrsListTab div.trans-container ul li").Each(func(i int, s *goquery.Selection) {
        blue("%s",s.Text())
	})
}

func printMeaningCN(doc *goquery.Document){
	doc.Find("div#phrsListTab div.trans-container p.wordGroup").Each(func(i int, s *goquery.Selection) {
		line:=""
		s.Find("span").Each(func(ii int, ss *goquery.Selection) {
            if ii==0{
				line+=fmt.Sprintf("%s ",ss.Text())
			}else {
				line+=fmt.Sprintf("%s ;",ss.Find("a").Text())
			}

		})
		blue("%s",line)
	})
}

func printSentence(doc *goquery.Document){
	doc.Find("div#bilingual ul.ol li").Each(func(i int, s *goquery.Selection) {
		fmt.Println("")
		s.Find(("p")).Each(func(ii int, ss *goquery.Selection) {
			sss:=strings.Replace(ss.Text(),"\n","",-1)
			sss=strings.TrimSpace(sss)
			if ii==0{
				magenta("%d. %s",i+1,sss)
			}else{
				if ii<2{
					magenta("%s",sss)
				}
			}

		})
	})
}