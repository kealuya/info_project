//系统实际发票类型

/**
 * 发票类型数据
 */
const INVTYPES = [{
		key: 'ELECINV',
		name: '电子发票',
		type: '0',
		ocrInvType: '6',
		verifyType: '10'
	},
	{
		key: 'QUOTAINV',
		name: '定额发票',
		type: '1',
		ocrInvType: '5',
		verifyType: '1'
	},
	// {
	//   key: 'METROINV',
	//   name: '一卡通',
	//   type: '2'
	// },
	// {
	// 	key: 'GASINV',
	// 	name: '油票',
	// 	type: '3',
	// 	ocrInvType: "8",
	// 	verifyType: '3'
	// },
	{
		key: 'TAXIINV',
		name: '出租车票',
		type: '4',
		ocrInvType: '2',
		verifyType: '4'
	},
	{
		key: 'FLIGHTINV',
		name: '飞机票',
		type: '5',
		ocrInvType: '3',
		verifyType: '5'
	},
	{
		key: 'SHIPINV',
		name: '轮船票',
		type: '6',
		ocrInvType: '7',
		verifyType: '6'
	},
	{
		key: 'CARINV',
		name: '汽车票',
		type: '7',
		ocrInvType: '10',
		verifyType: '7'
	},
	{
		key: 'TRAININV',
		name: '火车票',
		type: '8',
		ocrInvType: '1',
		verifyType: '8'
	},
	{
		key: 'ADDVTAXINV',
		name: '增值税普通发票',
		type: '9',
		ocrInvType: '0',
		verifyType: '04'
	},
	{
		key: 'ADDVSPECIALINV',
		name: '增值税专用发票',
		type: '10',
		ocrInvType: '12',
		verifyType: '01'
	},
	{
		key: 'REFUELINV',
		name: '(卷票)加油票',
		type: '11',
		ocrInvType: '8',
		verifyType: '11'
	},
	{
		key: 'PRINTEDINV',
		name: '机打发票',
		type: '12',
		ocrInvType: '18',
		verifyType: '12'
	},
	{
		key: 'TOLLINV',
		name: '过路费',
		type: '16',
		ocrInvType: '16',
		verifyType: '16'
	},
	{
		key: 'OTHERINV',
		name: '其他发票',
		type: '22',
		ocrInvType: '22',
		verifyType: '22'
	},
	{
		key: 'OTHERBILLS',
		name: '其他票据',
		type: '23',
		ocrInvType: '23',
		verifyType: '23'
	},
	{
		key: 'CENTRENONTAXINV',
		name: '中央非税票据',
		type: '27',
		ocrInvType: '27',
		verifyType: '27'
	},
	{
		key: 'FOREIGN',
		name: '国外发票',
		type: '28',
		ocrInvType: '28',
		verifyType: '28'
	}
]

/**
 * 通过type获取发票类型
 * @param {*} type 发票类型的type值
 * @returns 发票类型信息(存在返回数据，不存在返回null)
 */
export const getFPLXInfoByType = (type) => {
	let result = INVTYPES.find((item) => {
		return item.type == type
	})
	if (result == undefined) {
		result = null
	}
	return result
}
/**
 * 通过type获取发票类型
 * @param {*} key 发票类型的key值
 * @returns 发票类型信息(存在返回数据，不存在返回null)
 */
export const getFPLXInfoByKey = (key) => {
	let result = INVTYPES.find((item) => {
		return item.key == key
	})
	if (result == undefined) {
		result = null
	}
	return result
}
/**
 * 通过type获取发票类型
 * @param {*} ocrInvType //识别接口返回的发票类型
 * @returns 发票类型信息(存在返回数据，不存在返回null)
 */
export const getFPLXInfoByOcr = (ocrInvType) => {
	let result = INVTYPES.find((item) => {
		return item.ocrInvType == ocrInvType
	})
	if (result == undefined) {
		result = null
	}
	return result
}
/**
 * 查询含有type数组的发票类型数组
 * @param {*} array 发票类型的type数组
 * @returns 发票类型数组
 */
export const getFPLXInfoListFilterIncludeByTypeArray = (array, sort = false) => {

	if (sort) {
		let result = []
		array.forEach((item) => {
			const find = INVTYPES.find(function(it) {
				return it.type == item
			})
			if (find != null & find != undefined) {
				result.push(find)
			}
		})
		return result
	} else {
		const result = INVTYPES.filter((item) => {
			return array.indexOf(item.type) >= 0
		})
		return result
	}

}

/**
 * 查询除去type数组的发票类型数组
 * @param {*} array 发票类型的type数组
 * @returns 发票类型数组
 */
export const getFPLXInfoListFilterExceptByTypeArray = (array) => {
	const result = INVTYPES.filter((item) => {
		return array.indexOf(item.type) >= 0
	})
	return result
}
// export const getImageSrcByType = (type) => {
//   try {
//     return require(`@/assets/image/formexpenseselect/icon-${type}.png`)
//   } catch {
//     return require(`@/assets/image/formexpenseselect/icon-default.png`)
//   }

// }
let FPLX = {}
INVTYPES.forEach((item) => {
	FPLX[item.key] = Number(item.type)
})
export {
	FPLX
}

let FPLXKEY = {}
INVTYPES.forEach((item) => {
	FPLXKEY[item.key] = {
		"WSYYTYPE": item.type,
		"OCRTYPE": item.ocrInvType,
		"VERIFYTYPE": item.verifyType
	}
})

export const findByOcrType = (ocrType) => {
	let arr = Object.keys(FPLXKEY)
	for (let item of arr) {
		let val = FPLXKEY[item]
		if (val.OCRTYPE == ocrType) {
			return val
		}
	}
	return null
}
export const findByWsyyType = (wsyyType) => {
	let arr = Object.keys(FPLXKEY)
	for (let item of arr) {
		let val = FPLXKEY[item]
		if (val.WSYYTYPE == wsyyType) {
			return val
		}
	}
	return null
}
export const findByVerifyType = (verifyType) => {
	let arr = Object.keys(FPLXKEY)
	for (let item of arr) {
		let val = FPLXKEY[item]
		if (val.VERIFYTYPE == verifyType) {
			return val
		}
	}
	return null
}

export {
	FPLXKEY
}

export default {
	"TABLE_FPLX_DATA": INVTYPES,
    findByOcrType,
    findByWsyyType
}
