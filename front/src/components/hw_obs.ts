import moment from "moment";
import {Toast} from 'vant';

const server = "";
const Ak = "";
const Sk = "";
const bucket = "";

export async function UpLoadImg(imageList: Array<any>) {
    let phone = "no-phone";
    let imgUrlResult: string[] = [];
    // 登录人手机号
    phone = ""
    let df = moment().format("yyyyMMDDHHmmssSSS");
    for (let item of imageList) {
        let fileName = `${phone}-${df}${item.file.name.substring(item.file.name.lastIndexOf('.'))}`
        //@ts-ignore
        let obsClient: any = window.document.getObsClient(Ak, Sk, server);
        //上传
        await obsClient.putObject({
            Bucket: bucket,
            Key: fileName,
            SourceFile: item.file
        }).then(function (result: any) {
            if (result.CommonMsg.Status < 300) {
                imgUrlResult.push(`/${fileName}`)
            }
        })
    }
    return imgUrlResult;
}
