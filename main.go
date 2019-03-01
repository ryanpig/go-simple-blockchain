package main
import (
  "time"
  "log"
  "crypto/sha256"
  // "net/http"
  // "io"
  "encoding/hex"
  // "encoding/json"

  "github.com/davecgh/go-spew/spew"
  // "github.com/gorilla/mux"
  // "github.com/joho/godotenv"
)

// Task division
// - Create a new blockchain by using slices of blocks 
// - Generate new block w/ SHA-256 hash (package: "crypto/sha-256")
// - Solve conflict blockchains by choosing longer length
// - Build a webserver providing REST API that allows user to view blocks and add new block client browser. (package: "mux" , "net/http")

// Error record:
// 1. forget make the naming of struct member capital

type Block struct {
  Index int
  Timestamp string
  Hash string
  ProjName string
  ProjDes string
  ProjID int
}

func main() {
  log.Println("start blockchain application")
  h := sha256.New()
  h.Write([]byte("test string"))
  
  b := Block{0, time.Now().String(), hex.EncodeToString(h.Sum(nil)), "cpp_mt", "multithreading project in c++", 1}
  spew.Dump(b)
}



