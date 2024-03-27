package main

import "fmt"

//Array for list books
var booksName = []string{}
//Array for list writters
var writtersName = []string{}

//index hashon mosavi

func main(){
  var choose int

  fmt.Println("Welcome to library")
  fmt.Println("1- Add\n2- Remove\n3- Edit\n4- List\n5- Exit")
  fmt.Scan(&choose)
  
  switch choose{
    case 1:
      addBook()

    case 2:
      removeBook()

    case 3:
      editBooks()

    case 4:
      listBooks()
    
    case 5:
      fmt.Println("TNX for using :)")

    default:
      fmt.Println("Choose a option")
      main()
  }
  
}

func addBook(){
  var name string
  var writter string

  fmt.Println("Enter a title: withoutSpace")
  fmt.Scan(&name)
  fmt.Println("Enter writter's name: withoutSpace")
  fmt.Scan(&writter)
  
  booksName = append(booksName, name)
  writtersName = append(writtersName, writter)

  fmt.Println("Succes!")

  main()

}

//I got help to remove elements from this site:
//https://www.tutorialspoint.com/delete-elements-in-a-slice-in-golang#:~:text=package%20main%20import%20%22fmt%22%20func%20main()%20%7B%20slice,%3A%5D...)%20%7D
func removeBook(){
  var num int
  listBooks2()

  fmt.Println("Enter number of book: ")
  fmt.Scan(&num)
  num--

  booksName = append(booksName[:num], booksName[num+1:]...)
  writtersName = append(writtersName[:num], writtersName[num+1:]...)

  fmt.Println("Succes!")

  main()
}

func editBooks(){
  var num int
  var newName string
  var newWritter string
  listBooks2()

  fmt.Println("Enter number of book: ")
  fmt.Scan(&num)
  num--

  fmt.Println("Enter new name: ")
  fmt.Scan(&newName)
  fmt.Println("Enter new writter")
  fmt.Scan(&newWritter)

  booksName[num] = newName
  writtersName[num] = newWritter

  fmt.Println("Succes!")

  main()
}

func listBooks(){
  fmt.Println(" --- Library --- ")

  for index, value := range booksName{
    for index2, value2 := range writtersName{
      if index == index2{
        print(index + 1)
        fmt.Println("- " + value + "\nby " + value2)
        fmt.Println("----------------")
      }
    }
  }

  main() 

}


func listBooks2(){
  fmt.Println(" --- Library --- ")

  for index, value := range booksName{
    for index2, value2 := range writtersName{
      if index == index2{
        print(index + 1)
        fmt.Println("- " + value + "\nby " + value2)
        fmt.Println("----------------")
      }
    }
  }

}

