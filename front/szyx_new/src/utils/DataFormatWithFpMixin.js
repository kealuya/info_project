import {
    FPLXKEY,findByWsyyType,findByOcrType
} from './fplx.js'
import moment from 'moment';

function ocrDataToShowData(data,type) {
    let resultArr = []
    let fpdm = data.fpdm || ""
    let fphm = data.fphm || ""
    let kprq = "";
    if(data.kprq != null){
        let tmpKprq = data.kprq.trim().replace(/-/g, "").substr(0, 8)
        kprq = moment(Number(tmpKprq), "YYYYMMDD").format("YYYY-MM-DD")
    }
    let jym = ""
    let yzm = data.yzm||""
    if((data.fplx == 6 || data.fplx == 0 || data.fplx == 8) && type == "hand"){
        if(data.fpjym != null){
            jym = data.fpjym.substring(data.fpjym.length-6)
        }
    }else {
        if(data.fpjym != null){
            jym = data.fpjym.substring(data.fpjym.length-6)
        }
    }
    if(data.jym != null){
        jym = data.jym.substring(data.jym.length-6)
    }
    let ocrType = data.fplx//ocrType
    let invType = findByOcrType(ocrType).WSYYTYPE;
    let fylb = data.xflx || ''
    let kpnr = data.kpnr || ''
    let fpje;
    if(data.fplx == '1' && type!== 'hand'){
        fpje = data.je || ''
    }else {
        fpje = data.fpje || ''
    }

    let fpjetax = data.sqje || ''
    let zfcg = data.zfcg || ''
    let kpdwmc = data.kpdwmc || ''
    let xcsj = data.xcsj || ''
    let zjh = data.zjh || ''
    let xm = data.xm || ''
    let dddd = data.dddd || ''
    let ph = data.ph || ''
    let pjhm = data.pjhm || ''
    let scsj = data.scsj || ''
    let cfdd = data.cfdd || ''
    let hbh = data.hbh || ''
    let pb = data.pb || ''
    let wbje = data.wbje || ''
    let cph = data.cph || ''
    let cfrq = data.cfrq || ''
    let cc = data.cc || ''
    let bz = data.bz || ''
    let base64 = data.base64 || ''
    let rq = data.rq || ''
    let xflx = data.xflx || ''
    let tempReqult = {
        fpdm: fpdm,
        fphm: fphm,
        kprq: kprq,
        jym: jym,
        yzm:yzm,
        ocrType: ocrType,
        invType: invType,
        fylb: fylb,
        kpnr: kpnr,
        fpje: fpje,
        fpjetax: fpjetax,
        zfcg: zfcg,
        kpdwmc: kpdwmc,
        xcsj: xcsj,
        zjh: zjh,
        xm: xm,
        dddd: dddd,
        ph: ph,
        pjhm: pjhm,
        scsj: scsj,
        cfdd: cfdd,
        hbh: hbh,
        pb: pb,
        wbje: wbje,
        cph: cph,
        cfrq: cfrq,
        cc: cc,
        bz: bz,
        base64: base64
    }
    switch (invType) {
        case FPLXKEY.ELECINV.WSYYTYPE://电子发票
            resultArr.push(tempReqult)
            break;
        case FPLXKEY.QUOTAINV.WSYYTYPE://定额发票
            tempReqult.fylb = xflx.trim()
            resultArr.push(tempReqult)
            break;
        // case PLXKEY.METROINV.WSYYTYPE://一卡通
        // 	break;
        case FPLXKEY.TAXIINV.WSYYTYPE://出租车票
            tempReqult.cph = data.cph
            tempReqult.zjh = data.czh
            let scsj_h = data.scsj !== 'null' || data.scsj.length != 4 ? "" : data.scsj.substring(0,2)
            let scsj_m = data.scsj !== 'null' || data.scsj.length != 4 ? "" : data.scsj.substring(2)
            let xcsj_h = data.xcsj !== 'null' || data.xcsj.length != 4 ? "" : data.xcsj.substring(0,2)
            let xcsj_m = data.xcsj !== 'null' || data.xcsj.length != 4 ? "" : data.xcsj.substring(2)
            if( scsj_h !== 'null' && scsj_m !== 'null' ){
                scsj = scsj_h +":" + scsj_m
            }
            if( xcsj_h !== 'null' && xcsj_m !== 'null' ){
                xcsj = xcsj_h +":" + xcsj_m
            }
            tempReqult.scsj  = scsj
            tempReqult.xcsj  = xcsj
            resultArr.push(tempReqult)
            break;
        case FPLXKEY.FLIGHTINV.WSYYTYPE://飞机票
            tempReqult.fylb = xflx.trim()
            if(type == "hand"){
                //let flightList = list
                /* for(let flightItem of flightList){*/
                let temp = {...tempReqult}
                let xm = data.xm
                let tmpcfrq1 = data.cfrq
                let cfrq = data.cfrq
                let cfdd = data.cfdd
                let dddd = data.dddd
                let hbh = data.hbh
                let fpje = data.fpje
                let fpdm = data.ph
                let fphm = data.ph
                let zfcg = data.zfcg
                let pb = data.pb
                temp.xm = xm
                temp.cfrq = cfrq
                temp.cfdd = cfdd
                temp.dddd = dddd
                temp.hbh = hbh
                temp.fpje = fpje
                temp.fpdm = fpdm
                temp.fphm = fphm
                temp.zfcg = zfcg
                temp.pb = pb
                resultArr.push(temp)
            }else {
            let flightList = data.list
            for(let flightItem of flightList){
            let temp = {...tempReqult}
              let xm = data.cjrxm
              let tmpcfrq1 = flightItem.cjrq.trim().replace(/-/g, "").substr(0, 8)
              let cfrq = moment(Number(tmpcfrq1), "YYYYMMDD").format("YYYY-MM-DD")
              let cfdd = flightItem.fz
              let dddd = flightItem.dz
              let hbh = flightItem.cc
              let fpje = flightItem.je
              let fpdm = data.ph
              let fphm = flightItem.cjrq + flightItem.sj
              let zfcg = data.zfcgjgh
              let pb = flightItem.pb
              let ph = flightItem.ph
                temp.xm = xm
                temp.cfrq = cfrq
                temp.cfdd = cfdd
                temp.dddd = dddd
                temp.ph = ph
                temp.hbh = hbh
                temp.fpje = fpje
                temp.fpdm = fpdm
                temp.fphm = fphm
                temp.zfcg = zfcg
                temp.pb = pb
                resultArr.push(temp)
             }
        }
            break;
        case FPLXKEY.SHIPINV.WSYYTYPE://轮船票
            resultArr.push(tempReqult)
            break;
        case FPLXKEY.CARINV.WSYYTYPE://汽车票
            // tempReqult.fylb = data.xflx.trim();
            tempReqult.fylb ='交通费';
            resultArr.push(tempReqult)
            break;
        case FPLXKEY.TRAININV.WSYYTYPE://火车票
            let tmpcfrq2 = rq.trim().replace(/-/g, "").substr(0, 8)
            // let cfrqshH = moment(Number(tmpcfrq2), "YYYYMMDD").format("YYYY-MM-DD")
             let cfrqshH = data.cfrq
            let xmH = data.xm
            let cfddH = data.cfdd
            let ddddH = data.dddd
            let pbH = data.pb
            let ccH = data.cc
            let fpjeH;
            if(type!== 'hand'){
                fpjeH = data.je;
            }else {
                fpjeH = data.fpje || '';
            }
            tempReqult.cfrq = cfrqshH
            tempReqult.xm = xmH
            tempReqult.cfdd = cfddH
            tempReqult.dddd = ddddH
            tempReqult.pb = pbH
            tempReqult.cc = ccH
            tempReqult.fpje = fpjeH
            resultArr.push(tempReqult)
            break;
        case FPLXKEY.ADDVTAXINV.WSYYTYPE://增值税普通发票
            resultArr.push(tempReqult)
            break;
        case FPLXKEY.ADDVSPECIALINV.WSYYTYPE://增值税专用发票
            tempReqult.fpje = data.sqje || ''
            tempReqult.fpjetax = data.fpje || ''
            resultArr.push(tempReqult)
            break;
        case FPLXKEY.REFUELINV.WSYYTYPE://(卷票)加油票
            resultArr.push(tempReqult)
            break;
        case FPLXKEY.PRINTEDINV.WSYYTYPE://机打发票
            resultArr.push(tempReqult)
            break;
        case FPLXKEY.TOLLINV.WSYYTYPE://过路费
            resultArr.push(tempReqult)
            break;
        case FPLXKEY.OTHERINV.WSYYTYPE://其他发票
            tempReqult.fylb = xflx.trim()
            resultArr.push(tempReqult)
            break;
        case FPLXKEY.OTHERBILLS.WSYYTYPE://其他票据
            resultArr.push(tempReqult)
            break;
        case FPLXKEY.CENTRENONTAXINV.WSYYTYPE://中央非税票据
            // tempReqult.fylb = data.xflx.trim()
            tempReqult.fylb = '中央非税票据'
            resultArr.push(tempReqult)
            break;
        case FPLXKEY.FOREIGN.WSYYTYPE://国外发票
            resultArr.push(tempReqult)
            break;
    }
    return resultArr
}
function showDataToVerifyData(data) {
    let fpdm = data.fpdm || ''
    let fphm = data.fphm || ''
    let kprq = data.kprq || ''
    if(data.kprq != null && data.kprq != ""){
        //处理日期格式 去掉- 保存发票接口定义
        let kprqNew = data.kprq.replace('-','')
        kprq =  kprqNew.replace('-','')
    }


    let jym = data.jym || ''
    let yzm = data.yzm || ''
    let pickey2 = data.pickey2 || ''
    let pickey3 = data.pickey3 || ''
    let isfpyzm = data.isfpyzm || ''
    let zje = data.fpje || ''
    let zje_tax = data.fpjetax || ''
    let fplx = data.fplx || ''
    let funcClassify = data.funcClassify || ''

    let invType = data.invType
    if(data.invType==undefined||data.invType==null){
        invType = ""
    }
    let travellerName = data.travellerName || ''
    let departTime = data.departTime || ''
    let arriveTime = data.arriveTime || ''
    let departSite = data.departSite || ''
    let arriveSite = data.arriveSite || ''
    let vehicleNum = data.vehicleNum || ''
    let pb = data.pb || ''
    let ph = data.ph || ''
    let kpnr = data.kpnr || ''
    let kpdwmc = data.kpdwmc || ''
    let fwlb = data.fwlb || ''
    let fileBase64 = data.base64 || ''
    let fileName = data.ocrImageName || ''
    let ywlx = data.businessType || '113'
    let invSource = "0" //0 发票管理 1待稽核 2一键报销
    let fpfjuuid = data.fpfjuuid || '' //录入发票时关联的附件uuid
    let cph = data.cph || '' //车牌号（打车票）
    let zjh = data.zjh || '' //证件号（打车票）
    let scsj = data.scsj || ''
    let xcsj = data.xcsj || ''

    if (fplx == 0 || fplx == 9 || fplx == 10 || fplx == 11) {
        fplx = findByWsyyType(invType).VERIFYTYPE
    }
    let invTypeStr = invType + ""
    switch (invTypeStr) {
        case FPLXKEY.ELECINV.WSYYTYPE://电子发票
            break;
        case FPLXKEY.QUOTAINV.WSYYTYPE://定额发票
            fwlb = data.fylb
            break;
        // case PLXKEY.METROINV.WSYYTYPE://一卡通
        // 	break;
        case FPLXKEY.TAXIINV.WSYYTYPE://出租车票
            cph = data.cph || ''
            zjh = data.zjh || ''
            zje = data.fpje
            scsj = data.scsj
            xcsj = data.xcsj
            break;
        case FPLXKEY.FLIGHTINV.WSYYTYPE://飞机票
            travellerName = data.xm
            departTime = data.cfrq
            departSite = data.cfdd
            arriveSite = data.dddd
            pb = data.pb
            vehicleNum = data.hbh
            fpdm = departSite + departTime
            ph = data.ph
            // if(fphm != null){
            //     fphm = travellerName
            // }
            pickey3 = data.zfcg
            break;
        case FPLXKEY.SHIPINV.WSYYTYPE://轮船票
            break;
        case FPLXKEY.CARINV.WSYYTYPE://汽车票
            break;
        case FPLXKEY.TRAININV.WSYYTYPE://火车票
            travellerName = data.xm
            departTime = data.cfrq
            arriveTime = departTime
            departSite = data.cfdd
            arriveSite = data.dddd
            pb = data.pb
            vehicleNum = data.cc
            fpdm = vehicleNum;
            fphm = data.ph;
            break;
        case FPLXKEY.ADDVTAXINV.WSYYTYPE://增值税普通发票
            break;
        case FPLXKEY.ADDVSPECIALINV.WSYYTYPE://增值税专用发票
            break;
        case FPLXKEY.REFUELINV.WSYYTYPE://(卷票)加油票
            break;
        case FPLXKEY.PRINTEDINV.WSYYTYPE://机打发票
            break;
        case FPLXKEY.TOLLINV.WSYYTYPE://过路费
            break;
        case FPLXKEY.OTHERINV.WSYYTYPE://其他发票
            break;
        case FPLXKEY.OTHERBILLS.WSYYTYPE://其他票据
            fpdm = data.invType
            fphm = data.pjhm
            kpnr = data.kpnr
            kpdwmc = data.kpdwmc
            break;
        case FPLXKEY.CENTRENONTAXINV.WSYYTYPE://中央非税票据
            break;
        case FPLXKEY.FOREIGN.WSYYTYPE://国外发票
            break;
    }


    let result = {
        "fpdm": fpdm,
        "fphm": fphm,
        "kprq": kprq,
        "jym": jym,
        "yzm": yzm,
        "pickey2": pickey2,
        "pickey3": pickey3,
        "isfpyzm": isfpyzm,
        "zje": zje,
        "zje_tax": zje_tax,
        "fplx": fplx,
        "funcClassify": funcClassify,
        "invType": invType,
        "travellerName": travellerName,
        "departTime": departTime,
        "arriveTime": arriveTime,
        "departSite": departSite,
        "arriveSite": arriveSite,
        "vehicleNum": vehicleNum,
        "pb": pb,
        "ph": ph,
        "kpnr": kpnr,
        "kpdwmc": kpdwmc,
        "fwlb": fwlb,
        "fileBase64": '',
        "fileName": fileName,
        "ywlx": ywlx,
        "invSource": "0", //0 发票管理 1待稽核 2一键报销
        "fpfjuuid": fpfjuuid, //录入发票时关联的附件uuid
        "cph": cph, //车牌号（打车票）
        "zjh": zjh, //证件号（打车票）
        "scsj": scsj,
        "xcsj": xcsj
    }
    return result
}
function getUuid() {
    var len = 48; //48长度
    var radix = 16; //16进制
    var chars = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz'.split('');
    var uuid = [],
        i;
    radix = radix || chars.length;
    if (len) {
        for (i = 0; i < len; i++) uuid[i] = chars[0 | Math.random() * radix];
    } else {
        var r;
        uuid[8] = uuid[13] = uuid[18] = uuid[23] = '-';
        uuid[14] = '4';
        for (i = 0; i < 36; i++) {
            if (!uuid[i]) {
                r = 0 | Math.random() * 16;
                uuid[i] = chars[(i == 19) ? (r & 0x3) | 0x8 : r];
            }
        }
    }
    return uuid.join('');
}
export default {
    ocrDataToShowData,
    showDataToVerifyData,
    getUuid
}
