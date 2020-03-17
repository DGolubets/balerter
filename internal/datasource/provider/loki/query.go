package loki

import (
	"fmt"
	"github.com/prometheus/common/model"
	"github.com/yuin/gluamapper"
	lua "github.com/yuin/gopher-lua"
	"go.uber.org/zap"
	"strings"
)

type queryQueryOptions struct {
	Time string
}

func (m *Loki) doQuery(L *lua.LState) int {
	query := strings.TrimSpace(L.Get(1).String())
	if query == "" {
		L.Push(lua.LNil)
		L.Push(lua.LString("query must be not empty"))
		return 2
	}

	options := L.Get(2)
	queryOptions := queryQueryOptions{}
	if options.Type() == lua.LTTable {
		err := gluamapper.Map(options.(*lua.LTable), &queryOptions)
		if err != nil {
			m.logger.Error("error decode query query options", zap.Error(err))
			L.Push(lua.LNil)
			L.Push(lua.LString("error decode query query options"))
			return 2
		}
	}

	m.logger.Debug("call loki query", zap.String("name", m.name), zap.String("query", query))

	v, err := m.sendQuery(query, queryOptions)
	if err != nil {
		m.logger.Error("error send query to loki", zap.Error(err))
		L.Push(lua.LNil)
		L.Push(lua.LString("error send query to loki: " + err.Error()))
		return 2
	}

	switch v.Type() {
	case model.ValVector:
		vv := v.(model.Vector)

		tbl := &lua.LTable{}
		for _, s := range vv {
			row := &lua.LTable{}
			metrics := &lua.LTable{}
			for key, val := range s.Metric {
				metrics.RawSet(lua.LString(key), lua.LString(val))
			}
			row.RawSet(lua.LString("metrics"), metrics)

			value := &lua.LTable{}
			value.RawSet(lua.LString("timestamp"), lua.LNumber(s.Timestamp.Unix()))
			value.RawSet(lua.LString("value"), lua.LNumber(s.Value))

			row.RawSet(lua.LString("value"), value)
			tbl.Append(row)
		}

		L.Push(tbl)
	default:
		L.Push(lua.LNil)
		return m.luaError(L, fmt.Errorf("unexpected loki model type"))
	}

	L.Push(lua.LNil)

	return 2
}

type queryRangeOptions struct {
	Start string
	End   string
	Step  string
}

func (m *Loki) doRange(L *lua.LState) int {
	query := strings.TrimSpace(L.Get(1).String())
	if query == "" {
		L.Push(lua.LNil)
		L.Push(lua.LString("query must be not empty"))
		return 2
	}

	options := L.Get(2)
	rangeOptions := queryRangeOptions{}
	if options.Type() == lua.LTTable {
		err := gluamapper.Map(options.(*lua.LTable), &rangeOptions)
		if err != nil {
			m.logger.Error("error decode query range options", zap.Error(err))
			L.Push(lua.LNil)
			L.Push(lua.LString("error decode query range options"))
			return 2
		}
	}

	m.logger.Debug("call loki query range", zap.String("name", m.name), zap.String("query", query))

	v, err := m.sendRange(query, rangeOptions)
	if err != nil {
		m.logger.Error("error send query to loki", zap.Error(err))
		L.Push(lua.LNil)
		L.Push(lua.LString("error send query to loki: " + err.Error()))
		return 2
	}

	switch v.Type() {
	case model.ValMatrix:
		vv := v.(model.Matrix)

		tbl := &lua.LTable{}
		for _, s := range vv {
			row := &lua.LTable{}

			metrics := &lua.LTable{}
			for key, val := range s.Metric {
				metrics.RawSet(lua.LString(key), lua.LString(val))
			}

			values := &lua.LTable{}
			for _, val := range s.Values {
				value := &lua.LTable{}
				value.RawSet(lua.LString("timestamp"), lua.LNumber(val.Timestamp.Unix()))
				value.RawSet(lua.LString("value"), lua.LNumber(val.Value))

				values.Append(value)
			}

			row.RawSet(lua.LString("metrics"), metrics)
			row.RawSet(lua.LString("values"), values)
			tbl.Append(row)
		}

		L.Push(tbl)

	case model.ValVector:
		vv := v.(model.Vector)

		tbl := &lua.LTable{}
		for _, s := range vv {
			row := &lua.LTable{}
			metrics := &lua.LTable{}
			for key, val := range s.Metric {
				metrics.RawSet(lua.LString(key), lua.LString(val))
			}
			row.RawSet(lua.LString("metrics"), metrics)
			row.RawSet(lua.LString("value"), lua.LNumber(s.Value))
			tbl.Append(row)
		}

		L.Push(tbl)
	default:
		L.Push(lua.LNil)
		return m.luaError(L, fmt.Errorf("unexpected loki model type"))
	}

	L.Push(lua.LNil)

	return 2
}

func (m *Loki) luaError(L *lua.LState, err error) int {
	m.logger.Debug("query error", zap.Error(err))
	L.Push(lua.LNil)
	L.Push(lua.LString(err.Error()))
	return 2
}