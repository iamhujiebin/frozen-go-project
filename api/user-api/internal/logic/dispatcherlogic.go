package logic

import (
	"context"
	"frozen-go-project/common/codes/resp_codes"
	"frozen-go-project/rpc/base-rpc/baserpc"
	"strconv"
	"strings"

	"frozen-go-project/api/user-api/internal/svc"
	"frozen-go-project/api/user-api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type DispatcherLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDispatcherLogic(ctx context.Context, svcCtx *svc.ServiceContext) DispatcherLogic {
	return DispatcherLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DispatcherLogic) Dispatcher(req types.DispatcherRequest) (*types.CommonResponse, error) {
	res, err := l.svcCtx.BaseRpc.GetPkgConfig(l.ctx, &baserpc.GetPkgConfigReq{
		PkgName: req.Pkg,
		Section: "ChatWs",
		Key:     "Servers",
	})
	if err != nil {
		return &types.CommonResponse{
			Code:    resp_codes.RpcError,
			Message: err.Error(),
		}, nil
	}
	servers := parseChatServers(res.Value)
	index := DJBHash(req.Id) % uint(len(servers))
	server := servers[index]
	wssServer := server
	hostAndPort := strings.Split(server, ":")
	if len(hostAndPort) == 2 {
		port, err := strconv.Atoi(hostAndPort[1])
		if err == nil {
			port += 1000
			wssServer = hostAndPort[0] + ":" + strconv.Itoa(port)
		}
	}
	return &types.CommonResponse{
		Body: map[string]interface{}{
			"wss_server": wssServer,
			"server":     server,
			"path":       "/chat/ws",
		},
	}, nil
}

func parseChatServers(chatWsServer string) []string {
	if len(chatWsServer) == 0 {
		return nil
	}
	arr := make([]string, 0)
	parts := strings.Split(chatWsServer, ",")
	for _, part := range parts {
		list := strings.Split(part, ":")
		if len(list) != 3 {
			continue
		}
		ip := list[0]
		port, err := strconv.Atoi(list[1])
		if err != nil {
			continue
		}
		serverNums, err := strconv.Atoi(list[2])
		if err != nil {
			continue
		}
		for i := 0; i < serverNums; i++ {
			arr = append(arr, strings.Join([]string{ip, strconv.Itoa(port)}, ":"))
			port++
		}
	}
	return arr
}

//获取hash值
func DJBHash(str string) uint {
	if len(str) == 0 {
		return 0
	}
	hash := uint(5381)
	for i := 0; i < len(str); i++ {
		hash = ((hash << 5) + hash) + uint(str[i])
	}
	return hash
}
