package console

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ellcrys/elld/accountmgr"

	prompt "github.com/c-bata/go-prompt"

	"github.com/ellcrys/elld/crypto"
	"github.com/ellcrys/elld/rpc/jsonrpc"
	"github.com/ellcrys/elld/util"
	"github.com/ellcrys/elld/util/logger"
	"github.com/fatih/color"
	"github.com/fatih/structs"
	prettyjson "github.com/ncodes/go-prettyjson"
	"github.com/robertkrimen/otto"
)

// FuncCallError creates an error describing
// an issue with the way a function was called.
func FuncCallError(msg string) error {
	return fmt.Errorf("function call error: %s", msg)
}

// Executor is responsible for executing operations inside a
// javascript VM.
type Executor struct {

	// vm is an Otto instance for JS evaluation
	vm *otto.Otto

	// exit indicates a request to exit the executor
	exit bool

	// rpc holds rpc client and config
	rpc *RPCConfig

	// coinbase is the loaded account used
	// for signing blocks and transactions
	coinbase *crypto.Key

	// authToken is the token derived from the last login() invocation
	authToken string

	// log is a logger
	log logger.Logger

	acctMgr *accountmgr.AccountManager
}

// NewExecutor creates a new executor
func newExecutor(coinbase *crypto.Key, l logger.Logger) *Executor {
	e := new(Executor)
	e.vm = otto.New()
	e.log = l
	e.coinbase = coinbase
	return e
}

func (e *Executor) login(args ...interface{}) interface{} {

	// parse arguments.
	// App RPC functions can have zero or one argument
	var arg interface{}
	if len(args) > 0 {
		arg = args[0]
	}

	// Call the auth RPC method
	rpcResp, err := e.rpc.Client.call("auth", arg, "")
	if err != nil {
		e.log.Error(color.RedString(RPCClientError(err.Error()).Error()))
		v, _ := otto.ToValue(nil)
		return v
	}

	if !rpcResp.IsError() {
		e.authToken = rpcResp.Result.(string)
		return nil
	}

	// decode response object to a map
	s := structs.New(rpcResp)
	s.TagName = "json"
	return s.Map()
}

// PrepareContext adds objects and functions into the VM's global
// contexts allowing users to have access to pre-defined values and objects
func (e *Executor) PrepareContext() ([]prompt.Suggest, error) {

	var suggestions = []prompt.Suggest{}

	// Add some methods to the global namespace
	e.vm.Set("pp", e.pp)
	e.vm.Set("runScript", e.runScript)
	e.vm.Set("rs", e.runScript)
	e.vm.Set("tx", func() *TxBuilder {
		return NewTxBuilder(e)
	})

	// Get all the methods
	resp, err := e.rpc.Client.call("methods", nil, e.authToken)
	if err != nil {
		e.log.Error(color.RedString(RPCClientError(err.Error()).Error()))
	}

	// Create console suggestions and collect methods info
	var methodsInfo = []jsonrpc.MethodInfo{}
	for _, m := range resp.Result.([]interface{}) {
		var mInfo jsonrpc.MethodInfo
		util.MapDecode(m, &mInfo)
		suggestions = append(suggestions, prompt.Suggest{
			Text:        fmt.Sprintf("%s.%s", mInfo.Namespace, mInfo.Name),
			Description: mInfo.Description,
		})
		methodsInfo = append(methodsInfo, mInfo)
	}

	// Add supported methods to the global objects map
	var namespacesObj = map[string]map[string]interface{}{}
	if len(methodsInfo) > 0 {
		for _, methodInfo := range methodsInfo {
			var mName = methodInfo.Name
			var ns = methodInfo.Namespace
			if namespacesObj[ns] == nil {
				namespacesObj[ns] = map[string]interface{}{}
			}
			namespacesObj[ns][mName] = func(args ...interface{}) interface{} {

				// parse arguments.
				// App RPC functions can have zero or one argument
				var arg interface{}
				if len(args) > 0 {
					arg = args[0]
				}

				// Call the RPC method passing the RPC API params
				rpcResp, err := e.rpc.Client.call(mName, arg, e.authToken)
				if err != nil {
					e.log.Error(color.RedString(RPCClientError(err.Error()).Error()))
					v, _ := otto.ToValue(nil)
					return v
				}

				// decode response object to a map
				s := structs.New(rpcResp)
				s.TagName = "json"
				return s.Map()
			}
		}
	}

	for ns, objs := range namespacesObj {
		e.vm.Set(ns, objs)
	}

	// Add some methods to namespaces
	namespacesObj["personal"]["login"] = e.login
	namespacesObj["personal"]["loadAccount"] = e.loadAccount
	namespacesObj["personal"]["loadedAccount"] = e.loadedAccount

	// Add some methods to the suggestions
	suggestions = append(suggestions, prompt.Suggest{Text: "personal.login", Description: "Authenticate the console RPC session"})
	suggestions = append(suggestions, prompt.Suggest{Text: "personal.loadAccount", Description: "Load and set an account as the default"})
	suggestions = append(suggestions, prompt.Suggest{Text: "personal.loadedAccount", Description: "Gets the address of the loaded account"})

	return suggestions, nil
}

func (e *Executor) runScript(file string) {

	fullPath, err := filepath.Abs(file)
	if err != nil {
		panic(e.vm.MakeCustomError("ExecError", err.Error()))
	}

	script, err := ioutil.ReadFile(fullPath)
	if err != nil {
		panic(e.vm.MakeCustomError("ExecError", err.Error()))
	}

	_, err = e.vm.Run(string(script))
	if err != nil {
		panic(e.vm.MakeCustomError("ExecError", err.Error()))
	}
}

// loadAccount loads an account and
// sets it as the default account
func (e *Executor) loadAccount(address, password string) {

	// Get the account from the account manager
	sa, err := e.acctMgr.GetByAddress(address)
	if err != nil {
		panic(e.vm.MakeCustomError("AccountError", err.Error()))
	}

	if err := sa.Decrypt(password); err != nil {
		panic(e.vm.MakeCustomError("AccountError", err.Error()))
	}

	e.coinbase = sa.GetKey()
}

// loadedAccount returns the currently loaded account
func (e *Executor) loadedAccount() string {
	return e.coinbase.Addr()
}

// pp pretty prints a slice of arbitrary objects
func (e *Executor) pp(values ...interface{}) {
	var v interface{} = values
	if len(values) == 1 {
		v = values[0]
	}
	bs, err := prettyjson.Marshal(v)
	if err != nil {
		panic(e.vm.MakeCustomError("PrettyPrintError", err.Error()))
	}
	fmt.Println(string(bs))
}

// OnInput receives inputs and executes
func (e *Executor) OnInput(in string) {

	e.exit = false

	switch in {
	case ".exit":
		e.exitProgram(true)
	case ".help":
		e.help()
	default:

		e.exec(in)
	}
}

func (e *Executor) exitProgram(immediately bool) {
	if !immediately && !e.exit {
		fmt.Println("(To exit, press ^C again or type .exit)")
		e.exit = true
		return
	}
	os.Exit(0)
}

func (e *Executor) exec(in string) {

	// RecoverFunc recovers from panics.
	defer func() {
		if r := recover(); r != nil {
			color.Red("Panic: %s", r)
		}
	}()

	v, err := e.vm.Run(in)
	if err != nil {
		color.Red("%s", err.Error())
		return
	}

	if v.IsNull() || v.IsUndefined() {
		color.Magenta("%s", v)
		return
	}

	vExp, _ := v.Export()
	bs, err := prettyjson.Marshal(vExp)
	fmt.Println(string(bs))
}

func (e *Executor) help() {
	for _, f := range commonFunc {
		fmt.Println(fmt.Sprintf("%s\t\t%s", f[0], f[1]))
	}
}
