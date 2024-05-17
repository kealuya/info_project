package com.szhtjykj.speech.xfyun.speech.utils.kdxfResult;

import com.google.gson.Gson;
import lombok.Data;

import java.util.List;

/**
 * 获取语音转译结果  上传obs文件专用
 * @program: kdxf_speech
 * @description:
 * @packagename: utils.kdxfResult
 * @author: zhanbaohua
 * @date: 2024-05-17 15:55
 **/
public class kdxfAudioResult {

    private static final Gson gson = new Gson();

    public static String getContentAudioMeeting(String sb) {
        Result jsonParse = gson.fromJson(sb.toString(), Result.class);
        List<Lattice> latticeList = jsonParse.lattice;
        String allResult = "";
        for (int i = 0; i < latticeList.size(); i++) {
            Lattice tempLattice = latticeList.get(i);
            //Json_1best json_1best = gson.fromJson(tempLattice.json_1best, Json_1best.class);
//            List<Rt> rtList = tempLattice.st.rt;
            List<Rt> rtList = tempLattice.json_1best.st.rt;
            for (int j = 0; j < rtList.size(); j++) {
                Rt tempRt = rtList.get(j);
                List<Ws> wsList = tempRt.ws;
                for (int k = 0; k < wsList.size(); k++) {
                    Ws tempWs = wsList.get(k);
                    List<Cw> cwList = tempWs.cw;
                    for (int l = 0; l < cwList.size(); l++) {
                        Cw tempCw = cwList.get(l);
                        System.out.print(tempCw.w);
                        allResult = allResult + tempCw.w;
                    }
                }
            }
        }

        return allResult;
    }

}

@Data
class Result {
    List<Lattice> lattice;
}

@Data
class Lattice {
    Json_1best json_1best;
}
@Data
class Json_1best {
    St st;
}
@Data
class St {
    List<Rt> rt;
    String rl;
}
@Data
class Rt {
    List<Ws> ws;
}
@Data
class Ws {
    List<Cw> cw;
}
@Data
class Cw {
    String w;
}

