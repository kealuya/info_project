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