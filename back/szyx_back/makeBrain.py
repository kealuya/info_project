import json
import os
import sys
import zipfile
import pymysql
import xmind
from xmind.core.const import TOPIC_DETACHED
from xmind.core.markerref import MarkerId
from xmind.core.topic import TopicElement


def repair(fname):
    zip_file = zipfile.ZipFile(fname, 'a')
    zip_file.writestr('META-INF/manifest.xml',
                      '<?xml version="1.0" encoding="UTF-8" standalone="no"?><manifest xmlns="urn:xmind:xmap:xmlns:manifest:1.0" password-hint=""></manifest>')
    zip_file.close()


def gen_my_xmind_file(orderId):
    current_path = os.getcwd()
    # load an existing file or create a new workbook if nothing is found
    file = current_path + "/" + orderId + ".xmind"
    workbook = xmind.load(file)
    # get the first sheet(a new workbook has a blank sheet by default)
    sheet1 = workbook.getPrimarySheet()
    design_sheet1(sheet1, orderId)
    # create sheet2
    # gen_sheet2(workbook, sheet1)
    # now we save as test.xmind
    xmind.save(workbook)
    repair(file)


def makeBrain(items, node):
    for key in items:
        if isinstance(items[key], list):
            for item in items[key]:
                sub_topic = node.addSubTopic()
                sub_topic.setTitle(item)
        else:
            sub_topic = node.addSubTopic()
            sub_topic.setTitle(key)
            makeBrain(items[key], sub_topic)


def design_sheet1(sheet1, orderId):
    # ***** first sheet *****
    sheet1.setTitle("脑图")  # set its title
    # python_obj = json.loads(getJson())
    python_obj = json.loads(getBrainJsonFromDb(orderId))

    root_topic1 = sheet1.getRootTopic()
    root_topic1.setTitle("会议纪要")
    makeBrain(python_obj, root_topic1)


def getBrainJsonFromDb(orderId):
    # 连接数据库
    db = pymysql.connect(host='124.70.48.30',
                         port=3306,
                         user='root',
                         password='renhao666',
                         database='info_project',
                         charset='utf8')
    # 创建游标对象
    cur = db.cursor()
    sql = "select brain from kdxf_knowledge where order_id = %s"
    # 执行sql语句
    cur.execute(sql, [orderId])
    # 打印读取的内容

    markdown_text = cur.fetchone()[0]
    markdown_dict = {}
    current_key = ""
    for line in markdown_text.split('\n'):
        if line.startswith("##"):
            current_key = line.strip("##").strip()
            markdown_dict[current_key] = {"内容": []}
        elif line.startswith("-"):
            if current_key:
                markdown_dict[current_key]["内容"].append(line.strip("-").strip())

    # print(markdown_dict)
    # 将字典转换为 JSON 格式
    json_data = json.dumps(markdown_dict, indent=4)
    # 解码 Unicode 转义字符
    decoded_str = bytes(json_data, "utf-8").decode("unicode_escape")
    print(decoded_str)
    cur.close()
    db.close()
    return decoded_str


def getJson():
    s = '''
  
{
  "会议纪要": {
    "报销流程讨论": {
      "内容": [
        "报销员指定报销项目提交审批。",
        "项目负责人审批后进行报销。",
        "团委部门结算其负担的费用部分。"
      ]
    },
    "学生团队出行报销处理": {
      "内容": [
        "存在院系层面的报销。",
        "需要按每个学生的明细检索并发起二次报销。",
        "二次报销选择院系所负担的项目，提交审批。",
        "记录第一次与第二次报销信息，进行第三次报销。",
        "所有报销层级需附上发票信息。"
      ]
    },
    "报销流程复杂性讨论": {
      "内容": [
        "当前流程复杂，考虑创建单独模块处理团委报销。",
        "对于团队出行的多个报销处理方式未确定。",
        "讨论是否应将团委报销单独处理。"
      ]
    },
    "报销明细表设计讨论": {
      "内容": [
        "设计了一个差旅报销汇总表。",
        "领队填写实际报销金额，计算应报和实报金额。",
        "多次选择同一行程申请单的控制方法未确定。"
      ]
    },
    "功能开发与评估": {
      "内容": [
        "老师希望前台集成多项功能。",
        "担心功能多而使用情况不佳，需做前评估。",
        "先梳理逻辑框架和业务流程，确保易用性。"
      ]
    },
    "后续步骤讨论": {
      "内容": [
        "确定先后顺序，先开发效果显著的部分。",
        "讨论整个逻辑框架的合理性和业务匹配度。",
        "确保老师能按照流程完成任务。"
      ]
    },
    "行动计划": {
      "内容": [
        "参与者需再次审视文档，梳理思路。",
        "下一步聚焦于确定可执行的部分和设计问题。",
        "决定哪些部分现在可做，哪些需调整或延后开发。"
      ]
    }
  }
}
 
    '''

    return s



# 执行方法（数据库写死了，可以改）
# python3 mytest2.py DKHJQ20240401161506467K5ULHBSLFPsQdyzq
if __name__ == '__main__':
    # current_path = os.getcwd()
    # print(current_path)
    # 一次输入两个参数，以空格分隔
    # input_str = input("请输入orderId: ")  # DKHJQ20240401161506467K5ULHBSLFPsQdyzq
    input_str = sys.argv[1]
    print("传入的参数:", input_str)
    #
    # # orderId = "DKHJQ20240401161506467K5ULHBSLFPsQdyzq"
    #
    gen_my_xmind_file(input_str)
