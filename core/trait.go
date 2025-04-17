package core

import (
	"bytes"
	"github.com/google/uuid"
	"sync"
)

type CuteContext interface {
	Get(key uuid.UUID) CuteData
	Set(key uuid.UUID, value CuteData)
	Delete(key uuid.UUID)
}

type CuteData interface {
	Serialize() (bytes.Buffer, *CuteError)
	Deserialize(buf bytes.Buffer) *CuteError
	JsonSerialize() (string, *CuteError)
	JsonDeserialize(str string) *CuteError
}

type CuteError struct {
	Message string
	Code    CuteErrorCode
}

func (c *CuteError) SetError(err error) {
	c.Message = err.Error()
	c.Code = InternalError
}

func (c *CuteError) GetCode() CuteErrorCode {
	return c.Code
}
func (c *CuteError) GetMessage() string {
	return c.Message
}

type CuteCacheContext struct {
	CacheMap map[uuid.UUID]CuteData
	RwLock   sync.RWMutex
}

func CreateCacheContext() *CuteCacheContext {
	return &CuteCacheContext{CacheMap: make(map[uuid.UUID]CuteData)}
}

func (ctx *CuteCacheContext) Get(key uuid.UUID) CuteData {
	ctx.RwLock.RLock()
	defer ctx.RwLock.RUnlock()
	data, exists := ctx.CacheMap[key]
	if exists {
		return data
	} else {
		return nil
	}
}

func (ctx *CuteCacheContext) Set(key uuid.UUID, value CuteData) {
	ctx.RwLock.Lock()
	defer ctx.RwLock.Unlock()
	ctx.CacheMap[key] = value
}

func (ctx *CuteCacheContext) Delete(key uuid.UUID) {
	ctx.RwLock.Lock()
	defer ctx.RwLock.Unlock()
	delete(ctx.CacheMap, key)
}

func GetCuteDataUUID(data CuteData) (uuid.UUID, *CuteError) {
	if data == nil {
		return uuid.Nil, &CuteError{Message: "CuteData is nil", Code: InternalError}
	} else {
		serialize, err := data.Serialize()
		if err != nil {
			return uuid.Nil, err
		}
		id, err2 := uuid.ParseBytes(serialize.Bytes())
		if err2 != nil {
			return uuid.Nil, &CuteError{Message: "serialize is fail", Code: InternalError}
		} else {
			return id, nil
		}
	}
}
