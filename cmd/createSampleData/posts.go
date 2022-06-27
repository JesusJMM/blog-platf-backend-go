package main

import (
	"fmt"

	"github.com/JesusJMM/blog-plat-go/postgres"
)

func createFailPosts(userID int) [15]postgres.Article {
  out := [15]postgres.Article{}
  for i := 0; i < 15; i++{
    Desc := "Article for test propouse, this article must be just in development enviroment"
    Content := "Article Content"
    out[i] = postgres.Article{
      Title: "Article Title",
      Desc: &Desc,
      Content: &Content,
      Slug: fmt.Sprintf("test-article-n%d", i),
      UserID: userID,
    }
  }
  return out
}

var ImgArr [15]string = [15]string{
  "https://images.unsplash.com/photo-1655552360620-5212a1b24961?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwzfHx8ZW58MHx8fHw%3D&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1531297484001-80022131f5a1?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8dGVjaHxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1655438819488-69bfb10a6c21?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw0fHx8ZW58MHx8fHw%3D&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1496065187959-7f07b8353c55?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MTV8fHRlY2h8ZW58MHx8MHx8&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1655372501819-4c1261a50c55?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwzMnx8fGVufDB8fHx8&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1518770660439-4636190af475?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MTB8fHRlY2h8ZW58MHx8MHx8&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1655012325191-cbc22182fa9f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHx0b3BpYy1mZWVkfDl8Q0R3dXdYSkFiRXd8fGVufDB8fHx8&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1653669718797-5670d0b57ca2?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHx0b3BpYy1mZWVkfDM5fENEd3V3WEpBYkV3fHxlbnwwfHx8fA%3D%3D&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1488590528505-98d2b5aba04b?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Mnx8dGVjaHxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1526228653958-2fda2c7479eb?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHx0b3BpYy1mZWVkfDJ8NnNNVmpUTFNrZVF8fGVufDB8fHx8&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1655212874354-5dace1fdc6ce?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHx0b3BpYy1mZWVkfDV8NnNNVmpUTFNrZVF8fGVufDB8fHx8&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1638913975386-d61f0ec6500d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwxMXx8fGVufDB8fHx8&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1655566655625-4aadace23511?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwzNXx8fGVufDB8fHx8&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1655493623673-13f063cb78cf?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw2NXx8fGVufDB8fHx8&auto=format&fit=crop&w=200&q=60",
  "https://images.unsplash.com/photo-1655480530311-b3c63da9021e?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw3MXx8fGVufDB8fHx8&auto=format&fit=crop&w=200&q=60",
}

var BigImgArr [15]string = [15]string{
  "https://images.unsplash.com/photo-1655552360620-5212a1b24961?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwzfHx8ZW58MHx8fHw%3D&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1531297484001-80022131f5a1?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8dGVjaHxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1655438819488-69bfb10a6c21?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw0fHx8ZW58MHx8fHw%3D&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1496065187959-7f07b8353c55?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MTV8fHRlY2h8ZW58MHx8MHx8&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1655372501819-4c1261a50c55?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwzMnx8fGVufDB8fHx8&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1518770660439-4636190af475?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MTB8fHRlY2h8ZW58MHx8MHx8&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1655012325191-cbc22182fa9f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHx0b3BpYy1mZWVkfDl8Q0R3dXdYSkFiRXd8fGVufDB8fHx8&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1653669718797-5670d0b57ca2?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHx0b3BpYy1mZWVkfDM5fENEd3V3WEpBYkV3fHxlbnwwfHx8fA%3D%3D&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1488590528505-98d2b5aba04b?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Mnx8dGVjaHxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1526228653958-2fda2c7479eb?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHx0b3BpYy1mZWVkfDJ8NnNNVmpUTFNrZVF8fGVufDB8fHx8&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1655212874354-5dace1fdc6ce?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHx0b3BpYy1mZWVkfDV8NnNNVmpUTFNrZVF8fGVufDB8fHx8&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1638913975386-d61f0ec6500d?ixlib=rb-1.2.1&ixid=MnwxMjA3fDF8MHxlZGl0b3JpYWwtZmVlZHwxMXx8fGVufDB8fHx8&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1655566655625-4aadace23511?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHwzNXx8fGVufDB8fHx8&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1655493623673-13f063cb78cf?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw2NXx8fGVufDB8fHx8&auto=format&fit=crop&w=600&q=60",
  "https://images.unsplash.com/photo-1655480530311-b3c63da9021e?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxlZGl0b3JpYWwtZmVlZHw3MXx8fGVufDB8fHx8&auto=format&fit=crop&w=600&q=60",
}

