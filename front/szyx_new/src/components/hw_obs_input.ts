import moment from "moment";
import {Toast} from 'vant';

const server = "obs.cn-north-4.myhuaweicloud.com";
const Ak = "SNZGBWTDEF0IRJKXJGJF";
const Sk = "W3H3nbgxHU3zDAblqwvTjO18V6X9ZeIexyn7Ter1";
const bucket = "file-report-store";

export async function UpLoadImg1(imageList: Array<any>) {
    let phone = "no-phone";
    let imgUrlResult: string[] = [];
    // 登录人手机号
    phone = "15201618718"
    let df = moment().format("yyyyMMDDHHmmssSSS");
    for (let item of imageList) {
        let fileName = `${phone}-${df}${item.name.substring(item.name.lastIndexOf('.'))}`
        //@ts-ignore
        let obsClient: any = window.document.getObsClient(Ak, Sk, server);
        //上传
        await obsClient.putObject({
            Bucket: bucket,
            Key: fileName,
            SourceFile: item
        }).then(function (result: any) {
            if (result.CommonMsg.Status < 300) {
                imgUrlResult.push(`https://file-report-store.obs.cn-north-4.myhuaweicloud.com/${fileName}`)
            }
        })
    }
    return imgUrlResult;
}
