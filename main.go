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
// 2. slice doesn't increase length after `append` in the function, because it's already reached its capacity
// 3. read json data from file. Remember the format is {"key1":"value1", "key2":"value2"...}

type Block struct {
  BlockID int
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
func generateBlock(lastBlock Block, newdata ProjData) Block {
  //check current l 
  timestamp := time.Now().String()
  b := Block{lastBlock.BlockID + 1, timestamp, hashing(&newdata, timestamp), lastBlock.Hash, newdata}
  return b
}

// hashing data and returned a hash string
func hashing(data *ProjData, timestamp string) string {
  h := sha256.New()
  h.Write([]byte(data.ProjName + data.ProjDes + strconv.Itoa(data.ProjID) + timestamp))
  hashed := hex.EncodeToString(h.Sum(nil))
  return hashed
}

// read project data from file 
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
  }
  return resultData
}

func main() {
  log.Println("---Start blockchain application---")

  // create blockchain
  blockchain := make([]Block, 0)

  // create genesis block
  data := ProjData{"genesis", "This is genesis block", 0}
  t := time.Now().String()
  genesisBlock := Block{1, t, hashing(&data, t),"" , data}
  spew.Dump(genesisBlock)
  blockchain = append(blockchain, genesisBlock)

  // add blocks from the file
  p := parseData("testcase.txt")
  for _ , data_tmp := range p {
    lastBlock := blockchain[len(blockchain)-1]
    b := generateBlock(lastBlock, data_tmp)
    blockchain = append(blockchain, b)
    // debuging
    log.Println("Added block, the length of blockchain:", len(blockchain), "cap:", cap(blockchain))
    last_index := len(blockchain) - 1
    spew.Dump(blockchain[last_index])
  }
}



