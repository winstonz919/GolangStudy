'''
Author: WinstonRD
Date: 2025-03-31 20:15:53
LastEditors: WinstonRD && winstonz919@sina.com
LastEditTime: 2025-03-31 20:31:09
FilePath: /GolangStudy/9-map/args.py
Description: 

Copyright (c) 2025 by WinstonRD, All Rights Reserved. 
'''
# 生成一段python代码，解析以下数据，并输出data_struct下的两个数据jump_link和pic
# data = { "code": 0, "data": "\n![返回图片](https://static.shutu.cn/shutu/jpeg/open5a/2025/03/31/d2e444aba301f77f8545dd8c71ae4cc2.jpeg)\n\n[编辑](https://gapi.shutu.cn/ai/edit-mind-url?works_guid=open5a73b19322ad677e70b466681ef46565_coze)\n\n如果觉得这个思维导图还不够完美，或者你的想法需要更自由地表达，点击编辑按钮，将你的思维导图变形、变色、变内容、甚至可以添加新的元素，快来试试吧！。\n", "data_struct": { "jump_link": "https://gapi.shutu.cn/ai/edit-mind-url?works_guid=open5a73b19322ad677e70b466681ef46565_coze", "pic": "https://static.shutu.cn/shutu/jpeg/open5a/2025/03/31/d2e444aba301f77f8545dd8c71ae4cc2.jpeg" }, "log_id": "2025033120143470E885432E820654EAED", "msg": "success", "status_code": null, "type_for_model": 2, "errorBody": null }

def parse_data(data):
    data_struct = data.get('data_struct')
    if data_struct:
        jump_link = data_struct.get('jump_link')
        pic = data_struct.get('pic')
        return jump_link, pic
    return None, None

if __name__ == "__main__":
    data = { "code": 0, "data": "\n![返回图片](https://static.shutu.cn/shutu/jpeg/open5a/2025/03/31/d2e444aba301f77f8545dd8c71ae4cc2.jpeg)\n\n[编辑](https://gapi.shutu.cn/ai/edit-mind-url?works_guid=open5a73b19322ad677e70b466681ef46565_coze)\n\n如果觉得这个思维导图还不够完美，或者你的想法需要更自由地表达，点击编辑按钮，将你的思维导图变形、变色、变内容、甚至可以添加新的元素，快来试试吧！。\n", "data_struct": { "jump_link": "https://gapi.shutu.cn/ai/edit-mind-url?works_guid=open5a73b19322ad677e70b466681ef46565_coze", "pic": "https://static.shutu.cn/shutu/jpeg/open5a/2025/03/31/d2e444aba301f77f8545dd8c71ae4cc2.jpeg" }, "log_id": "2025033120143470E885432E820654EAED", "msg": "success", "status_code": null, "type_for_model": 2, "errorBody": null }
    print(data)
    ret1, ret2 = parse_data(data)
    print(ret1, ret2)