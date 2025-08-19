package main

import (
	"example.com/mod/fomat"
	"fmt"
	"github.com/disintegration/imaging"
	"time"
)

func main() {
	fomat.JsonFormat(`[
  {
    "code": 200,
    "data": {
      "driverDetail": {
        "driverName": "陈攀",
        "driverPhone": "14743638352",
        "vehicleColor": "灰",
        "vehicleModel": "秦plus ev",
        "vehicleNo": "湘ADR1775"
      },
      "orderDetail": {
        "bookingType": "实时",
        "canceled": false,
        "cityId": "0",
        "cityName": "Chang Sha Shi",
        "complainStatus": "toEvaluable",
        "countryId": 7,
        "countryName": "China",
        "createTime": 1746789029000,
        "createTimeYMD": "05月09日 19:10",
        "departureTime": 1746788906000,
        "dispatchedTime": 1746789036000,
        "endAddress": {
          "address": "尖山路",
          "addressDetail": "中国湖南省长沙市岳麓区尖山路 邮政编码: 410205",
          "cityName": "Chang Sha Shi",
          "latitude": 28.2347887,
          "longitude": 112.8841523
        },
        "estimateDuration": "7",
        "estimateTotalFee": 15.0000,
        "orderId": 1099537053,
        "orderType": "1",
        "orderTypeName": "即时用车",
        "passengerMobile": "15507384382",
        "passengerName": "Jek_771947",
        "paymentType": "公司支付",
        "startAddress": {
          "address": "岳麓大道",
          "addressDetail": "中国湖南省长沙市岳麓区岳麓大道 邮政编码: 410217",
          "cityName": "Chang Sha Shi",
          "latitude": 28.2282901,
          "longitude": 112.8758381
        },
        "status": "dispatched",
        "statusName": "等待接驾",
        "supplierId": "12216",
        "supplierLogo": "https://oss-now.heycars.cn/image/supplier/logo/rq@1x.png",
        "supplierName": "如祺出行",
        "updateTime": 1746789037000,
        "useTime": 1746788906000,
        "useTimeYMD": "05月09日 19:08",
        "vehicleGroupId": "1",
        "vehicleGroupName": "经济型",
        "yX_ServiceTel": "+862882122976"
      }
    },
    "msg": "成功"
  }
]`)

}

// Path := "/Users/developer/Downloads/"
// originalPath := Path + ""
// savePath := Path + ""
// ImageCutSquare(60, originalPath, savePath)
func apiSpeedLimiter(count float64) {
	speed := 1 / count

	fmt.Println("speed:")
	fmt.Println(speed)
	time.Sleep(time.Duration(speed) * time.Second)
}
func imageCutSquare(sidelength int, originalPath, savePath string) {
	// 1. 打开图片
	src, err := imaging.Open(originalPath)
	if err != nil {
		fmt.Println("打开图片失败")
	}

	// 2. 调整大小（固定 75x75）
	dst := imaging.Resize(src, sidelength, sidelength, imaging.Lanczos)

	// 3. 保存图片
	err = imaging.Save(dst, savePath)
	if err != nil {
		fmt.Println("保存图片失败")
	}

}
