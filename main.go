package main

import (
	"html/template"
	"os"
	"encoding/json"
	"log"
	//"strings"
	"unicode"
	"strings"
)

type Entry struct{
	Keys []string
	Name string
	URL template.URL
}
type Data struct{
	Entries []Entry
}
func main() {
	data,e:=readData("config.json")
	if e!=nil{
		log.Println(e)
		return
	}
	//Augment(&data)
	UseTemplate(data)
}
func getData()Data{
	data:=struct{
		Entries []Entry
	}{
		Entries: []Entry{
			Entry{
				Keys: []string{"d","D"},
				Name: "Duolingo",
				URL: template.URL("https://www.duolingo.com"),
			},
			Entry{
				Keys: []string{"s","S"},
				Name: "DuckDuckGo",
				URL: template.URL("https://duckduckgo.com/"),
			},
			Entry{
				Keys: []string{"w","W"},
				Name: "WhatsApp",
				URL: template.URL("https://web.whatsapp.com/"),
			},
		},
	}
	return data
}
func readData(filename string)(Data,error){
	f,e:=os.Open(filename)
	if e!=nil{
		return Data{}, e
	}
	defer f.Close()
	data:=Data{}
	e=json.NewDecoder(f).Decode(&data)
	if e!=nil{
		return Data{}, e
	}
	return data,nil
}
func UseTemplate(data Data){
	t := template.Must(template.New("template.html").Funcs(template.FuncMap{
    "toup":strings.ToUpper,
    "isup":func(input string)bool{
      return input == strings.ToUpper(input)
    },
    "tolow":strings.ToLower,
    "islow":func(input string)bool{
      return input == strings.ToLower(input)
    },
  }).ParseFiles("template.html"))
	t.Execute(os.Stdout,data)
}
func writeData(filename string, data Data) {
	f,e:=os.Create(filename)
	if e!=nil{
		log.Println(e)
	}
	defer f.Close()
	enc:=json.NewEncoder(f)
	enc.SetIndent(""," ")
	enc.Encode(data)
}

func Augment(input *Data){
	for i:=range input.Entries{
		input.Entries[i].Keys=AppendCase(input.Entries[i].Keys)
	}
}
func AppendCase(input []string)([]string){
	for _,v:=range input{
		if unicode.IsUpper(rune(v[0])){
			lower:=strings.ToLower(v)
			input=AppendIfNew(input,lower)
		}else{
			upper:=strings.ToUpper(v)
			input=AppendIfNew(input,upper)
		}
	}
	return input
}
func AppendIfNew(input []string,addition string)[]string{
			found:=false
			for _,v2:=range input{
				if v2==addition{
					found=true
				}
			}
			if !found{
				input=append(input,addition)
			}
			return input
}
