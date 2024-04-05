export const invoiceTypeList = [
    // 定额发票 type=1
    {
        "label": "定额发票",
        "ocrType": 5,
        "list": [
            {
                "fpType": 1,
                "fpTypeName": "定额发票",
                "ocrType": 5,
                "inputKey": "fpdm",
                "inputLabel": "发票代码",
                'maxLangth':12,
                "inputImportant": 1,
                "inputType": "code",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票代码",
                "selectData": []
            },
            {
                "fpType": 1,
                "fpTypeName": "定额发票",
                "ocrType": 5,
                "inputKey": "fphm",
                "inputLabel": "发票号码",
                'maxLangth':8,
                "inputImportant": 1,
                "inputType": "number",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票号码",
                "selectData": []
            },
            {
                "fpType": 1,
                "fpTypeName": "定额发票",
                "ocrType": 5,
                'maxLangth':8,
                "inputKey": "fpje",
                "inputLabel": "金额",
                "inputImportant": 1,
                "inputType": "money",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入金额",
                "selectData": []
            },
            {
                "fpType": 1,
                "fpTypeName": "定额发票",
                "ocrType": 5,
                "inputKey": "fylb",
                "inputLabel": "费用类别",
                "inputImportant": 0,
                'maxLangth':10,
                "inputType": "select",
                "other": "请选择#请选择;通用#通用;餐费#餐费;停车费#停车费;地铁#地铁;出租#出租",
                "inputLabel2": null,
                "placeholder2": "点击选择费用类别",
                "selectData": [
                    //   {
                    //   "text": "请选择",
                    //   "value": "请选择"
                    // },
                    {"text": "通用", "value": "通用"},
                    {"text": "餐费", "value": "餐费"},
                    {"text": "停车费", "value": "停车费"},
                    {"text": "地铁", "value": "地铁"},
                    {"text": "出租", "value": "出租"}
                ]
            }
        ],
        "value": 1,
    },
    // 电子发票 type=0
    {
        "label": "增值税电子发票",
        "ocrType": 6,
        "list": [
            {
                "fpType": 0,
                "fpTypeName": "增值税电子发票",
                "ocrType": 6,
                "inputKey": "fpdm",
                "inputLabel": "发票代码",
                'maxLangth':12,
                "inputImportant": 1,
                "inputType": "code",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票代码",
                "selectData": []
            },
            {
                "fpType": 0,
                "fpTypeName": "增值税电子发票",
                "ocrType": 6,
                "inputKey": "fphm",
                "inputLabel": "发票号码",
                'maxLangth':8,
                "inputImportant": 1,
                "inputType": "number",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票号码",
                "selectData": []
            },
            {
                "fpType": 0,
                "fpTypeName": "增值税电子发票",
                "ocrType": 6,
                "inputKey": "xsfmc",
                "inputLabel": "销售方名称",
                'maxLangth':30,
                "inputImportant": 0,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入销售方名称",
                "selectData": []
            },
            {
                "fpType": 0,
                "fpTypeName": "增值税电子发票",
                "ocrType": 6,
                "inputKey": "xsfnsrsbh",
                "inputLabel": "销售方税号",
                'maxLangth':30,
                "inputImportant": 0,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入销售方税号",
                "selectData": []
            },
            {
                "fpType": 0,
                "fpTypeName": "增值税电子发票",
                "ocrType": 6,
                "inputKey": "gmfmc",
                "inputLabel": "购买方名称",
                'maxLangth':30,
                "inputImportant": 0,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入销售方名称",
                "selectData": []
            },
            {
                "fpType": 0,
                "fpTypeName": "增值税电子发票",
                "ocrType": 6,
                "inputKey": "gmfnsrsbh",
                "inputLabel": "购买方税号",
                'maxLangth':30,
                "inputImportant": 0,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入销售方税号",
                "selectData": []
            },
            {
                "fpType": 0,
                "fpTypeName": "增值税电子发票",
                "ocrType": 6,
                "inputKey": "fwmc",
                "inputLabel": "商品明细",
                'maxLangth':120,
                "inputImportant": 0,
                "inputType": "textarea",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入商品明细",
                "selectData": []
            },
            {
                "fpType": 0,
                "fpTypeName": "增值税电子发票",
                "ocrType": 6,
                "inputKey": "fpjym",
                "inputLabel": "校验码(后六位)",
                'maxLangth':6,
                "inputImportant": 1,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入校验码(后六位)",
                "selectData": []
            },
            {
                "fpType": 0,
                "fpTypeName": "增值税电子发票",
                "ocrType": 6,
                "inputKey": "kprq",
                "inputLabel": "开票日期",
                'maxLangth':12,
                "inputImportant": 1,
                "inputType": "date",
                "other": "YYYYMMDD",
                "inputLabel2": null,
                "placeholder2": "请输入开票日期",
                "selectData": []
            },
            {
                "fpType": 0,
                "fpTypeName": "增值税电子发票",
                "ocrType": 6,
                "inputKey": "fpje",
                "inputLabel": "金额",
                'maxLangth':8,
                "inputImportant": 1,
                "inputType": "money",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入金额",
                "selectData": []
            }
        ],
        "value": 0
    },
    // 出租车票 type=4
    {
        "label": "出租车票",
        "ocrType": 2,
        "list": [
            {
                "fpType": 4,
                "fpTypeName": "出租车票",
                "ocrType": 2,
                "inputKey": "fpdm",
                "inputLabel": "发票代码",
                'maxLangth':12,
                "inputImportant": 1,
                "inputType": "code",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票代码",
                "selectData": []
            },
            {
                "fpType": 4,
                "fpTypeName": "出租车票",
                "ocrType": 2,
                "inputKey": "fphm",
                "inputLabel": "发票号码",
                "inputImportant": 1,
                'maxLangth':8,
                "inputType": "number",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票号码",
                "selectData": []
            },
            {
                "fpType": 4,
                "fpTypeName": "出租车票",
                "ocrType": 2,
                "inputKey": "kprq",
                'maxLangth':12,
                "inputLabel": "开票日期",
                "inputImportant": 1,
                "inputType": "date",
                "other": "YYYYMMDD",
                "inputLabel2": null,
                "placeholder2": "请输入开票日期",
                "selectData": []
            },
            {
                "fpType": 4,
                "fpTypeName": "出租车票",
                "ocrType": 2,
                "inputKey": "cph",
                "inputLabel": "车牌号",
                'maxLangth':8,
                "inputImportant": 0,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入车牌号",
                "selectData": []
            },
            {
                "fpType": 4,
                "fpTypeName": "出租车票",
                "ocrType": 2,
                "inputKey": "zjh",
                "inputLabel": "证件号",
                'maxLangth':18,
                "inputImportant": 0,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入证件号",
                "selectData": []
            },
            {
                "fpType": 4,
                "fpTypeName": "出租车票",
                "ocrType": 2,
                "inputKey": "scsj",
                "inputLabel": "上车时间",
                'maxLangth':12,
                "inputImportant": 0,
                "inputType": "time",
                "other": "HHmm",
                "inputLabel2": null,
                "placeholder2": "点击选择上车时间",
                "selectData": []
            },
            {
                "fpType": 4,
                "fpTypeName": "出租车票",
                "ocrType": 2,
                "inputKey": "xcsj",
                "inputLabel": "下车时间",
                'maxLangth':12,
                "inputImportant": 0,
                "inputType": "time",
                "other": "HHmm",
                "inputLabel2": null,
                "placeholder2": "点击选择下车时间",
                "selectData": []
            },
            {
                "fpType": 4,
                "fpTypeName": "出租车票",
                "ocrType": 2,
                "inputKey": "fpje",
                "inputLabel": "金额",
                'maxLangth':8,
                "inputImportant": 1,
                "inputType": "money",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入金额",
                "selectData": []
            }
        ],
        "value": 4
    },
    // 飞机票 type=5
    {
        "label": "飞机票",
        "ocrType": 3,
        "list": [
            {
                "fpType": 5,
                "fpTypeName": "飞机票",
                "ocrType": 3,
                "inputKey": "xm",
                "inputLabel": "姓名",
                "inputImportant": 1,
                "inputType": "text",
                'maxLangth':30,
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入姓名",
                "selectData": []
            },
            {
                "fpType": 5,
                "fpTypeName": "飞机票",
                "ocrType": 3,
                'maxLangth':30,
                "inputKey": "list",
                "inputLabel": "",
                "inputImportant": 1,
                "inputType": "map",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入姓名",
                "selectData": [],
            },
            {
                "fpType": 5,
                "fpTypeName": "飞机票",
                "ocrType": 3,
                "inputKey": "cfrq",
                'maxLangth':12,
                "inputLabel": "出发日期",
                "inputImportant": 1,
                "inputType": "date",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入出发日期",
                "selectData": []
            },
            {
                "fpType": 5,
                "fpTypeName": "飞机票",
                "ocrType": 3,
                "inputKey": "cfdd",
                "inputLabel": "出发地",
                'maxLangth':20,
                "inputImportant": 1,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入出发地",
                "selectData": []
            },
            {
                "fpType": 5,
                "fpTypeName": "飞机票",
                "ocrType": 3,
                "inputKey": "dddd",
                'maxLangth':20,
                "inputLabel": "目的地",
                "inputImportant": 1,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入目的地",
                "selectData": []
            },
            {
                "fpType": 5,
                "fpTypeName": "飞机票",
                "ocrType": 3,
                "inputKey": "pb",
                'maxLangth':12,
                "inputLabel": "票别",
                "inputImportant": 1,
                "inputType": "select",
                "other": "pbmc#飞机",
                "inputLabel2": null,
                "placeholder2": "点击选择票别",
                "selectData": [
                    {"text": "头等舱", "value": "头等舱"},
                    {"text": "商务舱", "value": "商务舱"},
                    {"text": "经济舱", "value": "经济舱"}
                ]
            },
            {
                "fpType": 5,
                "fpTypeName": "飞机票",
                "ocrType": 3,
                "inputKey": "ph",
                "inputLabel": "票号",
                'maxLangth':18,
                "inputImportant": 1,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入票号",
                "selectData": []
            },
            {
                "fpType": 5,
                "fpTypeName": "飞机票",
                "ocrType": 3,
                "inputKey": "hbh",
                'maxLangth':15,
                "inputLabel": "航班号",
                "inputImportant": 1,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入航班号",
                "selectData": []
            },
            {
                "fpType": 5,
                "fpTypeName": "飞机票",
                "ocrType": 3,
                "inputKey": "fpje",
                "inputLabel": "金额",
                'maxLangth':8,
                "inputImportant": 1,
                "inputType": "money",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入金额",
                "selectData": []
            },
            {
                "fpType": 5,
                "fpTypeName": "飞机票",
                "ocrType": 3,
                "inputKey": "zfcg",
                "inputLabel": "政府采购",
                'maxLangth':12,
                "inputImportant": 0,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请填写GP开头号码（非必填）",
                "selectData": []
            }
        ],
        "value": 5
    },
    // 船票 type=6
    {
        "label": "轮船票",
        "ocrType": 7,
        "list": [
            {
                "fpType": 6,
                "fpTypeName": "轮船票",
                "ocrType": 7,
                "inputKey": "fpdm",
                "inputLabel": "发票代码",
                'maxLangth':12,
                "inputImportant": 1,
                "inputType": "code",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票代码",
                "selectData": []
            },
            {
                "fpType": 6,
                "fpTypeName": "轮船票",
                "ocrType": 7,
                "inputKey": "fphm",
                "inputLabel": "发票号",
                "inputImportant": 1,
                'maxLangth':8,
                "inputType": "number",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票号",
                "selectData": []
            },
            {
                "fpType": 6,
                "fpTypeName": "轮船票",
                "ocrType": 7,
                "inputKey": "kprq",
                "inputLabel": "开票日期",
                'maxLangth':12,
                "inputImportant": 1,
                "inputType": "date",
                "other": "YYYYMMDD",
                "inputLabel2": null,
                "placeholder2": "请输入开票日期",
                "selectData": []
            },
            {
                "fpType": 6,
                "fpTypeName": "轮船票",
                "ocrType": 7,
                "inputKey": "cfrq",
                "inputLabel": "起始时间",
                'maxLangth':12,
                "inputImportant": 1,
                "inputType": "date",
                "other": "YYYYMMDD",
                "inputLabel2": null,
                "placeholder2": "请输入起始时间",
                "selectData": []
            },
            {
                "fpType": 6,
                "fpTypeName": "轮船票",
                "ocrType": 7,
                "inputKey": "fpje",
                "inputLabel": "金额",
                'maxLangth':8,
                "inputImportant": 1,
                "inputType": "money",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入金额",
                "selectData": []
            }
        ],
        "value": 6
    },
    // 汽车票 type=7
    {
        "label": "汽车票",
        "ocrType": 10,
        "list": [
            {
                "fpType": 7,
                "fpTypeName": "汽车票",
                "ocrType": 10,
                "inputKey": "fpdm",
                "inputLabel": "发票代码",
                'maxLangth':12,
                "inputImportant": 1,
                "inputType": "code",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票代码",
                "selectData": []
            },
            {
                "fpType": 7,
                "fpTypeName": "汽车票",
                "ocrType": 10,
                "inputKey": "fphm",
                "inputLabel": "发票号码",
                "inputImportant": 1,
                'maxLangth':8,
                "inputType": "number",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票号码",
                "selectData": []
            },
            {
                "fpType": 7,
                "fpTypeName": "汽车票",
                "ocrType": 10,
                "inputKey": "kprq",
                "inputLabel": "开票日期",
                "inputImportant": 1,
                "inputType": "date",
                "other": "YYYYMMDD",
                "inputLabel2": null,
                "placeholder2": "请输入开票日期",
                "selectData": []
            },
            {
                "fpType": 7,
                "fpTypeName": "汽车票",
                "ocrType": 10,
                "inputKey": "fpje",
                "inputLabel": "金额",
                "inputImportant": 1,
                "inputType": "money",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入金额",
                "selectData": []
            }
        ],
        "value": 7
    },
    // 火车票 type=8
    {
        "label": "火车票",
        "ocrType": 1,
        "list": [
            {
                "fpType": 8,
                "fpTypeName": "火车票",
                "ocrType": 1,
                "inputKey": "xm",
                "inputLabel": "姓名",
                "inputImportant": 1,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入姓名",
                "selectData": []
            },
            {
                "fpType": 8,
                "fpTypeName": "火车票",
                "ocrType": 1,
                "inputKey": "cfrq",
                "inputLabel": "出发日期",
                "inputImportant": 1,
                "inputType": "date",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入出发日期",
                "selectData": []
            },
            {
                "fpType": 8,
                "fpTypeName": "火车票",
                "ocrType": 1,
                "inputKey": "cfdd",
                "inputLabel": "出发地",
                "inputImportant": 1,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入出发地",
                "selectData": []
            },
            {
                "fpType": 8,
                "fpTypeName": "火车票",
                "ocrType": 1,
                "inputKey": "dddd",
                "inputLabel": "目的地",
                "inputImportant": 1,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入目的地",
                "selectData": []
            },
            {
                "fpType": 8,
                "fpTypeName": "火车票",
                "ocrType": 1,
                "inputKey": "pb",
                "inputLabel": "票别",
                "inputImportant": 1,
                "inputType": "select",
                "other": "pbmc#火车",
                "inputLabel2": null,
                "placeholder2": "点击选择票别",
                "selectData": [
                    {"text": "高铁一等座", "value": "高铁一等座"},
                    {"text": "高铁二等座", "value": "高铁二等座"},
                    {"text": "高铁特等座", "value": "高铁特等座"},
                    {"text": "普列硬座", "value": "普列硬座"},
                    {"text": "普列硬卧", "value": "普列硬卧"},
                    {"text": "普列软卧", "value": "普列软卧"},
                    {"text": "普列高级软卧", "value": "普列高级软卧"}
                ]
            },
            {
                "fpType": 8,
                "fpTypeName": "火车票",
                "ocrType": 1,
                "inputKey": "cc",
                "inputLabel": "车次",
                "inputImportant": 1,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入车次",
                "selectData": []
            },
            {
                "fpType": 8,
                "fpTypeName": "火车票",
                "ocrType": 1,
                "inputKey": "fpje",
                "inputLabel": "金额",
                "inputImportant": 1,
                "inputType": "money",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入金额",
                "selectData": []
            }
        ],
        "value": 8
    },
    // 增值税普通发票 type=9
    {
        "label": "增值税普通发票",
        "ocrType": 0,
        "list": [
            {
                "fpType": 9,
                "fpTypeName": "增值税普通发票",
                "ocrType": 0,
                "inputKey": "fpdm",
                "inputLabel": "发票代码",
                'maxLangth':12,
                "inputImportant": 1,
                "inputType": "code",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票代码",
                "selectData": []
            },
            {
                "fpType": 9,
                "fpTypeName": "增值税普通发票",
                "ocrType": 0,
                "inputKey": "fphm",
                "inputLabel": "发票号码",
                'maxLangth':8,
                "inputImportant": 1,
                "inputType": "number",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票号码",
                "selectData": []
            },
            {
                "fpType": 0,
                "fpTypeName": "增值税普通发票",
                "ocrType": 6,
                "inputKey": "xsfmc",
                "inputLabel": "销售方名称",
                "inputImportant": 0,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入销售方名称",
                "selectData": []
            },
            {
                "fpType": 0,
                "fpTypeName": "增值税普通发票",
                "ocrType": 6,
                "inputKey": "xsfnsrsbh",
                "inputLabel": "销售方税号",
                "inputImportant": 0,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入销售方税号",
                "selectData": []
            },
            {
                "fpType": 0,
                "fpTypeName": "增值税普通发票",
                "ocrType": 6,
                "inputKey": "gmfmc",
                "inputLabel": "购买方名称",
                "inputImportant": 0,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入销售方名称",
                "selectData": []
            },
            {
                "fpType": 0,
                "fpTypeName": "增值税普通发票",
                "ocrType": 6,
                "inputKey": "gmfnsrsbh",
                "inputLabel": "购买方税号",
                "inputImportant": 0,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入销售方税号",
                "selectData": []
            },
            {
                "fpType": 0,
                "fpTypeName": "增值税普通发票",
                "ocrType": 6,
                "inputKey": "fwmc",
                "inputLabel": "商品明细",
                "inputImportant": 0,
                "inputType": "textarea",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入商品明细",
                "selectData": []
            },
            {
                "fpType": 9,
                "fpTypeName": "增值税普通发票",
                "ocrType": 0,
                "inputKey": "jym",
                "inputLabel": "校验码(后六位)",
                "inputImportant": 1,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入校验码(后六位)",
                "selectData": []
            },
            {
                "fpType": 9,
                "fpTypeName": "增值税普通发票",
                "ocrType": 0,
                "inputKey": "kprq",
                "inputLabel": "开票日期",
                "inputImportant": 1,
                "inputType": "date",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入开票日期",
                "selectData": []
            },
            {
                "fpType": 9,
                "fpTypeName": "增值税普通发票",
                "ocrType": 0,
                "inputKey": "fpje",
                "inputLabel": "金额",
                "inputImportant": 1,
                "inputType": "money",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入金额",
                "selectData": []
            }
        ],
        "value": 9
    },
    // 增值税专用发票 type=10
    {
        "label": "增值税专用发票",
        "ocrType": 12,
        "list": [
            {
                "fpType": 10,
                "fpTypeName": "增值税专用发票",
                "ocrType": 12,
                "inputKey": "fpdm",
                "inputLabel": "发票代码",
                'maxLangth':12,
                "inputImportant": 1,
                "inputType": "code",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票代码",
                "selectData": []
            },
            {
                "fpType": 10,
                "fpTypeName": "增值税专用发票",
                "ocrType": 12,
                "inputKey": "fphm",
                "inputLabel": "发票号码",
                "inputImportant": 1,
                'maxLangth':8,
                "inputType": "number",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票号码",
                "selectData": []
            },
            {
                "fpType": 0,
                "fpTypeName": "增值税专用发票",
                "ocrType": 6,
                "inputKey": "xsfmc",
                "inputLabel": "销售方名称",
                "inputImportant": 0,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入销售方名称",
                "selectData": []
            },
            {
                "fpType": 0,
                "fpTypeName": "增值税专用发票",
                "ocrType": 6,
                "inputKey": "xsfnsrsbh",
                "inputLabel": "销售方税号",
                "inputImportant": 0,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入销售方税号",
                "selectData": []
            },
            {
                "fpType": 0,
                "fpTypeName": "增值税专用发票",
                "ocrType": 6,
                "inputKey": "gmfmc",
                "inputLabel": "购买方名称",
                "inputImportant": 0,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入销售方名称",
                "selectData": []
            },
            {
                "fpType": 0,
                "fpTypeName": "增值税专用发票",
                "ocrType": 6,
                "inputKey": "gmfnsrsbh",
                "inputLabel": "购买方税号",
                "inputImportant": 0,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入销售方税号",
                "selectData": []
            },
            {
                "fpType": 0,
                "fpTypeName": "增值税普通发票",
                "ocrType": 6,
                "inputKey": "fwmc",
                "inputLabel": "商品明细",
                "inputImportant": 0,
                "inputType": "textarea",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入商品明细",
                "selectData": []
            },
            {
                "fpType": 10,
                "fpTypeName": "增值税专用发票",
                "ocrType": 12,
                "inputKey": "kprq",
                "inputLabel": "开票日期",
                "inputImportant": 1,
                "inputType": "date",
                "other": "YYYYMMDD",
                "inputLabel2": null,
                "placeholder2": "请输入开票日期",
                "selectData": []
            },
            {
                "fpType": 10,
                "fpTypeName": "增值税专用发票",
                "ocrType": 12,
                "inputKey": "fpje",
                "inputLabel": "金额(不含税)",
                "inputImportant": 1,
                "inputType": "money",
                "other": null,
                "inputLabel2": "(不含税)",
                "placeholder2": "请输入金额",
                "selectData": []
            },
            {
                "fpType": 10,
                "fpTypeName": "增值税专用发票",
                "ocrType": 12,
                "inputKey": "fpjetax",
                "inputLabel": "总金额",
                "inputImportant": 1,
                "inputType": "money",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入总金额",
                "selectData": []
            }
        ],
        "value": 10
    },
    // 加油票（卷票） type=11
    {
        "label": "(卷票)加油票",
        "ocrType": 8,
        "list": [
            {
                "fpType": 11,
                "fpTypeName": "(卷票)加油票",
                "ocrType": 8,
                "inputKey": "fpdm",
                "inputLabel": "发票代码",
                "inputImportant": 1,
                'maxLangth':12,
                "inputType": "code",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票代码",
                "selectData": []
            },
            {
                "fpType": 11,
                "fpTypeName": "(卷票)加油票",
                "ocrType": 8,
                "inputKey": "fphm",
                "inputLabel": "发票号",
                "inputImportant": 1,
                'maxLangth':8,
                "inputType": "number",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票号",
                "selectData": []
            },
            {
                "fpType": 11,
                "fpTypeName": "(卷票)加油票",
                "ocrType": 8,
                "inputKey": "jym",
                "inputLabel": "校验码(后六位)",
                "inputImportant": 1,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入校验码(后六位)",
                "selectData": []
            },
            {
                "fpType": 11,
                "fpTypeName": "(卷票)加油票",
                "ocrType": 8,
                "inputKey": "kprq",
                "inputLabel": "开票日期",
                "inputImportant": 1,
                "inputType": "date",
                "other": "YYYYMMDD",
                "inputLabel2": null,
                "placeholder2": "请输入开票日期",
                "selectData": []
            },
            {
                "fpType": 11,
                "fpTypeName": "(卷票)加油票",
                "ocrType": 8,
                "inputKey": "fpje",
                "inputLabel": "金额",
                "inputImportant": 1,
                "inputType": "money",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入金额",
                "selectData": []
            }
        ],
        "value": 11
    },
    // 机打发票 type=12
    {
        "label": "机打发票",
        "ocrType": 18,
        "list": [
            {
                "fpType": 12,
                "fpTypeName": "机打发票",
                "ocrType": 18,
                "inputKey": "fpdm",
                "inputLabel": "发票代码",
                "inputImportant": 1,
                'maxLangth':12,
                "inputType": "code",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票代码",
                "selectData": []
            },
            {
                "fpType": 12,
                "fpTypeName": "机打发票",
                "ocrType": 18,
                "inputKey": "fphm",
                "inputLabel": "发票号",
                "inputImportant": 1,
                'maxLangth':8,
                "inputType": "number",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票号",
                "selectData": []
            },
            {
                "fpType": 12,
                "fpTypeName": "机打发票",
                "ocrType": 18,
                "inputKey": "kprq",
                "inputLabel": "开票日期",
                "inputImportant": 1,
                "inputType": "date",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入开票日期",
                "selectData": []
            },
            {
                "fpType": 12,
                "fpTypeName": "机打发票",
                "ocrType": 18,
                "inputKey": "fpje",
                "inputLabel": "金额",
                "inputImportant": 1,
                "inputType": "money",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入金额",
                "selectData": []
            }
        ],
        "value": 12
    },
    // 过路费 type=16
    {
        "label": "过路费",
        "ocrType": 16,
        "list": [
            {
                "fpType": 16,
                "fpTypeName": "过路费",
                "ocrType": 16,
                "inputKey": "fpdm",
                "inputLabel": "发票代码",
                'maxLangth':12,
                "inputImportant": 1,
                "inputType": "code",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票代码",
                "selectData": []
            },
            {
                "fpType": 16,
                "fpTypeName": "过路费",
                "ocrType": 16,
                "inputKey": "fphm",
                "inputLabel": "发票号码",
                'maxLangth':8,
                "inputImportant": 1,
                "inputType": "number",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票号码",
                "selectData": []
            },
            {
                "fpType": 16,
                "fpTypeName": "过路费",
                "ocrType": 16,
                "inputKey": "kprq",
                "inputLabel": "开票日期",
                "inputImportant": 1,
                "inputType": "date",
                "other": "YYYYMMDD",
                "inputLabel2": null,
                "placeholder2": "请输入开票日期",
                "selectData": []
            },
            {
                "fpType": 16,
                "fpTypeName": "过路费",
                "ocrType": 16,
                "inputKey": "fpje",
                "inputLabel": "金额",
                "inputImportant": 1,
                "inputType": "money",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入金额",
                "selectData": []
            }
        ],
        "value": 16
    },
    // 其他发票 type=22
    {
        "label": "其他发票",
        "ocrType": 22,
        "list": [
            {
                "fpType": 22,
                "fpTypeName": "其他发票",
                "ocrType": 22,
                "inputKey": "fpdm",
                "inputLabel": "发票代码",
                "inputImportant": 1,
                'maxLangth':12,
                "inputType": "code",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票代码",
                "selectData": []
            },
            {
                "fpType": 22,
                "fpTypeName": "其他发票",
                "ocrType": 22,
                "inputKey": "fphm",
                "inputLabel": "发票号码",
                'maxLangth':8,
                "inputImportant": 1,
                "inputType": "number",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票号码",
                "selectData": []
            },
            {
                "fpType": 22,
                "fpTypeName": "其他发票",
                "ocrType": 22,
                "inputKey": "kprq",
                "inputLabel": "开票日期",
                "inputImportant": 1,
                "inputType": "date",
                "other": "YYYYMMDD",
                "inputLabel2": null,
                "placeholder2": "请输入开票日期",
                "selectData": []
            },
            {
                "fpType": 22,
                "fpTypeName": "其他发票",
                "ocrType": 22,
                "inputKey": "fpje",
                "inputLabel": "金额",
                "inputImportant": 1,
                "inputType": "money",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入金额",
                "selectData": []
            }
        ],
        "value": 22
    },
    // 其他票据 type=23
    {
        "label": "其他票据",
        "ocrType": 23,
        "list": [
            {
                "fpType": 23,
                "fpTypeName": "其他票据",
                "ocrType": 23,
                "inputKey": "pjhm",
                "inputLabel": "票据号码",
                "inputImportant": 1,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入票据号码",
                "selectData": []
            },
            {
                "fpType": 23,
                "fpTypeName": "其他票据",
                "ocrType": 23,
                "inputKey": "kprq",
                "inputLabel": "开票日期",
                "inputImportant": 1,
                "inputType": "date",
                "other": "YYYYMMDD",
                "inputLabel2": null,
                "placeholder2": "请输入开票日期",
                "selectData": []
            },
            {
                "fpType": 23,
                "fpTypeName": "其他票据",
                "ocrType": 23,
                "inputKey": "fpje",
                "inputLabel": "金额",
                "inputImportant": 1,
                "inputType": "money",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入金额",
                "selectData": []
            },
            {
                "fpType": 23,
                "fpTypeName": "其他票据",
                "ocrType": 23,
                "inputKey": "kpnr",
                "inputLabel": "开票内容",
                "inputImportant": 1,
                "inputType": "textarea",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入开票内容",
                "selectData": []
            },
            {
                "fpType": 23,
                "fpTypeName": "其他票据",
                "ocrType": 23,
                "inputKey": "kpdwmc",
                "inputLabel": "开票单位名称",
                "inputImportant": 0,
                "inputType": "text",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入开票单位名称",
                "selectData": []
            }
        ],
        "value": 23
    },
    // 中央非税票据 type=17
    {
        "label": "中央非税票据",
        "ocrType": 27,
        "list": [
            {
                "fpType": 27,
                "fpTypeName": "中央非税票据",
                "ocrType": 27,
                "inputKey": "fpdm",
                "inputLabel": "发票代码",
                'maxLangth':12,
                "inputImportant": 1,
                "inputType": "code",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票代码",
                "selectData": []
            },
            {
                "fpType": 27,
                "fpTypeName": "中央非税票据",
                "ocrType": 27,
                "inputKey": "fphm",
                "inputLabel": "发票号码",
                'maxLangth':8,
                "inputImportant": 1,
                "inputType": "number",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票号码",
                "selectData": []
            },
            {
                "fpType": 27,
                "fpTypeName": "中央非税票据",
                "ocrType": 27,
                "inputKey": "kprq",
                "inputLabel": "开票日期",
                "inputImportant": 1,
                "inputType": "date",
                "other": "YYYYMMDD",
                "inputLabel2": null,
                "placeholder2": "请输入开票日期",
                "selectData": []
            },
            {
                "fpType": 27,
                "fpTypeName": "中央非税票据",
                "ocrType": 27,
                "inputKey": "fpje",
                "inputLabel": "金额",
                "inputImportant": 1,
                "inputType": "money",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入金额",
                "selectData": []
            }
        ],
        "value": 27,
    },
    // 国外票据 type=28
    {
        "label": "国外发票",
        "ocrType": 28,
        "list": [
            {
                "fpType": 28,
                "fpTypeName": "国外发票",
                "ocrType": 28,
                "inputKey": "fpdm",
                "inputLabel": "发票代码",
                'maxLangth':12,
                "inputImportant": 1,
                "inputType": "code",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票代码",
                "selectData": []
            },
            {
                "fpType": 28,
                "fpTypeName": "国外发票",
                "ocrType": 28,
                "inputKey": "fphm",
                'maxLangth':8,
                "inputLabel": "发票号",
                "inputImportant": 1,
                "inputType": "number",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入发票号",
                "selectData": []
            },
            {
                "fpType": 28,
                "fpTypeName": "国外发票",
                "ocrType": 28,
                "inputKey": "kprq",
                "inputLabel": "记账日期",
                "inputImportant": 1,
                "inputType": "date",
                "other": "YYYYMMDD",
                "inputLabel2": null,
                "placeholder2": "请输入记账日期",
                "selectData": []
            },
            {
                "fpType": 28,
                "fpTypeName": "国外发票",
                "ocrType": 28,
                "inputKey": "bz",
                "inputLabel": "币种",
                "inputImportant": 0,
                "inputType": "select",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "点击选择币种",
                "selectData": [
                    {"text": "比索", "value": "PHP"},
                    {"text": "兰特", "value": "ZAR"},
                    {"text": "新西兰元", "value": "NZD"},
                    {"text": "瑞士法郎", "value": "CHF"},
                    {"text": "福林", "value": "HUF"},
                    {"text": "瑞典克朗", "value": "SEK"},
                    {"text": "韩元", "value": "KRW"},
                    {"text": "林吉特", "value": "MYR"},
                    {"text": "新加坡元", "value": "SGD"},
                    {"text": "加元", "value": "CAD"},
                    {"text": "里亚尔", "value": "OMR"},
                    {"text": "卢布", "value": "SUR"},
                    {"text": "挪威克朗", "value": "NOK"},
                    {"text": "迪拉姆", "value": "AED"},
                    {"text": "里拉", "value": "ITL"},
                    {"text": "英镑", "value": "GBP"},
                    {"text": "欧元", "value": "EUR"},
                    {"text": "日元", "value": "JPY"},
                    {"text": "兹罗提", "value": "PLN"},
                    {"text": "丹麦克朗", "value": "DKK"},
                    {"text": "港元", "value": "HKD"},
                    {"text": "美元", "value": "USD"},
                    {"text": "澳元", "value": "AUD"},
                    {"text": "泰铢", "value": "THB"}
                ]
            },
            {
                "fpType": 28,
                "fpTypeName": "国外发票",
                "ocrType": 28,
                "inputKey": "wbje",
                "inputLabel": "外币金额",
                "inputImportant": 1,
                "inputType": "foreign-money",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入外币金额",
                "selectData": []
            },
            {
                "fpType": 28,
                "fpTypeName": "国外发票",
                "ocrType": 28,
                "inputKey": "fpje",
                "inputLabel": "人民币金额",
                "inputImportant": 1,
                "inputType": "money",
                "other": null,
                "inputLabel2": null,
                "placeholder2": "请输入人民币金额",
                "selectData": []
            }
        ],
        "value": 28
    }
]

export const typeList = [
    {
        "text": "定额发票",
        "value": "5"
    },
    {
        "text": "增值税电子发票",
        "value": "6"
    },
    {
        "text": "出租车票",
        "value": "2"
    },
    {
        "text": "飞机票",
        "value": "3"
    },
    {
        "text": "轮船票",
        "value": "7"
    },
    {
        "text": "汽车票",
        "value": "10"
    },
    {
        "text": "火车票",
        "value": "1"
    },
    {
        "text": "增值税普通发票",
        "value": "0"
    },
    {
        "text": "增值税专用发票",
        "value": "12"
    },
    {
        "text": "(卷票)加油票",
        "value": "8"
    },
    {
        "text": "机打发票",
        "value": "18"
    },
    {
        "text": "过路费",
        "value": "16"
    },
    {
        "text": "其他发票",
        "value": "22"
    },
    {
        "text": "其他票据",
        "value": "23"
    },
    {
        "text": "中央非税票据",
        "value": "27"
    },
    {
        "text": "国外发票",
        "value": "28"
    }
]
