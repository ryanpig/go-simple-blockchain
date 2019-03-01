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
  "strconv"
)

// Task division
// - Create a new blockchain by using slices of blocks 
// - Generate new block w/ SHA-256 hash 
// - Solve conflict blockchains by choosing longer length
// - Build a webserver providing REST API that allows user to view blocks and add new block client browser.

// Error record:
// 1. forget make the naming of struct member capital

type Block struct {
  Index int
  Timestamp string
  Hash string
  PrevHash string
  Data ProjData
}

type ProjData struct {
  ProjName string
  ProjDes string
  ProjID int
}

// func generateBlock(blockchain []Block, data ProjData) bool {

// }

func main() {
  log.Println("start blockchain application")
  // create genesis block
  data := ProjData{"cpp_mt", "multithreading project in c++", 1}
  h := sha256.New()
  h.Write([]byte(data.ProjName + data.ProjDes + strconv.Itoa(data.ProjID)))
  hashed := hex.EncodeToString(h.Sum(nil))
  genesisBlock := Block{0, time.Now().String(),hashed ,"", data}
  spew.Dump(genesisBlock)

  // create blockchain
  blockchain := make([]Block, 0)
  blockchain = append(blockchain, genesisBlock)
  // add new block
}



