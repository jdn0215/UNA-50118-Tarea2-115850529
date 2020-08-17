package main

import (
    "context"
    pb "github.com/jdn0215/UNA-50118-Tarea2-115850529/booksapp"
    "google.golang.org/grpc"
    "log"
    "os"
	"time"
	"fmt"
	"encoding/csv"
)

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

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
	fmt.Println("\n\n*************Ejericio #1 Borrado**************************")
	fmt.Println("*************Borrando libro ",book.Id,", " , book.Title,"....")
	r, err = c.DeleteBook(ctx,&pb.BookID{Value: r.Value}); 
	if err != nil{
		log.Fatalf("Could not add book: %v", err)
	}else{
		_, err := c.GetBook(ctx, &pb.BookID{Value: r.Value})
		if err != nil {
			log.Printf("\n***Libro ya no pudo ser encontrado porque se borró exitosamente: %v***", err)
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
		fmt.Println("\n*************Libro modificado correctamente",book.Id);	
		fmt.Println(book)
	}

	fmt.Println("\n\n*************Ejericio #3 Añadiendo y listando**************************")
	fmt.Println("*************Añadiendo registros desde CSV**************************")
	filePath := "books.csv"
	file, err1 := os.Open(filePath)
	checkError("Unable to read input file "+filePath, err1)
	defer file.Close()
	csvReader := csv.NewReader(file)
	records, err2 := csvReader.ReadAll()
	checkError("Unable to parse file as CSV for "+filePath, err2)
	defer file.Close()

	for _, record := range records {
		r, err = c.AddBook(ctx, &pb.Book{
			Id:        record[0],
			Title:     record[1],
			Edition:   record[2],
			Copyright: record[3],
			Language:  record[4],
			Pages:     record[5],
			Author:    record[6],
			Publisher: record[7]})
		if err != nil {
			log.Fatalf("Could not add book: %v", err)
		}
	}
	file.Close()
	fmt.Println("*************Listando registros que se añadieron desde el CSV**************************")
	books,err := c.GetAll(ctx,&pb.Nil{});
	if err != nil{
		log.Fatalf("Could not list books: %v ", err)
	}else{
		log.Printf("***Listado de libros: %v***:\n", books)
	}

}
