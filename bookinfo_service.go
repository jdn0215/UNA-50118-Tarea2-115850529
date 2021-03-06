package main

import (
	"fmt"
    "context"
    pb "github.com/jdn0215/UNA-50118-Tarea2-115850529/booksapp"
    "github.com/gofrs/uuid"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
)

type server struct {
    bookMap map[string]*pb.Book
}

func (s *server) AddBook(ctx context.Context, in *pb.Book) (*pb.BookID, error) {
    out, err := uuid.NewV4()
    if err != nil {
        return nil, status.Errorf(codes.Internal,
            "Error while generating Book ID", err)
    }
    in.Id = out.String()
    if s.bookMap == nil {
        s.bookMap = make(map[string]*pb.Book)
    }
    s.bookMap[in.Id] = in
    return &pb.BookID{Value: in.Id}, status.New(codes.OK, "").Err()
}

func (s *server) GetBook(ctx context.Context, in *pb.BookID) (*pb.Book, error) {
	value, exists := s.bookMap[in.Value]
    if exists {
        return value, status.New(codes.OK, "").Err()
    }
    return nil, status.Errorf(codes.NotFound, "Book does not exist.", in.Value)
}
/****Ejercicio #1 *************/
func (s * server) DeleteBook(ctx context.Context, in *pb.BookID) (*pb.BookID,error){
	_, exists := s.bookMap[in.Value];
	if !exists{
		return nil,status.Errorf(codes.NotFound,"Book does not exist.",in.Value);
	}
	fmt.Println("Libro encontrado...")
	delete(s.bookMap,in.Value)
	fmt.Println("Libro borrado exitosamente")
	return in, status.New(codes.OK, "").Err();
}

/*** Ejericio #2 ********************/
func (s * server) UpdateBook(ctx context.Context, in *pb.Book) (*pb.Book,error){
	_, exists := s.bookMap[in.Id];
	if !exists{
		fmt.Println("Libro no encontrado: ",in.Id,s.bookMap)
		return in,status.Errorf(codes.NotFound,"Book does not exist.",in.Id);
	}
	s.bookMap[in.Id] = in
	fmt.Println("Libro actualizado exitosamente")
	return in,nil;
}
/*** Ejericio #3 ********************/
func (s * server) GetAll(ctx context.Context, n *pb.Nil) (*pb.ArrayBook,error){
 result := &pb.ArrayBook{};
 fmt.Println("++++")
 fmt.Println(result.Books)
 fmt.Println("++++")
 for _, v := range s.bookMap { 
    result.Books = append(result.Books,v);
 }  
 return result,nil;
}
	
