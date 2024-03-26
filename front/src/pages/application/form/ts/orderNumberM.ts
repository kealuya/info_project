export  const orderNumberM=() =>{
    //时间日期
    let datetime = new Date();
    let year = datetime.getFullYear();
    let month =
        datetime.getMonth() + 1 < 10
            ? "0" + (datetime.getMonth() + 1)
            : datetime.getMonth() + 1;
    let date =
        datetime.getDate() < 10 ? "0" + datetime.getDate() : datetime.getDate();
    let hh =
        new Date().getHours() < 10
            ? "0" + new Date().getHours()
            : new Date().getHours();
    let mf =
        new Date().getMinutes() < 10
            ? "0" + new Date().getMinutes()
            : new Date().getMinutes();
    //8位随机数
    let arr = [1, 2, 3, 4, 5, 6, 7, 8, 9, 0];
    let code = "";
    for (let i = 0; i < 8; i++) {
        let index = Math.floor(Math.random() * arr.length);
        code += arr[index];
    }
    return (
        year.toString() +
        month.toString() +
        date.toString() +
        hh.toString() +
        mf.toString() +
        code
    );
}
