package main
import (
  "bufio"
  "context"
  "crypto/sha256"
  "encoding/hex"
  "encoding/json"
  "fmt"
  "log"
  "net/http"
  "os"
  "time"

  "github.com/davecgh/go-spew/spew"
  "github.com/gorilla/mux"
  "github.com/joho/godotenv"
  "golang.org/x/oauth2"
  "github.com/shurcooL/githubv4"
)

// global variable
var blockchain []Block
var query struct {
  Viewer struct {
    Repositories struct {
      Nodes []struct {
        Name         string
        Description  string
        Id           string
      }
    } `graphql:"repositories(first: 100)"`
  }
}

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
  ProjID string 
}

// generate a new block to blockchain
func generateBlock(lastBlock Block, newdata ProjData) Block {
  timestamp := time.Now().String()
  b := Block{lastBlock.BlockID + 1, timestamp, "", lastBlock.Hash, newdata}
  b.Hash = hashing(b)
  return b
}

// hash data and return a hash string
func hashing(b Block) string {
  strData := b.Data.ProjName + b.Data.ProjDes + b.Data.ProjID
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

  // use json.Unmarshal to decode each line and add it to Block slice
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

func makeQuery() {
  log.Println("making a query via Github API")
  src := oauth2.StaticTokenSource(
    &oauth2.Token{AccessToken: os.Getenv("GITHUB_TOKEN")},
  )
  httpClient := oauth2.NewClient(context.Background(), src)
  client := githubv4.NewClient(httpClient)
  err := client.Query(context.Background(), &query, nil)
  if err != nil {
    // err 
  }
  // GraphQL
  err = client.Query(context.Background(), &query, nil)
  if err != nil {
      fmt.Println("error", err)
  }
}

func parseDataFromAPI() []ProjData {
  log.Println("Parsing data from Github API v4")
  resultData := make([]ProjData, 0)

  // retrieve data from GraphQL
  makeQuery()
  spew.Dump(query)

  // save query result to project data
  for _, n := range query.Viewer.Repositories.Nodes {
      res := ProjData{n.Name, n.Description, n.Id}
      resultData = append(resultData, res)
  }
  log.Println("Parsing finished")
  return resultData
}

// initialize blockchain from the file. (TODO: use github API) 
func blockchain_initialization() []Block {

  blockchain := make([]Block, 0)
  // create genesis block
  data := ProjData{"genesis", "This is genesis block", "0"}
  t := time.Now().String()
  genesisBlock := Block{1, t,"", "", data}
  genesisBlock.Hash = hashing(genesisBlock)
  blockchain = append(blockchain, genesisBlock)

  // add blocks from the file
  // p := parseData("testcase.txt")
  p := parseDataFromAPI()
  for _ , data_tmp := range p {
    lastBlock := blockchain[len(blockchain)-1]
    b := generateBlock(lastBlock, data_tmp)
    blockchain = append(blockchain, b)
  }
  return blockchain
}

func makeSimpleHtml() string {

  var sum string =` 
  <style>
  table {
    font-family: arial, sans-serif;
    border-collapse: collapse;
    width: 100%;
  }

  td, th {
    border: 1px solid #dddddd;
    text-align: center;
    padding: 8px;
  }

  tr:nth-child(even) {
    background-color: #dddddd;
  }
  </style>
  <table>
  <tr>
  <th>BlockID</th>
  <th>Timestamp</th>
  <th>Hash </th>
  <th>PrevHash </th>
  <th>Project Name </th>
  <th>Project Description</th>
  <th>Project ID </th>
  </tr>
  `
  for _, b := range blockchain {
    sum  += fmt.Sprintf("<tr> <td>%d</td> <td>%s</td> <td>%s</td> <td>%s</td> <td>%s</td> <td>%s</td> <td>%s</td> </tr>", b.BlockID, b.Timestamp, b.Hash, b.PrevHash, b.Data.ProjName, b.Data.ProjDes, b.Data.ProjID)
  }

  sum +=  "</table>"
  return sum
}

// handling GET request, return blockchain data
func handlerGetBlockchain(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>%s</h1>", "Blockchain Info")
  fmt.Fprintf(w, "%s", makeSimpleHtml())
}
// handling POST request to add a new block 
func handlerAddBlock(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", "Add new block", "test")
}

// start a web server
func startWebserver() error {
  // read port configuration from a file ".env"
  err := godotenv.Load()
  if err != nil {
        log.Fatal(err)
  }

  mux := makeMuxRouter()
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



