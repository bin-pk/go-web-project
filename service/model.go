package service

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"go-web-project/core"
)

type HelloInput struct {
}

func (ptr *HelloInput) Serialize() (bytes.Buffer, *core.CuteError) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	cuteErr := &core.CuteError{}
	if err := enc.Encode(ptr); err != nil {
		cuteErr.SetError(err)
		return buf, cuteErr
	}
	return buf, nil
}
func (ptr *HelloInput) Deserialize(buf bytes.Buffer) *core.CuteError {
	dec := gob.NewDecoder(&buf)

	cuteErr := &core.CuteError{}
	if err := dec.Decode(ptr); err != nil {
		cuteErr.SetError(err)
		return cuteErr
	}
	return nil
}
func (ptr *HelloInput) JsonSerialize() (string, *core.CuteError) {

	jsonBytes, err := json.Marshal(ptr)
	cuteErr := &core.CuteError{}
	if err != nil {
		cuteErr.SetError(err)
		return "", cuteErr
	}
	return string(jsonBytes), nil
}
func (ptr *HelloInput) JsonDeserialize(str string) *core.CuteError {
	cuteErr := &core.CuteError{}
	err := json.Unmarshal([]byte(str), ptr)
	if err != nil {
		cuteErr.SetError(err)
		return cuteErr
	}
	return nil
}

type HelloOutput struct {
	Data string
}

func (ptr *HelloOutput) Serialize() (bytes.Buffer, *core.CuteError) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	cuteErr := &core.CuteError{}
	if err := enc.Encode(ptr); err != nil {
		cuteErr.SetError(err)
		return buf, cuteErr
	}
	return buf, nil
}
func (ptr *HelloOutput) Deserialize(buf bytes.Buffer) *core.CuteError {
	dec := gob.NewDecoder(&buf)

	cuteErr := &core.CuteError{}
	if err := dec.Decode(ptr); err != nil {
		cuteErr.SetError(err)
		return cuteErr
	}
	return nil
}
func (ptr *HelloOutput) JsonSerialize() (string, *core.CuteError) {

	jsonBytes, err := json.Marshal(ptr)
	cuteErr := &core.CuteError{}
	if err != nil {
		cuteErr.SetError(err)
		return "", cuteErr
	}
	return string(jsonBytes), nil
}
func (ptr *HelloOutput) JsonDeserialize(str string) *core.CuteError {
	cuteErr := &core.CuteError{}
	err := json.Unmarshal([]byte(str), ptr)
	if err != nil {
		cuteErr.SetError(err)
		return cuteErr
	}
	return nil
}

type MapSearcherInput struct {
	TableName string
	XMin      float64
	XMax      float64
	YMin      float64
	YMax      float64
}

func (ptr *MapSearcherInput) Serialize() (bytes.Buffer, *core.CuteError) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	cuteErr := &core.CuteError{}
	if err := enc.Encode(ptr); err != nil {
		cuteErr.SetError(err)
		return buf, cuteErr
	}
	return buf, nil
}
func (ptr *MapSearcherInput) Deserialize(buf bytes.Buffer) *core.CuteError {
	dec := gob.NewDecoder(&buf)

	cuteErr := &core.CuteError{}
	if err := dec.Decode(ptr); err != nil {
		cuteErr.SetError(err)
		return cuteErr
	}
	return nil
}
func (ptr *MapSearcherInput) JsonSerialize() (string, *core.CuteError) {

	jsonBytes, err := json.Marshal(ptr)
	cuteErr := &core.CuteError{}
	if err != nil {
		cuteErr.SetError(err)
		return "", cuteErr
	}
	return string(jsonBytes), nil
}
func (ptr *MapSearcherInput) JsonDeserialize(str string) *core.CuteError {
	cuteErr := &core.CuteError{}
	err := json.Unmarshal([]byte(str), ptr)
	if err != nil {
		cuteErr.SetError(err)
		return cuteErr
	}
	return nil
}
func (ptr *MapSearcherInput) GetSQL() string {

	return fmt.Sprintf("select * from %s", ptr.TableName)

}
