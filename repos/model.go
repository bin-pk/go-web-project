package repos

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-web-project/core"
	"net"
	"time"
)

type CuteDuckDB struct {
	db  *sql.DB
	ctx context.Context
}

type CuteRedis struct {
	db      *redis.Client
	ctx     context.Context
	Timeout time.Duration
}

type CuteDBConfig struct {
	DriverName string
	DBFilePath string
	Ipaddr     *net.IPAddr
	Port       uint16
	Username   string
	Password   string
	Timeout    time.Duration
}

func (ptr *CuteDuckDB) Create(config CuteDBConfig) *core.CuteError {
	ctx := context.Context(context.Background())
	db, err := sql.Open(config.DriverName, fmt.Sprintf("%s?access_mode=READ_WRITE&threads=4&memory_limit=1GB", config.DBFilePath))
	cuteErr := &core.CuteError{}
	if err != nil {
		cuteErr.SetError(err)
		return cuteErr
	}
	ptr.db = db
	ptr.ctx = ctx
	return nil
}

func (ptr *CuteDuckDB) Close() *core.CuteError {
	cuteErr := &core.CuteError{}
	if err := ptr.db.Close(); err != nil {
		cuteErr.SetError(err)
		return cuteErr
	}
	return nil
}

func (ptr *CuteDuckDB) Set(input CuteDBParameter) *core.CuteError {
	cuteErr := &core.CuteError{}
	if _, err := ptr.db.QueryContext(ptr.ctx, input.GetSQL()); err != nil {
		cuteErr.SetError(err)
		return cuteErr
	}
	return nil
}
func (ptr *CuteDuckDB) Get(input CuteDBParameter, refData *[]CuteDBOutput, factory func() CuteDBOutput) *core.CuteError {
	cuteErr := &core.CuteError{}
	if rows, err := ptr.db.QueryContext(ptr.ctx, input.GetSQL()); err != nil {
		cuteErr.SetError(err)
		return cuteErr
	} else {
		defer rows.Close()
		columns, err2 := rows.Columns()
		if err2 != nil {
			cuteErr.SetError(err2)
			return cuteErr
		}
		for rows.Next() {
			values := make([]interface{}, len(columns))
			valuePtrs := make([]interface{}, len(columns))
			for i := range values {
				valuePtrs[i] = &values[i]
			}
			if err := rows.Scan(valuePtrs...); err != nil {
				cuteErr.SetError(err)
				return cuteErr
			}
			rowMap := make(map[string]interface{})
			for i, colName := range columns {
				rowMap[colName] = values[i]
			}
			inst := factory()
			inst.SetOutput(rowMap)
			*refData = append(*refData, inst)
		}

		if err3 := rows.Err(); err3 != nil {
			cuteErr.SetError(err)
			return cuteErr
		}

		return nil
	}
}
func (ptr *CuteDuckDB) Delete(input CuteDBParameter) *core.CuteError {
	cuteErr := &core.CuteError{}
	if _, err := ptr.db.QueryContext(ptr.ctx, input.GetSQL()); err != nil {
		cuteErr.SetError(err)
		return cuteErr
	}
	return nil
}

func (ptr *CuteRedis) Create(config CuteDBConfig) *core.CuteError {
	ctx := context.Context(context.Background())

	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Ipaddr.String(),
		Password: config.Password,
		DB:       0,
	})

	if rdb == nil {
		return &core.CuteError{Code: core.InternalError, Message: "redis object is null!!!"}
	}
	ptr.ctx = ctx
	ptr.db = rdb
	ptr.Timeout = config.Timeout
	return nil
}
func (ptr *CuteRedis) Close() *core.CuteError {
	cuteErr := &core.CuteError{}
	if err := ptr.db.Close(); err != nil {
		cuteErr.SetError(err)
		return cuteErr
	}
	return nil
}

func (ptr *CuteRedis) Set(input CuteDBParameter) *core.CuteError {
	_, err := ptr.db.Set(ptr.ctx, input.GetKey(), input.GetValue(), ptr.Timeout).Result()
	cuteErr := &core.CuteError{}
	if err != nil {
		cuteErr.SetError(err)
		return cuteErr
	}
	return nil
}
func (ptr *CuteRedis) Get(input CuteDBParameter, refData *[]CuteDBOutput, factory func() CuteDBOutput) *core.CuteError {
	value, err := ptr.db.Get(ptr.ctx, input.GetKey()).Result()
	cuteErr := &core.CuteError{}
	if err != nil {
		cuteErr.SetError(err)
		return cuteErr
	}
	inst := factory()
	jsonMap := make(map[string]interface{})
	if err := json.Unmarshal([]byte(value), &jsonMap); err == nil {
		inst.SetOutput(jsonMap)
	} else {
		jsonMap[input.GetKey()] = value
		inst.SetOutput(jsonMap)
	}
	*refData = append(*refData, inst)
	return nil
}
func (ptr *CuteRedis) Delete(input CuteDBParameter) *core.CuteError {
	_, err := ptr.db.Del(ptr.ctx, input.GetKey()).Result()
	cuteErr := &core.CuteError{}
	if err != nil {
		cuteErr.SetError(err)
		return cuteErr
	}
	return nil
}
