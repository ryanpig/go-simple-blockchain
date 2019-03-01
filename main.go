package main
import (
  "time"
  "log"
  "crypto/sha256"
  // "net/http"
  // "io"
  "encoding/hex"
  "encoding/json"

  "github.com/davecgh/go-spew/spew"
  // "github.com/gorilla/mux"
  // "github.com/joho/godotenv"
  "strconv"
  "os"
  "bufio"
)

// Task division
// - Create a new blockchain by using slices of blocks 
// - Generate new block w/ SHA-256 hash 
// - Solve conflict blockchains by choosing longer length (TODO)
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

// generate a new block to blockchain
func generateBlock(blockchain []Block, newdata ProjData) bool {
  //check current l 
  last_index := len(blockchain) - 1
  b := blockchain[last_index]
  timestamp := time.Now().String()
  blockchain = append(blockchain, Block{b.Index, timestamp, hashing(&newdata, timestamp), b.Hash, newdata})
  log.Println("Added block, the length of blockchain:", len(blockchain), "cap:", cap(blockchain))
  last_index = len(blockchain) - 1
  spew.Dump(blockchain[last_index])
  return true
}

// hashing data and returned a hash string
func hashing(data *ProjData, timestamp string) string {
  h := sha256.New()
  h.Write([]byte(data.ProjName + data.ProjDes + strconv.Itoa(data.ProjID) + timestamp))
  hashed := hex.EncodeToString(h.Sum(nil))
  return hashed
}

func parseData(filename string) []ProjData {
  resultData := make([]ProjData, 0)

  file, err := os.Open(filename)
  if err != nil {
    log.Fatal(err)
  }
  defer file.Close()
  fileScanner := bufio.NewScanner(file)

  // using json.Unmarshal to decode each line and add it to Block slice
  for fileScanner.Scan() {
      log.Println(fileScanner.Text())
      line := fileScanner.Text()
      res := ProjData{}
      json.Unmarshal([]byte(line), &res)
      resultData = append(resultData, res)
      // Debuging
      // log.Println("parsing result...")
      // spew.Dump(res)
  }

  return resultData
}

func main() {
  log.Println("start blockchain application")
  // create genesis block
  data := ProjData{"cpp_mt", "multithreading project in c++", 1}
  t := time.Now().String()
  genesisBlock := Block{1, t, hashing(&data, t),"" , data}
  spew.Dump(genesisBlock)

  // create blockchain
  blockchain := make([]Block, 0)
  blockchain = append(blockchain, genesisBlock)

  // add new block
  data2 := ProjData{"goserver", "build a go webserver", 2}
  if generateBlock(blockchain, data2) {
    log.Println("success of generation of new block")
  } else {
    log.Println("failure of generation of new block")
  }

  // add more blocks from the file
  p := parseData("testcase.txt")
  for _ , d := range p {
    if generateBlock(blockchain, d) {
      log.Println("success of generation of new block")
    } else {
      log.Println("failure of generation of new block")
    }
  }
}



