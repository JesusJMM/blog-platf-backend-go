package main

import "github.com/JesusJMM/blog-plat-go/postgres"

const TEST_USERNAME = "TestUser"

func createFailUser() postgres.User {
  img := "https://images.unsplash.com/photo-1571771894821-ce9b6c11b08e?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8&auto=format&fit=crop&w=80&q=80"
  return postgres.User{
    Name: TEST_USERNAME,
    Img: &img,
  }
}
