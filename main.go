package main

import (
	"fmt"
	"net/http"
  "github.com/fatih/color"
)
  func main(){
    fmt.Println(`

   █████████   ███████████ ██████   ██████   █████████     █████████    █████████  
  ███░░░░░███ ░█░░░███░░░█░░██████ ██████   ███░░░░░███   ███░░░░░███  ███░░░░░███ 
 ░███    ░███ ░   ░███  ░  ░███░█████░███  ░███    ░███  ███     ░░░  ░███    ░███ 
 ░███████████     ░███     ░███░░███ ░███  ░███████████ ░███          ░███████████ 
 ░███░░░░░███     ░███     ░███ ░░░  ░███  ░███░░░░░███ ░███          ░███░░░░░███ 
 ░███    ░███     ░███     ░███      ░███  ░███    ░███ ░░███     ███ ░███    ░███ 
 █████   █████    █████    █████     █████ █████   █████ ░░█████████  █████   █████
░░░░░   ░░░░░    ░░░░░    ░░░░░     ░░░░░ ░░░░░   ░░░░░   ░░░░░░░░░  ░░░░░   ░░░░░ 
                                                                                 
                                                                                   
                                                                                   
    `)
    var name string
    fmt.Print("kullanıcı adı : ")
    fmt.Scanln(&name)
    fmt.Println()
    tiktok(name)
    vsco(name)
    instagram(name)
    twitter(name)
 
  }

  func tiktok(name string){
    url := "https://tiktok.com/@"+name
    color.Magenta("[?] %v için TikTok taraması yapılıyor.\n",name)
    res,_ := http.Get(url)

    if res.StatusCode !=200{
      color.Red("[-] Hesap bulunamadı")
      return
    }else{
      color.Cyan("[+] Hesap bulundu")
    }
  }

  func vsco(name string){
    
    url := "https://vsco.co/"+name+"/gallery"
    color.Magenta("[?] %v için VSCO taraması yapılıyor.\n",name)
    res,_ := http.Get(url)

    if res.StatusCode !=200{
      color.Red("[-] Hesap bulunamadı")
      return
    }else{
      color.Cyan("[+] Hesap bulundu")
    }
    }


    func instagram(name string){
    url := "https://instagram.com/"+name
    color.Magenta("[?] %v için Instagram taraması yapılıyor.\n",name)
    res,_ := http.Get(url)

    if res.StatusCode !=200{
      color.Red("[-] Hesap bulunamadı")
      return
    }else{
      color.Cyan("[+] Hesap bulundu")
    }
    }

    func twitter(name string){
    url := "https://twitter.com/"+name
    color.Magenta("[?] %v için Twitter taraması yapılıyor.\n",name)
    res,_ := http.Get(url)

    if res.StatusCode !=200{
      color.Red("[-] Hesap bulunamadı")
      return
    }else{
      color.Cyan("[+] Hesap bulundu")
    }
    }
