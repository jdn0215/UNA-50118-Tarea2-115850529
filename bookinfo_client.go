package main

import (
    "context"
    pb "github.com/jdn0215/UNA-50118-Tarea2-115850529/booksapp"
    "google.golang.org/grpc"
    "log"
    "os"
	"time"
	"fmt"
)

func main() {
    address := os.Getenv("ADDRESS")
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewBookInfoClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()
    r, err := c.AddBook(ctx, &pb.Book{
        Id:        "1",
        Title:     "Operating System Concepts",
        Edition:   "9th",
        Copyright: "2012",
        Language:  "ENGLISH",
        Pages:     "976",
        Author:    "Abraham Silberschatz",
        Publisher: "John Wiley & Sons"})
    if err != nil {
        log.Fatalf("Could not add book: %v", err)
    }

    log.Printf("Book ID: %s added successfully", r.Value)
    book, err := c.GetBook(ctx, &pb.BookID{Value: r.Value})
    if err != nil {
        log.Fatalf("Could not get book: %v", err)
    }
	fmt.Println("*************Ejericio #1 Borrado**************************")
	fmt.Println("*************Borrando libro ",book.Id,", " , book.Title,"....")
	r, err = c.DeleteBook(ctx,&pb.BookID{Value: r.Value}); 
	if err != nil{
		log.Fatalf("Could not add book: %v", err)
	}else{
		_, err := c.GetBook(ctx, &pb.BookID{Value: r.Value})
		if err != nil {
			log.Printf("***Libro ya no pudo ser encontrado porque se borró exitosamente: %v***", err)
		}
	}
	fmt.Println("\n\n*************Ejericio #2 Actualización**************************")
	r, err = c.AddBook(ctx, &pb.Book{
        Id:        "1",
        Title:     "Operating System Concepts",
        Edition:   "9th",
        Copyright: "2012",
        Language:  "ENGLISH",
        Pages:     "976",
        Author:    "Abraham Silberschatz",
        Publisher: "John Wiley & Sons"})
    if err != nil {
        log.Fatalf("Could not add book: %v", err)
	}
	book, err = c.GetBook(ctx, &pb.BookID{Value: r.Value})
	fmt.Println("*************Se vuelve a añadir el libro para modificarlo ",book.Id)
	fmt.Println(book)
	fmt.Println("*************Modificando libro ",book.Id,", " , book.Title,"....")
	book, err = c.UpdateBook(ctx, &pb.Book{
        Id:        book.Id,
        Title:     "Conceptos de sistemas operativos",
        Edition:   book.Edition,
        Copyright: "2020",
        Language:  "SPANISH",
        Pages:     "1000",
        Author:    book.Author,
		Publisher: "J&D"});
	if err != nil {
		log.Fatalf("Could not update book: %v ", err)
	}else{
		fmt.Println("*************Libro modificado ",book.Id);	
		fmt.Println(book)
	}
}
