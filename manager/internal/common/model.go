package common

import (
	"encoding/json"
)

// IdToArray
//
//	@Description: 将json格式的id转换为数组格式
//	@Author zzh
//	@param str
//	@return res
//	@return err
func IdToArray(str string) (res []uint64, err error) {
	if len(str) != 0 {
		err = json.Unmarshal([]byte(str), &res)
	}

	return
}

// ArrayToId
//
//	@Description: 将数组格式的id转换为json格式
//	@Author zzh
//	@param ids
//	@return res
//	@return err
func ArrayToId(ids []uint64) (res string, err error) {
	if len(ids) != 0 {
		var byteData []byte
		byteData, err = json.Marshal(ids)
		if err == nil {
			res = string(byteData)
		}
	}

	return
}

// CreateDeviceDataTable
//
//	@Description: 创建设备数据存储表
//	@Author zzh
//	@param ctx
//	@param db
//	@param tableName
//	@return error
//func CreateDeviceDataTable(ctx context.Context, db *gorm.DB, tableName string) error {
//
//	if db == nil {
//		return errors.New("db can not be null")
//	}
//
//	sql := "CREATE TABLE `" + tableName + "`  " +
//		"( `id` bigint NOT NULL AUTO_INCREMENT, " +
//		"`data` text CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL, " +
//		"`createTime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP," +
//		" PRIMARY KEY (`id`) USING BTREE) " +
//		"ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;"
//
//	return db.WithContext(ctx).Exec(sql).Error
//}
