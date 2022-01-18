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

  }

  func tiktok(name string){
    url := "https://tiktok.com/@"+name
    //fmt.Printf("[?] %v için tiktok taraması yapılıyor.\n",name)
    color.Magenta("[?] %v için tiktok taraması yapılıyor.\n",name)
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
    //fmt.Printf("[?] %v için VSCO taraması yapılıyor.\n",name)
    color.Magenta("[?] %v için VSCO taraması yapılıyor.\n",name)
    res,_ := http.Get(url)

    if res.StatusCode !=200{
      color.Red("[-] Hesap bulunamadı")
      return
    }else{
      color.Cyan("[+] Hesap bulundu")
    }



    }
  
