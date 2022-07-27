package main

import (
	"bytes"
	"encoding/binary"
	"errors"
	"time"

	"github.com/lovelacelee/clsgo/pkg/config"
	"github.com/lovelacelee/clsgo/pkg/log"
	"github.com/lovelacelee/clsgo/pkg/net"
)

var (
	ErrOk            = errors.New("ok")
	ErrBadCredential = errors.New("bad credentials, username or password invalid")
	ErrUndefined     = errors.New("undefined error")
	ErrBusy          = errors.New("busy")
	ErrParam         = errors.New("wrong parameters")
	ErrFormat        = errors.New("invalid message format")
	ErrInternal      = errors.New("internal system error")
	ErrUnkownID      = errors.New("unkown ID")
	ErrConnectLimit  = errors.New("connection count exceeded")
	ErrUnauthorized  = errors.New("unauthorized")
	ErrNotSupported  = errors.New("not supported")
	ErrNotExist      = errors.New("not exist")
)

type HMERR struct {
	Code int
	Msg  error
}

var (
	HMERR_OK            HMERR = HMERR{Code: 0x00000000, Msg: ErrOk}
	HMERR_Undefined     HMERR = HMERR{Code: 0x00000001, Msg: ErrUndefined}
	HMERR_Busy          HMERR = HMERR{Code: 0x00000002, Msg: ErrBusy}
	HMERR_Param         HMERR = HMERR{Code: 0x00000003, Msg: ErrParam}
	HMERR_Format        HMERR = HMERR{Code: 0x00000004, Msg: ErrFormat}
	HMERR_Internal      HMERR = HMERR{Code: 0x00000005, Msg: ErrInternal}
	HMERR_UnkownID      HMERR = HMERR{Code: 0x00000006, Msg: ErrUnkownID}
	HMERR_ConnectLimit  HMERR = HMERR{Code: 0x00000007, Msg: ErrConnectLimit}
	HMERR_Unauthorized  HMERR = HMERR{Code: 0x00000008, Msg: ErrUnauthorized}
	HMERR_BadCredential HMERR = HMERR{Code: 0x00000009, Msg: ErrBadCredential}
	HMERR_NotSupported  HMERR = HMERR{Code: 0x0000000A, Msg: ErrNotSupported}
	HMERR_NotExist      HMERR = HMERR{Code: 0x0000000B, Msg: ErrNotExist}
)

const (
	HMID_VideoReq            = 0x00000101 // Video
	HMID_VideoData           = 0x00000102
	HMID_VideoClose          = 0x00000103
	HMID_VideoIFrame         = 0x00000104
	HMID_AudioReq            = 0x00000201 // Audio
	HMID_AudioData           = 0x00000202
	HMID_AudioClose          = 0x00000203
	HMID_VoiceReq            = 0x00000301 // Voice
	HMID_VoiceData           = 0x00000302
	HMID_VoiceClose          = 0x00000303
	HMID_PTZ                 = 0x00000401 // PTZ
	HMID_Postion3D           = 0x00000402
	HMID_ParamSet            = 0x00000501 // Param & Status
	HMID_ParamGet            = 0x00000502
	HMID_CommParamSet        = 0x00000503
	HMID_CommParamGet        = 0x00000504
	HMID_StatusEnable        = 0x00000505
	HMID_StatusNotificatioin = 0x00000506
	HMID_PeopleCount         = 0x0000060B // Human detection
	HMID_AlarmPush           = 0x0000060C
	HMID_Login               = 0x0000060D // Login
	HMID_Reboot              = 0x0000060E
	HMID_Disconnect          = 0x0000060F
	HMID_TimeSync            = 0x00000610
	HMID_InConnectionAuth    = 0x00000611
	HMID_NVS301              = 0x00000612
	HMID_ReplayQuery         = 0x00000701 // Replay
	HMID_ReplayReq           = 0x00000702
	HMID_ReplayData          = 0x00000703
	HMID_ReplayEnd           = 0x00000704
	HMID_ReplayStop          = 0x00000705
	HMID_ReplayDelete        = 0x00000706
	HMID_ReplayPause         = 0x00000707
	HMID_ReplayAssume        = 0x00000708
	HMID_ReplayStep          = 0x00000709
	HMID_ReplaySpeed         = 0x00000710
	HMID_SnapshotManual      = 0x00000801 // Snapshot
	HMID_SnapshotQuery       = 0x00000802
	HMID_SnapshotDownload    = 0x00000803
	HMID_SnapshotData        = 0x80000804
	HMID_SnapshotCancel      = 0x00000805
	HMID_SnapshotDelete      = 0x00000806
	HMID_SnapshotCGI         = 0x00000807
	HMID_SnapshotFile        = 0x00000F01
	HMID_RecordDownload      = 0x00001001 // Record download
	HMID_RecordData          = 0x00001002
	HMID_RecordCancel        = 0x00001003
	HMID_WifiQuery           = 0x00001101 // Wireless network
	HMID_UpnpDectection      = 0x00001102
	HMID_StorageFormat       = 0x00001103
	HMID_Upgrade             = 0x00001104
	HMID_UpgradeData         = 0x00001105
	HMID_UpgradeCancel       = 0x00001106
	HMID_UpgradeResult       = 0x00001107
	HMID_FactoryReset        = 0x00001108
	HMID_SysInfo             = 0x00001109
	HMID_AlarmDeploy         = 0x00001201 // Alarm
	HMID_AlarmCancelDeploy   = 0x00001202
	HMID_AlarmArea           = 0x00001203
	HMID_AlarmAddArea        = 0x00001204
	HMID_AlarmDelArea        = 0x00001205
	HMID_AlarmUpdateArea     = 0x00001206
	HMID_AlarmAddSensor      = 0x00001207
	HMID_AlarmDelSensor      = 0x00001208
	HMID_AlarmUpdateSensor   = 0x00001209
	HMID_StartPair           = 0x00001210
	HMID_StopPair            = 0x00001211
	HMID_GetPairInfo         = 0x00001212
	HMID_AlarmPushEnable     = 0x00001213
	HMID_AlarmVoiceEnable    = 0x00001214
	HMID_AlarmVoiceDisable   = 0x00001215
	HMID_FacePushDisable     = 0x00001216 // Face infomation push switch
	HMID_Locker              = 0x00001301 // Lock or unlock
	HMID_PrivacyEnable       = 0x00001303
	HMID_UpgradeInfoQeury    = 0x00001401 // Upgrade infomation query
	HMID_StartUpgrade        = 0x00001402
	HMID_StopUpgrade         = 0x00001403
	HMID_UpgradeProgress     = 0x00001404
	HMID_HeartBeat           = 0x00000A01
	HMID_OnlineUsers         = 0x00000B01 // Who is logon web interface
	HMID_UploadReq           = 0x0000FE01
	HMID_UploadData          = 0x0000FE02
	HMID_UploadStop          = 0x0000FE03
	HMID_UploadError         = 0x0000FE04
	HMID_Extend              = 0x0000FF02
	HMID_FaceXml             = 0x0000FF03
	HMID_FaceBinary          = 0x0000FF04
	HMID_VechicleBase64      = 0x0000FF05
	HMID_HumanBodyBase64     = 0x0000FF06
)

type HMHeadV1 struct {
	ID  uint32
	Len uint32
	Err uint32
}

type HMHeadV2 struct {
	ID  uint32
	Len uint32
	Err uint32
	Sid uint32
}

type HMHead struct {
	V1  HMHeadV1
	V2  HMHeadV2
	len int
}

type HMClient struct {
	SN            string
	Remote        string
	LastKeepAlive time.Time
	Uptime        time.Duration
	BootTime      time.Time
	Authorized    bool
}

// Every client has separated protocol context, so protocol method should
// not be implemented using *HMProtocol
type HMProtocol struct {
	Head HMHead
	Err  HMERR
	Cli  HMClient
}

func (h HMHead) Id() uint32 {
	if h.len == 12 {
		return h.V1.ID
	} else {
		return h.V2.ID
	}
}

func (h HMHead) AckBufferWithHead(bodylen int) (*bytes.Buffer, error) {
	var err error = nil
	buffer := bytes.NewBufferString("")

	if h.len == 12 {
		h.V1.ID |= 0x80000000 // Mark ack
		h.V1.Len = (uint32)(bodylen)
		err = binary.Write(buffer, binary.BigEndian, &h.V1)
	} else {
		h.V2.ID |= 0x80000000 // Mark ack
		h.V2.Len = (uint32)(bodylen)
		err = binary.Write(buffer, binary.BigEndian, &h.V2)
	}
	return buffer, err
}

var timeout = net.Retry{Count: 3, Interval: 10}

func (p *HMProtocol) onLogin(conn *net.Conn, data []byte) ([]byte, error) {
	xv, err := config.XmlDecode(data)

	if err != nil {
		e := &p.Err
		*e = HMERR_Format
		log.Errorf("Xml decode failed: %s %s", err, p.Err.Msg.Error())
		return nil, err
	}
	msgReq := xv["Message"].(map[string]interface{})
	p.Cli.SN = msgReq["Sn"].(string)
	p.Cli.Authorized = true

	ackMsg := make(map[string]interface{})
	ackMsg["UserName"] = "guest"
	ackMsg["UserType"] = "0"
	ackMsg["Ver"] = "2.0"
	ackMsg["Right"] = "65535"
	ackBytes, err := config.XmlEncodeWithIndent(ackMsg, "Message")
	log.Debugf("Ack:\n%s", string(ackBytes))
	return ackBytes, err
}

func (p *HMProtocol) onHeatBeat(conn *net.Conn, data []byte) ([]byte, error) {
	ackMsg := make(map[string]interface{})
	ackMsg["Time"] = time.Now().Unix()
	ackBytes, err := config.XmlEncodeWithIndent(ackMsg, "Message")

	log.Debugf("Ack:\n%s", string(ackBytes))
	return ackBytes, err
}

// Return response message head & body in bytes
// Return nil while error occurred
func (p *HMProtocol) msgHandler(conn *net.Conn, data []byte) ([]byte, error) {
	var err error = nil
	var dataAck []byte
	switch p.Head.Id() {
	case HMID_Login:
		dataAck, err = p.onLogin(conn, data)
	case HMID_HeartBeat:
		dataAck, err = p.onHeatBeat(conn, data)
	default:
		err = HMERR_NotSupported.Msg
	}
	if err != nil {
		return nil, err
	}
	//Write head first
	buffer, err := p.Head.AckBufferWithHead(len(dataAck))
	if err != nil {
		return nil, err
	}
	// Write body if exist
	if len(dataAck) > 0 {
		buffer.Write(dataAck)
	}
	return buffer.Bytes(), err
}

func (p *HMProtocol) OnHead(conn *net.Conn, headlen int) ([]byte, error) {
	bodylen := 0
	msgid := 0
	if p.Head.len == 12 {
		bodylen = int(p.Head.V1.Len)
		msgid = int(p.Head.V1.ID)
	} else {
		bodylen = int(p.Head.V2.Len)
		msgid = int(p.Head.V2.ID)
	}
	log.Infof("%s [%d]0x%08X len:%d", p.Cli.Remote, p.Head.len, msgid, bodylen)
	// message without body
	if bodylen == 0 {
		return p.msgHandler(conn, nil)
	}
	return p.OnBody(conn, bodylen)
}

func (p *HMProtocol) OnBody(conn *net.Conn, bodylen int) ([]byte, error) {
	// message body receive
	data, err := conn.Recv(bodylen, timeout)
	if err != nil {
		return nil, err
	}

	log.Debugf("%s", string(data))
	return p.msgHandler(conn, data)
}

func (p *HMProtocol) HandleMessage(conn *net.Conn) ([]byte, error) {
	if p.Cli.BootTime.IsZero() {
		p.Cli.BootTime = time.Now()
		p.Cli.LastKeepAlive = p.Cli.BootTime
		p.Cli.Remote = conn.RemoteAddr().String()
		log.Info(p.Cli.Remote, " waiting for login message.")
	}
	// First connected, authorization required
	if !p.Cli.Authorized {
		p.Head.len = 12
		data, err := conn.Recv(p.Head.len, timeout)
		if err != nil {
			return nil, err
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.BigEndian, &p.Head.V1); err != nil {
			log.Error("Authentication head recv failed")
			return nil, err
		}
	} else {
		p.Head.len = 16
		data, err := conn.Recv(p.Head.len, timeout)
		if err != nil {
			return nil, err
		}
		buf := bytes.NewBuffer(data)
		if err := binary.Read(buf, binary.BigEndian, &p.Head.V2); err != nil {
			log.Error("V2 head recv failed")
			return nil, err
		}
	}
	return p.OnHead(conn, p.Head.len)
}

func (p *HMProtocol) Instance() net.TcpProtocol {
	return &HMProtocol{}
}
