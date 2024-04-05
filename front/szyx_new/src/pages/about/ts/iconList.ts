import planeIcon from "../../../assets/icon/icon_01.png";
import hotelIcon from "../../../assets/icon/icon_02.png";
import invoiceIcon from "../../../assets/icon/icon_07.png";
import trainIcon from "../../../assets/icon/icon_03.png";
import useCarIcon from "../../../assets/icon/icon_04.png";
import buyIcon from "../../../assets/icon/icon_06.png";
import emsIcon from "../../../assets/icon/icon_05.png";
import budgetIcon from "../../../assets/icon/icon_09.png";
import costCenterIcon from "../../../assets/icon/icon_21.png";
// 静态默认的图标数组
export const staticCodeIconList=[
    'plane','hotel','invoice','car','train','buy','log',
]
// 全部首页图标功能
export const staticIconListA=[
    {icon:planeIcon,title:'飞机票',code:'plane',path:'thirdPartyBusiness/plane'},
    {icon:hotelIcon,title:'住酒店',code:'hotel',path:'thirdPartyBusiness/hotel'},
    {icon:invoiceIcon,title:'发票夹',code:'invoice',path:'invoiceList'},
    {icon:trainIcon,title:'火车票',code:'train',path:'thirdPartyBusiness/train'},
    {icon:useCarIcon,title:'用车',code:'car',path:'expressDeliveryOrCarPage'},
    {icon:buyIcon,title:'心意选',code:'buy',path:'thirdPartyBusiness/shop'},
    {icon:emsIcon,title:'快递',code:'log',path:'expressDeliveryOrCarPage'},
    {icon:budgetIcon,title:'预算管理',code:'budget',path:'budget'},
    {icon:costCenterIcon,title:'成本中心',code:'costCenter',path:'costCenter'},
];
