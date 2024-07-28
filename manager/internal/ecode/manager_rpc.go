// Code generated by https://github.com/zhufuyi/sponge

package ecode

import (
	"github.com/zhufuyi/sponge/pkg/errcode"
)

// manager business-level rpc error codes.
// the _managerNO value range is 1~100, if the same number appears, it will cause a failure to start the service.
var (
	_managerNO       = 96
	_managerName     = "manager"
	_managerBaseCode = errcode.RCode(_managerNO)

	StatusLoginManager             = errcode.NewRPCStatus(_managerBaseCode+1, "failed to Login "+_managerName)
	StatusUserRegisterManager      = errcode.NewRPCStatus(_managerBaseCode+2, "failed to UserRegister "+_managerName)
	StatusUserGetManager           = errcode.NewRPCStatus(_managerBaseCode+3, "failed to UserGet "+_managerName)
	StatusUserUpdateManager        = errcode.NewRPCStatus(_managerBaseCode+4, "failed to UserUpdate "+_managerName)
	StatusResetPasswordManager     = errcode.NewRPCStatus(_managerBaseCode+5, "failed to ResetPassword "+_managerName)
	StatusUserDeleteManager        = errcode.NewRPCStatus(_managerBaseCode+6, "failed to UserDelete "+_managerName)
	StatusRoleCreateManager        = errcode.NewRPCStatus(_managerBaseCode+7, "failed to RoleCreate "+_managerName)
	StatusRoleGetManager           = errcode.NewRPCStatus(_managerBaseCode+8, "failed to RoleGet "+_managerName)
	StatusRoleUpdateManager        = errcode.NewRPCStatus(_managerBaseCode+9, "failed to RoleUpdate "+_managerName)
	StatusRoleDeleteManager        = errcode.NewRPCStatus(_managerBaseCode+10, "failed to RoleDelete "+_managerName)
	StatusApiCreateManager         = errcode.NewRPCStatus(_managerBaseCode+11, "failed to ApiCreate "+_managerName)
	StatusApiGetManager            = errcode.NewRPCStatus(_managerBaseCode+12, "failed to ApiGet "+_managerName)
	StatusApiUpdateManager         = errcode.NewRPCStatus(_managerBaseCode+13, "failed to ApiUpdate "+_managerName)
	StatusApiDeleteManager         = errcode.NewRPCStatus(_managerBaseCode+14, "failed to ApiDelete "+_managerName)
	StatusDeviceCreateManager      = errcode.NewRPCStatus(_managerBaseCode+15, "failed to DeviceCreate "+_managerName)
	StatusDeviceGetManager         = errcode.NewRPCStatus(_managerBaseCode+16, "failed to DeviceGet "+_managerName)
	StatusDeviceUpdateManager      = errcode.NewRPCStatus(_managerBaseCode+17, "failed to DeviceUpdate "+_managerName)
	StatusDeviceDeleteManager      = errcode.NewRPCStatus(_managerBaseCode+18, "failed to DeviceDelete "+_managerName)
	StatusDeviceTypeCreateManager  = errcode.NewRPCStatus(_managerBaseCode+19, "failed to DeviceTypeCreate "+_managerName)
	StatusDeviceTypeGetManager     = errcode.NewRPCStatus(_managerBaseCode+20, "failed to DeviceTypeGet "+_managerName)
	StatusDeviceTypeUpdateManager  = errcode.NewRPCStatus(_managerBaseCode+21, "failed to DeviceTypeUpdate "+_managerName)
	StatusDeviceTypeDeleteManager  = errcode.NewRPCStatus(_managerBaseCode+22, "failed to DeviceTypeDelete "+_managerName)
	StatusUserDeviceCreateManager  = errcode.NewRPCStatus(_managerBaseCode+23, "failed to UserDeviceCreate "+_managerName)
	StatusUserDeviceGetManager     = errcode.NewRPCStatus(_managerBaseCode+24, "failed to UserDeviceGet "+_managerName)
	StatusUserDeviceDeleteManager  = errcode.NewRPCStatus(_managerBaseCode+25, "failed to UserDeviceDelete "+_managerName)
	StatusDeviceDataReceiveManager = errcode.NewRPCStatus(_managerBaseCode+26, "failed to DeviceDataReceive "+_managerName)
	StatusDeviceDataGetManager     = errcode.NewRPCStatus(_managerBaseCode+27, "failed to DeviceDataGet "+_managerName)
	// error codes are globally unique, adding 1 to the previous error code
)
