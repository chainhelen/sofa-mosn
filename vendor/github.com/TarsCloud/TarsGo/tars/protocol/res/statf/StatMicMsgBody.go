//Package statf comment
// This file war generated by tars2go 1.1
// Generated from StatF.tars
package statf

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars/protocol/codec"
)

//StatMicMsgBody strcut implement
type StatMicMsgBody struct {
	Count         int32           `json:"count"`
	TimeoutCount  int32           `json:"timeoutCount"`
	ExecCount     int32           `json:"execCount"`
	IntervalCount map[int32]int32 `json:"intervalCount"`
	TotalRspTime  int64           `json:"totalRspTime"`
	MaxRspTime    int32           `json:"maxRspTime"`
	MinRspTime    int32           `json:"minRspTime"`
}

func (st *StatMicMsgBody) resetDefault() {
}

//ReadFrom reads  from _is and put into struct.
func (st *StatMicMsgBody) ReadFrom(_is *codec.Reader) error {
	var err error
	var length int32
	var have bool
	var ty byte
	st.resetDefault()

	err = _is.Read_int32(&st.Count, 0, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.TimeoutCount, 1, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.ExecCount, 2, true)
	if err != nil {
		return err
	}

	err, have = _is.SkipTo(codec.MAP, 3, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&length, 0, true)
	if err != nil {
		return err
	}
	st.IntervalCount = make(map[int32]int32)
	for i0, e0 := int32(0), length; i0 < e0; i0++ {
		var k0 int32
		var v0 int32

		err = _is.Read_int32(&k0, 0, false)
		if err != nil {
			return err
		}

		err = _is.Read_int32(&v0, 1, false)
		if err != nil {
			return err
		}

		st.IntervalCount[k0] = v0
	}

	err = _is.Read_int64(&st.TotalRspTime, 4, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.MaxRspTime, 5, true)
	if err != nil {
		return err
	}

	err = _is.Read_int32(&st.MinRspTime, 6, true)
	if err != nil {
		return err
	}

	_ = length
	_ = have
	_ = ty
	return nil
}

//ReadBlock reads struct from the given tag , require or optional.
func (st *StatMicMsgBody) ReadBlock(_is *codec.Reader, tag byte, require bool) error {
	var err error
	var have bool
	st.resetDefault()

	err, have = _is.SkipTo(codec.STRUCT_BEGIN, tag, require)
	if err != nil {
		return err
	}
	if !have {
		if require {
			return fmt.Errorf("require StatMicMsgBody, but not exist. tag %d", tag)
		}
		return nil

	}

	st.ReadFrom(_is)

	err = _is.SkipToStructEnd()
	if err != nil {
		return err
	}
	_ = have
	return nil
}

//WriteTo encode struct to buffer
func (st *StatMicMsgBody) WriteTo(_os *codec.Buffer) error {
	var err error

	err = _os.Write_int32(st.Count, 0)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.TimeoutCount, 1)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.ExecCount, 2)
	if err != nil {
		return err
	}

	err = _os.WriteHead(codec.MAP, 3)
	if err != nil {
		return err
	}
	err = _os.Write_int32(int32(len(st.IntervalCount)), 0)
	if err != nil {
		return err
	}
	for k1, v1 := range st.IntervalCount {

		err = _os.Write_int32(k1, 0)
		if err != nil {
			return err
		}

		err = _os.Write_int32(v1, 1)
		if err != nil {
			return err
		}
	}

	err = _os.Write_int64(st.TotalRspTime, 4)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.MaxRspTime, 5)
	if err != nil {
		return err
	}

	err = _os.Write_int32(st.MinRspTime, 6)
	if err != nil {
		return err
	}

	return nil
}

//WriteBlock encode struct
func (st *StatMicMsgBody) WriteBlock(_os *codec.Buffer, tag byte) error {
	var err error
	err = _os.WriteHead(codec.STRUCT_BEGIN, tag)
	if err != nil {
		return err
	}

	st.WriteTo(_os)

	err = _os.WriteHead(codec.STRUCT_END, 0)
	if err != nil {
		return err
	}
	return nil
}
