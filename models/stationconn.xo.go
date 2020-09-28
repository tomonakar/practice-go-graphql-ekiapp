// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"database/sql/driver"
	"encoding/csv"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// StationConn represents a row from '[custom station_conn]'.
type StationConn struct {
	LineName string // line_name
	LineNameH string // line_name_h
	LineCd int // line_cd
	StationCd int // station_cd
	StationGCd int // station_g_cd
	Address string // address
	StationName string // station_name
	BeforeLineName string // before_line_name
	BeforeStationCd int // before_station_cd
	BeforeStationName string // before_station_name
	BeforeAddress string // before_address
	AfterLineName string // after_line_name
	AfterStationCd int // after_station_cd
	AfterStationName string // after_station_name
	AfterAddress string // after_address
	TransferLineName string // transfer_line_name
	TransferStationCd int // transfer_station_cd
	TransferStationName string // transfer_station_name
	TransferAddress string // transfer_address
}

// StationConnsByStationCD runs a custom query, returning results as StationConn.
func StationConnsByStationCD (db XODB, stationCD int) ([]*StationConn, error) {
	var err error

	// sql query
	const sqlstr = `select li.line_name, ` +
	`li.line_name_h, ` +
	`li.line_cd, ` +
	`st.station_cd, ` +
	`st.station_g_cd, ` +
	`st.address, ` +
	`st.station_name, ` +
	`COALESCE(s2l.line_name, '')     as before_line_name, ` +
	`COALESCE(st2.station_cd, 0)    as before_station_cd, ` +
	`COALESCE(st2.station_name, '') as before_station_name, ` +
	`COALESCE(st2.address, '')      as before_address, ` +
	`COALESCE(s3l.line_name, '')     as after_line_name, ` +
	`COALESCE(st3.station_cd, 0)    as after_station_cd, ` +
	`COALESCE(st3.station_name, '') as after_station_name, ` +
	`COALESCE(st2.address, '')      as after_address, ` +
	`COALESCE(gli.line_name, '')    as transfer_line_name, ` +
	`COALESCE(gs.station_cd, 0)     as transfer_station_cd, ` +
	`COALESCE(gs.station_name, '')  as transfer_station_name, ` +
	`COALESCE(gs.address, '')       as transfer_address ` +
	`from station st ` +
	`inner join line li on st.line_cd = li.line_cd ` +
	`left outer join station_join sjb on st.line_cd = sjb.line_cd and st.station_cd = sjb.station_cd2 ` +
	`left outer join station_join sja on st.line_cd = sja.line_cd and st.station_cd = sja.station_cd1 ` +
	`left outer join station st2 on sjb.line_cd = st2.line_cd and sjb.station_cd1 = st2.station_cd ` +
	`left outer join line s2l on st2.line_cd = s2l.line_cd ` +
	`left outer join station st3 on sja.line_cd = st3.line_cd and sja.station_cd2 = st3.station_cd ` +
	`left outer join line s3l on st3.line_cd = s3l.line_cd ` +
	`left outer join station gs on st.station_g_cd = gs.station_g_cd and st.station_cd <> gs.station_cd ` +
	`left outer join line gli on gs.line_cd = gli.line_cd ` +
	`where st.station_cd = $1 ` +
	`and st.e_status = 0 ` +
	`order by st.e_sort`

	// run query
	XOLog(sqlstr, stationCD)
	q, err := db.Query(sqlstr, stationCD)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*StationConn{}
	for q.Next() {
		sc := StationConn{}

		// scan
		err = q.Scan(&sc.LineName, &sc.LineNameH, &sc.LineCd, &sc.StationCd, &sc.StationGCd, &sc.Address, &sc.StationName, &sc.BeforeLineName, &sc.BeforeStationCd, &sc.BeforeStationName, &sc.BeforeAddress, &sc.AfterLineName, &sc.AfterStationCd, &sc.AfterStationName, &sc.AfterAddress, &sc.TransferLineName, &sc.TransferStationCd, &sc.TransferStationName, &sc.TransferAddress)
		if err != nil {
			return nil, err
		}

		res = append(res, &sc)
	}

	return res, nil
}
