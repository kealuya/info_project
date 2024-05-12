export interface IAfterSaleQueryRequest {
    customerImportance: number;
    isInternational: number;
    oldEnterpriseId: string;
    oldProductType: number;
    orderNo: string;
    phone: string;
}
export interface paramsTye{
    currentPage: number,
    pageSize: number,
    corpCode: string | null | undefined,
    flag: string,
    userId: string
}

export interface LISTTYPE {
    taskTitle: string,
    taskId: string,
    createTime: string,
    flag: string,
    finishTime: string,
    taskData: string
}
export interface finishTaskType{
    taskId:	string,//	任务ID
    corpName:string,//	企业名称
    corpCode:string,//企业code
    userId:	string | undefined	//	用户ID
    userName:string | undefined,//用户姓名
    userMobile:string |undefined,	//	用户手机号
    meetingId:string	//	是	关联的会议ID，用于完成任务关联会议下的会议文件，多个会议用；割开
}
export interface meetingListType{
    createTime:string,
    creater:string,
    meetingAddress:string,
    meetingAudioFileUrl:string,
    meetingAudioMinutes:string,
    meetingAudioSummary:string,
    meetingBrainMapFileUrl:string,
    meetingCity:string,
    meetingDocumentMinutes:string,
    meetingDocumentSummary:string,
    meetingFile:string,
    meetingFileUrl:string,
    meetingFlag:string,
    meetingId:string,
    meetingMminutesFileUrl:string,
    meetingPeople:string,
    meetingTime:string,
    meetingTitle:string,
    meetingType:string,
    taskId:string,
    useStatus:string,
}