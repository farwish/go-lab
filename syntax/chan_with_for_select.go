package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func init() {
	LoadConfig()
}

/**
 * waitForServer example.
 */
func main() {
	if err := checkServerHeartbeat(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("cross check successful!")
	}
}
// ollama 一些子命令在 cobra.Command 的 PreRunE 中指定检测、当返回错误时会阻止程序继续运行。
func checkServerHeartbeat() error {
	client, _ := ClientFromEnvironment()
	var ctx context.Context

	if err := startApp(ctx, client); err != nil {
		fmt.Println(err.Error())
		return fmt.Errorf("could not connect to ollama app, is it running?")
	}

	return nil
}
func startApp(ctx context.Context, client *Client) error {
	// 执行自己的启动服务的命令
	// ....

	// wait for server to start
	return waitForServer(ctx, client)
}
// 重点：
// 该函数使用 context.Context 类型参数 ctx 和 Client 类型的指针参数 client，用于等待服务器启动。
// 它设置了一个 5 秒的超时时间，并每 0.5 秒检查一次服务器是否已启动。
// 如果超时，则返回一个错误，表示等待服务器启动超时。
// 如果服务器已启动，则返回 nil。
func waitForServer(ctx context.Context, client *Client) error {
	// wait for the server to start
	timeout := time.After(5 * time.Second)
	tick := time.Tick(500 * time.Millisecond)
	for {
		select {
		case <-timeout:
			return errors.New("timed out waiting for server to start")
		case <-tick:
			if err := client.Heartbeat(ctx); err == nil {
				return nil // server has started
			}
		}
	}
}

type OllamaHost struct {
	Scheme string
	Host   string
	Port   string
}

var (
	// Set via OLLAMA_ORIGINS in the environment
	AllowOrigins []string
	// Set via OLLAMA_DEBUG in the environment
	Debug bool
	// Set via OLLAMA_HOST in the environment
	Host *OllamaHost
	// Others...
)
var Version string = "0.2.1"
var ErrInvalidHostPort = errors.New("invalid port specified in OLLAMA_HOST")
var defaultAllowOrigins = []string{
	"localhost",
	"127.0.0.1",
	"0.0.0.0",
}

func LoadConfig() {
	if debug := clean("OLLAMA_DEBUG"); debug != "" {
		d, err := strconv.ParseBool(debug)
		if err == nil {
			Debug = d
		} else {
			Debug = true
		}
	}

	if origins := clean("OLLAMA_ORIGINS"); origins != "" {
		AllowOrigins = strings.Split(origins, ",")
	}
	for _, allowOrigin := range defaultAllowOrigins {
		AllowOrigins = append(AllowOrigins,
			fmt.Sprintf("http://%s", allowOrigin),
			fmt.Sprintf("https://%s", allowOrigin),
			fmt.Sprintf("http://%s", net.JoinHostPort(allowOrigin, "*")),
			fmt.Sprintf("https://%s", net.JoinHostPort(allowOrigin, "*")),
		)
	}
	AllowOrigins = append(AllowOrigins,
		"app://*",
		"file://*",
		"tauri://*",
	)

	var err error
	Host, err = getOllamaHost()
	if err != nil {
		slog.Error("invalid setting", "OLLAMA_HOST", Host, "error", err, "using default port", Host.Port)
	}
}
func ClientFromEnvironment() (*Client, error) {
	Host, _ = getOllamaHost()
	ollamaHost := Host
	return &Client{
		base: &url.URL{
			Scheme: ollamaHost.Scheme,
			Host:   net.JoinHostPort(ollamaHost.Host, ollamaHost.Port),
		},
		http: http.DefaultClient,
	}, nil
}
func getOllamaHost() (*OllamaHost, error) {
	defaultPort := "11434"

	hostVar := os.Getenv("OLLAMA_HOST")
	hostVar = strings.TrimSpace(strings.Trim(strings.TrimSpace(hostVar), "\"'"))

	scheme, hostport, ok := strings.Cut(hostVar, "://")
	switch {
	case !ok:
		scheme, hostport = "http", hostVar
	case scheme == "http":
		defaultPort = "80"
	case scheme == "https":
		defaultPort = "443"
	}

	// trim trailing slashes
	hostport = strings.TrimRight(hostport, "/")

	host, port, err := net.SplitHostPort(hostport)
	if err != nil {
		host, port = "127.0.0.1", defaultPort
		if ip := net.ParseIP(strings.Trim(hostport, "[]")); ip != nil {
			host = ip.String()
		} else if hostport != "" {
			host = hostport
		}
	}

	if portNum, err := strconv.ParseInt(port, 10, 32); err != nil || portNum > 65535 || portNum < 0 {
		return &OllamaHost{
			Scheme: scheme,
			Host:   host,
			Port:   defaultPort,
		}, ErrInvalidHostPort
	}

	return &OllamaHost{
		Scheme: scheme,
		Host:   host,
		Port:   port,
	}, nil
}
// Clean quotes and spaces from the value
func clean(key string) string {
	return strings.Trim(os.Getenv(key), "\"' ")
}

type Client struct {
	base *url.URL
	http *http.Client
}

// Hearbeat checks if the server has started and is responsive; if yes, it
// returns nil, otherwise an error.
func (c *Client) Heartbeat(ctx context.Context) error {
	if err := c.do(ctx, http.MethodHead, "/", nil, nil); err != nil {
		return err
	}
	return nil
}
func (c *Client) do(ctx context.Context, method, path string, reqData, respData any) error {
	var reqBody io.Reader
	var data []byte
	var err error

	switch reqData := reqData.(type) {
	case io.Reader:
		// reqData is already an io.Reader
		reqBody = reqData
	case nil:
		// noop
	default:
		data, err = json.Marshal(reqData)
		if err != nil {
			return err
		}

		reqBody = bytes.NewReader(data)
	}

	requestURL := c.base.JoinPath(path)
	fmt.Println(requestURL.String())
	request, err := http.NewRequestWithContext(ctx, method, requestURL.String(), reqBody)
	if err != nil {
		return err
	}

	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("User-Agent", fmt.Sprintf("ollama/%s (%s %s) Go/%s", Version, runtime.GOARCH, runtime.GOOS, runtime.Version()))

	respObj, err := c.http.Do(request)
	if err != nil {
		return err
	}
	defer respObj.Body.Close()

	respBody, err := io.ReadAll(respObj.Body)
	if err != nil {
		return err
	}

	if err := checkError(respObj, respBody); err != nil {
		return err
	}

	if len(respBody) > 0 && respData != nil {
		if err := json.Unmarshal(respBody, respData); err != nil {
			return err
		}
	}
	return nil
}

func checkError(resp *http.Response, body []byte) error {
	if resp.StatusCode < http.StatusBadRequest {
		return nil
	}

	apiError := StatusError{StatusCode: resp.StatusCode}

	err := json.Unmarshal(body, &apiError)
	if err != nil {
		// Use the full body as the message if we fail to decode a response.
		apiError.ErrorMessage = string(body)
	}

	return apiError
}

type StatusError struct {
	StatusCode   int
	Status       string
	ErrorMessage string `json:"error"`
}
func (e StatusError) Error() string {
	switch {
	case e.Status != "" && e.ErrorMessage != "":
		return fmt.Sprintf("%s: %s", e.Status, e.ErrorMessage)
	case e.Status != "":
		return e.Status
	case e.ErrorMessage != "":
		return e.ErrorMessage
	default:
		// this should not happen
		return "something went wrong, please see the ollama server logs for details"
	}
}