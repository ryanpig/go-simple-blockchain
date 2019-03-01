package main
import (
  "fmt"
  "time"
  "log"
  "crypto/sha256"
  "net/http"
  "io"
  "encoding/hex"
  "encoding/json"

  "github.com/davecgh/go-spew/spew"
  "github.com/gorilla/mux"
  "github.com/joho/godotenv"
  "os"
  "bufio"
)

// Task division
// - Create a new blockchain by using slices of blocks 
// - Generate new block w/ SHA-256 hash 
// - Build a webserver providing REST API that allows user to view blocks and add new block client browser.

// TODO:
// - Solve conflict blockchains by choosing longer length

// Error record:
// 1. forget make the naming of struct member capital
// 2. slice doesn't increase length after `append` in the function, because it's already reached its capacity
// 3. read json data from file. Remember the format is {"key1":"value1", "key2":"value2"...}

// global varialbs
var blockchain []Block

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
  b := Block{lastBlock.BlockID + 1, timestamp, "", lastBlock.Hash, newdata}
  b.Hash = hashing(b)
  return b
}

// hashing data and returned a hash string
func hashing(b Block) string {
  strData := b.Data.ProjName + b.Data.ProjDes + string(b.Data.ProjID)
  h := sha256.New()
  h.Write([]byte(string(b.BlockID) + b.Timestamp + b.PrevHash + strData))
  hashed := hex.EncodeToString(h.Sum(nil))
  return hashed
}

// read project data from file 
func parseData(filename string) []ProjData {
  log.Println("Parsing data from the file:", filename)
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
  log.Println("Parsing finished")
  return resultData
}

// initialize blockchain from the file. (TODO: use github API) 
func blockchain_initialization() []Block {
  blockchain := make([]Block, 0)
  // create genesis block
  data := ProjData{"genesis", "This is genesis block", 0}
  t := time.Now().String()
  genesisBlock := Block{1, t,"", "", data}
  genesisBlock.Hash = hashing(genesisBlock)
  blockchain = append(blockchain, genesisBlock)

  // add blocks from the file
  p := parseData("testcase.txt")
  for _ , data_tmp := range p {
    lastBlock := blockchain[len(blockchain)-1]
    b := generateBlock(lastBlock, data_tmp)
    blockchain = append(blockchain, b)
  }
  return blockchain
}
// handling GET request, return blockchain data
func handlerGetBlockchain(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Return blockchain", "test")
  bytes, err := json.MarshalIndent(blockchain, "", "  ")
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  io.WriteString(w, string(bytes))
}
// handling POST request to add a new block 
func handlerAddBlock(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Add new block", "test")
}

// start a webserver
func startWebserver() error {
  // read port configuration from a file ".env"
  err := godotenv.Load()
  if err != nil {
        log.Fatal(err)
  }

  mux := makeMuxRouter()
  // var httpAddr string = "8080"
  httpAddr := os.Getenv("PORT")
  s := &http.Server{
    Addr:           ":" + httpAddr,
    Handler:        mux,
    ReadTimeout:    10 * time.Second,
    WriteTimeout:   10 * time.Second,
    MaxHeaderBytes: 1 << 20,
  }

  log.Println("Start webserver listening port ", httpAddr)
  if err := s.ListenAndServe(); err != nil {
    return err
  }

  return nil
}

// routing to different handlers (The router wrapper allows us to specify either GET or POST method) 
func makeMuxRouter() http.Handler {
  muxRouter := mux.NewRouter()
  muxRouter.HandleFunc("/", handlerGetBlockchain).Methods("GET")
  muxRouter.HandleFunc("/", handlerAddBlock).Methods("POST")
  return muxRouter
}

func main() {
  log.Println("---Start blockchain application---")
  // create blockchain
  blockchain = blockchain_initialization()
  spew.Dump(blockchain)
  // run server
  log.Fatal(startWebserver())
}



