package utils

import (
	"bytes"
	"encoding/base64"
	"encoding/binary"
	"github.com/fernet/fernet-go"
	"github.com/fsnotify/fsnotify"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	uuidV4 "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"github.com/vmihailenco/msgpack/v5"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
	"work.ctyun.cn/git/GoStack/gostone/execption"
	"work.ctyun.cn/git/GoStack/gostone/mapper/project"
	"work.ctyun.cn/git/GoStack/gostone/mapper/role"
)

var keyList []string

type FernetToken struct {
	Encoder *encoding.Encoder
}

func NewFernetToken() Token {
	encoder := unicode.UTF8.NewEncoder()
	return &FernetToken{
		Encoder: encoder,
	}
}

func (f *FernetToken) Sign(claims AuthContext) (string, JSONRFC3339Milli, JSONRFC3339Milli) {
	t := time.Duration(expireTime) * time.Minute
	now := time.Now().UTC()
	claims.IssuedAtZ = JSONRFC3339Milli(now)
	claims.IssuedAt = now.Unix()
	claims.ExpiresAtZ = JSONRFC3339Milli(now.Add(t))
	claims.ExpiresAt = now.Add(t).Unix()
	keys := keyList
	k := fernet.MustDecodeKeys(keys...)
	userId, err := uuid.Parse(claims.UserId)
	if err != nil {
		panic(err)
	}
	projectId, err := uuid.Parse(claims.ProjectId)
	method := convertMethodListToInteger(claims.Method)
	item := Item{
		Version: 2,
		User: UserItem{
			IsStoredAsBytes: true,
			UserId:          userId,
		},
		Methods: method,
		Project: ProjectItem{
			IsStoredAsBytes: true,
			ProjectId:       projectId,
		},
		ExpiresAt: float64(claims.ExpiresAt),
		AuditIds:  buildAuditInfo(),
	}

	v, err := msgpack.Marshal(&item)
	if err != nil {
		panic(err)
	}
	tok, err := fernet.EncryptAndSign(v, k[0])
	if err != nil {
		panic(err)
	}
	return strings.ReplaceAll(string(tok), "=", ""), claims.IssuedAtZ, claims.ExpiresAtZ
}

func (f *FernetToken) SignByExpiration(claims AuthContext, expiration int64) (string, JSONRFC3339Milli, JSONRFC3339Milli) {
	t := time.Duration(expiration) * time.Minute
	now := time.Now().UTC()
	claims.IssuedAtZ = JSONRFC3339Milli(now)
	claims.IssuedAt = now.Unix()
	claims.ExpiresAtZ = JSONRFC3339Milli(now.Add(t))
	claims.ExpiresAt = now.Add(t).Unix()
	keys := keyList
	k := fernet.MustDecodeKeys(keys...)
	userId, err := uuid.Parse(claims.UserId)
	if err != nil {
		panic(err)
	}
	projectId, err := uuid.Parse(claims.ProjectId)
	method := convertMethodListToInteger(claims.Method)
	item := Item{
		Version: 2,
		User: UserItem{
			IsStoredAsBytes: true,
			UserId:          userId,
		},
		Methods: method,
		Project: ProjectItem{
			IsStoredAsBytes: true,
			ProjectId:       projectId,
		},
		ExpiresAt: float64(claims.ExpiresAt),
		AuditIds:  buildAuditInfo(),
	}

	v, err := msgpack.Marshal(&item)
	if err != nil {
		panic(err)
	}
	tok, err := fernet.EncryptAndSign(v, k[0])
	if err != nil {
		panic(err)
	}
	return strings.ReplaceAll(string(tok), "=", ""), claims.IssuedAtZ, claims.ExpiresAtZ

}

const (
	TIMESTAMP_START = 1
	TIMESTAMP_END   = 9
)

func (f *FernetToken) Name() string {
	return "fernet"
}

func (f *FernetToken) Validate(token string) *AuthContext {
	keys := keyList
	k := fernet.MustDecodeKeys(keys...)
	token = restoreToken(token)
	tokenByte, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		panic(err)
	}
	timestampBytes := tokenByte[TIMESTAMP_START:TIMESTAMP_END]
	timestampInt := BytesToInt(timestampBytes)
	msg := fernet.VerifyAndDecrypt([]byte(token), time.Duration(expireTime)*time.Minute, k)
	var data Item
	err = msgpack.Unmarshal(msg, &data)
	if err != nil {
		panic(execption.NewGoStoneError(execption.StatusUnauthorized, "validate token failed"))
	}
	var (
		userId    string
		projectId string
		expiresAt time.Time
		issueAt   time.Time
		methods   []string
	)
	if data.User.IsStoredAsBytes {
		userId = FormatUUIDFromString(data.User.UserId.String())
	}
	if data.Project.IsStoredAsBytes {
		projectId = FormatUUIDFromString(data.Project.ProjectId.String())
	}
	for _, audit := range data.AuditIds {
		id := base64.URLEncoding.EncodeToString(audit)
		id = strings.ReplaceAll(id, "=", "")
	}
	expiresAt = time.Unix(int64(data.ExpiresAt), 0)
	issueAt = time.Unix(int64(timestampInt), 0)
	methods = ConvertIntegerToMethodList(data.Methods)
	p := project.FindProjectById(projectId, http.StatusUnauthorized)
	return &AuthContext{
		UserId:     userId,
		ProjectId:  projectId,
		ExpiresAtZ: JSONRFC3339Milli(expiresAt),
		IssuedAtZ:  JSONRFC3339Milli(issueAt),
		Method:     methods,
		Role:       role.FindRoleNameByUserIdAndProjectId(userId, projectId, http.StatusUnauthorized),
		DomainId:   p.DomainId,
	}
}

func (f *FernetToken) GetAuthContext(ctx echo.Context) AuthContext {
	c := ctx.Get("user").(*AuthContext)
	return *c
}

const INT_MAX = int(^uint(0) >> 1)

func (f *FernetToken) ValidateCanExpired(token string) *AuthContext {
	keys := keyList
	k := fernet.MustDecodeKeys(keys...)
	token = restoreToken(token)
	tokenByte, err := base64.URLEncoding.DecodeString(token)
	if err != nil {
		panic(err)
	}
	timestampBytes := tokenByte[TIMESTAMP_START:TIMESTAMP_END]
	timestampInt := BytesToInt(timestampBytes)
	msg := fernet.VerifyAndDecrypt([]byte(token), 876000*time.Hour, k)
	var data Item
	err = msgpack.Unmarshal(msg, &data)
	if err != nil {
		panic(execption.NewGoStoneError(execption.StatusUnauthorized, "validate token failed"))
	}
	var (
		userId    string
		projectId string
		expiresAt time.Time
		issueAt   time.Time
		methods   []string
	)
	if data.User.IsStoredAsBytes {
		userId = FormatUUIDFromString(data.User.UserId.String())
	}
	if data.Project.IsStoredAsBytes {
		projectId = FormatUUIDFromString(data.Project.ProjectId.String())
	}
	for _, audit := range data.AuditIds {
		id := base64.URLEncoding.EncodeToString(audit)
		id = strings.ReplaceAll(id, "=", "")
	}
	expiresAt = time.Unix(int64(data.ExpiresAt), 0)
	issueAt = time.Unix(int64(timestampInt), 0)
	methods = ConvertIntegerToMethodList(data.Methods)
	p := project.FindProjectById(projectId, http.StatusUnauthorized)
	return &AuthContext{
		UserId:     userId,
		ProjectId:  projectId,
		ExpiresAtZ: JSONRFC3339Milli(expiresAt),
		IssuedAtZ:  JSONRFC3339Milli(issueAt),
		Method:     methods,
		Role:       role.FindRoleNameByUserIdAndProjectId(userId, projectId, http.StatusUnauthorized),
		DomainId:   p.DomainId,
	}
}

func BytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}

type Item struct {
	_msgpack  struct{} `msgpack:",asArray"`
	Version   int
	User      UserItem
	Methods   int
	Project   ProjectItem
	ExpiresAt float64
	AuditIds  [][]byte
}

type UserItem struct {
	_msgpack        struct{} `msgpack:",asArray"`
	IsStoredAsBytes bool
	UserId          uuid.UUID
}

type ProjectItem struct {
	_msgpack        struct{} `msgpack:",asArray"`
	IsStoredAsBytes bool
	ProjectId       uuid.UUID
}

func restoreToken(token string) string {
	modReturned := len(token) % 4
	if modReturned != 0 {
		missingPadding := 4 - modReturned
		for i := 0; i < missingPadding; i++ {
			token = token + "="
		}
	}
	return token
}

func loadKey() {
	files, err := ioutil.ReadDir(fernetPath)
	if err != nil {
		log.Warnf("fernetKey Get failed error:[%s]", err)
		return
	}
	value := make(map[int]string)
	var keys []int
	for _, f := range files {
		key, err := strconv.Atoi(f.Name())
		if err != nil {
			continue
		}
		secret, err := ioutil.ReadFile(fernetPath + "/" + f.Name())
		if err != nil {
			panic(execption.NewGoStoneError(execption.StatusUnauthorized, "fernet secret load error"))
		}
		secretStr := string(secret)
		secretStr = strings.ReplaceAll(secretStr, "#", "")
		value[key] = secretStr
		keys = append(keys, key)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	var results []string
	for _, k := range keys {
		results = append(results, value[k])
	}
	keyList = results
}

func watchKey() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Debug("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Debug("modified file:", event.Name)
				}
				loadKey()
			case err := <-watcher.Errors:
				log.Error("error:", err)
			}
		}
	}()

	err = watcher.Add(fernetPath)
	if err != nil {
		log.Warnf("fernet path not exists: [%s]", err)
	}
}

func ConvertIntegerToMethodList(method int) (methods []string) {
	if method == 0 {
		return
	}
	methodMap := ConstructMethodMapFormConfig()
	sortMap := sortMap(methodMap)
	var confirmMethods []int
	for _, pair := range sortMap {
		if method/pair.Key == 1 {
			confirmMethods = append(confirmMethods, pair.Key)
			method = method - pair.Key
		}
	}
	for _, index := range confirmMethods {
		methods = append(methods, methodMap[index])
	}
	return
}

func ConstructMethodMapFormConfig() map[int]string {
	methodMap := make(map[int]string)
	defaultMethodMap := []string{"external", "password", "token", "oauth1", "mapped", "application_credential"}
	methodIndex := 1
	for _, m := range defaultMethodMap {
		methodMap[methodIndex] = m
		methodIndex = methodIndex * 2
	}
	return methodMap
}

type Pair struct {
	Key   int
	Value string
}

func sortMap(m map[int]string) []Pair {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	var pairs []Pair
	for _, k := range keys {
		pairs = append(pairs, Pair{
			Key:   k,
			Value: m[k],
		})
	}
	return pairs
}

func convertMethodListToInteger(methods []string) int {
	methodMap := ConstructMethodMapFormConfig()
	var methodInts = 0
	for _, method := range methods {
		for k, v := range methodMap {
			if v == method {
				methodInts = methodInts + k
			}
		}
	}
	return methodInts
}

func buildAuditInfo() [][]byte {
	var result []byte
	id := uuidV4.NewV4()
	buf := new(bytes.Buffer)
	enc := base64.NewEncoder(base64.URLEncoding, buf)
	enc.Write(id.Bytes())
	enc.Close()
	result = buf.Bytes()
	result = result[:len(result)-2]
	return [][]byte{result}
}
